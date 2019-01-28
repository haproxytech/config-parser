package params

import "fmt"

type ErrParseBindOption interface {
	Error() string
}

type ErrParseServerOption interface {
	Error() string
}

//ErrNotFound struct for creating parse errors
type ErrNotFound struct {
	Have string
	Want string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("error: have [%s] want [%s]", e.Have, e.Want)
}

//ParseError struct for creating parse errors
type ErrNotEnoughParams struct {
}

func (e *ErrNotEnoughParams) Error() string {
	return fmt.Sprintf("error: not enough params")
}
