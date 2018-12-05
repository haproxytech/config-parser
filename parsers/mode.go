package parsers

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/helpers"
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

func (m *Mode) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, "mode ") {
		parts := helpers.StringSplitIgnoreEmpty(line, ' ')
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

func (m *Mode) String() []string {
	return []string{fmt.Sprintf("  mode %s", m.Value)}
}
