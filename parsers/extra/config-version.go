package extra

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type ConfigVersion struct {
	Name string
	data *types.ConfigVersion
}

func (p *ConfigVersion) Init() {
	p.Name = "# _version"
	p.data = nil
}

func (s *ConfigVersion) Get(createIfNotExist bool) (common.ParserData, error) {
	if s.data != nil {
		return s.data, nil
	} else if createIfNotExist {
		s.data = &types.ConfigVersion{
			Value: 1,
		}
		return s.data, nil
	}
	return nil, fmt.Errorf("No data")
}

//Parse see if we have version, since it is not haproxy keyword, it's in comments
func (s *ConfigVersion) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if strings.HasPrefix(comment, "_version") {
		data := common.StringSplitIgnoreEmpty(comment, '=')
		if len(data) < 2 {
			return "", &errors.ParseError{Parser: "ConfigVersion", Line: line}
		}
		if version, err := strconv.ParseInt(data[1], 10, 64); err == nil {
			s.data = &types.ConfigVersion{
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
