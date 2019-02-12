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

var FetchError error = errors.New("no data")

var ParserMissingErr error = errors.New("parser missing")

var SectionMissingErr error = errors.New("section missing")
var SectionAlreadyExistsErr error = errors.New("section already exists")
