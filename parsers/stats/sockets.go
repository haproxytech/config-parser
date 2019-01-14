package stats

import (
	"fmt"

	bindoptions "github.com/haproxytech/config-parser/bind-options"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/helpers"
)

type Socket struct {
	Path   string //can be address:port
	Params []bindoptions.BindOption
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

func (l *SocketLines) parseSocketLine(line string) (*Socket, error) {

	elements := helpers.StringSplitIgnoreEmpty(line, ' ')
	if len(elements) < 3 {
		return nil, &errors.ParseError{Parser: "SocketSingle", Line: line, Message: "Parse error"}
	}
	socket := &Socket{
		Path:   elements[2],
		Params: bindoptions.Parse(elements[3:]),
	}
	//s.value = elements[1:]
	return socket, nil
}

func (l *SocketLines) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if len(parts) > 1 && parts[0] == "stats" && parts[1] == "socket" {
		if nameserver, err := l.parseSocketLine(line); err == nil {
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

func (l *SocketLines) Result(AddComments bool) []string {
	result := make([]string, len(l.SocketLines))
	for index, socket := range l.SocketLines {
		result[index] = fmt.Sprintf(fmt.Sprintf("  stats socket %s %s", socket.Path, bindoptions.String(socket.Params)))
	}
	return result
}