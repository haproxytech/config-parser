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
	"testing"

	parser "github.com/haproxytech/config-parser/v4"
	"github.com/haproxytech/config-parser/v4/options"
	"github.com/haproxytech/config-parser/v4/types"
)

func TestParseFecthCommentLines(t *testing.T) {
	tests := []struct {
		Name, Config string
	}{
		{"configBasicWithComments", configBasicWithComments},
	}
	for _, config := range tests {
		t.Run(config.Name, func(t *testing.T) {
			p, err := parser.New(options.String(config.Config))
			if err != nil {
				t.Fatalf(err.Error())
			}
			result := p.String()
			if result != config.Config {
				compare(t, config.Config, result)
				t.Fatalf("configurations does not match")
			}

			data, err := p.GetPreComments(parser.Defaults, parser.DefaultSectionName, "log")
			if err != nil {
				t.Fatalf("err should be nil %v", err)
			}
			if len(data) != 1 {
				t.Fatalf("number of log lines should be 1 but its %d", len(data))
			}
			if data[0] != "line comment 1" {
				t.Fatalf("comment should be 'line comment 1' but its %s", data[0])
			}

			data, err = p.GetPreComments(parser.Frontends, "http", "bind")
			if err != nil {
				t.Fatalf("err should be nil %v", err)
			}
			if len(data) != 1 {
				t.Fatalf("number of log lines should be 1 but its %d", len(data))
			}
			if data[0] != "line comment 2" {
				t.Fatalf("comment should be 'line comment 2' but its %s", data[0])
			}

			data, err = p.GetPreComments(parser.Backends, "default_backend", "mode")
			if err != nil {
				t.Fatalf("err should be nil %v", err)
			}
			if len(data) != 1 {
				t.Fatalf("number of log lines should be 1 but its %d", len(data))
			}
			if data[0] != "line comment 3" {
				t.Fatalf("comment should be 'line comment 3' but its %s", data[0])
			}
		})
	}
}

func TestParseFecthCommentLinesWrite(t *testing.T) {
	tests := []struct {
		Name, Config, EndConfig string
	}{
		{"configBasic1", configBasic1, configBasicWithLineComments},
	}
	for _, config := range tests {
		t.Run(config.Name, func(t *testing.T) {
			p, err := parser.New(options.String(config.Config))
			if err != nil {
				t.Fatalf(err.Error())
			}
			result := p.String()
			if result != config.Config {
				compare(t, config.Config, result)
				t.Fatalf("configurations does not match")
			}

			err = p.SetPreComments(parser.Defaults, parser.DefaultSectionName, "log", []string{"line comment 1"})
			if err != nil {
				t.Fatalf("err should be nil %v", err)
			}
			err = p.SetPreComments(parser.Frontends, "http", "bind", []string{"line comment 2"})
			if err != nil {
				t.Fatalf("err should be nil %v", err)
			}

			err = p.SetPreComments(parser.Backends, "default_backend", "mode", []string{"line comment 3"})
			if err != nil {
				t.Fatalf("err should be nil %v", err)
			}

			result = p.String()
			if result != config.EndConfig {
				compare(t, config.Config, result)
				t.Fatalf("configurations does not match")
			}
		})
	}
}

func TestParseFecthCommentInline(t *testing.T) {
	tests := []struct {
		Name, Config string
	}{
		{"configBasicWithComments", configBasicWithComments},
	}
	for _, config := range tests {
		t.Run(config.Name, func(t *testing.T) {
			p, err := parser.New(options.String(config.Config))
			if err != nil {
				t.Fatalf(err.Error())
			}
			result := p.String()
			if result != config.Config {
				compare(t, config.Config, result)
				t.Fatalf("configurations does not match")
			}

			rawData, err := p.Get(parser.Frontends, "http", "mode")
			mode, ok := rawData.(*types.StringC)
			if !ok {
				t.Fatalf("wrong type %v", rawData)
			}
			if mode.Comment != "inline comment #1" {
				t.Fatalf("comment should be 'line comment 1' but its %s", mode.Comment)
			}

			rawData, err = p.Get(parser.Frontends, "http", "bind")
			bindList, ok := rawData.([]types.Bind)
			if !ok {
				t.Fatalf("wrong type %v", rawData)
				t.Fatalf("wrong type %v", bindList)
			}
			if len(bindList) != 2 {
				t.Fatalf("nexpected length should be 2 but its %d", len(bindList))
			}
			if bindList[1].Comment != "inline comment #2" {
				t.Fatalf("comment should be 'inline comment #2' but its %s", bindList[1].Comment)
			}
		})
	}
}
