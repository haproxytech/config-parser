package extra

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type UnProcessed struct {
	unProcessed []common.ReturnResultLine
}

func (u *UnProcessed) Init() {
	u.unProcessed = []common.ReturnResultLine{}
}

func (u *UnProcessed) Clear() {
	u.Init()
}

func (u *UnProcessed) GetParserName() string {
	return ""
}

func (u *UnProcessed) Get(createIfNotExist bool) (common.ParserData, error) {
	return u.unProcessed, nil
}

func (u *UnProcessed) Set(data common.ParserData) error {
	newData, ok := data.(common.ReturnResultLine)
	if ok {
		u.unProcessed = append(u.unProcessed, newData)
		return nil
	}
	return fmt.Errorf("casting error")
}

func (u *UnProcessed) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	u.Clear()
	_, err := u.Parse(data, parts, []string{}, comment)
	return err
}

func (u *UnProcessed) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	u.unProcessed = append(u.unProcessed, common.ReturnResultLine{
		Data: line, //do not save comments separatelly
	})
	return "", nil
}

func (u *UnProcessed) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(u.unProcessed) == 0 {
		return nil, &errors.FetchError{}
	}
	return u.unProcessed, nil
}
