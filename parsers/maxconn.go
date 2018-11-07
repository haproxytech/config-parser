package parsers

import (
	"fmt"
	"strconv"
	"strings"

	"config-parser/errors"
)

type MaxConn struct {
	value int64
	valid bool
}

func (m *MaxConn) Init() {
	m.valid = false
}

func (m *MaxConn) GetParserName() string {
	return "maxconn"
}

func (m *MaxConn) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, "maxconn") {
		elements := strings.SplitN(line, " ", 2)
		if len(elements) != 2 {
			return "", &errors.ParseError{Parser: "SectionName", Line: line, Message: "Parse error"}
		}
		var num int64
		if num, err = strconv.ParseInt(elements[1], 10, 64); err != nil {
			return "", &errors.ParseError{Parser: "SectionName", Line: line, Message: err.Error()}
		} else {
			m.valid = true
			m.value = num
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "SectionName", Line: line}
}

func (m *MaxConn) Valid() bool {
	if m.valid {
		return true
	}
	return false
}

func (m *MaxConn) String() []string {
	return []string{fmt.Sprintf("  maxconn %d", m.value)}
}
