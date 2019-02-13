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
