package userlist

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/helpers"
)

type User struct {
	Name       string
	Password   string
	IsInsecure bool
	Groups     []string
}

type UserLines struct {
	UserLines []User
}

func (l *UserLines) Init() {
	l.UserLines = []User{}
}

func (l *UserLines) GetParserName() string {
	return "user"
}

func (l *UserLines) parseUserLine(line string, parts []string) (User, error) {
	if len(parts) >= 2 {
		user := User{
			Name: parts[1],
		}
		//see if we have password
		index := 3
		if len(parts) > index {
			if parts[2] == "password" {
				user.Password = parts[3]
				index += 2
			}
			if parts[2] == "insecure-password" {
				user.Password = parts[3]
				user.IsInsecure = true
				index += 2
			}
		}
		if len(parts) > index {
			user.Groups = helpers.StringSplitIgnoreEmpty(parts[index], ',')
		}
		return user, nil
	}
	return User{}, &errors.ParseError{Parser: "UserLines", Line: line}
}

func (l *UserLines) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if parts[0] == "user" {
		user, err := l.parseUserLine(line, parts)
		if err != nil {
			return "", &errors.ParseError{Parser: "UserLines", Line: line}
		}
		l.UserLines = append(l.UserLines, user)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "UserLines", Line: line}
}

func (l *UserLines) Valid() bool {
	if len(l.UserLines) > 0 {
		return true
	}
	return false
}

func (l *UserLines) Result(AddComments bool) []string {
	result := make([]string, len(l.UserLines))
	for index, user := range l.UserLines {
		pwd := ""
		if user.Password != "" {
			if user.IsInsecure {
				pwd = fmt.Sprintf(" insecure-password %s", user.Password)
			} else {
				pwd = fmt.Sprintf(" password %s", user.Password)
			}
		}
		groups := ""
		if len(user.Groups) > 0 {
			var s strings.Builder
			s.WriteString(" groups ")
			first := true
			for _, user := range user.Groups {
				if !first {
					s.WriteString(",")
				} else {
					first = false
				}
				s.WriteString(user)
			}
			groups = s.String()
		}
		result[index] = fmt.Sprintf("  user %s%s%s", user.Name, pwd, groups)
	}
	return result
}
