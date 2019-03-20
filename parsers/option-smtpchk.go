package parsers

import (
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type OptionSmtpchk struct {
	data *types.OptionSmtpchk
}

/*
option smtpchk <hello> <domain>
*/
func (s *OptionSmtpchk) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 1 && parts[0] == "option" && parts[1] == "smtpchk" {
		data := &types.OptionSmtpchk{
			Comment: comment,
		}
		if len(parts) > 2 {
			if len(parts) != 4 {
				return "", errors.InvalidData
			}
			data.Hello = parts[2]
			data.Domain = parts[3]
			if data.Hello != "EHLO" {
				data.Hello = "HELO"
			}
		}
		s.data = data
		return "", nil
	}
	if len(parts) > 2 && parts[0] == "no" && parts[1] == "option" && parts[2] == "smtpchk" {
		data := &types.OptionSmtpchk{
			NoOption: true,
			Comment:  comment,
		}
		s.data = data
		return "", nil
	}
	return "", &errors.ParseError{Parser: "option smtpchk", Line: line}
}

func (s *OptionSmtpchk) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, errors.FetchError
	}
	var sb strings.Builder
	if s.data.NoOption {
		sb.WriteString("no ")
	}
	sb.WriteString("option smtpchk")
	if s.data.Hello != "" && !s.data.NoOption {
		sb.WriteString(" ")
		sb.WriteString(s.data.Hello)
		sb.WriteString(" ")
		sb.WriteString(s.data.Domain)
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    sb.String(),
			Comment: s.data.Comment,
		},
	}, nil
}
