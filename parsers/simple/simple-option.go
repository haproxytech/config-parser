package simple

import (
	"fmt"

	"github.com/haproxytech/config-parser/errors"
)

type SimpleOption struct {
	Enabled  bool
	NoOption bool
	Name     string
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
		o.NoOption = false // only last one parsed counts
		return "", nil
	}
	if len(parts) > 2 && parts[0] == "no" && parts[1] == "option" && parts[2] == o.Name {
		o.Enabled = true
		o.NoOption = true
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
	noOption := ""
	if o.NoOption {
		noOption = "no "
	}
	if o.Enabled {
		return []string{fmt.Sprintf("  %soption %s", noOption, o.Name)}
	}
	return []string{}
}
