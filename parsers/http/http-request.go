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
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/parsers/http/actions"
	"github.com/haproxytech/config-parser/types"
)

type Requests struct {
	Name string
	data []types.HTTPAction
}

func (h *Requests) Init() {
	h.Name = "http-request"
	h.data = []types.HTTPAction{}
}

func (h *Requests) ParseHTTPRequest(request types.HTTPAction, parts []string, comment string) error {
	err := request.Parse(parts, comment)
	if err != nil {
		return &errors.ParseError{Parser: "HTTPRequestLines", Line: ""}
	}
	h.data = append(h.data, request)
	return nil
}

func (h *Requests) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) >= 2 && parts[0] == "http-request" {
		var err error
		switch parts[1] {
		case "add-header":
			err = h.ParseHTTPRequest(&actions.AddHeader{}, parts, comment)
		case "allow":
			err = h.ParseHTTPRequest(&actions.Allow{}, parts, comment)
		case "auth":
			err = h.ParseHTTPRequest(&actions.Auth{}, parts, comment)
		case "del-header":
			err = h.ParseHTTPRequest(&actions.DelHeader{}, parts, comment)
		case "deny":
			err = h.ParseHTTPRequest(&actions.Deny{}, parts, comment)
		case "redirect":
			err = h.ParseHTTPRequest(&actions.Redirect{}, parts, comment)
		case "reject":
			err = h.ParseHTTPRequest(&actions.Reject{}, parts, comment)
		case "replace-header":
			err = h.ParseHTTPRequest(&actions.ReplaceHeader{}, parts, comment)
		case "replace-value":
			err = h.ParseHTTPRequest(&actions.ReplaceValue{}, parts, comment)
		case "send-spoe-group":
			err = h.ParseHTTPRequest(&actions.SendSpoeGroup{}, parts, comment)
		case "set-header":
			err = h.ParseHTTPRequest(&actions.SetHeader{}, parts, comment)
		case "set-log-level":
			err = h.ParseHTTPRequest(&actions.SetLogLevel{}, parts, comment)
		case "set-path":
			err = h.ParseHTTPRequest(&actions.SetPath{}, parts, comment)
		case "set-query":
			err = h.ParseHTTPRequest(&actions.SetQuery{}, parts, comment)
		case "set-uri":
			err = h.ParseHTTPRequest(&actions.SetURI{}, parts, comment)
		case "set-var":
			err = h.ParseHTTPRequest(&actions.SetVar{}, parts, comment)
		case "tarpit":
			err = h.ParseHTTPRequest(&actions.Tarpit{}, parts, comment)
		default:
			switch {
			case strings.HasPrefix(parts[1], "add-acl("):
				err = h.ParseHTTPRequest(&actions.AddACL{}, parts, comment)
			case strings.HasPrefix(parts[1], "del-acl("):
				err = h.ParseHTTPRequest(&actions.DelACL{}, parts, comment)
			case strings.HasPrefix(parts[1], "set-var("):
				err = h.ParseHTTPRequest(&actions.SetVar{}, parts, comment)
			default:
				return "", &errors.ParseError{Parser: "HTTPRequestLines", Line: line}
			}
		}
		if err != nil {
			return "", err
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "HTTPRequestLines", Line: line}
}

func (h *Requests) Result() ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.ErrFetch
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    "http-request " + req.String(),
			Comment: req.GetComment(),
		}
	}
	return result, nil
}
