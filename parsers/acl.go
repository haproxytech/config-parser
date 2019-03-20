package parsers

import (
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Acl struct {
	data []types.Acl
}

func (h *Acl) parse(line string, parts []string, comment string) (*types.Acl, error) {
	if len(parts) >= 3 {
		data := &types.Acl{
			Name:      parts[1],
			Criterion: parts[2],
			Value:     strings.Join(parts[3:], " "),
			Comment:   comment,
		}
		return data, nil
	}
	return nil, &errors.ParseError{Parser: "AclLines", Line: line}
}

func (h *Acl) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		var sb strings.Builder
		sb.WriteString("acl ")
		sb.WriteString(req.Name)
		sb.WriteString(" ")
		sb.WriteString(req.Criterion)
		sb.WriteString(" ")
		sb.WriteString(req.Value)
		result[index] = common.ReturnResultLine{
			Data:    sb.String(),
			Comment: req.Comment,
		}
	}
	return result, nil
}
