package httprequest

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type HTTPRequestDelAcl struct {
	Action  string
	Options []string
}

type HTTPRequestDelAcls struct {
	HTTPRequestDelAcls []HTTPRequestDelAcl
}

func (h *HTTPRequestDelAcls) Init() {
	h.HTTPRequestDelAcls = []HTTPRequestDelAcl{}
}

func (h *HTTPRequestDelAcls) GetParserName() string {
	return "http-request del-acl"
}

func (h *HTTPRequestDelAcls) parseHTTPRequestDelAclLine(line string, parts []string) (HTTPRequestDelAcl, error) {
	if len(parts) >= 3 {
		return HTTPRequestDelAcl{
			Action:  parts[1],
			Options: parts[2:],
		}, nil
	}
	return HTTPRequestDelAcl{}, &errors.ParseError{Parser: "HTTPRequestDelAclLines", Line: line}
}

func (h *HTTPRequestDelAcls) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if len(parts) > 2 && parts[0] == "http-request" && strings.HasPrefix(parts[1], "del-acl(") {
		request, err := h.parseHTTPRequestDelAclLine(line, parts)
		if err != nil {
			return "", &errors.ParseError{Parser: "HTTPRequestDelAclLines", Line: line}
		}
		h.HTTPRequestDelAcls = append(h.HTTPRequestDelAcls, request)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "HTTPRequestDelAclLines", Line: line}
}

func (h *HTTPRequestDelAcls) Valid() bool {
	if len(h.HTTPRequestDelAcls) > 0 {
		return true
	}
	return false
}

func (h *HTTPRequestDelAcls) Result(AddComments bool) []string {
	result := make([]string, len(h.HTTPRequestDelAcls))
	for index, req := range h.HTTPRequestDelAcls {
		options := ""
		if len(req.Options) > 0 {
			options = fmt.Sprintf(" %s", strings.Join(req.Options, " "))
		}
		result[index] = fmt.Sprintf("  http-request %s%s", req.Action, options)
	}
	return result
}
