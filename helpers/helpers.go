package helpers

import "strings"

func AddToBoolMap(data map[string]bool, items ...string) {
	for _, item := range items {
		data[item] = true
	}
}

//StringSplitIgnoreEmpty while spliting, removes empty items
func StringSplitIgnoreEmpty(s string, sep rune) []string {
	f := func(c rune) bool {
		return c == sep
	}
	return strings.FieldsFunc(s, f)
}
