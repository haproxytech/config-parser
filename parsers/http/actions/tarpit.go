package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type Tarpit struct {
	DenyStatus string
	Cond       string
	CondTest   string
	Comment    string
}

func (f *Tarpit) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 4 {
		command, condition := common.SplitRequest(parts[2:])
		if len(command) > 1 && command[0] == "deny_status" {
			f.DenyStatus = command[1]
		}
		if len(condition) > 1 {
			f.Cond = condition[0]
			f.CondTest = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *Tarpit) String() string {
	var result strings.Builder
	result.WriteString("tarpit")
	if f.DenyStatus != "" {
		result.WriteString(" deny_status ")
		result.WriteString(f.DenyStatus)
	}
	if f.Cond != "" {
		result.WriteString(" ")
		result.WriteString(f.Cond)
		result.WriteString(" ")
		result.WriteString(f.CondTest)
	}
	return result.String()
}

func (f *Tarpit) GetComment() string {
	return f.Comment
}
