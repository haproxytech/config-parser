package extra

import (
	"strings"

	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/helpers"
)

type SectionName struct {
	comment     string
	Name        string
	SectionName string
	Line        string
}

func (s *SectionName) Init() {
	s.comment = ""
	s.SectionName = ""
}

func (s *SectionName) GetParserName() string {
	return s.Name
}

func (s *SectionName) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, s.Name) {
		s.Line = line
		parts := helpers.StringSplitIgnoreEmpty(line, ' ')
		if len(parts) > 1 {
			s.SectionName = parts[1]
		}
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
