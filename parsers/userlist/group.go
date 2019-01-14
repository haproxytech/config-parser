package userlist

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/helpers"
)

type Group struct {
	Name  string
	Users []string
}

type GroupLines struct {
	GroupLines []Group
}

func (l *GroupLines) Init() {
	l.GroupLines = []Group{}
}

func (l *GroupLines) GetParserName() string {
	return "group"
}

func (l *GroupLines) parseGroupLine(line string, parts []string) (Group, error) {
	if len(parts) >= 2 {
		group := Group{
			Name: parts[1],
		}
		if len(parts) > 3 && parts[2] == "users" {
			group.Users = helpers.StringSplitIgnoreEmpty(parts[3], ',')
		}
		return group, nil
	}
	return Group{}, &errors.ParseError{Parser: "GroupLines", Line: line}
}

func (l *GroupLines) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if parts[0] == "group" {
		group, err := l.parseGroupLine(line, parts)
		if err != nil {
			return "", &errors.ParseError{Parser: "GroupLines", Line: line}
		}
		l.GroupLines = append(l.GroupLines, group)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "GroupLines", Line: line}
}

func (l *GroupLines) Valid() bool {
	if len(l.GroupLines) > 0 {
		return true
	}
	return false
}

func (l *GroupLines) Result(AddComments bool) []string {
	result := make([]string, len(l.GroupLines))
	for index, group := range l.GroupLines {
		users := ""
		if len(group.Users) > 0 {
			var s strings.Builder
			s.WriteString(" users ")
			first := true
			for _, user := range group.Users {
				if !first {
					s.WriteString(",")
				} else {
					first = false
				}
				s.WriteString(user)
			}
			users = s.String()
		}
		result[index] = fmt.Sprintf("  group %s%s", group.Name, users)
	}
	return result
}
