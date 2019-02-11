package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Mode struct {
	data *types.StringC
}

func (p *Mode) Init() {
	p.data = nil
}

func (p *Mode) GetParserName() string {
	return "mode"
}

func (p *Mode) Get(createIfNotExist bool) (common.ParserData, error) {
	if p.data == nil {
		if createIfNotExist {
			p.data = &types.StringC{}
			return p.data, nil
		}
		return p.data, nil
	}
	return nil, errors.FetchError
}

func (p *Mode) GetOne(index int) (common.ParserData, error) {
	if index != 0 {
		return nil, errors.FetchError
	}
	if p.data == nil {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *Mode) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.StringC:
		p.data = newValue
	case types.StringC:
		p.data = &newValue
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (p *Mode) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "mode" {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "Mode", Line: line, Message: "Parse error"}
		}
		if parts[1] == "http" || parts[1] == "tcp" || parts[1] == "health" {
			p.data = &types.StringC{
				Value:   parts[1],
				Comment: comment,
			}
			return "", nil
		}
		return "", &errors.ParseError{Parser: "Mode", Line: line}
	}
	return "", &errors.ParseError{Parser: "Mode", Line: line}
}

func (p *Mode) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if p.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("mode %s", p.data.Value),
			Comment: p.data.Comment,
		},
	}, nil
}
