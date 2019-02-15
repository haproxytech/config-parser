package simple

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type SimpleTimeTwoWords struct {
	Keywords []string
	Name     string
	data     *types.StringC
}

func (s *SimpleTimeTwoWords) Init() {
	s.data = nil
	s.Name = fmt.Sprintf(strings.Join(s.Keywords, " "))
}

func (s *SimpleTimeTwoWords) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) >= 2 && parts[0] == s.Keywords[0] && parts[1] == s.Keywords[1] {
		if len(parts) < 3 {
			return "", &errors.ParseError{Parser: "SimpleTimeTwoWords", Line: line, Message: "Parse error"}
		}
		s.data = &types.StringC{
			Value:   parts[2],
			Comment: comment,
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.Name, Line: line}
}

func (s *SimpleTimeTwoWords) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("%s %s", s.Name, s.data.Value),
			Comment: s.data.Comment,
		},
	}, nil
}
