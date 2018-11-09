package simple

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/helpers"

	"github.com/haproxytech/config-parser/errors"
)

type SimpleTimeout struct {
	enabled    bool
	Name       string
	Value      string
	searchName string
}

func (t *SimpleTimeout) Init() {
	t.enabled = false
	t.searchName = fmt.Sprintf("timeout %s", t.Name)
}

func (t *SimpleTimeout) GetParserName() string {
	return t.searchName
}

func (t *SimpleTimeout) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, t.searchName) {
		parts := helpers.StringSplitIgnoreEmpty(line, ' ')
		if len(parts) < 3 {
			return "", &errors.ParseError{Parser: t.searchName, Line: line}
		}
		t.Value = parts[2]
		t.enabled = true
		return "", nil
	}
	return "", &errors.ParseError{Parser: t.searchName, Line: line}
}

func (t *SimpleTimeout) Valid() bool {
	if t.enabled {
		return true
	}
	return false
}

func (t *SimpleTimeout) String() []string {
	if t.enabled {
		return []string{fmt.Sprintf("  %s %s", t.searchName, t.Value)}
	}
	return []string{}
}
