package extra

import (
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type UnProcessed struct {
	Name string
	data []types.UnProcessed
}

func (u *UnProcessed) Init() {
	u.Name = ""
	u.data = []types.UnProcessed{}
}

func (p *UnProcessed) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	p.data = append(p.data, types.UnProcessed{
		Value: strings.TrimSpace(line),
	})
	return "", nil
}

func (u *UnProcessed) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(u.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(u.data))
	for index, d := range u.data {
		result[index] = common.ReturnResultLine{
			Data:    d.Value,
			Comment: "",
		}
	}
	return result, nil
}
