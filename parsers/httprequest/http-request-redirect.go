package httprequest

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type HTTPRequestRedirect struct {
	Action        string
	Header        string
	Fmt           string
	Condition     string
	ConditionKind string
}

type HTTPRequestRedirects struct {
	HTTPRequestRedirects []HTTPRequestRedirect
}

func (h *HTTPRequestRedirects) Init() {
	h.HTTPRequestRedirects = []HTTPRequestRedirect{}
}

func (h *HTTPRequestRedirects) GetParserName() string {
	return "http-request redirect"
}

func (h *HTTPRequestRedirects) parseHTTPRequestRedirectLine(line string, parts []string) (HTTPRequestRedirect, error) {
	if len(parts) >= 4 {
		command, condition := common.SplitRequest(parts[3:])
		data := HTTPRequestRedirect{
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
	return HTTPRequestRedirect{}, &errors.ParseError{Parser: "HTTPRequestRedirectLines", Line: line}
}

func (h *HTTPRequestRedirects) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if len(parts) > 2 && parts[0] == "http-request" && parts[1] == "redirect" {
		request, err := h.parseHTTPRequestRedirectLine(line, parts)
		if err != nil {
			return "", &errors.ParseError{Parser: "HTTPRequestRedirectLines", Line: line}
		}
		h.HTTPRequestRedirects = append(h.HTTPRequestRedirects, request)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "HTTPRequestRedirectLines", Line: line}
}

func (h *HTTPRequestRedirects) Valid() bool {
	if len(h.HTTPRequestRedirects) > 0 {
		return true
	}
	return false
}

func (h *HTTPRequestRedirects) Result(AddComments bool) []string {
	result := make([]string, len(h.HTTPRequestRedirects))
	for index, req := range h.HTTPRequestRedirects {
		condition := ""
		if req.Condition != "" {
			condition = fmt.Sprintf(" %s %s", req.ConditionKind, req.Condition)
		}
		result[index] = fmt.Sprintf("  http-request %s %s %s%s", req.Action, req.Header, req.Fmt, condition)
	}
	return result
}
