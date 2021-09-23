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

package types

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/parsers/actions"
	tcp_actions "github.com/haproxytech/config-parser/v4/parsers/tcp/actions"
	"github.com/haproxytech/config-parser/v4/types"
)

type Connection struct {
	Action  types.Action
	Comment string
}

func (f *Connection) ParseAction(action types.Action, parts []string) error {
	if action.Parse(parts, types.TCP, "") != nil {
		return &errors.ParseError{Parser: "TCPRequestContent", Line: ""}
	}

	f.Action = action
	return nil
}

func (f *Connection) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) < 3 {
		return fmt.Errorf("not enough params")
	}
	var err error
	switch parts[2] {
	case "accept":
		err = f.ParseAction(&tcp_actions.Accept{}, parts)
	case "reject":
		err = f.ParseAction(&actions.Reject{}, parts)
	case "expect-proxy":
		err = f.ParseAction(&tcp_actions.ExpectProxy{}, parts)
	case "expect-netscaler-cip":
		err = f.ParseAction(&tcp_actions.ExpectNetscalerCip{}, parts)
	case "capture":
		err = f.ParseAction(&tcp_actions.Capture{}, parts)
	case "track-sc0":
		err = f.ParseAction(&actions.TrackSc{}, parts)
	case "track-sc1":
		err = f.ParseAction(&actions.TrackSc{}, parts)
	case "track-sc2":
		err = f.ParseAction(&actions.TrackSc{}, parts)
	case "set-src":
		err = f.ParseAction(&tcp_actions.SetSrc{}, parts)
	case "silent-drop":
		err = f.ParseAction(&actions.SilentDrop{}, parts)
	default:
		switch {
		case strings.HasPrefix(parts[2], "lua."):
			err = f.ParseAction(&actions.Lua{}, parts)
		case strings.HasPrefix(parts[2], "sc-inc-gpc0"):
			err = f.ParseAction(&actions.ScIncGpc0{}, parts)
		case strings.HasPrefix(parts[2], "sc-inc-gpc1"):
			err = f.ParseAction(&actions.ScIncGpc1{}, parts)
		case strings.HasPrefix(parts[2], "sc-set-gpt0"):
			err = f.ParseAction(&actions.ScSetGpt0{}, parts)
		}
	}
	return err
}

func (f *Connection) String() string {
	var result strings.Builder

	result.WriteString("connection")
	result.WriteString(" ")
	result.WriteString(f.Action.String())

	if f.Comment != "" {
		result.WriteString(" # ")
		result.WriteString(f.Comment)
	}

	return result.String()
}

func (f *Connection) GetComment() string {
	return f.Comment
}
