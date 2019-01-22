package filters

import (
	"strings"

	"github.com/haproxytech/config-parser/common"
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

func (f *Compression) Result() common.ReturnResultLine {
	var result strings.Builder
	result.WriteString("filter compression")
	return common.ReturnResultLine{
		Data:    result.String(),
		Comment: f.Comment,
	}
}
