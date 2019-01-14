package simple

import (
	"fmt"

	"github.com/haproxytech/config-parser/errors"
)

type SimpleTimeout struct {
	Enabled bool
	Name    string
	Value   string
}

func (t *SimpleTimeout) Init() {
	t.Enabled = false
}

func (t *SimpleTimeout) GetParserName() string {
	return fmt.Sprintf("timeout %s", t.Name)
}

func (t *SimpleTimeout) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if len(parts) > 2 && parts[0] == "timeout" && parts[1] == t.Name {
		t.Value = parts[2]
		t.Enabled = true
		return "", nil
	}
	return "", &errors.ParseError{Parser: fmt.Sprintf("timeout %s", t.Name), Line: line}
}

func (t *SimpleTimeout) Valid() bool {
	if t.Enabled {
		return true
	}
	return false
}

func (t *SimpleTimeout) Result(AddComments bool) []string {
	if t.Enabled {
		return []string{fmt.Sprintf("  timeout %s %s", t.Name, t.Value)}
	}
	return []string{}
}
