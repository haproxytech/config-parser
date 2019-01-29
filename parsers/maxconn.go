package parsers

import (
	"fmt"
	"strconv"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type MaxConn struct {
	data *types.Int64C
}

func (p *MaxConn) Init() {
	p.data = nil
}

func (p *MaxConn) GetParserName() string {
	return "maxconn"
}

func (p *MaxConn) Clear() {
	p.Init()
}

func (p *MaxConn) Get(createIfNotExist bool) (common.ParserData, error) {
	if p.data.Value < 1 {
		if createIfNotExist {
			p.data = &types.Int64C{}
			return p.data, nil
		}
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *MaxConn) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case *types.Int64C:
		p.data = newValue
	case types.Int64C:
		p.data = &newValue
	}
	return fmt.Errorf("casting error")
}

func (p *MaxConn) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := p.Get(false)
	p.Clear()
	_, err := p.Parse(data, parts, []string{}, comment)
	if err != nil {
		p.Set(oldData)
	}
	return err
}

func (p *MaxConn) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "maxconn" {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "SectionName", Line: line, Message: "Parse error"}
		}
		var num int64
		if num, err = strconv.ParseInt(parts[1], 10, 64); err != nil {
			return "", &errors.ParseError{Parser: "SectionName", Line: line, Message: err.Error()}
		} else {
			p.data = &types.Int64C{
				Value:   num,
				Comment: comment,
			}
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "SectionName", Line: line}
}

func (p *MaxConn) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if p.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("maxconn %d", p.data.Value),
			Comment: p.data.Comment,
		},
	}, nil
}
