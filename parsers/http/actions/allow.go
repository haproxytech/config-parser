package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type Allow struct {
	Condition     string
	ConditionKind string
	Comment       string
}

func (f *Allow) Parse(parts []string, comment string) error {
	//we have filter trace [name <name>] [random-parsing] [random-forwarding] [hexdump]
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 4 {
		_, condition := common.SplitRequest(parts[2:]) // 2 not 3 !
		if len(condition) > 0 {
			f.ConditionKind = condition[0]
			f.Condition = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *Allow) String() string {
	condition := ""
	if f.Condition != "" {
		condition = fmt.Sprintf(" %s %s", f.ConditionKind, f.Condition)
	}
	return fmt.Sprintf("allow%s", condition)
}

func (f *Allow) GetComment() string {
	return f.Comment
}
