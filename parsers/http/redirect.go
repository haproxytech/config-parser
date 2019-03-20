package http

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/parsers/http/actions"
	"github.com/haproxytech/config-parser/types"
)

type Redirect struct {
	Name string
	data []types.HTTPAction
}

func (h *Redirect) Init() {
	h.Name = "redirect"
	h.data = []types.HTTPAction{}
}

func (f *Redirect) ParseHTTPResponse(response types.HTTPAction, parts []string, comment string) error {
	err := response.Parse(parts, comment)
	if err != nil {
		return &errors.ParseError{Parser: "HTTPResponseLines", Line: ""}
	}
	f.data = append(f.data, response)
	return nil
}

func (h *Redirect) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) >= 2 && parts[0] == "redirect" {
		adjusted := append([]string{""}, parts...)
		err := h.ParseHTTPResponse(&actions.Redirect{}, adjusted, comment)
		if err != nil {
			return "", err
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "HTTPResponseLines", Line: line}
}

func (h *Redirect) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, res := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    res.String(),
			Comment: res.GetComment(),
		}
	}
	return result, nil
}
