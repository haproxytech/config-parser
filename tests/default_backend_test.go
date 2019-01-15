package tests

import (
	"testing"

	"github.com/haproxytech/config-parser/parsers"
)

func TestDefaultBackendNormal1(t *testing.T) {
	err := ProcessLine("default_backend some_name", &parsers.DefaultBackend{})
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestDefaultBackendErr(t *testing.T) {
	err := ProcessLine("default_backend", &parsers.DefaultBackend{})
	if err == nil {
		t.Errorf(err.Error())
	}
}
