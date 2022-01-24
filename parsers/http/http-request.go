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

type Requests struct {
	Name        string
	Mode        string
	data        []types.Action
	preComments []string // comments that appear before the the actual line
}

func (h *Requests) Init() {
	h.Name = "http-request"
	h.data = []types.Action{}
}

func (h *Requests) ParseHTTPRequest(request types.Action, parts []string, comment string) error {
	err := request.Parse(parts, types.HTTP, comment)
	if err != nil {
		return &errors.ParseError{Parser: "HTTPRequestLines", Line: ""}
	}
	h.data = append(h.data, request)
	return nil
}

func (h *Requests) Parse(line string, parts []string, comment string) (changeState string, err error) { //nolint:gocyclo,cyclop
	if len(parts) >= 2 && parts[0] == "http-request" {
		var err error
		switch parts[1] {
		case "add-header":
			err = h.ParseHTTPRequest(&http_actions.AddHeader{}, parts, comment)
		case "allow":
			err = h.ParseHTTPRequest(&http_actions.Allow{}, parts, comment)
		case "auth":
			err = h.ParseHTTPRequest(&http_actions.Auth{}, parts, comment)
		case "cache-use":
			err = h.ParseHTTPRequest(&http_actions.CacheUse{}, parts, comment)
		case "capture":
			err = h.ParseHTTPRequest(&http_actions.Capture{}, parts, comment)
		case "del-header":
			err = h.ParseHTTPRequest(&http_actions.DelHeader{}, parts, comment)
		case "deny":
			err = h.ParseHTTPRequest(&http_actions.Deny{}, parts, comment)
		case "disable-l7-retry":
			err = h.ParseHTTPRequest(&http_actions.DisableL7Retry{}, parts, comment)
		case "early-hint":
			err = h.ParseHTTPRequest(&http_actions.EarlyHint{}, parts, comment)
		case "normalize-uri":
			err = h.ParseHTTPRequest(&http_actions.NormalizeURI{}, parts, comment)
		case "redirect":
			err = h.ParseHTTPRequest(&http_actions.Redirect{}, parts, comment)
		case "reject":
			err = h.ParseHTTPRequest(&actions.Reject{}, parts, comment)
		case "replace-header":
			err = h.ParseHTTPRequest(&http_actions.ReplaceHeader{}, parts, comment)
		case "replace-path":
			err = h.ParseHTTPRequest(&http_actions.ReplacePath{}, parts, comment)
		case "replace-pathq":
			err = h.ParseHTTPRequest(&http_actions.ReplacePathQ{}, parts, comment)
		case "replace-uri":
			err = h.ParseHTTPRequest(&http_actions.ReplaceURI{}, parts, comment)
		case "replace-value":
			err = h.ParseHTTPRequest(&http_actions.ReplaceValue{}, parts, comment)
		case "return":
			err = h.ParseHTTPRequest(&http_actions.Return{}, parts, comment)
		case "send-spoe-group":
			err = h.ParseHTTPRequest(&actions.SendSpoeGroup{}, parts, comment)
		case "set-dst":
			err = h.ParseHTTPRequest(&actions.SetDst{}, parts, comment)
		case "set-dst-port":
			err = h.ParseHTTPRequest(&actions.SetDstPort{}, parts, comment)
		case "set-header":
			err = h.ParseHTTPRequest(&http_actions.SetHeader{}, parts, comment)
		case "set-log-level":
			err = h.ParseHTTPRequest(&http_actions.SetLogLevel{}, parts, comment)
		case "set-mark":
			err = h.ParseHTTPRequest(&http_actions.SetMark{}, parts, comment)
		case "set-method":
			err = h.ParseHTTPRequest(&http_actions.SetMethod{}, parts, comment)
		case "set-nice":
			err = h.ParseHTTPRequest(&http_actions.SetNice{}, parts, comment)
		case "set-path":
			err = h.ParseHTTPRequest(&http_actions.SetPath{}, parts, comment)
		case "set-pathq":
			err = h.ParseHTTPRequest(&http_actions.SetPathQ{}, parts, comment)
		case "set-priority-class":
			err = h.ParseHTTPRequest(&actions.SetPriorityClass{}, parts, comment)
		case "set-priority-offset":
			err = h.ParseHTTPRequest(&actions.SetPriorityOffset{}, parts, comment)
		case "set-query":
			err = h.ParseHTTPRequest(&http_actions.SetQuery{}, parts, comment)
		case "set-src":
			err = h.ParseHTTPRequest(&http_actions.SetSrc{}, parts, comment)
		case "set-src-port":
			err = h.ParseHTTPRequest(&http_actions.SetSrcPort{}, parts, comment)
		case "set-timeout":
			err = h.ParseHTTPRequest(&http_actions.SetTimeout{}, parts, comment)
		case "set-tos":
			err = h.ParseHTTPRequest(&http_actions.SetTos{}, parts, comment)
		case "set-uri":
			err = h.ParseHTTPRequest(&http_actions.SetURI{}, parts, comment)
		case "silent-drop":
			err = h.ParseHTTPRequest(&actions.SilentDrop{}, parts, comment)
		case "strict-mode":
			err = h.ParseHTTPRequest(&http_actions.StrictMode{}, parts, comment)
		case "tarpit":
			err = h.ParseHTTPRequest(&http_actions.Tarpit{}, parts, comment)
		case "track-sc0":
			err = h.ParseHTTPRequest(&actions.TrackSc{}, parts, comment)
		case "track-sc1":
			err = h.ParseHTTPRequest(&actions.TrackSc{}, parts, comment)
		case "track-sc2":
			err = h.ParseHTTPRequest(&actions.TrackSc{}, parts, comment)
		case "use-service":
			err = h.ParseHTTPRequest(&actions.UseService{}, parts, comment)
		case "wait-for-handshake":
			err = h.ParseHTTPRequest(&http_actions.WaitForHandshake{}, parts, comment)
		default:
			switch {
			case strings.HasPrefix(parts[1], "add-acl("):
				err = h.ParseHTTPRequest(&http_actions.AddACL{}, parts, comment)
			case strings.HasPrefix(parts[1], "del-acl("):
				err = h.ParseHTTPRequest(&http_actions.DelACL{}, parts, comment)
			case strings.HasPrefix(parts[1], "set-map("):
				err = h.ParseHTTPRequest(&http_actions.SetMap{}, parts, comment)
			case strings.HasPrefix(parts[1], "del-map("):
				err = h.ParseHTTPRequest(&http_actions.DelMap{}, parts, comment)
			case strings.HasPrefix(parts[1], "do-resolve("):
				err = h.ParseHTTPRequest(&actions.DoResolve{}, parts, comment)
			case strings.HasPrefix(parts[1], "lua."):
				err = h.ParseHTTPRequest(&actions.Lua{}, parts, comment)
			case strings.HasPrefix(parts[1], "sc-inc-gpc0("):
				err = h.ParseHTTPRequest(&actions.ScIncGpc0{}, parts, comment)
			case strings.HasPrefix(parts[1], "sc-inc-gpc1("):
				err = h.ParseHTTPRequest(&actions.ScIncGpc1{}, parts, comment)
			case strings.HasPrefix(parts[1], "sc-set-gpt0("):
				err = h.ParseHTTPRequest(&actions.ScSetGpt0{}, parts, comment)
			case strings.HasPrefix(parts[1], "set-var("):
				err = h.ParseHTTPRequest(&actions.SetVar{}, parts, comment)
			case strings.HasPrefix(parts[1], "unset-var("):
				err = h.ParseHTTPRequest(&actions.UnsetVar{}, parts, comment)
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
