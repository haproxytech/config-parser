package parsers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/params"
	"github.com/haproxytech/config-parser/v4/types"
)

type ServerTemplate struct {
	data        []types.ServerTemplate
	preComments []string // comments that appear before the the actual line
}

func (h *ServerTemplate) parse(line string, parts []string, comment string) (*types.ServerTemplate, error) {
	if len(parts) < 4 {
		return nil, &errors.ParseError{Parser: "ServerTemplate", Line: line}
	}

	data := &types.ServerTemplate{}
	data.Prefix = parts[1]
	data.NumOrRange = parts[2]
	data.Comment = comment

	address := common.StringSplitIgnoreEmpty(parts[3], ':')
	if len(address) == 2 {
		if port, err := strconv.ParseInt(address[1], 10, 64); err == nil {
			data.Fqdn = address[0]
			data.Port = port
		}
	} else {
		data.Fqdn = parts[3]
	}
	if len(parts) >= 4 {
		data.Params = params.ParseServerOptions(parts[4:])
	}
	return data, nil
}

func (h *ServerTemplate) Result() ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.ErrFetch
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, template := range h.data {
		var sb strings.Builder
		sb.WriteString("server-template ")
		sb.WriteString(template.Prefix)
		sb.WriteString(" ")
		sb.WriteString(template.NumOrRange)
		sb.WriteString(" ")
		sb.WriteString(template.Fqdn)
		if template.Port != 0 {
			sb.WriteString(fmt.Sprintf(":%d", template.Port))
		}
		params := params.ServerOptionsString(template.Params)
		if params != "" {
			sb.WriteString(" ")
			sb.WriteString(params)
		}
		result[index] = common.ReturnResultLine{
			Data:    sb.String(),
			Comment: template.Comment,
		}
	}
	return result, nil
}
