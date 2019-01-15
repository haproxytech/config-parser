package httprequest

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type HTTPRequest struct {
	Action  string
	Options []string
}

type HTTPRequests struct {
	HTTPRequests []HTTPRequest
}

func (h *HTTPRequests) Init() {
	h.HTTPRequests = []HTTPRequest{}
}

func (h *HTTPRequests) GetParserName() string {
	return "http-request"
}

func (h *HTTPRequests) parseHTTPRequestLine(line string, parts []string) (HTTPRequest, error) {
	if len(parts) >= 3 {
		return HTTPRequest{
			Action:  parts[1],
			Options: parts[2:],
		}, nil
	}
	if len(parts) >= 2 {
		return HTTPRequest{
			Action: parts[1],
		}, nil
	}
	return HTTPRequest{}, &errors.ParseError{Parser: "HTTPRequestLines", Line: line}
}

func (h *HTTPRequests) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if parts[0] == "http-request" {
		request, err := h.parseHTTPRequestLine(line, parts)
		if err != nil {
			return "", &errors.ParseError{Parser: "HTTPRequestLines", Line: line}
		}
		h.HTTPRequests = append(h.HTTPRequests, request)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "HTTPRequestLines", Line: line}
}

func (h *HTTPRequests) Valid() bool {
	if len(h.HTTPRequests) > 0 {
		return true
	}
	return false
}

func (h *HTTPRequests) Result(AddComments bool) []string {
	result := make([]string, len(h.HTTPRequests))
	for index, req := range h.HTTPRequests {
		options := ""
		if len(req.Options) > 0 {
			options = fmt.Sprintf(" %s", strings.Join(req.Options, " "))
		}
		result[index] = fmt.Sprintf("  http-request %s%s", req.Action, options)
	}
	return result
}
