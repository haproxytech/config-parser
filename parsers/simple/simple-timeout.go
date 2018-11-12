package simple

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/helpers"

	"github.com/haproxytech/config-parser/errors"
)

type SimpleTimeout struct {
	Enabled    bool
	Name       string
	Value      string
	SearchName string
}

func (t *SimpleTimeout) Init() {
	t.Enabled = false
	t.SearchName = fmt.Sprintf("timeout %s", t.Name)
}

func (t *SimpleTimeout) GetParserName() string {
	return t.SearchName
}

func (t *SimpleTimeout) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, t.SearchName) {
		parts := helpers.StringSplitIgnoreEmpty(line, ' ')
		if len(parts) < 3 {
			return "", &errors.ParseError{Parser: t.SearchName, Line: line}
		}
		t.Value = parts[2]
		t.Enabled = true
		return "", nil
	}
	return "", &errors.ParseError{Parser: t.SearchName, Line: line}
}

func (t *SimpleTimeout) Valid() bool {
	if t.Enabled {
		return true
	}
	return false
}

func (t *SimpleTimeout) String() []string {
	if t.Enabled {
		return []string{fmt.Sprintf("  %s %s", t.SearchName, t.Value)}
	}
	return []string{}
}
