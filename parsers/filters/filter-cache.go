package filters

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type Cache struct {
	Name    string
	Comment string
}

func (f *Cache) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 2 {
		f.Name = parts[2]
	} else {
		return fmt.Errorf("no cache name")
	}
	return nil
}

func (f *Cache) Result() common.ReturnResultLine {
	var result strings.Builder
	result.WriteString("filter cache")
	if f.Name != "" {
		result.WriteString(" ")
		result.WriteString(f.Name)
	}
	return common.ReturnResultLine{
		Data:    result.String(),
		Comment: f.Comment,
	}
}
