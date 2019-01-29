package stick

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Sticks struct {
	data []types.Stick
}

func (h *Sticks) Init() {
	h.data = []types.Stick{}
}

func (h *Sticks) GetParserName() string {
	return "stick"
}

func (h *Sticks) Clear() {
	h.Init()
}

func (h *Sticks) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(h.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return h.data, nil
}

func (h *Sticks) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case []types.Stick:
		h.data = newValue
	case types.Stick:
		h.data = append(h.data, newValue)
	case *types.Stick:
		h.data = append(h.data, *newValue)
	}
	return fmt.Errorf("casting error")
}

func (h *Sticks) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := h.Get(false)
	h.Clear()
	_, err := h.Parse(data, parts, []string{}, comment)
	if err != nil {
		h.Set(oldData)
	}
	return err
}

/*func (f *Sticks) ParseLine(request Stick, parts []string, comment string) error {
	err := request.Parse(parts, comment)
	if err != nil {
		return &errors.ParseError{Parser: "Sticks", Line: ""}
	}
	f.data = append(f.data, request)
	return nil
}*/

func (h *Sticks) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
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
			return "", &errors.ParseError{Parser: "Sticks", Line: line}
		}
		if err != nil {
			return "", err
		}
		h.data = append(h.data, data)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "Sticks", Line: line}
}

func (h *Sticks) Result(AddComments bool) ([]common.ReturnResultLine, error) {
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
