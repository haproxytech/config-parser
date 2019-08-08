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

package tcp

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/parsers/tcp/actions"
	"github.com/haproxytech/config-parser/types"
)

type Requests struct {
	Name string
	Mode string //frontent, backend
	data []types.TCPAction
}

func (h *Requests) Init() {
	h.Name = "tcp-request"
	h.data = []types.TCPAction{}
}

func (h *Requests) ParseTCPRequest(request types.TCPAction, parts []string, comment string) error {
	err := request.Parse(parts, comment)
	if err != nil {
		return &errors.ParseError{Parser: "HTTPRequestLines", Line: ""}
	}
	h.data = append(h.data, request)
	return nil
}

func (h *Requests) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) >= 2 && parts[0] == "tcp-request" {
		var err error
		switch parts[1] {
		case "connection":
			if h.Mode == "backend" {
				return "", &errors.ParseError{Parser: "HTTPRequestLines", Line: line}
			}
			err = h.ParseTCPRequest(&actions.Connection{}, parts, comment)
		case "session":
			if h.Mode == "backend" {
				return "", &errors.ParseError{Parser: "HTTPRequestLines", Line: line}
			}
			err = h.ParseTCPRequest(&actions.Session{}, parts, comment)
		case "content":
			err = h.ParseTCPRequest(&actions.Content{}, parts, comment)
		case "inspect-delay":
			err = h.ParseTCPRequest(&actions.InspectDelay{}, parts, comment)
		default:
			return "", &errors.ParseError{Parser: "HTTPRequestLines", Line: line}
		}
		if err != nil {
			return "", err
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "HTTPRequestLines", Line: line}
}

func (h *Requests) Result() ([]common.ReturnResultLine, error) {
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    "tcp-request " + req.String(),
			Comment: req.GetComment(),
		}
	}
	return result, nil
}
