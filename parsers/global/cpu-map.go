package global

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type CpuMap struct {
	Name    string
	Value   string
	Comment string
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

func (c *CpuMapLines) parseCpuMapLine(parts []string, comment string) (*CpuMap, error) {

	if len(parts) < 3 {
		return nil, &errors.ParseError{Parser: "CpuMapSingle", Line: strings.Join(parts, " "), Message: "Parse error"}
	}
	cpuMap := &CpuMap{
		Name:    parts[1],
		Value:   parts[2],
		Comment: comment,
	}
	return cpuMap, nil
}

func (c *CpuMapLines) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "cpu-map" {
		if nameserver, err := c.parseCpuMapLine(parts, comment); err == nil {
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

func (c *CpuMapLines) Result(AddComments bool) []common.ReturnResultLine {
	result := make([]common.ReturnResultLine, len(c.CpuMapLines))
	for index, cpuMap := range c.CpuMapLines {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("cpu-map %s %s", cpuMap.Name, cpuMap.Value),
			Comment: cpuMap.Comment,
		}
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
