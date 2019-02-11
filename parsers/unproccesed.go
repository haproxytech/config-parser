package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type UnProcessed struct {
	data []common.ReturnResultLine
}

func (u *UnProcessed) Init() {
	u.data = []common.ReturnResultLine{}
}

func (u *UnProcessed) GetParserName() string {
	return ""
}

func (u *UnProcessed) Get(createIfNotExist bool) (common.ParserData, error) {
	return u.data, nil
}

func (p *UnProcessed) GetOne(index int) (common.ParserData, error) {
	if len(p.data) == 0 {
		return nil, errors.FetchError
	}
	if index < 0 || index >= len(p.data) {
		return nil, errors.FetchError
	}
	return p.data[index], nil
}

func (u *UnProcessed) Set(data common.ParserData, index int) error {
	newData, ok := data.(common.ReturnResultLine)
	if ok {
		u.data = append(u.data, newData)
		return nil
	}
	return fmt.Errorf("casting error")
}

func (u *UnProcessed) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	u.Init()
	_, err := u.Parse(data, parts, []string{}, comment)
	return err
}

func (u *UnProcessed) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	u.data = append(u.data, common.ReturnResultLine{
		Data: line, //do not save comments separatelly
	})
	return "", nil
}

func (u *UnProcessed) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(u.data) == 0 {
		return nil, errors.FetchError
	}
	return u.data, nil
}
