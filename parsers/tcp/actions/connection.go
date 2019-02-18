package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type Connection struct {
	Action   []string
	Cond     string
	CondTest string
	Comment  string
}

func (f *Connection) Parse(parts []string, comment string) error {
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

func (f *Connection) String() string {
	var result strings.Builder
	result.WriteString("connection ")

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

func (f *Connection) GetComment() string {
	return f.Comment
}
