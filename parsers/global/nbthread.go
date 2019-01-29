package global

import (
	"fmt"
	"strconv"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type NbThread struct {
	data *types.Int64C
}

func (n *NbThread) Init() {
	n.data = nil
}

func (n *NbThread) GetParserName() string {
	return "nbthread"
}

func (n *NbThread) Get(createIfNotExist bool) (common.ParserData, error) {
	if n.data == nil {
		if createIfNotExist {
			n.data = &types.Int64C{}
			return n.data, nil
		}
		return nil, errors.FetchError
	}
	return n.data, nil
}

func (n *NbThread) Set(data common.ParserData) error {
	if data == nil {
		n.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.Int64C:
		n.data = newValue
	case types.Int64C:
		n.data = &newValue
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (n *NbThread) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := n.Get(false)
	n.Init()
	_, err := n.Parse(data, parts, []string{}, comment)
	if err != nil {
		n.Set(oldData)
	}
	return err
}

func (n *NbThread) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "nbthread" {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "NbThread", Line: line, Message: "Parse error"}
		}
		var num int64
		if num, err = strconv.ParseInt(parts[1], 10, 64); err != nil {
			return "", &errors.ParseError{Parser: "NbThread", Line: line, Message: err.Error()}
		} else {
			n.data = &types.Int64C{
				Value:   num,
				Comment: comment,
			}
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "nbthread", Line: line}
}

func (n *NbThread) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if n.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    fmt.Sprintf("nbthread %d", n.data.Value),
			Comment: n.data.Comment,
		},
	}, nil
}
