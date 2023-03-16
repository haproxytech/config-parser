/*
Copyright 2023 HAProxy Technologies

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

package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/types"
)

type ScAddGpc struct {
	Idx      string
	ID       string
	Integer  string
	Cond     string
	CondTest string
	Comment  string
}

func (f *ScAddGpc) Parse(parts []string, parserType types.ParserType, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	var data string
	var command []string
	var minLen, requiredLen int
	switch parserType {
	case types.HTTP:
		data = parts[1]
		f.Integer = parts[2]
		command = parts[3:]
		minLen = 3
		requiredLen = 5
	case types.TCP:
		data = parts[2]
		f.Integer = parts[3]
		command = parts[4:]
		minLen = 4
		requiredLen = 6
	}
	idIdx := strings.TrimPrefix(data, "sc-add-gpc(")
	idIdx = strings.TrimRight(idIdx, ")")
	idIdxValues := strings.SplitN(idIdx, ",", 2)
	f.Idx, f.ID = idIdxValues[0], idIdxValues[1]
	if len(parts) == minLen {
		return nil
	}
	if len(parts) < requiredLen {
		return fmt.Errorf("not enough params")
	}
	if len(command) == 0 {
		return fmt.Errorf("no command found")
	}
	_, condition := common.SplitRequest(command)
	if len(condition) > 1 {
		f.Cond = condition[0]
		f.CondTest = strings.Join(condition[1:], " ")
	}
	return nil
}

func (f *ScAddGpc) String() string {
	var result strings.Builder
	result.WriteString("sc-add-gpc(")
	result.WriteString(f.Idx)
	result.WriteString(",")
	result.WriteString(f.ID)
	result.WriteString(")")
	result.WriteString(" ")
	result.WriteString(f.Integer)
	if f.Cond != "" {
		result.WriteString(" ")
		result.WriteString(f.Cond)
		result.WriteString(" ")
		result.WriteString(f.CondTest)
	}
	return result.String()
}

func (f *ScAddGpc) GetComment() string {
	return f.Comment
}
