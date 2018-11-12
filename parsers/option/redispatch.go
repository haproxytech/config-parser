package options

import (
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type OptionRedispatch struct {
	Enabled bool
}

func (o *OptionRedispatch) Init() {
	o.Enabled = false
}

func (o *OptionRedispatch) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, "option redispatch") {
		o.Enabled = true
		return "", nil
	}
	return "", &errors.ParseError{Parser: "option redispatch", Line: line}
}

func (o *OptionRedispatch) Valid() bool {
	if o.Enabled {
		return true
	}
	return false
}

func (o *OptionRedispatch) String() []string {
	if o.Enabled {
		return []string{"  option redispatch"}
	}
	return []string{}
}
