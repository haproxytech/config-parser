package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type Session struct {
	Action   []string
	Cond     string
	CondTest string
	Comment  string
}

func (f *Session) Parse(parts []string, comment string) error {
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
		if len(condition) > 1 {
			f.Cond = condition[0]
			f.CondTest = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *Session) String() string {
	var result strings.Builder
	result.WriteString("content ")

	result.WriteString(strings.Join(f.Action, " "))

	if f.Cond != "" {
		result.WriteString(" ")
		result.WriteString(f.Cond)
		result.WriteString(" ")
		result.WriteString(f.CondTest)
	}
	if f.Comment != "" {
		result.WriteString(" # ")
		result.WriteString(f.Comment)
	}
	return result.String()
}

func (f *Session) GetComment() string {
	return f.Comment
}
