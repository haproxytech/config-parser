package http

import (
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/parsers/http/actions"
	"github.com/haproxytech/config-parser/types"
)

type HTTPResponses struct {
	Name string
	data []types.HTTPAction
}

func (h *HTTPResponses) Init() {
	h.Name = "http-response"
	h.data = []types.HTTPAction{}
}

func (f *HTTPResponses) ParseHTTPResponse(response types.HTTPAction, parts []string, comment string) error {
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
		case "allow":
			err = h.ParseHTTPResponse(&actions.Allow{}, parts, comment)
		case "auth":
			err = h.ParseHTTPResponse(&actions.Auth{}, parts, comment)
		case "del-header":
			err = h.ParseHTTPResponse(&actions.DelHeader{}, parts, comment)
		case "deny":
			err = h.ParseHTTPResponse(&actions.Deny{}, parts, comment)
		case "redirect":
			err = h.ParseHTTPResponse(&actions.Redirect{}, parts, comment)
		case "replace-header":
			err = h.ParseHTTPResponse(&actions.ReplaceHeader{}, parts, comment)
		case "replace-value":
			err = h.ParseHTTPResponse(&actions.ReplaceValue{}, parts, comment)
		case "send-spoe-group":
			err = h.ParseHTTPResponse(&actions.SendSpoeGroup{}, parts, comment)
		case "set-header":
			err = h.ParseHTTPResponse(&actions.SetHeader{}, parts, comment)
		case "set-log-level":
			err = h.ParseHTTPResponse(&actions.SetLogLevel{}, parts, comment)
		case "set-status":
			err = h.ParseHTTPResponse(&actions.SetStatus{}, parts, comment)
		case "set-var":
			err = h.ParseHTTPResponse(&actions.SetVar{}, parts, comment)
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
	for index, res := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    "http-response " + res.String(),
			Comment: res.GetComment(),
		}
	}
	return result, nil
}
