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
	name     string
	data     *types.StringC
}

func (s *SimpleTimeTwoWords) Init() {
	s.data = nil
	s.name = fmt.Sprintf(strings.Join(s.Keywords, " "))
}

func (s *SimpleTimeTwoWords) GetParserName() string {
	return s.name
}

func (s *SimpleTimeTwoWords) Clear() {
	s.Init()
}

func (s *SimpleTimeTwoWords) Get(createIfNotExist bool) (common.ParserData, error) {
	if s.data == nil {
		if createIfNotExist {
			s.data = &types.StringC{}
			return s.data, nil
		}
		return nil, &errors.FetchError{}
	}
	return s.data, nil
}

func (s *SimpleTimeTwoWords) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case *types.StringC:
		s.data = newValue
	case types.StringC:
		s.data = &newValue
	}
	return fmt.Errorf("casting error")
}

func (s *SimpleTimeTwoWords) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := s.Get(false)
	s.Clear()
	_, err := s.Parse(data, parts, []string{}, comment)
	if err != nil {
		s.Set(oldData)
	}
	return err
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
	return "", &errors.ParseError{Parser: s.name, Line: line}
}

func (s *SimpleTimeTwoWords) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, &errors.FetchError{}
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("%s %s", s.name, s.data.Value),
			Comment: s.data.Comment,
		},
	}, nil
}
