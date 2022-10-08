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
	tcpActions "github.com/haproxytech/config-parser/v4/parsers/tcp/actions"
	"github.com/haproxytech/config-parser/v4/types"
)

type Session struct {
	Action  types.Action
	Comment string
}

func (f *Session) ParseAction(action types.Action, parts []string) error {
	if action.Parse(parts, types.TCP, "") != nil {
		return &errors.ParseError{Parser: "TCPRequestSession", Line: ""}
	}

	f.Action = action
	return nil
}

func (f *Session) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) < 3 {
		return fmt.Errorf("not enough params")
	}
	var err error
	switch parts[2] {
	case "accept":
		err = f.ParseAction(&tcpActions.Accept{}, parts)
	case "reject":
		err = f.ParseAction(&actions.Reject{}, parts)
	case "track-sc0":
		err = f.ParseAction(&actions.TrackSc{}, parts)
	case "track-sc1":
		err = f.ParseAction(&actions.TrackSc{}, parts)
	case "track-sc2":
		err = f.ParseAction(&actions.TrackSc{}, parts)
	case "silent-drop":
		err = f.ParseAction(&actions.SilentDrop{}, parts)
	default:
		switch {
		case strings.HasPrefix(parts[2], "sc-inc-gpc0"):
			err = f.ParseAction(&actions.ScIncGpc0{}, parts)
		case strings.HasPrefix(parts[2], "sc-inc-gpc1"):
			err = f.ParseAction(&actions.ScIncGpc1{}, parts)
		case strings.HasPrefix(parts[2], "sc-set-gpt0"):
			err = f.ParseAction(&actions.ScSetGpt0{}, parts)
		case strings.HasPrefix(parts[2], "set-var("):
			err = f.ParseAction(&actions.SetVar{}, parts)
		case strings.HasPrefix(parts[2], "set-var-fmt"):
			err = f.ParseAction(&actions.SetVarFmt{}, parts)
		case strings.HasPrefix(parts[2], "unset-var"):
			err = f.ParseAction(&actions.UnsetVar{}, parts)
		}
	}
	return err
}

func (f *Session) String() string {
	var result strings.Builder

	result.WriteString("session")
	result.WriteString(" ")
	result.WriteString(f.Action.String())

	if f.Comment != "" {
		result.WriteString(" # ")
		result.WriteString(f.Comment)
	}

	return result.String()
}

func (f *Session) GetComment() string {
	return f.Comment
}
