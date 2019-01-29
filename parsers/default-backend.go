package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type DefaultBackend struct {
	data *types.StringC
}

func (s *DefaultBackend) Init() {
	s.data = nil
}

func (s *DefaultBackend) GetParserName() string {
	return "default_backend"
}

func (s *DefaultBackend) Clear() {
	s.Init()
}

func (s *DefaultBackend) Get(createIfNotExist bool) (common.ParserData, error) {
	if s.data == nil {
		if createIfNotExist {
			s.data = &types.StringC{}
			return s.data, nil
		}
		return nil, errors.FetchError
	}
	return s.data, nil
}

func (s *DefaultBackend) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case *types.StringC:
		s.data = newValue
	case types.StringC:
		s.data = &newValue
	}
	return fmt.Errorf("casting error")
}

func (s *DefaultBackend) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := s.Get(false)
	s.Clear()
	_, err := s.Parse(data, parts, []string{}, comment)
	if err != nil {
		s.Set(oldData)
	}
	return err
}

func (s *DefaultBackend) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "default_backend" {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "DefaultBackend", Line: line, Message: "Parse error"}
		}
		s.data = &types.StringC{
			Comment: comment,
			Value:   parts[1],
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "default_backend", Line: line}
}

func (s *DefaultBackend) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("default_backend %s", s.data.Value),
			Comment: s.data.Comment,
		},
	}, nil
}
