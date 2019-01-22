package simple

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type SimpleStringMultiple struct {
	Enabled    bool
	Value      []string
	Name       string
	SearchName string
	Comment    string
}

func (s *SimpleStringMultiple) Init() {
	s.Enabled = false
	s.SearchName = s.Name
}

func (s *SimpleStringMultiple) GetParserName() string {
	return s.SearchName
}

func (s *SimpleStringMultiple) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == s.SearchName {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "SimpleStringMultiple", Line: line, Message: "Parse error"}
		}
		s.Enabled = true
		s.Value = parts[1:]
		s.Comment = comment
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.SearchName, Line: line}
}

func (s *SimpleStringMultiple) Valid() bool {
	if s.Enabled {
		return true
	}
	return false
}

func (s *SimpleStringMultiple) Result(AddComments bool) []common.ReturnResultLine {
	if s.Enabled {
		return []common.ReturnResultLine{
			common.ReturnResultLine{
				Data:    fmt.Sprintf("%s %s", s.SearchName, strings.Join(s.Value, " ")),
				Comment: s.Comment,
			},
		}
	}
	return []common.ReturnResultLine{}
}
