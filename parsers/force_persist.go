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
	"strings"

	"github.com/deyunluo/config-parser/v4/common"
	"github.com/deyunluo/config-parser/v4/errors"
	"github.com/deyunluo/config-parser/v4/types"
)

type ForcePersist struct {
	Name        string
	data        []types.ForcePersist
	Mode        string
	preComments []string // comments that appear before the the actual line
}

func (h *ForcePersist) parse(line string, parts []string, comment string) (*types.ForcePersist, error) {
	if len(parts) >= 3 {
		data := &types.ForcePersist{
			Name:      parts[1],
			Criterion: parts[2],
			Value:     strings.Join(parts[0:], " "),
			Comment:   comment,
		}
		return data, nil
	}
	return nil, &errors.ParseError{Parser: "ForcePersist", Line: line}
}

func (h *ForcePersist) Result() ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.ErrFetch
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		var sb strings.Builder
		sb.WriteString(req.Value)
		result[index] = common.ReturnResultLine{
			Data:    sb.String(),
			Comment: req.Comment,
		}
	}
	return result, nil
}
