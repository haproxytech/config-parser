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

type Responses struct {
	Name string
	data []types.HTTPAction
}

func (h *Responses) Init() {
	h.Name = "http-response"
	h.data = []types.HTTPAction{}
}

func (h *Responses) ParseHTTPResponse(response types.HTTPAction, parts []string, comment string) error {
	err := response.Parse(parts, comment)
	if err != nil {
		return &errors.ParseError{Parser: "HTTPResponseLines", Line: ""}
	}
	h.data = append(h.data, response)
	return nil
}

func (h *Responses) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) >= 2 && parts[0] == "http-response" {
		var err error
		switch parts[1] {
		case "add-header":
			err = h.ParseHTTPResponse(&actions.AddHeader{}, parts, comment)
		case "allow":
			err = h.ParseHTTPResponse(&actions.Allow{}, parts, comment)
		case "auth":
			err = h.ParseHTTPResponse(&actions.Auth{}, parts, comment)
		case "del-header":
			err = h.ParseHTTPResponse(&actions.DelHeader{}, parts, comment)
		case "deny":
			err = h.ParseHTTPResponse(&actions.Deny{}, parts, comment)
		case "redirect":
			err = h.ParseHTTPResponse(&actions.Redirect{}, parts, comment)
		case "replace-header":
			err = h.ParseHTTPResponse(&actions.ReplaceHeader{}, parts, comment)
		case "replace-value":
			err = h.ParseHTTPResponse(&actions.ReplaceValue{}, parts, comment)
		case "send-spoe-group":
			err = h.ParseHTTPResponse(&actions.SendSpoeGroup{}, parts, comment)
		case "set-header":
			err = h.ParseHTTPResponse(&actions.SetHeader{}, parts, comment)
		case "set-log-level":
			err = h.ParseHTTPResponse(&actions.SetLogLevel{}, parts, comment)
		case "set-status":
			err = h.ParseHTTPResponse(&actions.SetStatus{}, parts, comment)
		case "set-var":
			err = h.ParseHTTPResponse(&actions.SetVar{}, parts, comment)
		default:
			switch {
			case strings.HasPrefix(parts[1], "add-acl("):
				err = h.ParseHTTPResponse(&actions.AddACL{}, parts, comment)
			case strings.HasPrefix(parts[1], "del-acl("):
				err = h.ParseHTTPResponse(&actions.DelACL{}, parts, comment)
			case strings.HasPrefix(parts[1], "set-var("):
				err = h.ParseHTTPResponse(&actions.SetVar{}, parts, comment)
			default:
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

func (h *Responses) Result() ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.ErrFetch
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, res := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    "http-response " + res.String(),
			Comment: res.GetComment(),
		}
	}
	return result, nil
}
