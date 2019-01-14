package tests

import (
	"testing"

	"github.com/haproxytech/config-parser/parsers"
)

func TestDaemonNormal(t *testing.T) {
	err := ProcessLine("daemon", &parsers.Daemon{})
	if err != nil {
		t.Errorf(err.Error())
	}

	err = ProcessLine("daemon # some comment", &parsers.Daemon{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
