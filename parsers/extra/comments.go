package extra

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type Comments struct {
	data []string
}

func (p *Comments) Init() {
	p.data = []string{}
}

func (p *Comments) GetParserName() string {
	return "#"
}

func (p *Comments) Get(createIfNotExist bool) (common.ParserData, error) {
	return p.data, nil
}

func (p *Comments) GetOne(index int) (common.ParserData, error) {
	if len(p.data) == 0 {
		return nil, errors.FetchError
	}
	if index < 0 || index >= len(p.data) {
		return nil, errors.FetchError
	}
	return p.data[index], nil
}

func (p *Comments) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case []string:
		p.data = newValue
	case *string:
		p.data = append(p.data, *newValue)
	case string:
		p.data = append(p.data, newValue)
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (p *Comments) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if line[0] == '#' {
		p.data = append(p.data, line)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "Comments", Line: line}
}

func (p *Comments) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(p.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(p.data))
	for index, comment := range p.data {
		result[index] = common.ReturnResultLine{
			Data:    comment,
			Comment: "",
		}
	}
	return result, nil
}
