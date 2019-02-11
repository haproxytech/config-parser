package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/params"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Server struct {
	data []types.Server
}

func (h *Server) parse(line string, parts []string, comment string) (*types.Server, error) {
	if len(parts) >= 3 {
		data := &types.Server{
			Name:    parts[1],
			Address: parts[2],
			Params:  params.ParseServerOptions(parts[3:]),
			Comment: comment,
		}
		return data, nil
	}
	return nil, &errors.ParseError{Parser: "Server", Line: line}
}

func (h *Server) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("server %s %s %s", req.Name, req.Address, params.ServerOptionsString(req.Params)),
			Comment: req.Comment,
		}
	}
	return result, nil
}
