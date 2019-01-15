package httprequest

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type HTTPRequestAddAcl struct {
	Action  string
	Options []string
}

type HTTPRequestAddAcls struct {
	HTTPRequestAddAcls []HTTPRequestAddAcl
}

func (h *HTTPRequestAddAcls) Init() {
	h.HTTPRequestAddAcls = []HTTPRequestAddAcl{}
}

func (h *HTTPRequestAddAcls) GetParserName() string {
	return "http-request add-acl"
}

func (h *HTTPRequestAddAcls) parseHTTPRequestAddAclLine(line string, parts []string) (HTTPRequestAddAcl, error) {
	if len(parts) >= 3 {
		return HTTPRequestAddAcl{
			Action:  parts[1],
			Options: parts[2:],
		}, nil
	}
	return HTTPRequestAddAcl{}, &errors.ParseError{Parser: "HTTPRequestAddAclLines", Line: line}
}

func (h *HTTPRequestAddAcls) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if len(parts) > 2 && parts[0] == "http-request" && strings.HasPrefix(parts[1], "add-acl(") {
		request, err := h.parseHTTPRequestAddAclLine(line, parts)
		if err != nil {
			return "", &errors.ParseError{Parser: "HTTPRequestAddAclLines", Line: line}
		}
		h.HTTPRequestAddAcls = append(h.HTTPRequestAddAcls, request)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "HTTPRequestAddAclLines", Line: line}
}

func (h *HTTPRequestAddAcls) Valid() bool {
	if len(h.HTTPRequestAddAcls) > 0 {
		return true
	}
	return false
}

func (h *HTTPRequestAddAcls) Result(AddComments bool) []string {
	result := make([]string, len(h.HTTPRequestAddAcls))
	for index, req := range h.HTTPRequestAddAcls {
		options := ""
		if len(req.Options) > 0 {
			options = fmt.Sprintf(" %s", strings.Join(req.Options, " "))
		}
		result[index] = fmt.Sprintf("  http-request %s%s #add-acl", req.Action, options)
	}
	return result
}
