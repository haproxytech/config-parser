package extra

import "github.com/haproxytech/config-parser/common"

type UnProcessed struct {
	unProcessed []common.ReturnResultLine
}

func (u *UnProcessed) Init() {
	u.unProcessed = []common.ReturnResultLine{}
}

func (u *UnProcessed) GetParserName() string {
	return ""
}

func (u *UnProcessed) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	u.unProcessed = append(u.unProcessed, common.ReturnResultLine{
		Data: line, //do not save comments separatelly
	})
	return "", nil
}

func (u *UnProcessed) Valid() bool {
	if len(u.unProcessed) > 0 {
		return true
	}
	return false
}

func (u *UnProcessed) Result(AddComments bool) []common.ReturnResultLine {
	return u.unProcessed
}
