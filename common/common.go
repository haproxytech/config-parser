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

package common

import "strings"

func AddToBoolMap(data map[string]bool, items ...string) {
	for _, item := range items {
		data[item] = true
	}
}

//StringSplitIgnoreEmpty while spliting, removes empty items
func StringSplitIgnoreEmpty(s string, separators ...rune) []string {
	f := func(c rune) bool {
		willSplit := false
		for _, sep := range separators {
			if c == sep {
				willSplit = true
				break
			}
		}
		return willSplit
	}
	return strings.FieldsFunc(s, f)
}

//StringSplitWithCommentIgnoreEmpty while spliting, removes empty items, if we have comment, separate it
func StringSplitWithCommentIgnoreEmpty(s string, separators ...rune) (data []string, comment string) {
	tmp := strings.SplitN(s, "#", 2)
	comment = ""
	if len(tmp) > 1 {
		comment = strings.TrimSpace(tmp[1])
	}
	f := func(c rune) bool {
		willSplit := false
		for _, sep := range separators {
			if c == sep {
				willSplit = true
				break
			}
		}
		return willSplit
	}
	return strings.FieldsFunc(tmp[0], f), comment
}

//StringExtractComment checks if comment is added
func StringExtractComment(s string) string {
	p := StringSplitIgnoreEmpty(s, '#')
	if len(p) > 1 {
		return p[len(p)-1]
	}
	return ""
}

//searches for "if" or "unless" and returns result
func SplitRequest(parts []string) (command, condition []string) {
	if len(parts) == 0 {
		return []string{}, []string{}
	}
	index := 0
	found := false
	for index < len(parts) {
		switch parts[index] {
		case "if", "unless":
			found = true
		}
		if found {
			break
		}
		index++
	}
	command = parts[:index]
	condition = parts[index:]
	return command, condition
}
