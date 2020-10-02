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
	"github.com/haproxytech/config-parser/v3/common"
	"github.com/haproxytech/config-parser/v3/errors"
	"github.com/haproxytech/config-parser/v3/types"
)

type SslModeAsync struct {
	data        *types.SslModeAsync
	preComments []string
}

func (p *SslModeAsync) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) == 1 && parts[0] == "ssl-mode-async" {
		p.data = &types.SslModeAsync{}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "ssl-mode-async", Line: line}
}

func (p *SslModeAsync) Result() (res []common.ReturnResultLine, err error) {
	if p.data == nil {
		return nil, errors.ErrFetch
	}
	return []common.ReturnResultLine{
		{
			Data: "ssl-mode-async",
		},
	}, nil
}
