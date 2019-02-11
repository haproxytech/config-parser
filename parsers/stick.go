package parsers

import (
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Stick struct {
	data []types.Stick
}

func (h *Stick) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) >= 2 && parts[0] == "stick" {
		var err error
		command, condition := common.SplitRequest(parts[2:])
		data := types.Stick{
			Pattern: command[0],
			Comment: comment,
		}
		if len(command) > 2 {
			data.Table = command[2]
		}
		if len(condition) > 1 {
			data.ConditionType = condition[0]
			data.Condition = strings.Join(condition[1:], " ")
		}
		switch parts[1] {
		case "match", "on", "store-request", "store-response":
			data.Name = parts[1]
		default:
			return "", &errors.ParseError{Parser: "Stick", Line: line}
		}
		if err != nil {
			return "", err
		}
		h.data = append(h.data, data)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "Stick", Line: line}
}

func (h *Stick) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		var data strings.Builder
		data.WriteString("stick ")
		data.WriteString(req.Name)
		data.WriteString(" ")
		data.WriteString(req.Pattern)
		if req.Table != "" {
			data.WriteString(" table ")
			data.WriteString(req.Table)
		}
		if req.Condition != "" {
			data.WriteString(" ")
			data.WriteString(req.ConditionType)
			data.WriteString(" ")
			data.WriteString(req.Condition)
		}
		result[index] = common.ReturnResultLine{
			Data:    data.String(),
			Comment: req.Comment,
		}
	}
	return result, nil
}
