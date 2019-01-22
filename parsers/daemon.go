package parsers

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type Daemon struct {
	Enabled bool
	Comment string
}

func (d *Daemon) Init() {
	d.Enabled = false
}

func (d *Daemon) GetParserName() string {
	return "daemon"
}

func (d *Daemon) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "daemon" {
		d.Comment = comment
		d.Enabled = true
		return "", nil
	}
	return "", &errors.ParseError{Parser: "Daemon", Line: line}
}

func (d *Daemon) Valid() bool {
	if d.Enabled {
		return true
	}
	return false
}

func (d *Daemon) Result(AddComments bool) []common.ReturnResultLine {
	if d.Enabled {
		return []common.ReturnResultLine{
			common.ReturnResultLine{
				Data: "daemon",
			},
		}
	}
	return []common.ReturnResultLine{}
}
