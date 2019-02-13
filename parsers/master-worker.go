package parsers

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type MasterWorker struct {
	data *types.Enabled
}

func (m *MasterWorker) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "master-worker" {
		m.data = &types.Enabled{
			Comment: comment,
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "MasterWorker", Line: line}
}

func (m *MasterWorker) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if m.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    "master-worker",
			Comment: m.data.Comment,
		},
	}, nil
}
