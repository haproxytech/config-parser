package simple

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/types"

	"github.com/haproxytech/config-parser/errors"
)

type SimpleFlag struct {
	Name string
	data *types.Enabled
}

func (s *SimpleFlag) Init() {
	s.data = nil
}

func (s *SimpleFlag) GetParserName() string {
	return s.Name
}

func (s *SimpleFlag) Clear() {
	s.Init()
}

func (s *SimpleFlag) Get(createIfNotExist bool) (common.ParserData, error) {
	if s.data == nil {
		if createIfNotExist {
			s.data = &types.Enabled{}
			return s.data, nil
		}
		return nil, &errors.FetchError{}
	}
	return s.data, nil
}

func (s *SimpleFlag) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case *types.Enabled:
		s.data = newValue
	case types.Enabled:
		s.data = &newValue
	}
	return fmt.Errorf("casting error")
}

func (s *SimpleFlag) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := s.Get(false)
	s.Clear()
	_, err := s.Parse(data, parts, []string{}, comment)
	if err != nil {
		s.Set(oldData)
	}
	return err
}

func (s *SimpleFlag) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == s.Name {
		s.data = &types.Enabled{
			Comment: comment,
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.Name, Line: line}
}

func (s *SimpleFlag) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, &errors.FetchError{}
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("%s", s.Name),
			Comment: s.data.Comment,
		},
	}, nil
}
