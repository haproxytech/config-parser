package stats

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/bind-options"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/helpers"
)

type Socket struct {
	Enabled    bool
	Path       string //can be address:port
	Params     []bindoptions.BindOption
	Name       string
	SearchName string
}

func (s *Socket) Init() {
	s.Enabled = false
	s.SearchName = "stats socket"
}

func (s *Socket) GetParserName() string {
	return s.SearchName
}

func (s *Socket) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, s.SearchName) {
		elements := helpers.StringSplitIgnoreEmpty(line, ' ')
		if len(elements) < 3 {
			return "", &errors.ParseError{Parser: "Socket", Line: line, Message: "Parse error"}
		}
		s.Enabled = true
		s.Path = elements[2]
		s.Params = bindoptions.Parse(elements[3:])
		//s.value = elements[1:]
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.SearchName, Line: line}
}

func (s *Socket) Valid() bool {
	if s.Enabled {
		return true
	}
	return false
}

func (s *Socket) String() []string {
	if s.Enabled {
		return []string{fmt.Sprintf("  %s %s %s", s.SearchName, s.Path, bindoptions.String(s.Params))}
	}
	return []string{}
}
