package httprequest

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type HTTPRequestAddHeader struct {
	Action        string
	Header        string
	Fmt           string
	Condition     string
	ConditionKind string
}

type HTTPRequestAddHeaders struct {
	HTTPRequestAddHeaders []HTTPRequestAddHeader
}

func (h *HTTPRequestAddHeaders) Init() {
	h.HTTPRequestAddHeaders = []HTTPRequestAddHeader{}
}

func (h *HTTPRequestAddHeaders) GetParserName() string {
	return "http-request add-header"
}

func (h *HTTPRequestAddHeaders) parseHTTPRequestAddHeaderLine(line string, parts []string) (HTTPRequestAddHeader, error) {
	if len(parts) >= 4 {
		command, condition := common.SplitRequest(parts[3:])
		data := HTTPRequestAddHeader{
			Action: parts[1],
			Header: parts[2],
			Fmt:    strings.Join(command, " "),
		}
		if len(condition) > 0 {
			data.ConditionKind = condition[0]
			data.Condition = strings.Join(condition[1:], " ")
		}
		return data, nil
	}
	return HTTPRequestAddHeader{}, &errors.ParseError{Parser: "HTTPRequestAddHeaderLines", Line: line}
}

func (h *HTTPRequestAddHeaders) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if len(parts) > 2 && parts[0] == "http-request" && parts[1] == "add-header" {
		request, err := h.parseHTTPRequestAddHeaderLine(line, parts)
		if err != nil {
			return "", &errors.ParseError{Parser: "HTTPRequestAddHeaderLines", Line: line}
		}
		h.HTTPRequestAddHeaders = append(h.HTTPRequestAddHeaders, request)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "HTTPRequestAddHeaderLines", Line: line}
}

func (h *HTTPRequestAddHeaders) Valid() bool {
	if len(h.HTTPRequestAddHeaders) > 0 {
		return true
	}
	return false
}

func (h *HTTPRequestAddHeaders) Result(AddComments bool) []string {
	result := make([]string, len(h.HTTPRequestAddHeaders))
	for index, req := range h.HTTPRequestAddHeaders {
		condition := ""
		if req.Condition != "" {
			condition = fmt.Sprintf(" %s %s", req.ConditionKind, req.Condition)
		}
		result[index] = fmt.Sprintf("  http-request %s %s %s%s", req.Action, req.Header, req.Fmt, condition)
	}
	return result
}
