package parsers

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Daemon struct {
	data *types.Enabled
}

func (d *Daemon) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "daemon" {
		d.data = &types.Enabled{
			Comment: comment,
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "Daemon", Line: line}
}

func (d *Daemon) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if d.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    "daemon",
			Comment: d.data.Comment,
		},
	}, nil
}
