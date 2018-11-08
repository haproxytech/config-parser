package simple

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type SimpleString struct {
	enabled    bool
	value      string
	Name       string
	searchName string
}

func (s *SimpleString) Init() {
	s.enabled = false
	s.searchName = s.Name
}

func (s *SimpleString) GetParserName() string {
	return s.searchName
}

func (s *SimpleString) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, s.searchName) {
		elements := strings.SplitN(line, " ", 2)
		if len(elements) != 2 {
			return "", &errors.ParseError{Parser: "SimpleString", Line: line, Message: "Parse error"}
		}
		s.enabled = true
		s.value = elements[1]
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.searchName, Line: line}
}

func (s *SimpleString) Valid() bool {
	if s.enabled {
		return true
	}
	return false
}

func (s *SimpleString) String() []string {
	if s.enabled {
		return []string{fmt.Sprintf("  %s %s", s.searchName, s.value)}
	}
	return []string{}
}
