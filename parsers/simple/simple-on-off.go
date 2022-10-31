/*
Copyright 2022 HAProxy Technologies

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

package simple

import (
	"fmt"

	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/types"
)

type OnOff struct {
	Name        string
	data        *types.StringC
	preComments []string // comments that appear before the actual line
}

func (s *OnOff) Parse(line string, parts []string, comment string) (string, error) {
	if parts[0] == s.Name {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "OnOff", Line: line, Message: "Parse error"}
		}
		switch parts[1] {
		case "on":
			s.data = &types.StringC{
				Value:   "on",
				Comment: comment,
			}
		case "off":
			s.data = &types.StringC{
				Value:   "off",
				Comment: comment,
			}
		default:
			return "", &errors.ParseError{Parser: "OnOff", Line: line, Message: fmt.Sprintf("Option %s not supported", parts[1])}
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.Name, Line: line}
}

func (s *OnOff) Result() ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, errors.ErrFetch
	}
	return []common.ReturnResultLine{
		{
			Data:    fmt.Sprintf("%s %s", s.Name, s.data.Value),
			Comment: s.data.Comment,
		},
	}, nil
}
