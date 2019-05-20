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

type OptionHttpchk struct {
	data *types.OptionHttpchk
}

/*
option httpchk <uri>
option httpchk <method> <uri>
option httpchk <method> <uri> <version>
*/
func (s *OptionHttpchk) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 1 && parts[0] == "option" && parts[1] == "httpchk" {
		if len(parts) == 2 {
			s.data = &types.OptionHttpchk{
				Comment: comment,
			}
		}
		if len(parts) == 3 {
			s.data = &types.OptionHttpchk{
				Uri:     parts[2],
				Comment: comment,
			}
		} else if len(parts) == 4 {
			s.data = &types.OptionHttpchk{
				Method:  parts[2],
				Uri:     parts[3],
				Comment: comment,
			}
		} else if len(parts) == 5 {
			s.data = &types.OptionHttpchk{
				Method:  parts[2],
				Uri:     parts[3],
				Version: parts[4],
				Comment: comment,
			}
		} else {
			s.data = &types.OptionHttpchk{
				Method:  parts[2],
				Uri:     parts[3],
				Version: strings.Join(parts[4:], " "),
				Comment: comment,
			}
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "option httpchk", Line: line}
}

func (s *OptionHttpchk) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, errors.FetchError
	}
	var data string
	if s.data.Version != "" {
		data = fmt.Sprintf("option httpchk %s %s %s", s.data.Method, s.data.Uri, s.data.Version)
	} else if s.data.Method != "" {
		data = fmt.Sprintf("option httpchk %s %s", s.data.Method, s.data.Uri)
	} else if s.data.Uri != "" {
		data = fmt.Sprintf("option httpchk %s", s.data.Uri)
	} else {
		data = fmt.Sprintf("option httpchk")
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    data,
			Comment: s.data.Comment,
		},
	}, nil
}
