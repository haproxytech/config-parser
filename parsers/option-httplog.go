package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type OptionHTTPLog struct {
	data *types.OptionHTTPLog
}

func (o *OptionHTTPLog) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 2 && parts[0] == "option" && parts[1] == "httplog" && parts[2] == "clf" {
		o.data = &types.OptionHTTPLog{
			Comment: comment,
			Clf:     true,
		}
		return "", nil
	}
	if len(parts) > 1 && parts[0] == "option" && parts[1] == "httplog" {
		o.data = &types.OptionHTTPLog{
			Comment: comment,
		}
		return "", nil
	}
	if len(parts) > 3 && parts[0] == "no" && parts[1] == "option" && parts[2] == "httplog" && parts[3] == "clf" {
		o.data = &types.OptionHTTPLog{
			NoOption: true,
			Comment:  comment,
			Clf:      true,
		}
		return "", nil
	}
	if len(parts) > 2 && parts[0] == "no" && parts[1] == "option" && parts[2] == "httplog" {
		o.data = &types.OptionHTTPLog{
			NoOption: true,
			Comment:  comment,
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "option httplog", Line: line}
}

func (o *OptionHTTPLog) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if o.data == nil {
		return nil, errors.FetchError
	}
	clf := ""
	if o.data.Clf {
		clf = " clf"
	}
	noOption := ""
	if o.data.NoOption {
		noOption = "no "
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("%soption httplog%s", noOption, clf),
			Comment: o.data.Comment,
		},
	}, nil
}
