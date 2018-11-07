package simple

import (
	"fmt"
	"strings"

	"config-parser/errors"
)

type SimpleStringMultiple struct {
	enabled    bool
	value      []string
	Name       string
	searchName string
}

func (s *SimpleStringMultiple) Init() {
	s.enabled = false
	s.searchName = s.Name
}

func (s *SimpleStringMultiple) GetParserName() string {
	return s.searchName
}

func (s *SimpleStringMultiple) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, s.searchName) {
		elements := strings.SplitN(line, " ", 2)
		if len(elements) < 2 {
			return "", &errors.ParseError{Parser: "SimpleStringMultiple", Line: line, Message: "Parse error"}
		}
		s.enabled = true
		s.value = elements[1:]
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.searchName, Line: line}
}

func (s *SimpleStringMultiple) Valid() bool {
	if s.enabled {
		return true
	}
	return false
}

func (s *SimpleStringMultiple) String() []string {
	if s.enabled {
		return []string{fmt.Sprintf("  %s %s", s.searchName, strings.Join(s.value, " "))}
	}
	return []string{}
}
