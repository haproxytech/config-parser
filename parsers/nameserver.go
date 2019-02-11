package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Nameserver struct {
	data []types.Nameserver
}

func (l *Nameserver) parse(line string, parts []string, comment string) (*types.Nameserver, error) {
	if len(parts) >= 3 {
		return &types.Nameserver{
			Name:    parts[1],
			Address: parts[2],
			Comment: comment,
		}, nil
	}
	return nil, &errors.ParseError{Parser: "Nameserver", Line: line}
}

func (l *Nameserver) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(l.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(l.data))
	for index, nameserver := range l.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("nameserver %s %s", nameserver.Name, nameserver.Address),
			Comment: nameserver.Comment,
		}
	}
	return result, nil
}
