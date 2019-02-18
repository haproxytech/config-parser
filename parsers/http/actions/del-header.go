package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type DelHeader struct {
	Name     string
	Cond     string
	CondTest string
	Comment  string
}

func (f *DelHeader) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 4 {
		_, condition := common.SplitRequest(parts[3:])
		f.Name = parts[2]
		if len(condition) > 1 {
			f.Cond = condition[0]
			f.CondTest = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *DelHeader) String() string {
	condition := ""
	if f.Cond != "" {
		condition = fmt.Sprintf(" %s %s", f.Cond, f.CondTest)
	}
	return fmt.Sprintf("del-header %s%s", f.Name, condition)
}

func (f *DelHeader) GetComment() string {
	return f.Comment
}
