package simple

import (
	"fmt"
	"strconv"
	"strings"

	"config-parser/errors"
)

type SimpleNumber struct {
	enabled    bool
	value      int64
	Name       string
	searchName string
}

func (s *SimpleNumber) Init() {
	s.enabled = false
	s.searchName = s.Name
}

func (s *SimpleNumber) GetParserName() string {
	return s.searchName
}

func (s *SimpleNumber) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, s.searchName) {
		elements := strings.SplitN(line, " ", 2)
		if len(elements) != 2 {
			return "", &errors.ParseError{Parser: "SimpleNumber", Line: line, Message: "Parse error"}
		}
		var num int64
		if num, err = strconv.ParseInt(elements[1], 10, 64); err != nil {
			return "", &errors.ParseError{Parser: "SimpleNumber", Line: line, Message: err.Error()}
		} else {
			s.enabled = true
			s.value = num
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.searchName, Line: line}
}

func (s *SimpleNumber) Valid() bool {
	if s.enabled {
		return true
	}
	return false
}

func (s *SimpleNumber) String() []string {
	if s.enabled {
		return []string{fmt.Sprintf("  %s %d", s.searchName, s.value)}
	}
	return []string{}
}
