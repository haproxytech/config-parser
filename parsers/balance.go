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

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Balance struct {
	data *types.Balance
}

func (p *Balance) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "balance" {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "Balance", Line: line, Message: "Parse error"}
		}
		data := &types.Balance{
			Arguments: []string{},
			Comment:   comment,
		}

		switch parts[1] {
		case "roundrobin", "static-rr", "leastconn", "first", "source", "random":
			data.Algorithm = parts[1]
			p.data = data
			return "", nil
		case "uri", "url_param":
			p.data = data
			data.Algorithm = parts[1]
			if len(parts) > 2 {
				p.data.Arguments = parts[2:]
				return "", nil
			}
			return "", nil
		}
		if strings.HasPrefix(parts[1], "hdr(") && strings.HasSuffix(parts[1], ")") {
			p.data = data
			data.Algorithm = parts[1]
			return "", nil
		}
		if strings.HasPrefix(parts[1], "rdp-cookie(") && strings.HasSuffix(parts[1], ")") {
			p.data = data
			data.Algorithm = parts[1]
			return "", nil
		}
		return "", &errors.ParseError{Parser: "Balance", Line: line}
	}
	return "", &errors.ParseError{Parser: "Balance", Line: line}
}

func (p *Balance) Result() ([]common.ReturnResultLine, error) {
	if p.data == nil {
		return nil, errors.ErrFetch
	}
	params := ""
	if len(p.data.Arguments) > 0 {
		params = fmt.Sprintf(" %s", strings.Join(p.data.Arguments, " "))
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data: fmt.Sprintf("balance %s%s", p.data.Algorithm, params),
		},
	}, nil
}
