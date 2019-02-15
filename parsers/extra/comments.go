package extra

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Comments struct {
	Name string
	data []types.Comments
}

func (p *Comments) Init() {
	p.Name = "#"
	p.data = []types.Comments{}
}

func (p *Comments) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if line[0] == '#' {
		p.data = append(p.data, types.Comments{
			Value: comment,
		})
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
			Data:    "# " + comment.Value,
			Comment: "",
		}
	}
	return result, nil
}
