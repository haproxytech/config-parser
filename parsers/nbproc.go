package parsers

import (
	"fmt"
	"strconv"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type NbProc struct {
	data *types.Int64C
}

func (n *NbProc) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "nbproc" {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "NbProc", Line: line, Message: "Parse error"}
		}
		var num int64
		if num, err = strconv.ParseInt(parts[1], 10, 64); err != nil {
			return "", &errors.ParseError{Parser: "NbProc", Line: line, Message: err.Error()}
		} else {
			n.data = &types.Int64C{
				Value:   num,
				Comment: comment,
			}
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "nbproc", Line: line}
}

func (n *NbProc) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if n.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("nbproc %d", n.data.Value),
			Comment: n.data.Comment,
		},
	}, nil
}
