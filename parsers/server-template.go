package parsers

import (
	"strings"

	"github.com/haproxytech/config-parser/v3/common"
	"github.com/haproxytech/config-parser/v3/errors"
	"github.com/haproxytech/config-parser/v3/params"
	"github.com/haproxytech/config-parser/v3/types"
)

type ServerTemplate struct {
	data        *types.ServerTemplate
	preComments []string // comments that appear before the the actual line
}

func (h *ServerTemplate) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) < 4 {
		return "", &errors.ParseError{Parser: "ServerTemplate", Line: line}
	}
	h.data = &types.ServerTemplate{
		Prefix:     parts[1],
		NumOrRange: parts[2],
		Fqdn:       parts[3],
		Comment:    comment,
	}
	if len(parts) >= 4 {
		h.data.Params = params.ParseServerOptions(parts[4:])
	}
	return "", nil
}

func (h *ServerTemplate) Result() ([]common.ReturnResultLine, error) {
	if h.data == nil {
		return nil, errors.ErrFetch
	}
	var sb strings.Builder
	sb.WriteString("server-template ")
	sb.WriteString(h.data.Prefix)
	sb.WriteString(" ")
	sb.WriteString(h.data.NumOrRange)
	sb.WriteString(" ")
	sb.WriteString(h.data.Fqdn)
	params := params.ServerOptionsString(h.data.Params)
	if params != "" {
		sb.WriteString(" ")
		sb.WriteString(params)
	}
	return []common.ReturnResultLine{
		{
			Data:    sb.String(),
			Comment: h.data.Comment,
		},
	}, nil
}
