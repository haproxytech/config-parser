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

package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

type AliasTestData struct {
	Alias string
	Test  string
}

type Data struct { //nolint:maligned
	ParserMultiple     bool
	ParserSections     []string
	ParserName         string
	ParserSecondName   string
	StructName         string
	ParserType         string
	ParserTypeOverride string
	NoInit             bool
	NoName             bool
	NoParse            bool
	NoGet              bool
	NoSections         bool
	IsInterface        bool
	Dir                string
	ModeOther          bool
	TestOK             []string
	TestOKDefaults     []string
	TestOKFrontend     []string
	TestOKBackend      []string
	TestOKEscaped      []string
	TestFail           []string
	TestFailEscaped    []string
	TestAliasOK        []AliasTestData
	TestAliasFail      []AliasTestData
	TestSkip           bool
	DataDir            string
	Deprecated         bool
	HasAlias           bool
}

type ConfigFile struct {
	Section map[string][]string
	Tests   strings.Builder
}

func (c *ConfigFile) AddParserData(parser Data) { //nolint:gocognit,gocyclo,cyclop
	sections := parser.ParserSections
	testOK := parser.TestOK
	testOKDefaults := parser.TestOKDefaults
	testOKFrontend := parser.TestOKFrontend
	testOKBackend := parser.TestOKBackend
	TestOKEscaped := parser.TestOKEscaped
	if len(sections) == 0 && !parser.NoSections {
		log.Fatalf("parser %s does not have any section defined", parser.ParserName)
	}
	var lines []string
	var linesDefaults []string
	var linesFrontend []string
	var linesBackend []string
	var lines2 []string
	for _, section := range sections {
		_, ok := c.Section[section]
		if !ok {
			c.Section[section] = []string{}
		}
		// line = testOK[0]
		if parser.ParserMultiple {
			lines = testOK
			//nolint:gosimple
			for _, line := range testOK {
				c.Section[section] = append(c.Section[section], line)
			}
			// c.Section[s] = append(c.Section[s], testOK...)
		} else {
			lines = []string{testOK[0]}
			c.Section[section] = append(c.Section[section], testOK[0])
		}
		if section == "defaults" {
			if parser.ParserMultiple {
				linesDefaults = testOKDefaults
				//nolint:gosimple
				for _, line := range testOKDefaults {
					c.Section[section] = append(c.Section[section], line)
				}
			} else if len(testOKDefaults) > 0 {
				linesDefaults = []string{testOKDefaults[0]}
				c.Section[section] = append(c.Section[section], testOKDefaults[0])
			}
		}
		if section == "frontend" {
			if parser.ParserMultiple {
				linesFrontend = testOKFrontend
				//nolint:gosimple
				for _, line := range testOKFrontend {
					c.Section[section] = append(c.Section[section], line)
				}
			} else if len(testOKFrontend) > 0 {
				linesFrontend = []string{testOKFrontend[0]}
				c.Section[section] = append(c.Section[section], testOKFrontend[0])
			}
		}
		if section == "backend" {
			if parser.ParserMultiple {
				linesBackend = testOKBackend
				//nolint:gosimple
				for _, line := range testOKBackend {
					c.Section[section] = append(c.Section[section], line)
				}
			} else if len(testOKBackend) > 0 {
				linesBackend = []string{testOKBackend[0]}
				c.Section[section] = append(c.Section[section], testOKBackend[0])
			}
		}
		if parser.ParserMultiple {
			lines2 = TestOKEscaped
			//nolint:gosimple
			for _, line := range TestOKEscaped {
				c.Section[section] = append(c.Section[section], line)
			}
			// c.Section[s] = append(c.Section[s], TestOKEscaped...)
		} else if len(TestOKEscaped) > 0 {
			lines2 = []string{TestOKEscaped[0]}
			c.Section[section] = append(c.Section[section], TestOKEscaped[0])
		}
	}
	if len(lines) == 0 && len(lines2) == 0 {
		if parser.NoSections {
			return
		} else if !parser.Deprecated {
			log.Fatalf("parser %s does not have any tests defined", parser.ParserName)
		}
	}
	if !parser.NoSections {
		for _, line := range lines {
			c.Tests.WriteString(fmt.Sprintf("  {`  %s\n`, %d},\n", line, len(sections)))
		}
		for _, line := range linesDefaults {
			c.Tests.WriteString(fmt.Sprintf("  {`  %s\n`, 1},\n", line))
		}
		for _, line := range linesFrontend {
			c.Tests.WriteString(fmt.Sprintf("  {`  %s\n`, 1},\n", line))
		}
		for _, line := range linesBackend {
			c.Tests.WriteString(fmt.Sprintf("  {`  %s\n`, 1},\n", line))
		}
	}
}

func (c *ConfigFile) String() string {
	var result strings.Builder

	result.WriteString(license)
	result.WriteString("package configs\n\n")
	result.WriteString("const generatedConfig = `# _version=1\n# HAProxy Technologies\n# https://www.haproxy.com/\n# sections are in alphabetical order (except global & default) for code generation\n\n")

	first := true
	sectionNames := make([]string, len(c.Section)-2)
	index := 0
	for sectionName := range c.Section {
		if sectionName == "global" || sectionName == "defaults" {
			continue
		}
		sectionNames[index] = sectionName
		index++
	}
	sort.Strings(sectionNames)

	writeSection := func(sectionName string) {
		if !first {
			result.WriteString("\n")
		} else {
			first = false
		}
		result.WriteString(sectionName)
		result.WriteString(" test\n")
		lines := c.Section[sectionName]
		for _, line := range lines {
			result.WriteString("  ")
			result.WriteString(line)
			result.WriteString("\n")
		}
	}

	writeSection("global")
	writeSection("defaults")
	for _, sectionName := range sectionNames {
		writeSection(sectionName)
	}
	result.WriteString("`\n\n")

	result.WriteString("var configTests = []configTest{")
	result.WriteString(c.Tests.String())
	result.WriteString("}")

	result.WriteString("\n")
	return result.String()
}
