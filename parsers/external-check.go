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
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type ExternalCheck struct {
	data *types.Enabled
}

func (m *ExternalCheck) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "external-check" {
		m.data = &types.Enabled{
			Comment: comment,
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "ExternalCheck", Line: line}
}

func (m *ExternalCheck) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if m.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    "external-check",
			Comment: m.data.Comment,
		},
	}, nil
}
