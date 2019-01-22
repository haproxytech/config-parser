package parsers

import (
	"fmt"
	"strconv"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type MaxConn struct {
	Value   int64
	Comment string
}

func (m *MaxConn) Init() {
	m.Value = 0
}

func (m *MaxConn) GetParserName() string {
	return "maxconn"
}

func (m *MaxConn) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "maxconn" {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "SectionName", Line: line, Message: "Parse error"}
		}
		var num int64
		if num, err = strconv.ParseInt(parts[1], 10, 64); err != nil {
			m.Value = 0
			return "", &errors.ParseError{Parser: "SectionName", Line: line, Message: err.Error()}
		} else {
			m.Value = num
			m.Comment = comment
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "SectionName", Line: line}
}

func (m *MaxConn) Valid() bool {
	if m.Value > 0 {
		return true
	}
	return false
}

func (m *MaxConn) Result(AddComments bool) []common.ReturnResultLine {
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("maxconn %d", m.Value),
			Comment: m.Comment,
		},
	}
}
