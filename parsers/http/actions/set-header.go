package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type SetHeader struct {
	Name          string
	Fmt           string
	Condition     string
	ConditionKind string
	Comment       string
}

func (f *SetHeader) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 4 {
		command, condition := common.SplitRequest(parts[3:])
		f.Name = parts[2]
		f.Fmt = strings.Join(command, " ")
		if len(condition) > 0 {
			f.ConditionKind = condition[0]
			f.Condition = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *SetHeader) String() string {
	condition := ""
	if f.Condition != "" {
		condition = fmt.Sprintf(" %s %s", f.ConditionKind, f.Condition)
	}
	return fmt.Sprintf("set-header %s %s%s", f.Name, f.Fmt, condition)
}
