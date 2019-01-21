package http

import (
	"strings"

	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/parsers/http/actions"
)

type HTTPResponses struct {
	HTTPResponses []HTTPAction
}

func (h *HTTPResponses) Init() {
	h.HTTPResponses = []HTTPAction{}
}

func (h *HTTPResponses) GetParserName() string {
	return "http-response"
}

func (f *HTTPResponses) ParseHTTPResponse(response HTTPAction, parts []string, comment string) error {
	err := response.Parse(parts, "")
	if err != nil {
		return &errors.ParseError{Parser: "HTTPResponseLines", Line: ""}
	}
	f.HTTPResponses = append(f.HTTPResponses, response)
	return nil
}

func (h *HTTPResponses) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	comment := ""
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

func (h *HTTPResponses) Valid() bool {
	if len(h.HTTPResponses) > 0 {
		return true
	}
	return false
}

func (h *HTTPResponses) Result(AddComments bool) []string {
	result := make([]string, len(h.HTTPResponses))
	for index, req := range h.HTTPResponses {
		result[index] = "  http-response " + req.String()
	}
	return result
}
