package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/errors"
)

type Mode struct {
	Value string
}

func (m *Mode) Init() {
	m.Value = ""
}

func (m *Mode) GetParserName() string {
	return "mode"
}

func (m *Mode) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if parts[0] == "mode" {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "Mode", Line: line, Message: "Parse error"}
		}
		if parts[1] == "http" || parts[1] == "tcp" || parts[1] == "health" {
			m.Value = parts[1]
			return "", nil
		}
		return "", &errors.ParseError{Parser: "Mode", Line: line}
	}
	return "", &errors.ParseError{Parser: "Mode", Line: line}
}

func (m *Mode) Valid() bool {
	return m.Value != ""
}

func (m *Mode) Result(AddComments bool) []string {
	return []string{fmt.Sprintf("  mode %s", m.Value)}
}
