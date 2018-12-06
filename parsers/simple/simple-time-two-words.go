package simple

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/helpers"
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

func (s *SimpleTimeTwoWords) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, s.SearchName) {
		parts := helpers.StringSplitIgnoreEmpty(line, ' ')
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

func (s *SimpleTimeTwoWords) String() []string {
	if s.Enabled {
		return []string{fmt.Sprintf("  %s %s", s.SearchName, s.Value)}
	}
	return []string{}
}
