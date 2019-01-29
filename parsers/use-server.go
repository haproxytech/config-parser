package parsers

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type UseServers struct {
	data []types.UseServer
}

func (l *UseServers) Init() {
	l.data = []types.UseServer{}
}

func (l *UseServers) GetParserName() string {
	return "use-server"
}

func (l *UseServers) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(l.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return l.data, nil
}

func (l *UseServers) Set(data common.ParserData) error {
	if data == nil {
		l.Init()
		return nil
	}
	switch newValue := data.(type) {
	case []types.UseServer:
		l.data = newValue
	case *types.UseServer:
		l.data = append(l.data, *newValue)
	case types.UseServer:
		l.data = append(l.data, newValue)
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (l *UseServers) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := l.Get(false)
	l.Init()
	_, err := l.Parse(data, parts, []string{}, comment)
	if err != nil {
		l.Set(oldData)
	}
	return err
}

func (l *UseServers) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 3 && parts[0] == "use-server" {
		data := types.UseServer{
			Name:      parts[1],
			Condition: strings.Join(parts[3:], " "),
			Comment:   comment,
		}
		switch parts[2] {
		case "if", "unless":
			data.ConditionType = parts[2]
		default:
			return "", &errors.ParseError{Parser: "UseServers", Line: line}
		}
		l.data = append(l.data, data)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "UseServers", Line: line}
}

func (l *UseServers) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(l.data) == 0 {
		return nil, errors.FetchError
	}
	//use-server
	result := make([]common.ReturnResultLine, len(l.data))
	for index, data := range l.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("use-server %s %s %s", data.Name, data.ConditionType, data.Condition),
			Comment: data.Comment,
		}
	}
	return result, nil
}
