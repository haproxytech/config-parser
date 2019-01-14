package parsers

import (
	"github.com/haproxytech/config-parser/errors"
)

type MasterWorker struct {
	Enabled bool
}

func (m *MasterWorker) Init() {
	m.Enabled = false
}

func (m *MasterWorker) GetParserName() string {
	return "master-worker"
}

func (m *MasterWorker) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if parts[0] == "master-worker" {
		m.Enabled = true
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

func (m *MasterWorker) Result(AddComments bool) []string {
	if m.Enabled {
		return []string{"  master-worker"}
	}
	return []string{}
}
