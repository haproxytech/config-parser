package stats

import (
	"fmt"

	bindoptions "github.com/haproxytech/config-parser/bind-options"
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type Socket struct {
	Path    string //can be address:port
	Params  []bindoptions.BindOption
	Comment string
}

type SocketLines struct {
	SocketLines []*Socket
}

func (l *SocketLines) Init() {
	l.SocketLines = []*Socket{}
}

func (l *SocketLines) GetParserName() string {
	return "stats socket"
}

func (l *SocketLines) parseSocketLine(line string, parts []string, comment string) (*Socket, error) {

	if len(parts) < 3 {
		return nil, &errors.ParseError{Parser: "SocketSingle", Line: line, Message: "Parse error"}
	}
	socket := &Socket{
		Path:    parts[2],
		Params:  bindoptions.Parse(parts[3:]),
		Comment: comment,
	}
	//s.value = elements[1:]
	return socket, nil
}

func (l *SocketLines) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 1 && parts[0] == "stats" && parts[1] == "socket" {
		if nameserver, err := l.parseSocketLine(line, parts, comment); err == nil {
			l.SocketLines = append(l.SocketLines, nameserver)
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "SocketLines", Line: line}
}

func (l *SocketLines) Valid() bool {
	if len(l.SocketLines) > 0 {
		return true
	}
	return false
}

func (l *SocketLines) Result(AddComments bool) []common.ReturnResultLine {
	result := make([]common.ReturnResultLine, len(l.SocketLines))
	for index, socket := range l.SocketLines {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf(fmt.Sprintf("stats socket %s %s", socket.Path, bindoptions.String(socket.Params))),
			Comment: socket.Comment,
		}
	}
	return result
}
