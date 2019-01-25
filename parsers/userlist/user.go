package userlist

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type UserLines struct {
	data []types.User
}

func (l *UserLines) Init() {
	l.data = []types.User{}
}

func (l *UserLines) GetParserName() string {
	return "user"
}

func (l *UserLines) Clear() {
	l.Init()
}

func (l *UserLines) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(l.data) == 0 && createIfNotExist {
		return nil, &errors.FetchError{}
	}
	return l.data, nil
}

func (l *UserLines) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case []types.User:
		l.data = newValue
	case *types.User:
		l.data = append(l.data, *newValue)
	case types.User:
		l.data = append(l.data, newValue)
	}
	return fmt.Errorf("casting error")
}

func (l *UserLines) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := l.Get(false)
	l.Clear()
	_, err := l.Parse(data, parts, []string{}, comment)
	if err != nil {
		l.Set(oldData)
	}
	return err
}

func (l *UserLines) parseUserLine(line string, parts []string, comment string) (*types.User, error) {
	if len(parts) >= 2 {
		user := types.User{
			Name:    parts[1],
			Comment: comment,
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
			user.Groups = common.StringSplitIgnoreEmpty(parts[index], ',')
		}
		return &user, nil
	}
	return nil, &errors.ParseError{Parser: "UserLines", Line: line}
}

func (l *UserLines) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "user" {
		user, err := l.parseUserLine(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "UserLines", Line: line}
		}
		l.data = append(l.data, *user)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "UserLines", Line: line}
}

func (l *UserLines) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(l.data) == 0 {
		return nil, &errors.FetchError{}
	}
	result := make([]common.ReturnResultLine, len(l.data))
	for index, user := range l.data {
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
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("user %s%s%s", user.Name, pwd, groups),
			Comment: user.Comment,
		}
	}
	return result, nil
}
