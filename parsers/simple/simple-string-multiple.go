package simple

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type SimpleStringMultiple struct {
	Name string
	data *types.StringSliceC
}

func (s *SimpleStringMultiple) Init() {
	s.data = nil
}

func (s *SimpleStringMultiple) GetParserName() string {
	return s.Name
}

func (s *SimpleStringMultiple) Get(createIfNotExist bool) (common.ParserData, error) {
	if s.data == nil {
		if createIfNotExist {
			s.data = &types.StringSliceC{}
			return s.data, nil
		}
		return nil, errors.FetchError
	}
	return s.data, nil
}

func (p *SimpleStringMultiple) GetOne(index int) (common.ParserData, error) {
	if index != 0 {
		return nil, errors.FetchError
	}
	if p.data == nil {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (s *SimpleStringMultiple) Set(data common.ParserData, index int) error {
	if data == nil {
		s.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.StringSliceC:
		s.data = newValue
	case types.StringSliceC:
		s.data = &newValue
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (s *SimpleStringMultiple) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == s.Name {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "SimpleStringMultiple", Line: line, Message: "Parse error"}
		}
		s.data = &types.StringSliceC{
			Value:   parts[1:],
			Comment: comment,
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.Name, Line: line}
}

func (s *SimpleStringMultiple) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("%s %s", s.Name, strings.Join(s.data.Value, " ")),
			Comment: s.data.Comment,
		},
	}, nil
}
