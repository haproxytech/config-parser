// Code generated by go generate; DO NOT EDIT.
/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package http

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/errors"
	parsersactions "github.com/haproxytech/config-parser/v4/parsers/actions"
	"github.com/haproxytech/config-parser/v4/parsers/http/actions"
	"github.com/haproxytech/config-parser/v4/types"
)

type AfterResponses struct {
	Name string
	// Mode string
	data        []types.Action
	preComments []string // comments that appear before the actual line
}

func (p *AfterResponses) Init() {
	p.Name = "http-after-response"
	p.data = []types.Action{}
	p.preComments = []string{}
}

func (h *AfterResponses) Parse(line string, parts []string, comment string) (string, error) {
	var err error
	if len(parts) == 0 {
		return "", &errors.ParseError{Parser: parseErrorLines(h.Name), Line: line, Message: "missing attribute"}
	}

	if parts[0] != h.Name {
		return "", &errors.ParseError{Parser: parseErrorLines(h.Name), Line: line, Message: "expected attribute http-after-response"}
	}

	if len(parts) == 1 {
		return "", &errors.ParseError{Parser: parseErrorLines(h.Name), Line: line, Message: "expected action for http-after-response"}
	}

	action := parts[1]

	switch {
	case action == "add-header":
		err = h.ParseHTTPRequest(&actions.AddHeader{}, parts, comment)
	case action == "allow":
		err = h.ParseHTTPRequest(&actions.Allow{}, parts, comment)
	case action == "del-header":
		err = h.ParseHTTPRequest(&actions.DelHeader{}, parts, comment)
	case action == "replace-header":
		err = h.ParseHTTPRequest(&actions.ReplaceHeader{}, parts, comment)
	case action == "replace-value":
		err = h.ParseHTTPRequest(&actions.ReplaceValue{}, parts, comment)
	case action == "set-header":
		err = h.ParseHTTPRequest(&actions.SetHeader{}, parts, comment)
	case action == "set-status":
		err = h.ParseHTTPRequest(&actions.SetStatus{}, parts, comment)
	case action == "strict-mode":
		err = h.ParseHTTPRequest(&actions.StrictMode{}, parts, comment)
	case strings.HasPrefix(parts[1], "unset-var("):
		err = h.ParseHTTPRequest(&parsersactions.UnsetVar{}, parts, comment)
	case strings.HasPrefix(parts[1], "set-var("):
		err = h.ParseHTTPRequest(&parsersactions.SetVar{}, parts, comment)
	default:
		err = fmt.Errorf("unsupported action %s", action)
	}

	if err != nil {
		return "", err
	}

	return "", nil
}

func (h *AfterResponses) ParseHTTPRequest(request types.Action, parts []string, comment string) error {
	err := request.Parse(parts, types.HTTP, comment)
	if err != nil {
		return &errors.ParseError{Parser: "HTTPAfterResponses", Line: ""}
	}

	h.data = append(h.data, request)

	return nil
}

func (h *AfterResponses) Result() ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.ErrFetch
	}

	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("%s %s", h.Name, req.String()),
			Comment: req.GetComment(),
		}
	}
	return result, nil
}

func parseErrorLines(s string) string {
	var r string
	parts := strings.Split(s, "-")
	if len(parts) == 1 {
		r = strings.Title(parts[0])
	} else {
		r = strings.Title(parts[0]) + strings.Title(parts[1])
	}
	return r + "Lines"
}
