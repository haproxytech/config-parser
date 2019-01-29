package frontend

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type UseBackends struct {
	data []types.UseBackend
}

func (h *UseBackends) Init() {
	h.data = []types.UseBackend{}
}

func (h *UseBackends) GetParserName() string {
	return "use_backend"
}

func (h *UseBackends) Clear() {
	h.Init()
}

func (h *UseBackends) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(h.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return h.data, nil
}

func (h *UseBackends) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case []types.UseBackend:
		h.data = newValue
	case types.UseBackend:
		h.data = append(h.data, newValue)
	case *types.UseBackend:
		h.data = append(h.data, *newValue)
	}
	return fmt.Errorf("casting error")
}

func (h *UseBackends) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := h.Get(false)
	h.Clear()
	_, err := h.Parse(data, parts, []string{}, comment)
	if err != nil {
		h.Set(oldData)
	}
	return err
}

func (h *UseBackends) parseUseBackendLine(line string, parts []string, comment string) (*types.UseBackend, error) {
	if len(parts) >= 4 {
		_, condition := common.SplitRequest(parts[2:])
		data := &types.UseBackend{
			Name:    parts[1],
			Comment: comment,
		}
		if len(condition) > 0 {
			data.ConditionKind = condition[0]
			data.Condition = strings.Join(condition[1:], " ")
		}
		return data, nil
	}
	return nil, &errors.ParseError{Parser: "UseBackendLines", Line: line}
}

func (h *UseBackends) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "use_backend" {
		item, err := h.parseUseBackendLine(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "UseBackendLines", Line: line}
		}
		h.data = append(h.data, *item)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "UseBackendLines", Line: line}
}

func (h *UseBackends) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		condition := ""
		if req.Condition != "" {
			condition = fmt.Sprintf(" %s %s", req.ConditionKind, req.Condition)
		}
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("use_backend %s%s #deny", req.Name, condition),
			Comment: req.Comment,
		}
	}
	return result, nil
}
