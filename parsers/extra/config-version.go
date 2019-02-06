package extra

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type ConfigVersion struct {
	data *types.Int64C
}

func (s *ConfigVersion) Init() {
	s.data = nil
}

func (s *ConfigVersion) GetParserName() string {
	return "# _version"
}

func (s *ConfigVersion) Get(createIfNotExist bool) (common.ParserData, error) {
	if s.data != nil {
		return s.data, nil
	} else if createIfNotExist {
		s.data = &types.Int64C{
			Value: 1,
		}
		return s.data, nil
	}
	return nil, fmt.Errorf("No data")
}

func (s *ConfigVersion) Set(data common.ParserData) error {
	if data == nil {
		s.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.Int64C:
		s.data = newValue
	case types.Int64C:
		s.data = &newValue
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (s *ConfigVersion) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	s.Init()
	_, err := s.Parse(data, parts, []string{}, comment)
	return err
}

//Parse see if we have version, since it is not haproxy keyword, it's in comments
func (s *ConfigVersion) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if strings.HasPrefix(comment, "_version") {
		data := common.StringSplitIgnoreEmpty(comment, '=')
		log.Println(data)
		if len(data) < 2 {
			return "", &errors.ParseError{Parser: "ConfigVersion", Line: line}
		}
		if version, err := strconv.ParseInt(data[1], 10, 64); err == nil {
			s.data = &types.Int64C{
				Value: version,
			}
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "ConfigVersion", Line: line}
}

func (s *ConfigVersion) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, errors.FetchError
	}

	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("# _version=%d", s.data.Value),
			Comment: "",
		},
	}, nil
}
