package extra

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Section struct {
	Name string
	data *types.Section
}

func (s *Section) Init() {
	s.data = &types.Section{}
}

//Parse see if we have section name
func (s *Section) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == s.Name {
		if len(parts) > 1 {
			s.data.Name = parts[1]
		}
		if len(previousParts) > 1 && previousParts[0] == "#" {
			s.data.Comment = previousParts[1]
		}
		return s.Name, nil
	}
	return "", &errors.ParseError{Parser: "Section", Line: line}
}

func (s *Section) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	return nil, fmt.Errorf("Not valid")
}
