package parsers

import (
	"fmt"
	"strconv"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type NbThread struct {
	data *types.Int64C
}

func (n *NbThread) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "nbthread" {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "NbThread", Line: line, Message: "Parse error"}
		}
		var num int64
		if num, err = strconv.ParseInt(parts[1], 10, 64); err != nil {
			return "", &errors.ParseError{Parser: "NbThread", Line: line, Message: err.Error()}
		} else {
			n.data = &types.Int64C{
				Value:   num,
				Comment: comment,
			}
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "nbthread", Line: line}
}

func (n *NbThread) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if n.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("nbthread %d", n.data.Value),
			Comment: n.data.Comment,
		},
	}, nil
}
