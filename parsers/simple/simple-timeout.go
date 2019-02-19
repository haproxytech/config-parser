package simple

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type SimpleTimeout struct {
	Name string
	name string
	data *types.SimpleTimeout
}

func (t *SimpleTimeout) Init() {
	if !strings.HasPrefix(t.Name, "timeout") {
		t.name = t.Name
		t.Name = fmt.Sprintf("timeout %s", t.Name)
	}
	t.data = nil
}

func (t *SimpleTimeout) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 2 && parts[0] == "timeout" && parts[1] == t.name {
		t.data = &types.SimpleTimeout{
			Value:   parts[2],
			Comment: comment,
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: fmt.Sprintf("timeout %s", t.name), Line: line}
}

func (t *SimpleTimeout) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if t.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("timeout %s %s", t.name, t.data.Value),
			Comment: t.data.Comment,
		},
	}, nil
}
