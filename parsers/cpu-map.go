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

type CpuMap struct {
	data []types.CpuMap
}

func (c *CpuMap) parse(line string, parts []string, comment string) (*types.CpuMap, error) {

	if len(parts) < 3 {
		return nil, &errors.ParseError{Parser: "CpuMap", Line: line, Message: "Parse error"}
	}
	cpuMap := &types.CpuMap{
		Process: parts[1],
		CpuSet:  strings.Join(parts[2:], " "),
		Comment: comment,
	}
	return cpuMap, nil
}

func (c *CpuMap) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(c.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(c.data))
	for index, cpuMap := range c.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("cpu-map %s %s", cpuMap.Process, cpuMap.CpuSet),
			Comment: cpuMap.Comment,
		}
	}
	return result, nil
}

func (c *CpuMap) Equal(b *CpuMap) bool {
	if b == nil {
		return false
	}
	if b.data == nil {
		return false
	}
	if len(c.data) != len(b.data) {
		return false
	}
	for _, cCpuMap := range c.data {
		found := false
		for _, bCpuMap := range b.data {
			if cCpuMap.Process == bCpuMap.Process {
				if cCpuMap.CpuSet != bCpuMap.CpuSet {
					return false
				}
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
