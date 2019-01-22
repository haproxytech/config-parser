package simple

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type SimpleTimeTwoWords struct {
	Enabled    bool
	Value      string
	Keywords   []string
	SearchName string
	Comment    string
}

func (s *SimpleTimeTwoWords) Init() {
	s.Enabled = false
	s.SearchName = fmt.Sprintf(strings.Join(s.Keywords, " "))
}

func (s *SimpleTimeTwoWords) GetParserName() string {
	return s.SearchName
}

func (s *SimpleTimeTwoWords) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) >= 2 && parts[0] == s.Keywords[0] && parts[1] == s.Keywords[1] {
		if len(parts) < 3 {
			return "", &errors.ParseError{Parser: "SimpleTimeTwoWords", Line: line, Message: "Parse error"}
		}
		s.Enabled = true
		s.Value = parts[2]
		s.Comment = comment
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

func (s *SimpleTimeTwoWords) Result(AddComments bool) []common.ReturnResultLine {
	if s.Enabled {
		return []common.ReturnResultLine{
			common.ReturnResultLine{
				Data:    fmt.Sprintf("%s %s", s.SearchName, s.Value),
				Comment: s.Comment,
			},
		}
	}
	return []common.ReturnResultLine{}
}
