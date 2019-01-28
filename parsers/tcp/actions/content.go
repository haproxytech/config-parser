package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type Content struct {
	Action        []string
	Condition     string
	ConditionKind string
	Comment       string
}

func (f *Content) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 3 {
		command, condition := common.SplitRequest(parts[2:])
		if len(command) > 0 {
			f.Action = command
		} else {
			return fmt.Errorf("not enough params")
		}
		if len(condition) > 0 {
			f.ConditionKind = condition[0]
			f.Condition = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *Content) String() string {
	var result strings.Builder
	result.WriteString("content ")

	result.WriteString(strings.Join(f.Action, " "))

	if f.Condition != "" {
		result.WriteString(" ")
		result.WriteString(f.ConditionKind)
		result.WriteString(" ")
		result.WriteString(f.Condition)
	}
	if f.Comment != "" {
		result.WriteString(" # ")
		result.WriteString(f.Comment)
	}
	return result.String()
}

func (f *Content) GetComment() string {
	return f.Comment
}
