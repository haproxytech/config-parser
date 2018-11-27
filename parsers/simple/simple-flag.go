package simple

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type SimpleFlag struct {
	Enabled    bool
	Name       string
	SearchName string
}

func (s *SimpleFlag) Init() {
	s.Enabled = false
	s.SearchName = fmt.Sprintf("%s", s.Name)
}

func (s *SimpleFlag) GetParserName() string {
	return s.SearchName
}

func (s *SimpleFlag) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, s.SearchName) {
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

func (s *SimpleFlag) String() []string {
	if s.Enabled {
		return []string{fmt.Sprintf("  %s", s.SearchName)}
	}
	return []string{}
}
