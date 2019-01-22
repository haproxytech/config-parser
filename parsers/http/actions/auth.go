package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
)

type Auth struct {
	Realm         string
	Condition     string
	ConditionKind string
	Comment       string
}

func (f *Auth) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 4 {
		command, condition := common.SplitRequest(parts[2:])
		if len(command) > 1 && command[0] == "realm" {
			f.Realm = command[1]
		}
		if len(condition) > 0 {
			f.ConditionKind = condition[0]
			f.Condition = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *Auth) String() string {
	var result strings.Builder
	result.WriteString("auth")
	if f.Realm != "" {
		result.WriteString(" realm ")
		result.WriteString(f.Realm)
	}
	if f.Condition != "" {
		result.WriteString(" ")
		result.WriteString(f.ConditionKind)
		result.WriteString(" ")
		result.WriteString(f.Condition)
	}
	return result.String()
}

func (f *Auth) GetComment() string {
	return f.Comment
}
