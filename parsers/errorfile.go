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

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type ErrorFile struct {
	data        []types.ErrorFile
	allowedCode map[string]bool
}

func (l *ErrorFile) Init() {
	l.data = []types.ErrorFile{}
	l.allowedCode = map[string]bool{}
	common.AddToBoolMap(l.allowedCode,
		"200", "400", "403", "405", "408", "425", "429",
		"500", "502", "503", "504")
}

func (l *ErrorFile) parse(line string, parts []string, comment string) (*types.ErrorFile, error) {
	if len(parts) < 3 {
		return nil, &errors.ParseError{Parser: "ErrorFile", Line: line}
	}
	errorfile := &types.ErrorFile{
		File:    parts[2],
		Comment: comment,
	}
	code := parts[1]
	if _, ok := l.allowedCode[code]; !ok {
		return errorfile, nil
	}
	errorfile.Code = code
	return errorfile, nil
}

func (l *ErrorFile) Result() ([]common.ReturnResultLine, error) {
	if len(l.data) == 0 {
		return nil, errors.ErrFetch
	}
	result := make([]common.ReturnResultLine, len(l.data))
	for index, data := range l.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("errorfile %s %s", data.Code, data.File),
			Comment: data.Comment,
		}
	}
	return result, nil
}
