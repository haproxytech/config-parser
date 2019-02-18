package errors

import (
	"errors"
	"fmt"
)

//ParseError struct for creating parse errors
type ParseError struct {
	Parser  string
	Line    string
	Message string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%s: Parse error on %s", e.Parser, e.Message, e.Line)
}

var AttributeNotFoundErr error = errors.New("attribute not found")

var FetchError error = errors.New("no data")

var IndexOutOfRange error = errors.New("index out of range")

var InvalidData error = errors.New("invalid data")

var ParserMissingErr error = errors.New("parser missing")

var SectionAlreadyExistsErr error = errors.New("section already exists")

var SectionMissingErr error = errors.New("section missing")
