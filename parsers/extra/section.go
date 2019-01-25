package extra

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type SectionName struct {
	Name    string
	section *types.Section
}

func (s *SectionName) Init() {
	s.section = &types.Section{}
}

func (s *SectionName) Clear() {
	s.Init()
}

func (s *SectionName) GetParserName() string {
	return s.Name
}

func (s *SectionName) Get(createIfNotExist bool) (common.ParserData, error) {
	if s.section.Name != "" {
		return s.section, nil
	} else if createIfNotExist {
		s.Clear()
		return s.section, nil
	}
	return nil, fmt.Errorf("No data")
}

func (s *SectionName) Set(data common.ParserData) error {
	newData, ok := data.(types.Section)
	if ok {
		s.section = &newData
		return nil
	}
	return fmt.Errorf("casting error")
}

func (s *SectionName) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	s.Clear()
	_, err := s.Parse(data, parts, []string{}, comment)
	return err
}

//Parse see if we have section name
func (s *SectionName) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == s.Name {
		if len(parts) > 1 {
			s.section.Name = parts[1]
		}
		if len(previousParts) > 1 && previousParts[0] == "#" {
			s.section.Comment = previousParts[1]
		}
		return s.Name, nil
	}
	return "", &errors.ParseError{Parser: "SectionName", Line: line}
}

func (s *SectionName) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	return nil, fmt.Errorf("Not valid")
}
