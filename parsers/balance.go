package parsers

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Balance struct {
	data *types.Balance
}

func (p *Balance) Init() {
	p.data = nil
}

func (p *Balance) GetParserName() string {
	return "balance"
}

func (p *Balance) Clear() {
	p.Init()
}

func (p *Balance) Get(createIfNotExist bool) (common.ParserData, error) {
	if p.data == nil {
		if createIfNotExist {
			p.data = &types.Balance{
				Arguments: []string{},
			}
			return p.data, nil
		}
		return p.data, nil
	}
	return nil, &errors.FetchError{}
}

func (p *Balance) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case *types.Balance:
		p.data = newValue
	case types.Balance:
		p.data = &newValue
	}
	return fmt.Errorf("casting error")
}

func (p *Balance) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := p.Get(false)
	p.Clear()
	_, err := p.Parse(data, parts, []string{}, comment)
	if err != nil {
		p.Set(oldData)
	}
	return err
}

func (p *Balance) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "balance" {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "Balance", Line: line, Message: "Parse error"}
		}
		p.data = &types.Balance{
			Arguments: []string{},
			Comment:   comment,
		}

		switch parts[1] {
		case "roundrobin", "static-rr", "leastconn", "first", "source", "random":
			p.data.Algorithm = parts[1]
			return "", nil
		case "uri", "url_param":
			p.data.Algorithm = parts[1]
			if len(parts) > 2 {
				p.data.Arguments = parts[2:]
				return "", nil
			}
			return "", &errors.ParseError{Parser: "Balance", Line: line}
		}
		if strings.HasPrefix(parts[1], "hdr(") && strings.HasSuffix(parts[1], ")") {
			p.data.Algorithm = parts[1]
			return "", nil
		}
		if strings.HasPrefix(parts[1], "rdp-cookie(") && strings.HasSuffix(parts[1], ")") {
			p.data.Algorithm = parts[1]
			return "", nil
		}
		return "", &errors.ParseError{Parser: "Balance", Line: line}
	}
	return "", &errors.ParseError{Parser: "Balance", Line: line}
}

func (p *Balance) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if p.data == nil {
		return nil, &errors.FetchError{}
	}
	params := ""
	if len(p.data.Arguments) > 0 {
		params = fmt.Sprintf(" %s", strings.Join(p.data.Arguments, " "))
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data: fmt.Sprintf("balance %s%s", p.data.Algorithm, params),
		},
	}, nil
}
