package http

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/parsers/http/actions"
)

type HTTPResponses struct {
	data []HTTPAction
}

func (h *HTTPResponses) Init() {
	h.data = []HTTPAction{}
}

func (h *HTTPResponses) Clear() {
	h.Init()
}

func (h *HTTPResponses) GetParserName() string {
	return "http-response"
}

func (h *HTTPResponses) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(h.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return h.data, nil
}

func (h *HTTPResponses) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case []HTTPAction:
		h.data = newValue
	case HTTPAction:
		h.data = append(h.data, newValue)
	case *HTTPAction:
		h.data = append(h.data, *newValue)
	}
	return fmt.Errorf("casting error")
}

func (h *HTTPResponses) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := h.Get(false)
	h.Clear()
	_, err := h.Parse(data, parts, []string{}, comment)
	if err != nil {
		h.Set(oldData)
	}
	return err
}

func (f *HTTPResponses) ParseHTTPResponse(response HTTPAction, parts []string, comment string) error {
	err := response.Parse(parts, comment)
	if err != nil {
		return &errors.ParseError{Parser: "HTTPResponseLines", Line: ""}
	}
	f.data = append(f.data, response)
	return nil
}

func (h *HTTPResponses) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) >= 2 && parts[0] == "http-response" {
		var err error
		switch parts[1] {
		case "add-header":
			err = h.ParseHTTPResponse(&actions.AddHeader{}, parts, comment)
		case "del-header":
			err = h.ParseHTTPResponse(&actions.DelHeader{}, parts, comment)
		case "set-header":
			err = h.ParseHTTPResponse(&actions.SetHeader{}, parts, comment)
		case "set-var":
			err = h.ParseHTTPResponse(&actions.SetVar{}, parts, comment)
		case "allow":
			err = h.ParseHTTPResponse(&actions.Allow{}, parts, comment)
		case "deny":
			err = h.ParseHTTPResponse(&actions.Deny{}, parts, comment)
		case "redirect":
			err = h.ParseHTTPResponse(&actions.Redirect{}, parts, comment)
		case "auth":
			err = h.ParseHTTPResponse(&actions.Auth{}, parts, comment)
		default:
			if strings.HasPrefix(parts[1], "add-acl(") {
				err = h.ParseHTTPResponse(&actions.AddAcl{}, parts, comment)
			} else if strings.HasPrefix(parts[1], "del-acl(") {
				err = h.ParseHTTPResponse(&actions.DelAcl{}, parts, comment)
			} else if strings.HasPrefix(parts[1], "set-var(") {
				err = h.ParseHTTPResponse(&actions.SetVar{}, parts, comment)
			} else {
				return "", &errors.ParseError{Parser: "HTTPResponseLines", Line: line}
			}
		}
		if err != nil {
			return "", err
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "HTTPResponseLines", Line: line}
}

func (h *HTTPResponses) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    "http-request " + req.String(),
			Comment: req.GetComment(),
		}
	}
	return result, nil
}
