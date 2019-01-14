package tests

import (
	"testing"

	"github.com/haproxytech/config-parser/parsers"
)

func TestModeNormal(t *testing.T) {
	err := ProcessLine("mode http", &parsers.Mode{})
	if err != nil {
		t.Errorf(err.Error())
	}

	err = ProcessLine("mode tcp", &parsers.Mode{})
	if err != nil {
		t.Errorf(err.Error())
	}

	err = ProcessLine("mode health", &parsers.Mode{})
	if err != nil {
		t.Errorf(err.Error())
	}

	err = ProcessLine("mode health # some comment after", &parsers.Mode{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestModeMissing(t *testing.T) {
	err := ProcessLine("mode", &parsers.Mode{})
	if err == nil {
		t.Errorf(err.Error())
	}
}

func TestModeDifferent(t *testing.T) {
	err := ProcessLine("haproxy http", &parsers.NameserverLines{})
	if err == nil {
		t.Errorf(err.Error())
	}
}

func TestModeErr(t *testing.T) {
	err := ProcessLine("mode ns1", &parsers.Mode{})
	if err == nil {
		t.Errorf("Does not thow error")
	}
}
