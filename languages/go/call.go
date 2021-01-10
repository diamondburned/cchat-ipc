package cchatipc

import "github.com/tinylib/msgp/msgp"

//go:generate go run github.com/tinylib/msgp

// CallID is the ID of a call. How the ID is generated is up to the
// implementation.
type CallID uint64

// ContextID is the ID of a context. How the ID is generated is up to the
// implementation.
type ContextID uint64

// StopHandle is the handle ID to stop or cancel an object. This could be
// understood as a destructor callback to destroy an object. It is not to be
// confused with ContextID, which is used to cancel an ongoing call.
type StopHandle uint64

// ObjectID describes an object ID number that identifies an object instance.
type ObjectID uint64

// Object describes an object over the wire.
type Object struct {
	// ID is the ID of this object.
	ID ObjectID `msgp:"id"`
	// Name is the name of this object. It could be used to look up the schema
	// for a list of method names.
	Name string `msgp:"name"`

	// StopHandle is the handle ID to stop or destroy the object.
	StopHandle StopHandle `msgp:"s,omitempty"`
}

// Call is the object type for a single IPC call.
type Call struct {
	// ObjectID describes the instance that this call was intended for. In the
	// usual terms, it's similar to a method receiver pointer.
	ObjectID ObjectID `msgp:"oid"`

	// CallID is used to identify (as a nonce) each method call and its response
	// on an object. Note that call IDs are managed solely by the side that
	// invokes the call, so in a way, both the frontend and backend manages its
	// own registry of call IDs.
	//
	// All synchronous calls must have a CallID; synchronous calls without an ID
	// is considered erroneous. An asynchronous call does not need a CallID.
	CallID CallID `msgp:"cid,omitempty"`

	// Method is the method name to be called. It is not optional.
	Method string `msgp:"m"`

	// Arguments is the field for an arbitrary argument type. This type is
	// usually an array, but it doesn't have to be; it depends on the method and
	// object that the call is for.
	Arguments msgp.Raw `msgp:"args"`
}

// Reply is the object type for a single IPC reply; it is followed after an IPC
// call with the same CallID. This is used for synchronous calls only.
type Reply struct {
	// CallID is the same ID as the ID used in a call.
	CallID CallID `msgp:"cid"`

	// Error is an additional error in the reply. It is only used for calls that
	// can error out.
	Error Error `msgp:"e,omitempty"`

	// Return is the return value of a call. For example, this type could return
	// a new Object. The type of this field should be determined based on the
	// parent object ID of the call that this reply was associated with.
	Return msgp.Raw `msgp:"ret,omitempty"`
}

// CancelCall is the object type for requesting cancellation of an active IPC
// call. This request is done asynchronously; there is no reply. In Go, this is
// often implemented using the context.Context API. Methods that don't accept a
// context.Context will not be cancellable, and cancelling such call is
// erroneous. The server can ignore erroneous cancel calls.
type CancelCall struct {
	// CallID is the same ID as the ID used in a call.
	CallID CallID `msgp:"cid"`
}
