package simple

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"

	"github.com/haproxytech/config-parser/errors"
)

type SimpleFlag struct {
	Enabled    bool
	Name       string
	SearchName string
	Comment    string
}

func (s *SimpleFlag) Init() {
	s.Enabled = false
	s.SearchName = fmt.Sprintf("%s", s.Name)
}

func (s *SimpleFlag) GetParserName() string {
	return s.SearchName
}

func (s *SimpleFlag) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if parts[0] == s.SearchName {
		s.Comment = common.StringExtractComment(line)
		s.Enabled = true
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.SearchName, Line: line}
}

func (s *SimpleFlag) Valid() bool {
	if s.Enabled {
		return true
	}
	return false
}

func (s *SimpleFlag) Result(AddComments bool) []string {
	if s.Enabled {
		if s.Comment != "" {
			return []string{fmt.Sprintf("  %s # %s", s.SearchName, s.Comment)}
		}
		return []string{fmt.Sprintf("  %s", s.SearchName)}
	}
	return []string{}
}

func (s *SimpleFlag) Annotation() string {
	return s.Comment
}
