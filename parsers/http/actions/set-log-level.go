package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type SetLogLevel struct {
	Level    string
	Cond     string
	CondTest string
	Comment  string
}

func (f *SetLogLevel) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 4 {
		command, condition := common.SplitRequest(parts[2:])
		if len(command) == 0 {
			return errors.InvalidData
		}
		f.Level = command[0]
		if len(condition) > 1 {
			f.Cond = condition[0]
			f.CondTest = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *SetLogLevel) String() string {
	condition := ""
	if f.Cond != "" {
		condition = fmt.Sprintf(" %s %s", f.Cond, f.CondTest)
	}
	return fmt.Sprintf("set-log-level %s%s", f.Level, condition)
}

func (f *SetLogLevel) GetComment() string {
	return f.Comment
}
