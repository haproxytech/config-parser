package simple

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type SimpleTimeout struct {
	Name string
	data *types.SimpleTimeout
}

func (t *SimpleTimeout) Init() {
	t.data = nil
}

func (t *SimpleTimeout) GetParserName() string {
	return fmt.Sprintf("timeout %s", t.Name)
}

func (p *SimpleTimeout) Get(createIfNotExist bool) (common.ParserData, error) {
	if p.data == nil {
		if createIfNotExist {
			p.data = &types.SimpleTimeout{}
			return p.data, nil
		}
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *SimpleTimeout) Set(data common.ParserData) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.SimpleTimeout:
		p.data = newValue
	case types.SimpleTimeout:
		p.data = &newValue
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (p *SimpleTimeout) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := p.Get(false)
	p.Init()
	_, err := p.Parse(data, parts, []string{}, comment)
	if err != nil {
		p.Set(oldData)
	}
	return err
}

func (t *SimpleTimeout) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 2 && parts[0] == "timeout" && parts[1] == t.Name {
		t.data = &types.SimpleTimeout{
			Value:   parts[2],
			Comment: comment,
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: fmt.Sprintf("timeout %s", t.Name), Line: line}
}

func (t *SimpleTimeout) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if t.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("timeout %s %s", t.Name, t.data.Value),
			Comment: t.data.Comment,
		},
	}, nil
}
