/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package configs //nolint:testpackage

import (
	"slices"
	"strings"
	"testing"

	parser "github.com/haproxytech/config-parser/v5"
	"github.com/haproxytech/config-parser/v5/options"
)

func TestParseGetAttributesExistingOnly(t *testing.T) {
	tests := []struct {
		Name, Config string
	}{
		{"configBasic1", configBasic1},
	}
	for _, config := range tests {
		t.Run(config.Name, func(t *testing.T) {
			p, err := parser.New(options.String(config.Config))
			if err != nil {
				t.Fatalf(err.Error())
			}
			result, err := p.SectionAttributesGet(parser.Frontends, "http", true)
			if err != nil {
				t.Fatalf("err should be nil %v", err)
			}
			if 0 != slices.Compare(result, []string{
				"mode", "bind", "default_backend",
			}) {
				t.Fatalf("retrieved attributes do not match %s", strings.Join(result, ", "))
			}
		})
	}
}

func TestParseGetAttributesComplete(t *testing.T) {
	tests := []struct {
		Name, Config string
	}{
		{"configBasic1", configBasic1},
	}
	for _, config := range tests {
		t.Run(config.Name, func(t *testing.T) {
			p, err := parser.New(options.String(config.Config))
			if err != nil {
				t.Fatalf(err.Error())
			}
			result, err := p.SectionAttributesGet(parser.Backends, "default_backend", false)
			if err != nil {
				t.Fatalf("err should be nil %v", err)
			}
			if !slices.Contains(result, "option tcp-smart-connect") {
				t.Fatalf("should contain registered attribute 'option tcp-smart-connect'")
			}
		})
	}
}
