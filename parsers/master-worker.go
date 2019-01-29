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

func (m *MasterWorker) Clear() {
	m.Init()
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

func (m *MasterWorker) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case *types.Enabled:
		m.data = newValue
	case types.Enabled:
		m.data = &newValue
	}
	return fmt.Errorf("casting error")
}

func (m *MasterWorker) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := m.Get(false)
	m.Clear()
	_, err := m.Parse(data, parts, []string{}, comment)
	if err != nil {
		m.Set(oldData)
	}
	return err
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
