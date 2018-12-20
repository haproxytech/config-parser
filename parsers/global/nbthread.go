package global

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/helpers"
)

type NbThread struct {
	Enabled bool
	Value   int64
}

func (n *NbThread) Init() {
	n.Enabled = false
}

func (n *NbThread) GetParserName() string {
	return "nbthread"
}

func (n *NbThread) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, "nbthread") {
		parts := helpers.StringSplitIgnoreEmpty(line, ' ')
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "NbThread", Line: line, Message: "Parse error"}
		}
		var num int64
		if num, err = strconv.ParseInt(parts[1], 10, 64); err != nil {
			return "", &errors.ParseError{Parser: "NbThread", Line: line, Message: err.Error()}
		} else {
			n.Enabled = true
			n.Value = num
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

func (n *NbThread) String() []string {
	if n.Enabled {
		return []string{fmt.Sprintf("  nbthread %d", n.Value)}
	}
	return []string{}
}
