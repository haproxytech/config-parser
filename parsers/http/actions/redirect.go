package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type Redirect struct { // http-request redirect location <loc> [code <code>] [<option>] [<condition>]
	Type     string
	Value    string
	Code     string
	Option   string
	Cond     string
	CondTest string
	Comment  string
}

func (f *Redirect) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	/*
	  redirect location <loc> [code <code>] <option> [{if | unless} <condition>]
	  redirect prefix   <pfx> [code <code>] <option> [{if | unless} <condition>]
	  redirect scheme   <sch> [code <code>] <option> [{if | unless} <condition>]
	*/
	if len(parts) >= 4 {
		command, condition := common.SplitRequest(parts[2:])
		if len(command) < 2 {
			return errors.InvalidData
		}
		f.Type = command[0]
		f.Value = command[1]
		index := 2
		if index < len(command) && command[index] == "code" {
			index++
			if index == len(command) {
				return fmt.Errorf("not enough params")
			}
			f.Code = command[index]
			index++
		}
		if index < len(command) {
			f.Option = command[index]
		}
		if len(condition) > 1 {
			f.Cond = condition[0]
			f.CondTest = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *Redirect) String() string {
	var result strings.Builder
	result.WriteString("redirect ")
	result.WriteString(f.Type)
	result.WriteString(" ")
	result.WriteString(f.Value)
	if f.Code != "" {
		result.WriteString(" code ")
		result.WriteString(f.Code)
	}
	if f.Option != "" {
		result.WriteString(" ")
		result.WriteString(f.Option)
	}
	if f.Cond != "" {
		result.WriteString(" ")
		result.WriteString(f.Cond)
		result.WriteString(" ")
		result.WriteString(f.CondTest)
	}
	return result.String()
}

func (f *Redirect) GetComment() string {
	return f.Comment
}
