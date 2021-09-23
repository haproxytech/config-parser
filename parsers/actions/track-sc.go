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

package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/types"
)

type TrackScT string

const (
	TrackSc0 TrackScT = "track-sc0"
	TrackSc1 TrackScT = "track-sc1"
	TrackSc2 TrackScT = "track-sc2"
)

type TrackSc struct {
	Type     TrackScT
	Key      string
	Table    string
	Comment  string
	Cond     string
	CondTest string
}

func (f *TrackSc) Parse(parts []string, parserType types.ParserType, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) < 3 {
		return fmt.Errorf("not enough params")
	}
	var data string
	var command []string
	var minLen, requiredLen int
	switch parserType {
	case types.HTTP:
		data = parts[1]
		command = parts[2:]
		minLen = 3
		requiredLen = 5
	case types.TCP:
		data = parts[2]
		command = parts[3:]
		minLen = 4
		requiredLen = 6
	}
	f.Type = TrackScT(data)
	if len(parts) == minLen {
		f.Key = parts[minLen-1]
		return nil
	}
	if len(parts) < requiredLen {
		return fmt.Errorf("not enough params")
	}
	command, condition := common.SplitRequest(command)
	if len(command) > 1 && command[1] == "table" {
		if len(command) < 3 {
			return fmt.Errorf("not enough params")
		}
		f.Key = command[0]
		f.Table = command[2]
	}
	if len(command) == 1 {
		f.Key = command[0]
	}
	if len(condition) > 1 {
		f.Cond = condition[0]
		f.CondTest = strings.Join(condition[1:], " ")
	}
	return nil
}

func (f *TrackSc) String() string {
	var result strings.Builder
	result.WriteString(string(f.Type))
	result.WriteString(" ")
	result.WriteString(f.Key)
	if f.Table != "" {
		result.WriteString(" table ")
		result.WriteString(f.Table)
	}
	if f.Cond != "" {
		result.WriteString(" ")
		result.WriteString(f.Cond)
		result.WriteString(" ")
		result.WriteString(f.CondTest)
	}
	return result.String()
}

func (f *TrackSc) GetComment() string {
	return f.Comment
}
