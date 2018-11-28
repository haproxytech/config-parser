package global

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/helpers"
)

type CpuMap struct {
	Name  string
	Value string
}

type CpuMapLines struct {
	CpuMapLines []*CpuMap
}

func (c *CpuMapLines) Init() {
	c.CpuMapLines = []*CpuMap{}
}

func (c *CpuMapLines) GetParserName() string {
	return "cpu-map"
}

func (c *CpuMapLines) parseCpuMapLine(line string) (*CpuMap, error) {

	elements := helpers.StringSplitIgnoreEmpty(line, ' ')
	if len(elements) < 3 {
		return nil, &errors.ParseError{Parser: "CpuMapSingle", Line: line, Message: "Parse error"}
	}
	cpuMap := &CpuMap{
		Name:  elements[1],
		Value: elements[2],
	}
	return cpuMap, nil
}

func (c *CpuMapLines) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, "cpu-map") {
		if nameserver, err := c.parseCpuMapLine(line); err == nil {
			c.CpuMapLines = append(c.CpuMapLines, nameserver)
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "CpuMapLines", Line: line}
}

func (c *CpuMapLines) Valid() bool {
	if len(c.CpuMapLines) > 0 {
		return true
	}
	return false
}

func (c *CpuMapLines) String() []string {
	result := make([]string, len(c.CpuMapLines))
	for index, cpuMap := range c.CpuMapLines {
		result[index] = fmt.Sprintf("  cpu-map %s %s", cpuMap.Name, cpuMap.Value)
	}
	return result
}

func (c *CpuMapLines) Equal(b *CpuMapLines) bool {
	if b == nil {
		return false
	}
	if b.CpuMapLines == nil {
		return false
	}
	if len(c.CpuMapLines) != len(b.CpuMapLines) {
		return false
	}
	for _, cCpuMap := range c.CpuMapLines {
		found := false
		for _, bCpuMap := range b.CpuMapLines {
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
