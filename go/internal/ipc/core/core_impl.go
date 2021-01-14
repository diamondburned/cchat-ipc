package core

import "io"

// RegularError is a regular text error with ErrorType(0).
type RegularError string

// Error satisfies the error interface.
func (err RegularError) Error() string { return string(err) }

func (err Error) Convert() error {
	switch err.Type {
	case RegularErrorType:
		var regularError RegularError

		_, err := regularError.UnmarshalMsg(err.Data)
		if err != nil {
			return nil
		}

		return regularError
	default:
		panic("TODO: implement enum generation")
	}
}

// Convert TODO
func (r Reader) Convert() io.Reader {
	panic("a")
	return nil
}
