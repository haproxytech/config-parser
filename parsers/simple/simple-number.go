package simple

import (
	"fmt"
	"strconv"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type SimpleNumber struct {
	Name string
	data *types.Int64C
}

func (s *SimpleNumber) Init() {
	s.data = nil
}

func (s *SimpleNumber) GetParserName() string {
	return s.Name
}

func (s *SimpleNumber) Clear() {
	s.Init()
}

func (s *SimpleNumber) Get(createIfNotExist bool) (common.ParserData, error) {
	if s.data == nil {
		if createIfNotExist {
			s.data = &types.Int64C{}
			return s.data, nil
		}
		return nil, errors.FetchError
	}
	return s.data, nil
}

func (s *SimpleNumber) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case *types.Int64C:
		s.data = newValue
	case types.Int64C:
		s.data = &newValue
	}
	return fmt.Errorf("casting error")
}

func (s *SimpleNumber) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := s.Get(false)
	s.Clear()
	_, err := s.Parse(data, parts, []string{}, comment)
	if err != nil {
		s.Set(oldData)
	}
	return err
}

func (s *SimpleNumber) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == s.Name {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "SimpleNumber", Line: line, Message: "Parse error"}
		}
		var num int64
		if num, err = strconv.ParseInt(parts[1], 10, 64); err != nil {
			return "", &errors.ParseError{Parser: "SimpleNumber", Line: line, Message: err.Error()}
		} else {
			s.data = &types.Int64C{
				Value:   num,
				Comment: comment,
			}
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.Name, Line: line}
}

func (s *SimpleNumber) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, errors.FetchError
	}

	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("%s %d", s.Name, s.data.Value),
			Comment: s.data.Comment,
		},
	}, nil
}
