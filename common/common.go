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
			break
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
