package simple

import (
	"fmt"
	"strings"

	"config-parser/errors"
)

type SimpleOption struct {
	enabled    bool
	Name       string
	searchName string
}

func (o *SimpleOption) Init() {
	o.enabled = false
	o.searchName = fmt.Sprintf("option %s", o.Name)
}

func (o *SimpleOption) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, o.searchName) {
		o.enabled = true
		return "", nil
	}
	return "", &errors.ParseError{Parser: o.searchName, Line: line}
}

func (o *SimpleOption) Valid() bool {
	if o.enabled {
		return true
	}
	return false
}

func (o *SimpleOption) String() []string {
	if o.enabled {
		return []string{fmt.Sprintf("  %s", o.searchName)}
	}
	return []string{}
}
