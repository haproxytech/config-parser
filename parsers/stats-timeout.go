package parsers

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type StatsTimeout struct {
	Enabled    bool
	Value      []string
	Name       string
	SearchName string
}

func (s *StatsTimeout) Init() {
	s.Enabled = false
	s.Name = "stats timeout"
	s.SearchName = s.Name
}

func (s *StatsTimeout) GetParserName() string {
	return s.SearchName
}

func (s *StatsTimeout) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, s.SearchName) {
		elements := strings.SplitN(line, " ", 3)
		if len(elements) < 3 {
			return "", &errors.ParseError{Parser: "StatsTimeout", Line: line, Message: "Parse error"}
		}
		s.Enabled = true
		s.Value = elements[2:]
		//todo add validation with simple timeouts
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.SearchName, Line: line}
}

func (s *StatsTimeout) Valid() bool {
	if s.Enabled {
		return true
	}
	return false
}

func (s *StatsTimeout) String() []string {
	if s.Enabled {
		return []string{fmt.Sprintf("  %s %s", s.SearchName, strings.Join(s.Value, " "))}
	}
	return []string{}
}
