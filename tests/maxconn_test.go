package tests

import (
	"testing"

	"github.com/haproxytech/config-parser/parsers"
)

func TestMaxconnEmpty(t *testing.T) {
	err := ProcessLine("", &parsers.MaxConn{})
	if err == nil {
		t.Errorf("Does not throw error")
	}
}
func TestMaxconnNormal(t *testing.T) {
	err := ProcessLine("maxconn 3000", &parsers.MaxConn{})
	if err != nil {
		t.Errorf(err.Error())
	}
	err = ProcessLine("maxconn 3000 # some comment after", &parsers.MaxConn{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestMaxconnMissing(t *testing.T) {
	err := ProcessLine("maxconn", &parsers.MaxConn{})
	if err == nil {
		t.Errorf(err.Error())
	}
}

func TestMaxconnDifferent(t *testing.T) {
	err := ProcessLine("haproxy 3000", &parsers.NameserverLines{})
	if err == nil {
		t.Errorf(err.Error())
	}
}

func TestMaxconnErr(t *testing.T) {
	err := ProcessLine("mode ns1", &parsers.MaxConn{})
	if err == nil {
		t.Errorf("Does not thow error")
	}
}
