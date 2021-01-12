package ipc

// RegularError is a regular text error with ErrorType(0).
type RegularError string

// Error satisfies the error interface.
func (err RegularError) Error() string { return string(err) }
