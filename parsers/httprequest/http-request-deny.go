package httprequest

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type HTTPRequestDeny struct {
	Action string
	//Header        string
	//Fmt           string
	DenyStatus    string
	Condition     string
	ConditionKind string
}

type HTTPRequestDenials struct {
	HTTPRequestDenials []HTTPRequestDeny
}

func (h *HTTPRequestDenials) Init() {
	h.HTTPRequestDenials = []HTTPRequestDeny{}
}

func (h *HTTPRequestDenials) GetParserName() string {
	return "http-request deny"
}

func (h *HTTPRequestDenials) parseHTTPRequestDenyLine(line string, parts []string) (HTTPRequestDeny, error) {
	if len(parts) >= 4 {
		command, condition := common.SplitRequest(parts[2:])
		data := HTTPRequestDeny{
			Action: parts[1],
			//Header: parts[2],
			//Fmt:    strings.Join(command, " "),
		}
		if len(command) > 1 && command[0] == "deny_status" {
			data.DenyStatus = command[1]
		}
		if len(condition) > 0 {
			data.ConditionKind = condition[0]
			data.Condition = strings.Join(condition[1:], " ")
		}
		return data, nil
	}
	return HTTPRequestDeny{}, &errors.ParseError{Parser: "HTTPRequestDenyLines", Line: line}
}

func (h *HTTPRequestDenials) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if len(parts) > 2 && parts[0] == "http-request" && parts[1] == "deny" {
		request, err := h.parseHTTPRequestDenyLine(line, parts)
		if err != nil {
			return "", &errors.ParseError{Parser: "HTTPRequestDenyLines", Line: line}
		}
		h.HTTPRequestDenials = append(h.HTTPRequestDenials, request)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "HTTPRequestDenyLines", Line: line}
}

func (h *HTTPRequestDenials) Valid() bool {
	if len(h.HTTPRequestDenials) > 0 {
		return true
	}
	return false
}

func (h *HTTPRequestDenials) Result(AddComments bool) []string {
	result := make([]string, len(h.HTTPRequestDenials))
	for index, req := range h.HTTPRequestDenials {
		denyStatus := ""
		if req.DenyStatus != "" {
			denyStatus = fmt.Sprintf(" deny_status %s", req.DenyStatus)
		}
		condition := ""
		if req.Condition != "" {
			condition = fmt.Sprintf(" %s %s", req.ConditionKind, req.Condition)
		}
		result[index] = fmt.Sprintf("  http-request %s%s%s #deny", req.Action, denyStatus, condition)
	}
	return result
}
