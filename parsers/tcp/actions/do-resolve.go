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
)

type DoResolve struct {
	Var       string
	Resolvers string
	Protocol  string
	Expr      common.Expression
}

func (f *DoResolve) Parse(parts []string) error {
	if len(parts) < 2 {
		return fmt.Errorf("not enough params")
	}
	data := strings.TrimPrefix(parts[0], "do-resolve(")
	data = strings.TrimRight(data, ")")
	d := strings.Split(data, ",")
	if len(d) < 2 {
		return fmt.Errorf("not enough params")
	}
	f.Var = d[0]
	f.Resolvers = d[1]
	if len(d) > 2 {
		f.Protocol = d[2]
	}

	expr := common.Expression{}
	if expr.Parse(parts[1:]) != nil {
		return fmt.Errorf("not enough params")
	}
	f.Expr = expr
	return nil
}

func (f *DoResolve) String() string {
	if f.Protocol != "" {
		return fmt.Sprintf("do-resolve(%s,%s,%s) %s", f.Var, f.Resolvers, f.Protocol, f.Expr.String())
	}
	return fmt.Sprintf("do-resolve(%s,%s) %s", f.Var, f.Resolvers, f.Expr.String())
}
