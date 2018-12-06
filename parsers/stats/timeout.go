package stats

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/helpers"
)

type Timeout struct {
	Enabled    bool
	Value      []string
	Name       string
	SearchName string
}

func (s *Timeout) Init() {
	s.Enabled = false
	s.Name = "stats timeout"
	s.SearchName = s.Name
}

func (s *Timeout) GetParserName() string {
	return s.SearchName
}

func (s *Timeout) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, s.SearchName) {
		parts := helpers.StringSplitIgnoreEmpty(line, ' ')
		if len(parts) < 3 {
			return "", &errors.ParseError{Parser: "Timeout", Line: line, Message: "Parse error"}
		}
		s.Enabled = true
		s.Value = parts[2:]
		//todo add validation with simple timeouts
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.SearchName, Line: line}
}

func (s *Timeout) Valid() bool {
	if s.Enabled {
		return true
	}
	return false
}

func (s *Timeout) String() []string {
	if s.Enabled {
		return []string{fmt.Sprintf("  %s %s", s.SearchName, strings.Join(s.Value, " "))}
	}
	return []string{}
}
