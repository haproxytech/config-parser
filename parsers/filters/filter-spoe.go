package filters

import (
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type Spoe struct { //filter spoe [engine <name>] config <file>
	Engine  string
	Config  string
	Comment string
}

func (f *Spoe) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	index := 2
	for index < len(parts) {
		switch parts[index] {
		case "engine":
			index++
			f.Engine = parts[index]
		case "config":
			index++
			f.Config = parts[index]
		}
		index++
	}
	if f.Config == "" {
		return errors.InvalidData
	}
	return nil
}

func (f *Spoe) Result() common.ReturnResultLine {
	var result strings.Builder
	result.WriteString("filter spoe")
	if f.Engine != "" {
		result.WriteString(" engine ")
		result.WriteString(f.Engine)
	}
	result.WriteString(" config ")
	result.WriteString(f.Config)
	return common.ReturnResultLine{
		Data:    result.String(),
		Comment: f.Comment,
	}
}
