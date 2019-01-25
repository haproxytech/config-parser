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

func (p *Mode) Clear() {
	p.Init()
}

func (p *Mode) Get(createIfNotExist bool) (common.ParserData, error) {
	if p.data == nil {
		if createIfNotExist {
			p.data = &types.StringC{}
			return p.data, nil
		}
		return p.data, nil
	}
	return nil, &errors.FetchError{}
}

func (p *Mode) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case *types.StringC:
		p.data = newValue
	case types.StringC:
		p.data = &newValue
	}
	return fmt.Errorf("casting error")
}

func (p *Mode) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := p.Get(false)
	p.Clear()
	_, err := p.Parse(data, parts, []string{}, comment)
	if err != nil {
		p.Set(oldData)
	}
	return err
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
		return nil, &errors.FetchError{}
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("mode %s", p.data.Value),
			Comment: p.data.Comment,
		},
	}, nil
}
