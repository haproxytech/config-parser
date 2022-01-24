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

	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/parsers/actions"
	http_actions "github.com/haproxytech/config-parser/v4/parsers/http/actions"
	"github.com/haproxytech/config-parser/v4/types"
)

type Responses struct {
	Name        string
	Mode        string
	data        []types.Action
	preComments []string // comments that appear before the the actual line
}

func (h *Responses) Init() {
	h.Name = "http-response"
	h.data = []types.Action{}
}

func (h *Responses) ParseHTTPResponse(response types.Action, parts []string, comment string) error {
	err := response.Parse(parts, types.HTTP, comment)
	if err != nil {
		return &errors.ParseError{Parser: "HTTPResponseLines", Line: ""}
	}
	h.data = append(h.data, response)
	return nil
}

func (h *Responses) Parse(line string, parts []string, comment string) (changeState string, err error) { //nolint:gocyclo,cyclop,cyclop
	if len(parts) >= 2 && parts[0] == "http-response" {
		var err error
		switch parts[1] {
		case "add-header":
			err = h.ParseHTTPResponse(&http_actions.AddHeader{}, parts, comment)
		case "allow":
			err = h.ParseHTTPResponse(&http_actions.Allow{}, parts, comment)
		case "cache-store":
			err = h.ParseHTTPResponse(&http_actions.CacheStore{}, parts, comment)
		case "capture":
			err = h.ParseHTTPResponse(&http_actions.Capture{}, parts, comment)
		case "del-header":
			err = h.ParseHTTPResponse(&http_actions.DelHeader{}, parts, comment)
		case "deny":
			err = h.ParseHTTPResponse(&http_actions.Deny{}, parts, comment)
		case "redirect":
			err = h.ParseHTTPResponse(&http_actions.Redirect{}, parts, comment)
		case "replace-header":
			err = h.ParseHTTPResponse(&http_actions.ReplaceHeader{}, parts, comment)
		case "replace-value":
			err = h.ParseHTTPResponse(&http_actions.ReplaceValue{}, parts, comment)
		case "return":
			err = h.ParseHTTPResponse(&http_actions.Return{}, parts, comment)
		case "send-spoe-group":
			err = h.ParseHTTPResponse(&actions.SendSpoeGroup{}, parts, comment)
		case "set-header":
			err = h.ParseHTTPResponse(&http_actions.SetHeader{}, parts, comment)
		case "set-log-level":
			err = h.ParseHTTPResponse(&http_actions.SetLogLevel{}, parts, comment)
		case "set-mark":
			err = h.ParseHTTPResponse(&http_actions.SetMark{}, parts, comment)
		case "set-nice":
			err = h.ParseHTTPResponse(&http_actions.SetNice{}, parts, comment)
		case "set-status":
			err = h.ParseHTTPResponse(&http_actions.SetStatus{}, parts, comment)
		case "set-tos":
			err = h.ParseHTTPResponse(&http_actions.SetTos{}, parts, comment)
		case "silent-drop":
			err = h.ParseHTTPResponse(&actions.SilentDrop{}, parts, comment)
		case "strict-mode":
			err = h.ParseHTTPResponse(&http_actions.StrictMode{}, parts, comment)
		case "track-sc0":
			err = h.ParseHTTPResponse(&actions.TrackSc{}, parts, comment)
		case "track-sc1":
			err = h.ParseHTTPResponse(&actions.TrackSc{}, parts, comment)
		case "track-sc2":
			err = h.ParseHTTPResponse(&actions.TrackSc{}, parts, comment)
		case "wait-for-body":
			err = h.ParseHTTPResponse(&http_actions.WaitForBody{}, parts, comment)
		default:
			switch {
			case strings.HasPrefix(parts[1], "add-acl("):
				err = h.ParseHTTPResponse(&http_actions.AddACL{}, parts, comment)
			case strings.HasPrefix(parts[1], "del-acl("):
				err = h.ParseHTTPResponse(&http_actions.DelACL{}, parts, comment)
			case strings.HasPrefix(parts[1], "lua."):
				err = h.ParseHTTPResponse(&actions.Lua{}, parts, comment)
			case strings.HasPrefix(parts[1], "sc-inc-gpc0("):
				err = h.ParseHTTPResponse(&actions.ScIncGpc0{}, parts, comment)
			case strings.HasPrefix(parts[1], "sc-inc-gpc1("):
				err = h.ParseHTTPResponse(&actions.ScIncGpc1{}, parts, comment)
			case strings.HasPrefix(parts[1], "sc-set-gpt0("):
				err = h.ParseHTTPResponse(&actions.ScSetGpt0{}, parts, comment)
			case strings.HasPrefix(parts[1], "set-map("):
				err = h.ParseHTTPResponse(&http_actions.SetMap{}, parts, comment)
			case strings.HasPrefix(parts[1], "del-map("):
				err = h.ParseHTTPResponse(&http_actions.DelMap{}, parts, comment)
			case strings.HasPrefix(parts[1], "set-var("):
				err = h.ParseHTTPResponse(&actions.SetVar{}, parts, comment)
			case strings.HasPrefix(parts[1], "unset-var("):
				err = h.ParseHTTPResponse(&actions.UnsetVar{}, parts, comment)
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
