package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/params"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type DefaultServer struct {
	data []types.DefaultServer
}

func (h *DefaultServer) parse(line string, parts []string, comment string) (*types.DefaultServer, error) {
	if len(parts) >= 2 {
		data := &types.DefaultServer{
			Params:  params.ParseServerOptions(parts[1:]),
			Comment: comment,
		}
		return data, nil
	}
	return nil, &errors.ParseError{Parser: "DefaultServer", Line: line}
}

func (h *DefaultServer) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("default-server %s", params.ServerOptionsString(req.Params)),
			Comment: req.Comment,
		}
	}
	return result, nil
}
