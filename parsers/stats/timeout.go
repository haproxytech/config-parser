package stats

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Timeout struct {
	name string
	data *types.StringSliceC
}

func (s *Timeout) Init() {
	s.name = "stats timeout"
	s.data = nil
}

func (s *Timeout) GetParserName() string {
	return s.name
}

func (s *Timeout) Clear() {
	s.Init()
}

func (s *Timeout) Get(createIfNotExist bool) (common.ParserData, error) {
	if s.data == nil {
		if createIfNotExist {
			s.data = &types.StringSliceC{}
			return s.data, nil
		}
		return nil, errors.FetchError
	}
	return s.data, nil
}

func (s *Timeout) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case *types.StringSliceC:
		s.data = newValue
	case types.StringSliceC:
		s.data = &newValue
	}
	return fmt.Errorf("casting error")
}

func (s *Timeout) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := s.Get(false)
	s.Clear()
	_, err := s.Parse(data, parts, []string{}, comment)
	if err != nil {
		s.Set(oldData)
	}
	return err
}

func (s *Timeout) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 1 && parts[0] == "stats" && parts[1] == "timeout" {
		if len(parts) < 3 {
			return "", &errors.ParseError{Parser: "StatsTimeout", Line: line, Message: "Parse error"}
		}
		s.data = &types.StringSliceC{
			Value:   parts[2:],
			Comment: comment,
		}
		//todo add validation with simple timeouts
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.name, Line: line}
}

func (s *Timeout) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("%s %s", s.name, strings.Join(s.data.Value, " ")),
			Comment: s.data.Comment,
		},
	}, nil
}
