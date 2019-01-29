package userlist

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type GroupLines struct {
	data []types.Group
}

func (l *GroupLines) Init() {
	l.data = []types.Group{}
}

func (l *GroupLines) GetParserName() string {
	return "group"
}

func (l *GroupLines) Clear() {
	l.Init()
}

func (l *GroupLines) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(l.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return l.data, nil
}

func (l *GroupLines) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case []types.Group:
		l.data = newValue
	case types.Group:
		l.data = append(l.data, newValue)
	}
	return fmt.Errorf("casting error")
}

func (l *GroupLines) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := l.Get(false)
	l.Clear()
	_, err := l.Parse(data, parts, []string{}, comment)
	if err != nil {
		l.Set(oldData)
	}
	return err
}

func (l *GroupLines) parseGroupLine(line string, parts []string, comment string) (*types.Group, error) {
	if len(parts) >= 2 {
		group := &types.Group{
			Name:    parts[1],
			Comment: comment,
		}
		if len(parts) > 3 && parts[2] == "users" {
			group.Users = common.StringSplitIgnoreEmpty(parts[3], ',')
		}
		return group, nil
	}
	return nil, &errors.ParseError{Parser: "GroupLines", Line: line}
}

func (l *GroupLines) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "group" {
		group, err := l.parseGroupLine(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "GroupLines", Line: line}
		}
		l.data = append(l.data, *group)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "GroupLines", Line: line}
}

func (l *GroupLines) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(l.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(l.data))
	for index, group := range l.data {
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
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("group %s%s", group.Name, users),
			Comment: group.Comment,
		}
	}
	return result, nil
}
