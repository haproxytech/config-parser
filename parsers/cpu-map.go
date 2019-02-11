package parsers

import (
	"fmt"

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
		Name:    parts[1],
		Value:   parts[2],
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
			Data:    fmt.Sprintf("cpu-map %s %s", cpuMap.Name, cpuMap.Value),
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
			if cCpuMap.Name == bCpuMap.Name {
				if cCpuMap.Value != bCpuMap.Value {
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
