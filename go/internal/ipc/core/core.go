// Package core describes the primitives of the cchat IPC protocol.
//

package core

//go:generate go run github.com/tinylib/msgp -tests=false -o=msgp.go -file=.

import "github.com/tinylib/msgp/msgp"

// ErrorType is the type of the error, which determines its Data type.
type ErrorType uint8

// RegularError defines an error whose type is of a string.
const RegularErrorType ErrorType = 0

// Error describes an error over the wire.
type Error struct {
	Type ErrorType `msg:"t"`
	Data msgp.Raw  `msg:"d"`
}

// Time describes a timestamp without a timezone; it is in UnixNano format.
type Time int64

// Reader describes possible readers. This structure should be treated like a
// union, that is, only one of those fields should be used.
type Reader struct {
	Stream *Stream `msg:"stream,omitempty"`
	File   *File   `msg:"file,omitempty"`
}

// File describes a physical file on disk. Backends that read this file must
// not modify nor delete it.
type File struct {
	Filename string `msg:"filename"`
}

// Stream describes the information of a stream ticket along with extraneous
// metadata.
type Stream struct {
	Ticket StreamTicket `msg:"t"`
}

// StreamTicket is a unique identifier for each stream.
type StreamTicket uint64

// StreamPacket is a possible event that is an incoming piece of a stream. It
// contains the byte array of arbitrary length, usually segmented using a
// constant in the implementation, to send over to the client.
type StreamPacket struct {
	Ticket StreamTicket `msg:"t"`
	// Bytes contains the segment of data belonging to the stream.
	// Implementations should keep this slice small and prefer sending multiple
	// data packets over the connection to avoid clogging other events.
	//
	// The length of the byte arrays must cumulatively add up to the requested
	// amount at maximum; the sender should ensure this. The receiver will wait
	// until either an error occurs or the requested amount is filled, so the
	// sender must accomodate for either scenario accordingly.
	Bytes []byte `msg:"b,omitempty"`
	// Error is the optional error that a stream may emit. The stream is
	// considered to be invalid after an error has been emitted and must not be
	// used again.
	//
	// An EOF error must be sent over the wire with the associated ticket to
	// indicate the end of a stream. The error can be attached with the bytes.
	Error *Error `msg:"e,omitempty"`
}

// MaxStreamAmount defines the maximum number of bytes that the receiver could
// request the sender to send over. Both the receiver and the sender should
// adhere to this constand and error out when the receiver requests a number
// higher than this.
const MaxStreamAmount = 128 * 1024 * 1024

// StreamRequest is sent by the stream receiver to request more byte packets.
// The receiver should call this when the buffer is drained and block until more
// data arrives.
type StreamRequest struct {
	Ticket StreamTicket `msg:"t"`
	// Amount is the requested amount of bytes to send over. Note that the
	// sender must only send maximum this amount of bytes or less; the receiver
	// must ignore extraneous bytes sent over until it requests for more. This
	// is to allow the use of fixed buffers.
	//
	// If a zero amount is sent, the sender must interrupt or close the stream
	// entirely.
	Amount uint32 `msg:"amount"`
}

// CallID is the ID of a call. How the ID is generated is up to the
// implementation.
type CallID uint64

// StopHandle is the handle ID to stop or cancel an object. This could be
// understood as a destructor callback to destroy an object. It is not to be
// confused with ContextID, which is used to cancel an ongoing call.
type StopHandle uint64

// InstanceID describes an object ID number that identifies an object instance.
type InstanceID uint64

// InstanceType is the enumeration of all possible instance types.
// Implementations should automatically generate the enumeration values from the
// main schema.
type InstanceType uint16

// Instance describes an interface instance over the wire. Note that struct and
// primitive types are marshalled as-is and are not considered objects; only
// interfaces are treated as objects with each having an object ID.
//
// Asserters
//
// Underneath, asserters are implemented as a struct containing interfaces, each
// corresponding to a possible children instance. If an asserted interface is
// nil and a method of that instance is called, then the server must attempt to
// reassert the interface by calling the As() method again. If that is still
// nil, then the call is erroneous.
//
// Children Instances
//
// For the most part, a children instance should keep its top-most parent's
// instance ID. By doing this, the children instance's lifespan is inherently
// dependent on its parent: if the parent is destroyed or is no longer valid,
// then so are all its children instances.
//
// However, there are circumstances when this is not strictly true. If an
// instance has explicit disposers/destructors such as an IOMethod marked as
// Disposer or a ContainerMethod with a stop handle, then that child instance
// (and all of its children instances) can be explicitly destroyed independently
// of the parent. Any call using the parent instance into its children's
// children is therefore invalid, except on calls to the children itself (the
// second level, parent being the first), then the server must attempt to
// reassert the children accordingly. Note that this does not mean that this
// specific child instance can outlive its parent; it simply means the child
// instance can be destroyed and created independently of the parent's lifetime.
//
// During code generation, a method path unique to the parent instance type can
// be generated in the call function without needing to store it during runtime,
// as the method paths will be constant.
//
// Pure Instances
//
// All methods within an instance that is of type GetterMethod and has no
// arguments must be resolved within before being sent over the wire. Structs
// and instances with only these methods can consider its type to be "pure" and
// therefore do not need instance IDs, since it doesn't need to be refered to
// again in the future.
//
// Instances within a Struct
//
// To ease memory management, structs must not have methods that require
// arguments. In other words, they must be pure. With this assumption, both ends
// need not store the struct inside a registry for future calls, as all values
// can be prefetched once and loaded all over the wire.
//
// Code generators should be aware of this assumption and fail appropriately
// when it detects a violation. Any violation of this assumption is considered a
// cchat bug and should be reported.
type Instance struct {
	// ID is the ID of this object. It is omitted if the instance type is pure.
	ID InstanceID `msg:"id,omitempty"`
	// Type is the type of this object. The IPC implementation (such as this
	// one) can use this information to infer the name of the instance.
	Type InstanceType `msg:"typ"`

	// StopHandle is the handle ID to stop or destroy the object.
	StopHandle StopHandle `msg:"s,omitempty"`
}

