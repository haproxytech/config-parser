package parsers

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type UseBackend struct {
	data []types.UseBackend
}

func (h *UseBackend) parse(line string, parts []string, comment string) (*types.UseBackend, error) {
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
	return nil, &errors.ParseError{Parser: "UseBackend", Line: line}
}

func (h *UseBackend) Result(AddComments bool) ([]common.ReturnResultLine, error) {
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
