package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type AddAcl struct {
	FileName string
	KeyFmt   string
	Cond     string
	CondTest string
	Comment  string
}

func (f *AddAcl) Parse(parts []string, comment string) error {
	//we have filter trace [name <name>] [random-parsing] [random-forwarding] [hexdump]
	if comment != "" {
		f.Comment = comment
	}
	f.FileName = strings.TrimPrefix(parts[1], "add-acl(")
	f.FileName = strings.TrimRight(f.FileName, ")")
	if len(parts) >= 4 {
		command, condition := common.SplitRequest(parts[2:]) // 2 not 3 !
		if len(command) > 0 {
			f.KeyFmt = command[0]
		}
		if len(condition) > 1 {
			f.Cond = condition[0]
			f.CondTest = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *AddAcl) String() string {
	keyfmt := ""
	condition := ""
	if f.KeyFmt != "" {
		keyfmt = " " + f.KeyFmt
	}
	if f.Cond != "" {
		condition = fmt.Sprintf(" %s %s", f.Cond, f.CondTest)
	}
	return fmt.Sprintf("add-acl(%s)%s%s", f.FileName, keyfmt, condition)
}

func (f *AddAcl) GetComment() string {
	return f.Comment
}