// MethodID is the enumerated method identifier that's specific to each instance
// type.
type MethodID uint8

// MethodPath is the list of fields to traverse and call, with the last ID in
// the list to be the method ID of the instance within the parent. This is most
// commonly used for instances with nested asserts.
//
// Example
//
// Take for example these types:
//
//    type Parent interface {
//        UnrelatedMethod()
//        AsChild() Child
//    }
//
//    type Child interface {
//        Method(param Parameter)
//    }
//
// Given these types, we know beforehand from just the declaration that type
// Parent has 2 methods, one is the AsChild asserter, and type Child also has 1
// method, which is Method. If the Parent instance has an ID of 2, then we can
// call Parent.AsChild().Method() like so:
//
//    Call{
//        InstanceID: 2,
//        CallID:     0,
//        MethodPath: MethodPath{1, 0},
//        Parameters: MarshalMsgp(Parameter{}),
//    }
//
// To visualize it better, this call convention can be represented using the
// following diagram:
//
//    - Parent (id: 2)          <- InstanceID = 2
//        - UnrelatedMethod: 0
//        - AsChild: 1          <- MethodPath[0] = 1
//            - Method: 0       <- MethodPath[1] = 0
//
type MethodPath []MethodID

// Call is the header of a single IPC call. It indicates which instance
// the call belongs to. A body of parameters must follow after.
type Call struct {
	// InstanceID describes the instance that this call was intended for. In the
	// usual terms, it's similar to a method receiver pointer.
	InstanceID InstanceID `msg:"oid"`

	// CallID is used to identify (as a nonce) each method call and its response
	// on an object. Note that call IDs are managed solely by the side that
	// invokes the call, so in a way, both the frontend and backend manages its
	// own registry of call IDs.
	//
	// All synchronous calls must have a CallID; synchronous calls without an ID
	// is considered erroneous. An asynchronous call does not need a CallID.
	CallID CallID `msg:"cid,omitempty"`

	// MethodPath is the path to the method to be called from the parent's
	// instance type. Refer to type MethodPath for more information.
	MethodPath MethodPath `msg:"mp"`

	// Parameters is the field for an arbitrary parameter type. This type is
	// usually an array, but it doesn't have to be; it depends on the method and
	// object that the call is for.
	//
	// If the method being called can take multiple parameters, that is they're
	// one of the types GetterMethod, SetterMethod or IOMethod, then Parameters
	// will be a tuple of parameters that must match the length required.
	//
	// If the method being called is of type ContainerMethod, then the type is
	// of whatever ContainerType is.
	Parameters msgp.Raw `msg:"p"`
}

// Reply is the object type for a single IPC reply; it is followed after an IPC
// call with the same CallID. This is used for synchronous calls only.
//
// A reply should remove the call ID from the contexts registry. Refer to
// CancelCall for more information.
type Reply struct {
	// CallID is the same ID as the ID used in a call.
	CallID CallID `msg:"cid"`

	// Error is an additional error in the reply. It is only used for calls that
	// can error out.
	Error *Error `msg:"e,omitempty"`

	// StopHandle is the handle ID of a stop request. This field is only present
	// for ContainerMethods with the HasStopFn field set to true.
	StopHandle StopHandle `msg:"sh,omitempty"`

	// Return is the return value of a call. For example, this type could return
	// a new Instance. The type of this field should be determined based on the
	// parent object ID of the call that this reply was associated with.
	Return msgp.Raw `msg:"ret,omitempty"`
}

// CancelCall is the object type for requesting cancellation of an active IPC
// call. This request is done asynchronously; there is no reply. In Go, this is
// often implemented using the context.Context API. Methods that don't accept a
// context.Context will not be cancellable, and cancelling such call is
// erroneous. The server can ignore erroneous cancel calls.
//
// Note that a reply will invalidate itself so that any future CancelCall with
// the same ID will either be invalid or not be used for the old call anymore;
// as such, careful synchronization must be done in case a context is held for
// longer than the call.
type CancelCall struct {
	// CallID is the same ID as the ID used in a call.
	CallID CallID `msg:"cid"`
}
