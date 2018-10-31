package errors

import "fmt"

//ParseError struct for creating parse errors
type ParseError struct {
	Parser  string
	Line    string
	Message string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%s: Parse error on %s", e.Parser, e.Message, e.Line)
}
