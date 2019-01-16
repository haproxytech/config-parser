package frontend

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type UseBackend struct {
	Name          string
	Condition     string
	ConditionKind string
}

type UseBackends struct {
	UseBackends []UseBackend
}

func (h *UseBackends) Init() {
	h.UseBackends = []UseBackend{}
}

func (h *UseBackends) GetParserName() string {
	return "use_backend"
}

func (h *UseBackends) parseUseBackendLine(line string, parts []string) (UseBackend, error) {
	if len(parts) >= 4 {
		_, condition := common.SplitRequest(parts[2:])
		data := UseBackend{
			Name: parts[1],
		}
		if len(condition) > 0 {
			data.ConditionKind = condition[0]
			data.Condition = strings.Join(condition[1:], " ")
		}
		return data, nil
	}
	return UseBackend{}, &errors.ParseError{Parser: "UseBackendLines", Line: line}
}

func (h *UseBackends) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if parts[0] == "use_backend" {
		request, err := h.parseUseBackendLine(line, parts)
		if err != nil {
			return "", &errors.ParseError{Parser: "UseBackendLines", Line: line}
		}
		h.UseBackends = append(h.UseBackends, request)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "UseBackendLines", Line: line}
}

func (h *UseBackends) Valid() bool {
	if len(h.UseBackends) > 0 {
		return true
	}
	return false
}

func (h *UseBackends) Result(AddComments bool) []string {
	result := make([]string, len(h.UseBackends))
	for index, req := range h.UseBackends {
		condition := ""
		if req.Condition != "" {
			condition = fmt.Sprintf(" %s %s", req.ConditionKind, req.Condition)
		}
		result[index] = fmt.Sprintf("  use_backend %s%s #deny", req.Name, condition)
	}
	return result
}
