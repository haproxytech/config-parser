package simple

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type SimpleTime struct {
	Name string
	data *types.StringC
}

func (s *SimpleTime) Init() {
	s.data = nil
}

func (s *SimpleTime) GetParserName() string {
	return s.Name
}

func (s *SimpleTime) Clear() {
	s.Init()
}

func (s *SimpleTime) Get(createIfNotExist bool) (common.ParserData, error) {
	if s.data == nil {
		if createIfNotExist {
			s.data = &types.StringC{}
			return s.data, nil
		}
		return nil, &errors.FetchError{}
	}
	return s.data, nil
}

func (s *SimpleTime) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case *types.StringC:
		s.data = newValue
	case types.StringC:
		s.data = &newValue
	}
	return fmt.Errorf("casting error")
}

func (s *SimpleTime) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := s.Get(false)
	s.Clear()
	_, err := s.Parse(data, parts, []string{}, comment)
	if err != nil {
		s.Set(oldData)
	}
	return err
}

func (s *SimpleTime) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == s.Name {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "SimpleTime", Line: line, Message: "Parse error"}
		}
		s.data = &types.StringC{
			Value:   parts[1],
			Comment: comment,
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.Name, Line: line}
}

func (s *SimpleTime) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, &errors.FetchError{}
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("%s %s", s.Name, s.data.Value),
			Comment: s.data.Comment,
		},
	}, nil
}
