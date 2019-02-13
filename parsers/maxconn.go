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
