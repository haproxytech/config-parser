package simple

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type SimpleOption struct {
	Name string
	data *types.SimpleOption
}

func (o *SimpleOption) Init() {
	o.data = nil
}

func (o *SimpleOption) GetParserName() string {
	return fmt.Sprintf("option %s", o.Name)
}

func (p *SimpleOption) Get(createIfNotExist bool) (common.ParserData, error) {
	if p.data == nil {
		if createIfNotExist {
			p.data = &types.SimpleOption{}
			return p.data, nil
		}
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *SimpleOption) Set(data common.ParserData) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.SimpleOption:
		p.data = newValue
	case types.SimpleOption:
		p.data = &newValue
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (p *SimpleOption) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := p.Get(false)
	p.Init()
	_, err := p.Parse(data, parts, []string{}, comment)
	if err != nil {
		p.Set(oldData)
	}
	return err
}

func (o *SimpleOption) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 1 && parts[0] == "option" && parts[1] == o.Name {
		o.data = &types.SimpleOption{
			Comment: comment,
		}
		return "", nil
	}
	if len(parts) > 2 && parts[0] == "no" && parts[1] == "option" && parts[2] == o.Name {
		o.data = &types.SimpleOption{
			NoOption: true,
			Comment:  comment,
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: fmt.Sprintf("option %s", o.Name), Line: line}
}

func (o *SimpleOption) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if o.data == nil {
		return nil, errors.FetchError
	}
	noOption := ""
	if o.data.NoOption {
		noOption = "no "
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("%soption %s", noOption, o.Name),
			Comment: o.data.Comment,
		},
	}, nil
}
