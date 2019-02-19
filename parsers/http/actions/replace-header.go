package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type ReplaceHeader struct {
	Name       string
	MatchRegex string
	ReplaceFmt string
	Cond       string
	CondTest   string
	Comment    string
}

func (f *ReplaceHeader) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 4 {
		command, condition := common.SplitRequest(parts[2:])
		if len(command) < 3 {
			return errors.InvalidData
		}
		f.Name = command[0]
		f.MatchRegex = command[1]
		f.ReplaceFmt = command[2]
		if len(condition) > 1 {
			f.Cond = condition[0]
			f.CondTest = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return errors.InvalidData
}

func (f *ReplaceHeader) String() string {
	condition := ""
	if f.Cond != "" {
		condition = fmt.Sprintf(" %s %s", f.Cond, f.CondTest)
	}
	return fmt.Sprintf("replace-header %s %s %s%s", f.Name, f.MatchRegex, f.ReplaceFmt, condition)
}

func (f *ReplaceHeader) GetComment() string {
	return f.Comment
}
