package parsers

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type StatsTimeout struct {
	enabled    bool
	value      []string
	Name       string
	searchName string
}

func (s *StatsTimeout) Init() {
	s.enabled = false
	s.Name = "stats timeout"
	s.searchName = s.Name
}

func (s *StatsTimeout) GetParserName() string {
	return s.searchName
}

func (s *StatsTimeout) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, s.searchName) {
		elements := strings.SplitN(line, " ", 3)
		if len(elements) < 3 {
			return "", &errors.ParseError{Parser: "StatsTimeout", Line: line, Message: "Parse error"}
		}
		s.enabled = true
		s.value = elements[2:]
		//todo add validation with simple timeouts
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.searchName, Line: line}
}

func (s *StatsTimeout) Valid() bool {
	if s.enabled {
		return true
	}
	return false
}

func (s *StatsTimeout) String() []string {
	if s.enabled {
		return []string{fmt.Sprintf("  %s %s", s.searchName, strings.Join(s.value, " "))}
	}
	return []string{}
}
