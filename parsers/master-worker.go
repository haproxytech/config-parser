package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type MasterWorker struct {
	data *types.Enabled
}

func (m *MasterWorker) Init() {
	m.data = nil
}

func (m *MasterWorker) GetParserName() string {
	return "master-worker"
}

func (m *MasterWorker) Get(createIfNotExist bool) (common.ParserData, error) {
	if m.data == nil {
		if createIfNotExist {
			m.data = &types.Enabled{}
			return m.data, nil
		}
		return nil, errors.FetchError
	}
	return m.data, nil
}

func (p *MasterWorker) GetOne(index int) (common.ParserData, error) {
	if index != 0 {
		return nil, errors.FetchError
	}
	if p.data == nil {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (m *MasterWorker) Set(data common.ParserData, index int) error {
	if data == nil {
		m.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.Enabled:
		m.data = newValue
	case types.Enabled:
		m.data = &newValue
	default:
		return fmt.Errorf("casting error")
	}
	return nil
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
