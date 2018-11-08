package options

import (
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type OptionRedispatch struct {
	enabled bool
}

func (o *OptionRedispatch) Init() {
	o.enabled = false
}

func (o *OptionRedispatch) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, "option redispatch") {
		o.enabled = true
		return "", nil
	}
	return "", &errors.ParseError{Parser: "option redispatch", Line: line}
}

func (o *OptionRedispatch) Valid() bool {
	if o.enabled {
		return true
	}
	return false
}

func (o *OptionRedispatch) String() []string {
	if o.enabled {
		return []string{"  option redispatch"}
	}
	return []string{}
}
