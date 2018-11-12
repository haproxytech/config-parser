package simple

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type SimpleOption struct {
	Enabled    bool
	Name       string
	SearchName string
}

func (o *SimpleOption) Init() {
	o.Enabled = false
	o.SearchName = fmt.Sprintf("option %s", o.Name)
}

func (o *SimpleOption) GetParserName() string {
	return o.SearchName
}

func (o *SimpleOption) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, o.SearchName) {
		o.Enabled = true
		return "", nil
	}
	return "", &errors.ParseError{Parser: o.SearchName, Line: line}
}

func (o *SimpleOption) Valid() bool {
	if o.Enabled {
		return true
	}
	return false
}

func (o *SimpleOption) String() []string {
	if o.Enabled {
		return []string{fmt.Sprintf("  %s", o.SearchName)}
	}
	return []string{}
}
