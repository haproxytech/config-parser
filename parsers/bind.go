package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/params"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Bind struct {
	data []types.Bind
}

func (h *Bind) parse(line string, parts []string, comment string) (*types.Bind, error) {
	if len(parts) >= 2 {
		data := &types.Bind{
			Path:    parts[1],
			Comment: comment,
		}
		if len(parts) > 2 {
			data.Params = params.ParseBindOptions(parts[2:])
		}
		return data, nil
	}
	return nil, &errors.ParseError{Parser: "BindLines", Line: line}
}

func (h *Bind) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("bind %s %s", req.Path, params.BindOptionsString(req.Params)),
			Comment: req.Comment,
		}
	}
	return result, nil
}
