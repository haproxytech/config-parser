package tests

import (
	"testing"

	"github.com/haproxytech/config-parser/parsers"
)

func TestBalanceNormal1(t *testing.T) {
	err := ProcessLine("balance roundrobin", &parsers.Balance{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestBalanceNormal2(t *testing.T) {
	err := ProcessLine("balance static-rr # some comment after", &parsers.Balance{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestBalanceNormal3(t *testing.T) {
	err := ProcessLine("balance url_param userid", &parsers.Balance{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestBalanceMissing(t *testing.T) {
	err := ProcessLine("balance url_param", &parsers.Balance{})
	if err == nil {
		t.Errorf(err.Error())
	}
}

func TestBalanceErr(t *testing.T) {
	err := ProcessLine("balance nonexistent", &parsers.NameserverLines{})
	if err == nil {
		t.Errorf(err.Error())
	}
}
