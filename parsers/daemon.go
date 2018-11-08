package parsers

import (
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type Daemon struct {
	Enabled bool
}

func (d *Daemon) Init() {
	d.Enabled = false
}

func (d *Daemon) GetParserName() string {
	return "daemon"
}

func (d *Daemon) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, "daemon") {
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

func (d *Daemon) String() []string {
	if d.Enabled {
		return []string{"  daemon"}
	}
	return []string{}
}
