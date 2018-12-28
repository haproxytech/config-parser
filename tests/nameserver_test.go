package tests

import (
	"testing"

	"github.com/haproxytech/config-parser/parsers"
)

func TestNameserverEmpty(t *testing.T) {
	err := ProcessLine("", &parsers.NameserverLines{})
	if err == nil {
		t.Errorf("Does not throw error")
	}
}
func TestNameserverNormal(t *testing.T) {
	err := ProcessLine("nameserver ns1 127.0.0.1", &parsers.NameserverLines{})
	if err != nil {
		t.Errorf(err.Error())
	}
	err = ProcessLine("nameserver ns1 127.0.0.1 #some comment after", &parsers.NameserverLines{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestNameserverMissing(t *testing.T) {
	err := ProcessLine("nameserver ns1", &parsers.NameserverLines{})
	if err == nil {
		t.Errorf(err.Error())
	}
}

func TestNameserverDifferent(t *testing.T) {
	err := ProcessLine("haproxy ns1", &parsers.NameserverLines{})
	if err == nil {
		t.Errorf(err.Error())
	}
}
