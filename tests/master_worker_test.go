package tests

import (
	"testing"

	"github.com/haproxytech/config-parser/parsers"
)

func TestMasterWorkerNormal(t *testing.T) {
	err := ProcessLine("master-worker", &parsers.MasterWorker{})
	if err != nil {
		t.Errorf(err.Error())
	}
	err = ProcessLine("master-worker # some comment after", &parsers.MasterWorker{})
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestMasterWorkerDifferent(t *testing.T) {
	err := ProcessLine("haproxy", &parsers.NameserverLines{})
	if err == nil {
		t.Errorf(err.Error())
	}
}
