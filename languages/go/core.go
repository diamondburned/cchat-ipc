package cchatipc

//go:generate go run github.com/tinylib/msgp

// Error describes an error over the wire.
type Error struct {
	Error string `msgp:"error"`
}

// Time describes a timestamp without a timezone; it is in UnixNano format.
type Time int64

// Reader describes possible readers. This structure should be treated like a
// union, that is, only one of those fields should be used.
type Reader struct {
	Stream *Stream `msgp:"stream,omitempty"`
	File   *File   `msgp:"file,omitempty"`
}

// File describes a physical file on disk. Backends that read this file must
// not modify nor delete it.
type File struct {
	Filename string
}

// Stream describes the information of a stream ticket along with extraneous
// metadata.
type Stream struct {
	Ticket StreamTicket `msgp:"t"`

	// BufferSize suggests to the receiver the size of each data chunk to be
	// sent using StreamBytes. This suggestion is not required.
	BufferSize int64 `json:"buffer_size"`
}

// StreamTicket is a unique identifier for each stream.
type StreamTicket uint64

// StreamData is a possible event that is an incoming piece of a stream.
type StreamData struct {
	Ticket StreamTicket `msgp:"t"`
	Bytes  []byte       `msgp:"b,omitempty"`

	// Error is the optional error that a stream may emit. The stream is
	// considered to be invalid after an error has been emitted and must not be
	// used again.
	Error Error `msgp:"e,omitempty"`
}
