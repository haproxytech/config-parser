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
package configs

import (
	"bytes"
	"strings"
	"testing"

	parser "github.com/haproxytech/config-parser/v4"
	"github.com/haproxytech/config-parser/v4/options"
)

func TestWholeConfigs(t *testing.T) {
	tests := []struct {
		Name, Config string
	}{
		{"configBasic1", configBasic1},
		{"configBasic2", configBasic2},
		{"configFull", configFull},
	}
	for _, config := range tests {
		t.Run(config.Name, func(t *testing.T) {
			var buffer bytes.Buffer
			buffer.WriteString(config.Config)
			p, err := parser.New(options.Reader(&buffer))
			if err != nil {
				t.Fatalf(err.Error())
			}
			result := p.String()
			if result != config.Config {
				compare(t, config.Config, result)
				t.Fatalf("configurations does not match")
			}
		})
	}
}

func compare(t *testing.T, configOriginal, configResult string) {
	original := strings.Split(configOriginal, "\n")
	result := strings.Split(configResult, "\n")
	if len(original) != len(result) {
		t.Logf("not the same size: original: %d, result: %d", len(original), len(result))
		return
	}
	for index, line := range original {
		if line != result[index] {
			t.Logf("line %d: '%s' != '%s'", index+3, line, result[index])
		}
	}
}

func TestGeneratedConfig(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString(generatedConfig)
	p, err := parser.New(options.DisableUnProcessed, options.Reader(&buffer))
	if err != nil {
		t.Fatalf(err.Error())
	}
	result := p.String()
	for _, configLine := range configTests {
		count := strings.Count(result, configLine.Line)
		if count != configLine.Count {
			t.Fatalf("line '%s' found %d times, expected %d times", configLine.Line, count, configLine.Count)
		}
	}
}

func TestHashConfig(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString(configBasicHash)
	p, err := parser.New(options.UseMd5Hash, options.Reader(&buffer))
	if err != nil {
		t.Fatalf(err.Error())
	}
	result, err := p.StringWithHash()
	if err != nil {
		t.Fatalf(err.Error())
	}
	if result != configBasicHash {
		compare(t, configBasicHash, result)
		t.Fatalf("configurations does not match")
	}
}

func TestConfigUseV2HTTPCheck(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString(configBasicUseV2HTTPCheck)
	p, err := parser.New(options.UseV2HTTPCheck, options.Reader(&buffer))
	if err != nil {
		t.Fatalf(err.Error())
	}
	result := p.String()
	if result != configBasicUseV2HTTPCheck {
		compare(t, configBasicUseV2HTTPCheck, result)
		t.Fatalf("configurations does not match")
	}
}
