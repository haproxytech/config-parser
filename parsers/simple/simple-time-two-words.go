package simple

import (
	"fmt"

	"github.com/haproxytech/config-parser/errors"
)

type SimpleTimeTwoWords struct {
	Enabled    bool
	Value      string
	Name       string
	SearchName string
}

func (s *SimpleTimeTwoWords) Init() {
	s.Enabled = false
	s.SearchName = s.Name
}

func (s *SimpleTimeTwoWords) GetParserName() string {
	return s.SearchName
}

func (s *SimpleTimeTwoWords) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if parts[0] == s.SearchName {
		if len(parts) < 3 {
			return "", &errors.ParseError{Parser: "SimpleTimeTwoWords", Line: line, Message: "Parse error"}
		}
		s.Enabled = true
		s.Value = parts[2]
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.SearchName, Line: line}
}

func (s *SimpleTimeTwoWords) Valid() bool {
	if s.Enabled {
		return true
	}
	return false
}

func (s *SimpleTimeTwoWords) Result(AddComments bool) []string {
	if s.Enabled {
		return []string{fmt.Sprintf("  %s %s", s.SearchName, s.Value)}
	}
	return []string{}
}
