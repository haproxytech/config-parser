package stats

import (
	"fmt"

	"github.com/haproxytech/config-parser/params"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type SocketLines struct {
	data []types.Socket
}

func (l *SocketLines) Init() {
	l.data = []types.Socket{}
}

func (l *SocketLines) GetParserName() string {
	return "stats socket"
}

func (l *SocketLines) Clear() {
	l.Init()
}

func (l *SocketLines) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(l.data) == 0 && !createIfNotExist {
		return nil, &errors.FetchError{}
	}
	return l.data, nil
}

func (l *SocketLines) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case []types.Socket:
		l.data = newValue
	case *types.Socket:
		l.data = append(l.data, *newValue)
	case types.Socket:
		l.data = append(l.data, newValue)
	}
	return fmt.Errorf("casting error")
}

func (l *SocketLines) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := l.Get(false)
	l.Clear()
	_, err := l.Parse(data, parts, []string{}, comment)
	if err != nil {
		l.Set(oldData)
	}
	return err
}

func (l *SocketLines) parseSocketLine(line string, parts []string, comment string) (*types.Socket, error) {
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

func (l *SocketLines) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 1 && parts[0] == "stats" && parts[1] == "socket" {
		if socket, err := l.parseSocketLine(line, parts, comment); err == nil {
			l.data = append(l.data, *socket)
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "SocketLines", Line: line}
}

func (l *SocketLines) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(l.data) == 0 {
		return nil, &errors.FetchError{}
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
