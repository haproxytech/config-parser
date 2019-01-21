package common

import "strings"

//Expression is standard HAProxy expression formed by a sample-fetch followed by some converters.
type Expression struct {
	Expr []string
}

func (e *Expression) Parse(expression []string) error {
	e.Expr = expression
	return nil
}

func (e *Expression) String() string {
	return strings.Join(e.Expr, " ")
}
