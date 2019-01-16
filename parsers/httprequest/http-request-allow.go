package httprequest

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type HTTPRequestAllow struct {
	Action        string
	Header        string
	Fmt           string
	Condition     string
	ConditionKind string
}

type HTTPRequestAllows struct {
	HTTPRequestAllows []HTTPRequestAllow
}

func (h *HTTPRequestAllows) Init() {
	h.HTTPRequestAllows = []HTTPRequestAllow{}
}

func (h *HTTPRequestAllows) GetParserName() string {
	return "http-request allow"
}

func (h *HTTPRequestAllows) parseHTTPRequestAllowLine(line string, parts []string) (HTTPRequestAllow, error) {
	if len(parts) >= 4 {
		_, condition := common.SplitRequest(parts[2:]) // 2 not 3 !
		data := HTTPRequestAllow{
			Action: parts[1],
			//Header: parts[2], allow does not have this
			//Fmt:    strings.Join(command, " "),
		}
		if len(condition) > 0 {
			data.ConditionKind = condition[0]
			data.Condition = strings.Join(condition[1:], " ")
		}
		return data, nil
	}
	return HTTPRequestAllow{}, &errors.ParseError{Parser: "HTTPRequestAllowLines", Line: line}
}

func (h *HTTPRequestAllows) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if len(parts) > 2 && parts[0] == "http-request" && parts[1] == "allow" {
		request, err := h.parseHTTPRequestAllowLine(line, parts)
		if err != nil {
			return "", &errors.ParseError{Parser: "HTTPRequestAllowLines", Line: line}
		}
		h.HTTPRequestAllows = append(h.HTTPRequestAllows, request)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "HTTPRequestAllowLines", Line: line}
}

func (h *HTTPRequestAllows) Valid() bool {
	if len(h.HTTPRequestAllows) > 0 {
		return true
	}
	return false
}

func (h *HTTPRequestAllows) Result(AddComments bool) []string {
	result := make([]string, len(h.HTTPRequestAllows))
	for index, req := range h.HTTPRequestAllows {
		condition := ""
		if req.Condition != "" {
			condition = fmt.Sprintf("%s %s", req.ConditionKind, req.Condition) //no extra space!
		}
		result[index] = fmt.Sprintf("  http-request %s %s", req.Action, condition)
	}
	return result
}
