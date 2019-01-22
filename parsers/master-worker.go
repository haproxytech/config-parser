package parsers

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type MasterWorker struct {
	Enabled bool
	Comment string
}

func (m *MasterWorker) Init() {
	m.Enabled = false
}

func (m *MasterWorker) GetParserName() string {
	return "master-worker"
}

func (m *MasterWorker) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "master-worker" {
		m.Enabled = true
		m.Comment = comment
		return "", nil
	}
	return "", &errors.ParseError{Parser: "MasterWorker", Line: line}
}

func (m *MasterWorker) Valid() bool {
	if m.Enabled {
		return true
	}
	return false
}

func (m *MasterWorker) Result(AddComments bool) []common.ReturnResultLine {
	if m.Enabled {
		return []common.ReturnResultLine{
			common.ReturnResultLine{
				Data:    "master-worker",
				Comment: m.Comment,
			},
		}
	}
	return []common.ReturnResultLine{}
}
