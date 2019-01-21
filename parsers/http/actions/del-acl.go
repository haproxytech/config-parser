package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type DelAcl struct {
	FileName      string
	KeyFmt        string
	Condition     string
	ConditionKind string
	Comment       string
}

func (f *DelAcl) Parse(parts []string, comment string) error {
	//we have filter trace [name <name>] [random-parsing] [random-forwarding] [hexdump]
	if comment != "" {
		f.Comment = comment
	}
	f.FileName = strings.TrimPrefix(parts[1], "del-acl(")
	f.FileName = strings.TrimRight(f.FileName, ")")
	if len(parts) >= 4 {
		command, condition := common.SplitRequest(parts[2:]) // 2 not 3 !
		if len(command) > 0 {
			f.KeyFmt = command[0]
		}
		if len(condition) > 0 {
			f.ConditionKind = condition[0]
			f.Condition = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *DelAcl) String() string {
	keyfmt := ""
	condition := ""
	comment := ""
	if f.KeyFmt != "" {
		keyfmt = " " + f.KeyFmt
	}
	if f.Condition != "" {
		condition = fmt.Sprintf(" %s %s", f.ConditionKind, f.Condition)
	}
	if f.Comment != "" {
		comment = " # " + f.Comment
	}
	return fmt.Sprintf("del-acl(%s)%s%s%s", f.FileName, keyfmt, condition, comment)
}
