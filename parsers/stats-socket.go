package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/params"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Socket struct {
	data []types.Socket
}

func (l *Socket) parse(line string, parts []string, comment string) (*types.Socket, error) {
	if len(parts) < 3 {
		return nil, &errors.ParseError{Parser: "SocketSingle", Line: line, Message: "Parse error"}
	}
	socket := &types.Socket{
		Path:    parts[2],
		Params:  params.ParseBindOptions(parts[3:]),
		Comment: comment,
	}
	//s.value = elements[1:]
	return socket, nil
}

func (l *Socket) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(l.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(l.data))
	for index, socket := range l.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf(fmt.Sprintf("stats socket %s %s", socket.Path, params.BindOptionsString(socket.Params))),
			Comment: socket.Comment,
		}
	}
	return result, nil
}
