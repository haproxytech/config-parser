package parsers

import (
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type OptionMysqlCheck struct {
	data *types.OptionMysqlCheck
}

/*
option mysql-check [ user <username> [ post-41 ] ]
*/
func (s *OptionMysqlCheck) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 1 && parts[0] == "option" && parts[1] == "mysql-check" {
		data := &types.OptionMysqlCheck{
			Comment: comment,
		}
		if len(parts) > 2 {
			if len(parts) < 6 {
				if parts[2] != "user" {
					return "", errors.InvalidData
				}
				if len(parts) < 4 {
					return "", errors.InvalidData
				}
				data.User = parts[3]
				if len(parts) == 5 {
					if parts[4] != "post-41" {
						return "", errors.InvalidData
					}
					data.Post41 = true
				}
			} else {
				return "", errors.InvalidData
			}
		}
		s.data = data
		return "", nil
	}
	return "", &errors.ParseError{Parser: "option mysql-check", Line: line}
}

func (s *OptionMysqlCheck) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, errors.FetchError
	}
	var sb strings.Builder
	sb.WriteString("option mysql-check")
	if s.data.User != "" {
		sb.WriteString(" user ")
		sb.WriteString(s.data.User)
	}
	if s.data.Post41 {
		sb.WriteString(" post-41")
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    sb.String(),
			Comment: s.data.Comment,
		},
	}, nil
}
