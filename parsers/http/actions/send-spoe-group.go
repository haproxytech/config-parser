package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type SendSpoeGroup struct {
	Engine   string
	Group    string
	Cond     string
	CondTest string
	Comment  string
}

func (f *SendSpoeGroup) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 4 {
		command, condition := common.SplitRequest(parts[2:])
		if len(command) < 2 {
			return errors.InvalidData
		}
		f.Engine = command[0]
		f.Group = command[1]
		if len(condition) > 1 {
			f.Cond = condition[0]
			f.CondTest = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *SendSpoeGroup) String() string {
	condition := ""
	if f.Cond != "" {
		condition = fmt.Sprintf(" %s %s", f.Cond, f.CondTest)
	}
	return fmt.Sprintf("send-spoe-group %s %s%s", f.Engine, f.Group, condition)
}

func (f *SendSpoeGroup) GetComment() string {
	return f.Comment
}
