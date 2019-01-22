package options

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type OptionRedispatch struct {
	Enabled bool
	Comment string
}

func (o *OptionRedispatch) Init() {
	o.Enabled = false
}

func (o *OptionRedispatch) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 1 && parts[0] == "option" && parts[1] == "redispatch" {
		o.Enabled = true
		o.Comment = comment
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

func (o *OptionRedispatch) Result(AddComments bool) []common.ReturnResultLine {
	if o.Enabled {
		return []string{"option redispatch"}
	}
	return []string{}
}
