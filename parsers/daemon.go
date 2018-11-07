package parsers

import (
	"strings"

	"config-parser/errors"
)

type Daemon struct {
	enabled bool
}

func (d *Daemon) Init() {
	d.enabled = false
}

func (d *Daemon) GetParserName() string {
	return "daemon"
}

func (d *Daemon) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, "daemon") {
		d.enabled = true
		return "", nil
	}
	return "", &errors.ParseError{Parser: "Daemon", Line: line}
}

func (d *Daemon) Valid() bool {
	if d.enabled {
		return true
	}
	return false
}

func (d *Daemon) String() []string {
	if d.enabled {
		return []string{"  daemon"}
	}
	return []string{}
}
