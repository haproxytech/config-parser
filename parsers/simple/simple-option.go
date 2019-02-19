package simple

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type SimpleOption struct {
	Name string
	name string
	data *types.SimpleOption
}

func (o *SimpleOption) Init() {
	if !strings.HasPrefix(o.Name, "option") {
		o.name = o.Name
		o.Name = fmt.Sprintf("option %s", o.Name)
	}
	o.data = nil
}

func (o *SimpleOption) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 1 && parts[0] == "option" && parts[1] == o.name {
		o.data = &types.SimpleOption{
			Comment: comment,
		}
		return "", nil
	}
	if len(parts) > 2 && parts[0] == "no" && parts[1] == "option" && parts[2] == o.name {
		o.data = &types.SimpleOption{
			NoOption: true,
			Comment:  comment,
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: fmt.Sprintf("option %s", o.name), Line: line}
}

func (o *SimpleOption) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if o.data == nil {
		return nil, errors.FetchError
	}
	noOption := ""
	if o.data.NoOption {
		noOption = "no "
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("%soption %s", noOption, o.name),
			Comment: o.data.Comment,
		},
	}, nil
}
