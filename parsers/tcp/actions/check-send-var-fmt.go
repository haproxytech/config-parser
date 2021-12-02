/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/types"
)

// tcp-check set-var-fmt(<var-name>) <fmt>
type CheckSetVarFmt struct {
	VarScope string
	VarName  string
	Format   string
	Cond     string
	CondTest string
	Comment  string
}

func (c *CheckSetVarFmt) Parse(parts []string, parserType types.ParserType, comment string) error {
	if comment != "" {
		c.Comment = comment
	}
	if len(parts) < 3 {
		return fmt.Errorf("not enough params")
	}
	var data string
	var format []string
	data = parts[1]
	format = parts[2:]
	data = strings.TrimPrefix(data, "set-var-fmt(")
	data = strings.TrimRight(data, ")")
	d := strings.SplitN(data, ".", 2)
	c.VarScope = d[0]
	c.VarName = d[1]
	format, condition := common.SplitRequest(format)
	c.Format = strings.Join(format[0:], "")
	if len(condition) > 1 {
		c.Cond = condition[0]
		c.CondTest = strings.Join(condition[1:], " ")
	}
	return nil
}

func (c *CheckSetVarFmt) String() string {
	condition := ""
	if c.Cond != "" {
		condition = fmt.Sprintf(" %s %s", c.Cond, c.CondTest)
	}
	return fmt.Sprintf("set-var-fmt(%s.%s) %s%s", c.VarScope, c.VarName, c.Format, condition)
}

func (c *CheckSetVarFmt) GetComment() string {
	return c.Comment
}
