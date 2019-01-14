package simple

import (
	"fmt"
	"strconv"

	"github.com/haproxytech/config-parser/errors"
)

type SimpleNumber struct {
	Enabled    bool
	Value      int64
	Name       string
	SearchName string
}

func (s *SimpleNumber) Init() {
	s.Enabled = false
	s.SearchName = s.Name
}

func (s *SimpleNumber) GetParserName() string {
	return s.SearchName
}

func (s *SimpleNumber) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if parts[0] == s.SearchName {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "SimpleNumber", Line: line, Message: "Parse error"}
		}
		var num int64
		if num, err = strconv.ParseInt(parts[1], 10, 64); err != nil {
			return "", &errors.ParseError{Parser: "SimpleNumber", Line: line, Message: err.Error()}
		} else {
			s.Enabled = true
			s.Value = num
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.SearchName, Line: line}
}

func (s *SimpleNumber) Valid() bool {
	if s.Enabled {
		return true
	}
	return false
}

func (s *SimpleNumber) Result(AddComments bool) []string {
	if s.Enabled {
		return []string{fmt.Sprintf("  %s %d", s.SearchName, s.Value)}
	}
	return []string{}
}
