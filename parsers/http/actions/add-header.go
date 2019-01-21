package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type AddHeader struct {
	Name          string
	Fmt           string
	Condition     string
	ConditionKind string
	Comment       string
}

func (f *AddHeader) Parse(parts []string, comment string) error {
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

func (f *AddHeader) String() string {
	condition := ""
	if f.Condition != "" {
		condition = fmt.Sprintf(" %s %s", f.ConditionKind, f.Condition)
	}
	return fmt.Sprintf("add-header %s %s%s", f.Name, f.Fmt, condition)
}
