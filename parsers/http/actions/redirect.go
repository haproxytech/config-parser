package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type Redirect struct {
	Header        string
	Fmt           string
	Condition     string
	ConditionKind string
	Comment       string
}

func (f *Redirect) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 4 {
		command, condition := common.SplitRequest(parts[3:])
		f.Header = parts[2]
		f.Fmt = strings.Join(command, " ")
		if len(condition) > 0 {
			f.ConditionKind = condition[0]
			f.Condition = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *Redirect) String() string {
	var result strings.Builder
	result.WriteString("redirect ")
	result.WriteString(f.Header)
	result.WriteString(" ")
	result.WriteString(f.Fmt)
	if f.Condition != "" {
		result.WriteString(" ")
		result.WriteString(f.ConditionKind)
		result.WriteString(" ")
		result.WriteString(f.Condition)
	}
	return result.String()
}

func (f *Redirect) GetComment() string {
	return f.Comment
}
