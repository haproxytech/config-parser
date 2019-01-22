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

func (s *SimpleFlag) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == s.SearchName {
		s.Comment = comment
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

func (s *SimpleFlag) Result(AddComments bool) []common.ReturnResultLine {
	if s.Enabled {
		return []common.ReturnResultLine{
			common.ReturnResultLine{
				Data:    fmt.Sprintf("%s", s.SearchName),
				Comment: s.Comment,
			},
		}
	}
	return []common.ReturnResultLine{}
}

func (s *SimpleFlag) Annotation() string {
	return s.Comment
}
