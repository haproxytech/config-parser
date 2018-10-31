package extra

import (
	"config-parser/errors"
	"strings"
)

type SectionName struct {
	comment string
	Name    string
	Line    string
	valid   bool
}

func (s *SectionName) Init() {
	s.comment = ""
	s.valid = false
}

func (s *SectionName) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, s.Name) {
		s.valid = true
		s.Line = line
		if len(previousLine) > 0 && previousLine[0] == '#' {
			s.comment = previousLine
		}
		return s.Name, nil
	}
	return "", &errors.ParseError{Parser: "SectionName", Line: line}
}

func (s *SectionName) Valid() bool {
	return false
}

func (s *SectionName) String() []string {
	return []string{s.comment, s.Line}
}
