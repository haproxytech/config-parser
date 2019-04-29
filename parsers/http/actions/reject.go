package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type Reject struct {
	Cond     string
	CondTest string
	Comment  string
}

func (f *Reject) Parse(parts []string, comment string) error {
	//we have filter trace [name <name>] [random-parsing] [random-forwarding] [hexdump]
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 4 {
		_, condition := common.SplitRequest(parts[2:]) // 2 not 3 !
		if len(condition) > 1 {
			f.Cond = condition[0]
			f.CondTest = strings.Join(condition[1:], " ")
		}
		return nil
	} else if len(parts) == 2 {
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *Reject) String() string {
	condition := ""
	if f.Cond != "" {
		condition = fmt.Sprintf(" %s %s", f.Cond, f.CondTest)
	}
	return fmt.Sprintf("reject%s", condition)
}

func (f *Reject) GetComment() string {
	return f.Comment
}
