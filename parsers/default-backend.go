package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type DefaultBackend struct {
	Value   string
	Comment string
}

func (s *DefaultBackend) Init() {
}

func (s *DefaultBackend) GetParserName() string {
	return "default_backend"
}

func (s *DefaultBackend) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "default_backend" {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "DefaultBackend", Line: line, Message: "Parse error"}
		}
		s.Comment = comment
		s.Value = parts[1]
		return "", nil
	}
	return "", &errors.ParseError{Parser: "default_backend", Line: line}
}

func (s *DefaultBackend) Valid() bool {
	if s.Value != "" {
		return true
	}
	return false
}

func (s *DefaultBackend) Result(AddComments bool) []common.ReturnResultLine {
	if s.Value != "" {
		return []common.ReturnResultLine{
			common.ReturnResultLine{
				Data:    fmt.Sprintf("default_backend %s", s.Value),
				Comment: s.Comment,
			},
		}
	}
	return []common.ReturnResultLine{}
}
