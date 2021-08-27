/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package parsers

import (
	"fmt"

	"github.com/deyunluo/config-parser/v4/common"
	"github.com/deyunluo/config-parser/v4/errors"
	"github.com/deyunluo/config-parser/v4/spoe/types"
)

type SPOESection struct {
	Name        string
	data        *types.SPOESection
	preComments []string // comments that appear before the the actual line
}

func (s *SPOESection) Init() {
	s.data = &types.SPOESection{}
}

// GetParserName return name of the section
func (s *SPOESection) GetParserName() string {
	return s.Name
}

// Parse see if we have section name
func (s *SPOESection) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
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

func (s *SPOESection) Result() ([]common.ReturnResultLine, error) {
	return nil, fmt.Errorf("not valid")
}
