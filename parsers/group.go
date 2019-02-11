package parsers

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Group struct {
	data []types.Group
}

func (l *Group) parse(line string, parts []string, comment string) (*types.Group, error) {
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
	return nil, &errors.ParseError{Parser: "Group", Line: line}
}

func (l *Group) Result(AddComments bool) ([]common.ReturnResultLine, error) {
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
