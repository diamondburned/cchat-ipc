# cchat-ipc

## DISCARD! Service Architecture

When generating IPC code for languages, it is recommended to be mindful about
the separation of "backend" and "frontend".

### Core

Core defines all core types that are commonly shared across the backend and
frontend. It serves as language-agnostic implementations of core Go types, such
as `time.Time` or `error`, as well as supplement types .

### Common

Common contains all the common types that are unidentifiable and
indistinguishable from either directions of communication. Such types include
`Identifier` or `Namer`.

### Backend

Backend contains all interfaces that services will implement. This usually
implies any interfaces that do not have a "Container" as its suffix.

### Frontend

Frontend contains all services that the frontend implementation has. Such
services mostly have "Container" as its suffix.

All frontend services must not return anything after its call. Backends must
assume that any operation sent to the frontend is asynchronous and will always
succeed.

## Type Rules

- `A`
- For `ContainerMethod`:
	- Arguments type is `ContainerType`.
- 

## Scratchboard

```go
// generated from all error types.
const (
	AuthenticateErrorType ErrorType = iota + 1
)

// Authenticator has method Authenticate which is an IOMethod, so it is treated
// as an Instance.
type Authenticator struct {
	inst Instance

	authenticator cchat.Authenticator

	AuthenticateForm []AuthenticateEntry `msg:"authenticate_form"`
	Description      text.Rich           `msg:"description"`
	Name             text.Rich           `msg:"name"`
}

// AuthenticateEntry is directly translated 
type AuthenticateEntry struct {
	Name        string `msg:"name"`
	Placeholder string `msg:"placeholder"`
	Description string `msg:"description"`
	Secret      bool   `msg:"secret"`
	Multiline   bool   `msg:"multiline"`
}

// AuthenticateError does not have any methods that require explicit calling,
// therefore it is not treated as an Instance.
type AuthenticateError struct {
	NextStage []Authenticator `msg:"next_stage"`
	Error     string          `msg:"error"`
}

// wrapTYPE is used for wrapping return types. Reuse the same
// interface-to-instance marshal code.
//
// wrapTYPE : cchat.TYPE -> msgp TYPE

func (inst *Instance) wrapAuthenticateError(v cchat.AuthenticateError) AuthenticateError {
	// array generation code
	//    v: their variable
	//    _: our variable

	// preprocessing

	vNextStage := v.NextStage()
	_nextStage := make([]Authenticator, len(vNextStage))
	for i, authenticator := range vNextStage {
		_nextStage[i] = c.wrapAuthenticator(authenticator)
	}

	authenticateError := AuthenticateError{
		NextStage: _nextStage,
		Error:     v.Error(),
	}

	// postprocessing

	return authenticateError
}

func (inst *Instance) wrapAuthenticator(v cchat.Authenticator) Authenticator {
	// preprocessing
	vAuthenticateForm := v.AuthenticateForm()
	_authenticateForm := make([]AuthenticateForm, len(vAuthenticateForm))
	for i, authenticateForm := range vAuthenticateForm {
		_authenticateForm[i] = c.wrapAuthenticateForm(authenticateForm)
	}

	authenticator := Authenticator{
		authenticator:    v,
		AuthenticateForm: _authenticateForm,
		Description:      c.wrapTextRich(v.Description()),
		Name:             c.wrapTextRich(v.Name()),
	}

	// postprocessing

	// Add this because Authenticator has IOMethod.
	authenticator.inst = c.conn.registerInstance(authenticator)

	return authenticator
}

func (inst *Instance) wrapAuthenticateForm(v cchat.AuthenticateForm) AuthenticateForm {
	var authenticateForm AuthenticateForm
	unsafefuckery.Cast(v, &authenticateForm)
	return authenticateForm
}

// wrapError is useless because we already know the error type to use in the
// schema.

func (inst *Instance) wrapError(err error) msgp.Encodable {
	if err == nil {
		return nil
	}

	switch err := err.(type) {
	case cchat.AuthenticateError:
		return c.wrapAuthenticateError(err)
	default:
		return RegularError(err.Error())
	}
}

// InstanceType is an enum type that signifies the type of the other end's
// instance. It is implemented this way rather than a typed structure to
// be efficient in memory usage.
type InstanceType uint64

const (
	InvalidInstanceType InstanceType = iota
	MessagesContainerInstance
)

type contextMap struct {
	mutex    sync.Mutex
	contexts map[CallID]context.CancelFunc
}

type Conn struct {
	Connector
	reg *InstanceRegistry
}

type InstanceRegistry struct {
	// TODO: separate these to InstanceRegistry
	instances slab.SlabMutex // InstanceID -> interface{}
	stopFuncs slab.SlabMutex // StopHandle -> func()

	contexts contextMap // CallID -> context.CancelFunc
}

type Instance struct {
	reg  *InstanceRegistry
	conn *Conn
	id   *uint64
}

func (inst *Instance) ID() InstanceID {
	return InstanceID(atomic.LoadUint64(inst.id))
}

func (inst *Instance) Invalidate() {
	atomic.StoreUint64(inst.id, 0)
}

func (conn *Conn) registerInstance(v interface{}) Instance {
	id := new(uint64)
	*id = conn.reg.instances.Put(v)

	return Instance{
		conn:   conn,
		reg:    conn.reg,
		instID: id,
	}
}

// Backend

// onCall is called on a Call OP.
func (conn *Conn) onCall(call Call) (interface{}, error) {
	v := conn.instances.Get(uint64(call.InstanceID))
	if v != nil {
		return nil, ErrUnknownInstance
	}

	switch v := v.(type) {
	case cchat.Messenger:
		return conn.callMessenger(call)
	default:
		return nil, ErrUnknownInstanceType // BUG
	}
}

func (conn *Conn) newContext(callID CallID) (context.Context, context.CancelFunc) {
	conn.contexts.
	ctx, cancel := context.Background()
	id := conn.contexts.Put(cancel)
	return ctx, ContextID(id)
}

func (conn *Conn) cancelContext(callID CallID) {
	cancel, ok := conn.contexts.Pop(uint64(id)).(context.CancelFunc)
	if !ok {
		return
	}
	cancel()
}

func (conn *Conn) registerStop(stop func()) StopHandle {
	id := conn.stopFuncs.Put(stop)
	return StopHandle(id)
}

// messenger.go

// Equivalent of only giving the instance.
type _MessengerJoinServerArgs InstanceID

const (
	MessengerJoinServer MethodID = iota
)

// errors returned from callMessenger are local (probably) fatal errors that
// shouldn't be echoed back. The caller will automatically set Reply's CallID.
func (conn *Conn) callMessenger(m cchat.Messenger, call *Call, ret *Return) (*Reply, error) {
	var err error

	switch call.Method {
	case MessengerJoinServer:
		var args MessengerJoinServerArgs
		call.Arguments, err = args.UnmarshalMsg([]byte(call.Arguments))
		if err != nil {
			return nil, msgp.WrapError(err, "JoinServer")
		}

		ctx, ctxID := conn.newContext(call.ID)

		stop, err := m.JoinServer(
			ctx,
			MessagesContainer{
				inst: conn.instect(InstanceID(args)),
			},
		)
		if err != nil {
			conn.stopContext(ctxID)
			return &Reply{
				Error: &Error{
					Type: RegularErrorType,
					Data: RegularError(err.Error()),
				},
			}, nil
		}

		return &Reply{
			// If this StopHandle is called, then InstanceID(args) is
			// automatically invalid. The backend implementation must not call
			// after stopping, else it's undefined behavior.
			//
			// Nonetheless, should we provide safety guarantees? That would
			// require keeping a pointer to the Instance and an atomic ID.
			StopHandle: conn.registerStop(stop),
		}, nil
	}
}

func (conn *Conn) callAsync(call *Call) {
	if call.InstanceID == 0 {
		return
	}
}

func (conn *Conn) call(call *Call) *Reply {}

type Noncer struct {
	Nonce string `msg:"nonce"`
}

type MessageCreate struct {
	MessageHeader
	Noncer

	Mentioned bool      `msg:"mentioned"`
	Content   text.Rich `msg:"content"`
	Author    Author    `msg:"author"`
}

type MessagesContainer struct {
	inst Instance
}

const (
	MessagesContainerCreateMessage MethodID = iota
	MessagesContainerUpdateMessage
	MessagesContainerDeleteMessage
)

func (m MessagesContainer) CreateMessage(v cchat.MessageCreate) {
	messageCreate := MessageCreate{
		MessageHeader: m.inst.wrapMessageHeader(v),
		Noncer:        m.inst.wrapNoncer(v),
		Mentioned:     v.Mentioned(),
		Content:       m.inst.wrapTextRich(v.Content()),
		Author:        m.inst.wrapAuthor(v.Author()),
	}

	var err error
	__encoded := make([]byte, messageCreate.Msgsize())
	__encoded, err = messageCreate.MarshalMsg(__encoded)
	if err != nil {
		m.handleError(err)
		return
	}

	// SetterMethod, therefore synchronous call and no CallID needed.
	m.inst.conn.callAsync(&Call{
		InstanceID: m.inst.ID(),
		Method:     MessagesContainerCreateMessage,
		Parameters: msgp.Raw(__encoded),
	})
}
```
