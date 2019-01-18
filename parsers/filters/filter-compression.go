package filters

import (
	"fmt"
	"strings"
)

type Compression struct {
	Enabled bool
	Comment string
}

func (f *Compression) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	f.Enabled = true
	return nil
}

func (f *Compression) String() string {
	var result strings.Builder
	result.WriteString("  filter compression")
	if f.Comment != "" {
		result.WriteString(fmt.Sprintf(" # %s", f.Comment))
	}
	return result.String()
}
