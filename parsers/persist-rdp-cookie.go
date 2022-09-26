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
package parsers

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/types"
)

type PersistRdpCookie struct {
	data        *types.PersistRdpCookie
	preComments []string // comments that appear before the the actual line
}

func (f *PersistRdpCookie) Parse(line string, parts []string, comment string) (changeState string, err error) {
	if len(parts) != 2 {
		return "", &errors.ParseError{Parser: "PersistRDPCookie", Line: line}
	}
	if parts[0] != "persist" {
		return "", &errors.ParseError{Parser: "PersistRDPCookie", Line: line}
	}
	if parts[1] == "rdp-cookie" {
		f.data = &types.PersistRdpCookie{Name: "", Comment: comment}
		return "", nil
	}
	var data string
	data = parts[1]
	data = strings.TrimPrefix(data, "rdp-cookie(")
	data = strings.TrimRight(data, ")")
	f.data = &types.PersistRdpCookie{Name: data, Comment: comment}
	return "", nil
}

func (f *PersistRdpCookie) Result() ([]common.ReturnResultLine, error) {
	if f.data == nil {
		return nil, errors.ErrFetch
	}
	line := "persist rdp-cookie"
	if f.data.Name != "" {
		line = fmt.Sprintf("%s(%s)", line, f.data.Name)
	}
	return []common.ReturnResultLine{
		{
			Data:    line,
			Comment: f.data.Comment,
		},
	}, nil
}
