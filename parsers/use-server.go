package parsers

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type UseServer struct {
	data []types.UseServer
}

func (l *UseServer) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
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
			return "", &errors.ParseError{Parser: "UseServer", Line: line}
		}
		l.data = append(l.data, data)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "UseServer", Line: line}
}

func (l *UseServer) Result(AddComments bool) ([]common.ReturnResultLine, error) {
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
