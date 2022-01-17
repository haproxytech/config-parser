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
	"strings"

	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/parsers/actions"
	tcp_actions "github.com/haproxytech/config-parser/v4/parsers/tcp/actions"
	"github.com/haproxytech/config-parser/v4/types"
)

type Checks struct {
	Name        string
	Mode        string
	data        []types.Action
	preComments []string // comments that appear before the the actual line
}

func (h *Checks) Init() {
	h.Name = "tcp-check"
	h.data = []types.Action{}
}

func (h *Checks) parseTCPCheck(request types.Action, parts []string, comment string) error {
	err := request.Parse(parts, types.TCP, comment)
	if err != nil {
		return &errors.ParseError{Parser: "TCPCheck", Line: "", Message: err.Error()}
	}
	h.data = append(h.data, request)
	return nil
}

func (h *Checks) Parse(line string, parts []string, comment string) (string, error) {
	if len(parts) < 2 {
		return "", &errors.ParseError{Parser: "TCPCheck", Line: line, Message: "tcp-check type not provided"}
	}

	if parts[0] != h.Name {
		return "", &errors.ParseError{Parser: "TCPCheck", Line: line, Message: "name is not tcp-check"}
	}

	if h.Mode == "frontend" {
		return "", &errors.ParseError{Parser: "TCPCheck", Line: line, Message: "tcp-check cannot be used in frontend section"}
	}

	var err error

	switch {
	case parts[1] == "comment":
		err = h.parseTCPCheck(&tcp_actions.CheckComment{}, parts, comment)
	case parts[1] == "connect":
		err = h.parseTCPCheck(&actions.CheckConnect{}, parts, comment)
	case parts[1] == "expect":
		err = h.parseTCPCheck(&actions.CheckExpect{}, parts, comment)
	case parts[1] == "send":
		err = h.parseTCPCheck(&tcp_actions.CheckSend{}, parts, comment)
	case parts[1] == "send-lf":
		err = h.parseTCPCheck(&tcp_actions.CheckSendLf{}, parts, comment)
	case parts[1] == "send-binary":
		err = h.parseTCPCheck(&tcp_actions.CheckSendBinary{}, parts, comment)
	case parts[1] == "send-binary-lf":
		err = h.parseTCPCheck(&tcp_actions.CheckSendBinaryLf{}, parts, comment)
	case strings.HasPrefix(parts[1], "set-var("):
		err = h.parseTCPCheck(&actions.SetVarCheck{}, parts, comment)
	case strings.HasPrefix(parts[1], "set-var-fmt("):
		err = h.parseTCPCheck(&tcp_actions.SetVarFmtCheck{}, parts, comment)
	case strings.HasPrefix(parts[1], "unset-var("):
		err = h.parseTCPCheck(&actions.UnsetVarCheck{}, parts, comment)
	default:
		err = &errors.ParseError{Parser: "TCPCheck", Line: line, Message: "invalid tcp-check type provided"}
	}

	return "", err
}

func (h *Checks) Result() ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.ErrFetch
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    "tcp-check " + req.String(),
			Comment: req.GetComment(),
		}
	}
	return result, nil
}
