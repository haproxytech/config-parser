package option

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type OptionHttpchk struct {
	data *types.OptionHttpchk
}

func (s *OptionHttpchk) Init() {
	s.data = nil
}

func (s *OptionHttpchk) GetParserName() string {
	return "option httpchk"
}

func (s *OptionHttpchk) Get(createIfNotExist bool) (common.ParserData, error) {
	if s.data == nil {
		if createIfNotExist {
			s.data = &types.OptionHttpchk{}
			return s.data, nil
		}
		return nil, errors.FetchError
	}
	return s.data, nil
}

func (s *OptionHttpchk) Set(data common.ParserData) error {
	if data == nil {
		s.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.OptionHttpchk:
		s.data = newValue
	case types.OptionHttpchk:
		s.data = &newValue
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (s *OptionHttpchk) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := s.Get(false)
	s.Init()
	_, err := s.Parse(data, parts, []string{}, comment)
	if err != nil {
		s.Set(oldData)
	}
	return err
}

/*
option httpchk <uri>
option httpchk <method> <uri>
option httpchk <method> <uri> <version>
*/
func (s *OptionHttpchk) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 1 && parts[0] == "option" && parts[1] == "httpchk" {
		if len(parts) == 3 {
			s.data = &types.OptionHttpchk{
				Uri:     parts[2],
				Comment: comment,
			}
		} else if len(parts) == 4 {
			s.data = &types.OptionHttpchk{
				Method:  parts[2],
				Uri:     parts[3],
				Comment: comment,
			}
		} else if len(parts) == 5 {
			s.data = &types.OptionHttpchk{
				Method:  parts[2],
				Uri:     parts[3],
				Version: parts[4],
				Comment: comment,
			}
		} else {
			s.data = &types.OptionHttpchk{
				Method:  parts[2],
				Uri:     parts[3],
				Version: strings.Join(parts[4:], " "),
				Comment: comment,
			}
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "option httpchk", Line: line}
}

func (s *OptionHttpchk) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, errors.FetchError
	}
	var data string
	if s.data.Version != "" {
		data = fmt.Sprintf("option httpchk %s %s %s", s.data.Method, s.data.Uri, s.data.Version)
	} else if s.data.Method != "" {
		data = fmt.Sprintf("option httpchk %s %s", s.data.Method, s.data.Uri)
	} else {
		data = fmt.Sprintf("option httpchk %s", s.data.Uri)
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    data,
			Comment: s.data.Comment,
		},
	}, nil
}
