package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type SetVar struct {
	VarScope      string
	VarName       string
	Expr          common.Expression
	Cond          string
	ConditionTest string
	Comment       string
}

func (f *SetVar) Parse(parts []string, comment string) error {
	//we have filter trace [name <name>] [random-parsing] [random-forwarding] [hexdump]
	if comment != "" {
		f.Comment = comment
	}
	data := strings.TrimPrefix(parts[1], "set-var(")
	data = strings.TrimRight(data, ")")
	d := strings.SplitN(data, ".", 2)
	f.VarScope = d[0]
	f.VarName = d[1]
	if len(parts) >= 3 {
		command, condition := common.SplitRequest(parts[2:]) // 2 not 3 !
		if len(command) > 0 {
			expr := common.Expression{}
			err := expr.Parse(command)
			if err != nil {
				return fmt.Errorf("not enough params")
			}
			f.Expr = expr
		}
		if len(condition) > 1 {
			f.Cond = condition[0]
			f.ConditionTest = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *SetVar) String() string {
	condition := ""
	if f.Cond != "" {
		condition = fmt.Sprintf(" %s %s", f.Cond, f.ConditionTest)
	}
	return fmt.Sprintf("set-var(%s.%s) %s%s", f.VarScope, f.VarName, f.Expr.String(), condition)
}

func (f *SetVar) GetComment() string {
	return f.Comment
}
