package global

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type CpuMapLines struct {
	data []types.CpuMap
}

func (c *CpuMapLines) Init() {
	c.data = []types.CpuMap{}
}

func (c *CpuMapLines) GetParserName() string {
	return "cpu-map"
}

func (c *CpuMapLines) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(c.data) == 0 {
		return nil, errors.FetchError
	}
	return c.data, nil
}

func (c *CpuMapLines) Set(data common.ParserData) error {
	if data == nil {
		c.Init()
		return nil
	}
	switch newValue := data.(type) {
	case []types.CpuMap:
		c.data = newValue
	case *types.CpuMap:
		c.data = append(c.data, *newValue)
	case types.CpuMap:
		c.data = append(c.data, newValue)
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (c *CpuMapLines) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := c.Get(false)
	c.Init()
	_, err := c.Parse(data, parts, []string{}, comment)
	if err != nil {
		c.Set(oldData)
	}
	return err
}

func (c *CpuMapLines) parseCpuMapLine(parts []string, comment string) (*types.CpuMap, error) {

	if len(parts) < 3 {
		return nil, &errors.ParseError{Parser: "CpuMapSingle", Line: strings.Join(parts, " "), Message: "Parse error"}
	}
	cpuMap := &types.CpuMap{
		Name:    parts[1],
		Value:   parts[2],
		Comment: comment,
	}
	return cpuMap, nil
}

func (c *CpuMapLines) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "cpu-map" {
		if item, err := c.parseCpuMapLine(parts, comment); err == nil {
			c.data = append(c.data, *item)
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "CpuMapLines", Line: line}
}

func (c *CpuMapLines) Result(AddComments bool) ([]common.ReturnResultLine, error) {
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

func (c *CpuMapLines) Equal(b *CpuMapLines) bool {
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
