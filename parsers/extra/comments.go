package extra

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type Comments struct {
	comments []string
}

func (p *Comments) Init() {
	p.comments = []string{}
}

func (p *Comments) Clear() {
	p.Init()
}

func (p *Comments) GetParserName() string {
	return "#"
}

func (p *Comments) Get(createIfNotExist bool) (common.ParserData, error) {
	return p.comments, nil
}

func (p *Comments) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case []string:
		p.comments = newValue
	case string:
		p.comments = append(p.comments, newValue)
	}
	return fmt.Errorf("casting error")
}

func (p *Comments) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := p.Get(false)
	p.Clear()
	_, err := p.Parse(data, parts, []string{}, comment)
	if err != nil {
		p.Set(oldData)
	}
	return err
}

func (p *Comments) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if line[0] == '#' {
		p.comments = append(p.comments, line)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "Comments", Line: line}
}

func (p *Comments) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(p.comments) == 0 {
		return nil, &errors.FetchError{}
	}
	result := make([]common.ReturnResultLine, len(p.comments))
	for index, comment := range p.comments {
		result[index] = common.ReturnResultLine{
			Data:    comment,
			Comment: "",
		}
	}
	return result, nil
}
