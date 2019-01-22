package global

import (
	"fmt"
	"strconv"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type NbThread struct {
	Enabled bool
	Value   int64
	Comment string
}

func (n *NbThread) Init() {
	n.Enabled = false
}

func (n *NbThread) GetParserName() string {
	return "nbthread"
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
			n.Enabled = true
			n.Value = num
			n.Comment = comment
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "nbthread", Line: line}
}

func (n *NbThread) Valid() bool {
	if n.Enabled {
		return true
	}
	return false
}

func (n *NbThread) Result(AddComments bool) []common.ReturnResultLine {
	if n.Enabled {
		return []common.ReturnResultLine{
			common.ReturnResultLine{
				Data:    fmt.Sprintf("nbthread %d", n.Value),
				Comment: n.Comment,
			},
		}
	}
	return []common.ReturnResultLine{}
}
