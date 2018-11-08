package parsers

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/bind-options"
	"github.com/haproxytech/config-parser/errors"
)

type StatsSocket struct {
	enabled    bool
	path       string //can be address:port
	params     []bindoptions.BindOption
	Name       string
	searchName string
}

func (s *StatsSocket) Init() {
	s.enabled = false
	s.searchName = "stats socket"
}

func (s *StatsSocket) GetParserName() string {
	return s.searchName
}

func (s *StatsSocket) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, s.searchName) {
		elements := strings.Split(line, " ")
		if len(elements) < 3 {
			return "", &errors.ParseError{Parser: "StatsSocket", Line: line, Message: "Parse error"}
		}
		s.enabled = true
		s.path = elements[2]
		s.params = bindoptions.Parse(elements[3:])
		//s.value = elements[1:]
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.searchName, Line: line}
}

func (s *StatsSocket) Valid() bool {
	if s.enabled {
		return true
	}
	return false
}

func (s *StatsSocket) String() []string {
	if s.enabled {
		return []string{fmt.Sprintf("  %s %s %s", s.searchName, s.path, bindoptions.String(s.params))}
	}
	return []string{}
}
