package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type Deny struct {
	DenyStatus    string
	Condition     string
	ConditionKind string
	Comment       string
}

func (f *Deny) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 4 {
		command, condition := common.SplitRequest(parts[2:])
		if len(command) > 1 && command[0] == "deny_status" {
			f.DenyStatus = command[1]
		}
		if len(condition) > 0 {
			f.ConditionKind = condition[0]
			f.Condition = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *Deny) String() string {
	var result strings.Builder
	result.WriteString("deny")
	if f.DenyStatus != "" {
		result.WriteString(" deny_status ")
		result.WriteString(f.DenyStatus)
	}
	if f.Condition != "" {
		result.WriteString(" ")
		result.WriteString(f.ConditionKind)
		result.WriteString(" ")
		result.WriteString(f.Condition)
	}
	return result.String()
}

func (f *Deny) GetComment() string {
	return f.Comment
}
