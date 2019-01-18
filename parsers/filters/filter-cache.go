package filters

import (
	"fmt"
	"strings"
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

func (f *Cache) String() string {
	var result strings.Builder
	result.WriteString("  filter cache")
	if f.Name != "" {
		result.WriteString(" ")
		result.WriteString(f.Name)
	}
	if f.Comment != "" {
		result.WriteString(fmt.Sprintf(" # %s", f.Comment))
	}
	return result.String()
}
