package simple

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type SimpleString struct {
	Enabled    bool
	Value      string
	Name       string
	SearchName string
}

func (s *SimpleString) Init() {
	s.Enabled = false
	s.SearchName = s.Name
}

func (s *SimpleString) GetParserName() string {
	return s.SearchName
}

func (s *SimpleString) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, s.SearchName) {
		elements := strings.SplitN(line, " ", 2)
		if len(elements) != 2 {
			return "", &errors.ParseError{Parser: "SimpleString", Line: line, Message: "Parse error"}
		}
		s.Enabled = true
		s.Value = elements[1]
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.SearchName, Line: line}
}

func (s *SimpleString) Valid() bool {
	if s.Enabled {
		return true
	}
	return false
}

func (s *SimpleString) String() []string {
	if s.Enabled {
		return []string{fmt.Sprintf("  %s %s", s.SearchName, s.Value)}
	}
	return []string{}
}
