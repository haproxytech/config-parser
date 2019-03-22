package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type StatsTimeout struct {
	data *types.StringC
}

func (s *StatsTimeout) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 1 && parts[0] == "stats" && parts[1] == "timeout" {
		if len(parts) < 3 {
			return "", &errors.ParseError{Parser: "StatsTimeout", Line: line, Message: "Parse error"}
		}
		s.data = &types.StringC{
			Value:   parts[2],
			Comment: comment,
		}
		//todo add validation with simple timeouts
		return "", nil
	}
	return "", &errors.ParseError{Parser: "StatsTimeout", Line: line}
}

func (s *StatsTimeout) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("stats timeout %s", s.data.Value),
			Comment: s.data.Comment,
		},
	}, nil
}
