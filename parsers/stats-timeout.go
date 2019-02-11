package parsers

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type StatsTimeout struct {
	name string
	data *types.StringSliceC
}

func (s *StatsTimeout) Init() {
	s.name = "stats timeout"
	s.data = nil
}

func (s *StatsTimeout) GetParserName() string {
	return s.name
}

func (s *StatsTimeout) Get(createIfNotExist bool) (common.ParserData, error) {
	if s.data == nil {
		if createIfNotExist {
			s.data = &types.StringSliceC{}
			return s.data, nil
		}
		return nil, errors.FetchError
	}
	return s.data, nil
}

func (p *StatsTimeout) GetOne(index int) (common.ParserData, error) {
	if index != 0 {
		return nil, errors.FetchError
	}
	if p.data == nil {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (s *StatsTimeout) Set(data common.ParserData, index int) error {
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

func (s *StatsTimeout) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 1 && parts[0] == "stats" && parts[1] == "timeout" {
		if len(parts) < 3 {
			return "", &errors.ParseError{Parser: "StatsStatsTimeout", Line: line, Message: "Parse error"}
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

func (s *StatsTimeout) Result(AddComments bool) ([]common.ReturnResultLine, error) {
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
