package simple

import (
	"fmt"

	"github.com/haproxytech/config-parser/errors"
)

type SimpleOption struct {
	Enabled bool
	Name    string
}

func (o *SimpleOption) Init() {
	o.Enabled = false
}

func (o *SimpleOption) GetParserName() string {
	return fmt.Sprintf("option %s", o.Name)
}

func (o *SimpleOption) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if len(parts) > 1 && parts[0] == "option" && parts[1] == o.Name {
		o.Enabled = true
		return "", nil
	}
	return "", &errors.ParseError{Parser: fmt.Sprintf("option %s", o.Name), Line: line}
}

func (o *SimpleOption) Valid() bool {
	if o.Enabled {
		return true
	}
	return false
}

func (o *SimpleOption) Result(AddComments bool) []string {
	if o.Enabled {
		return []string{fmt.Sprintf("  option %s", o.Name)}
	}
	return []string{}
}
