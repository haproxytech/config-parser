package extra

import (
	"github.com/haproxytech/config-parser/errors"
)

type SectionName struct {
	Comment     string
	Name        string
	SectionName string
}

func (s *SectionName) Init() {
	s.Comment = ""
	s.SectionName = ""
}

func (s *SectionName) GetParserName() string {
	return s.Name
}

//Parse see if we have section name
func (s *SectionName) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if parts[0] == s.Name {
		if len(parts) > 1 {
			s.SectionName = parts[1]
		}
		if len(previousParts) > 1 && previousParts[0] == "#" {
			s.Comment = previousParts[1]
		}
		return s.Name, nil
	}
	return "", &errors.ParseError{Parser: "SectionName", Line: line}
}

func (s *SectionName) Valid() bool {
	return false
}

func (s *SectionName) Result(AddComments bool) []string {
	return []string{s.Comment, s.SectionName}
}
