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

// run this as go run go-generate.go $(pwd)

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/google/renameio"
	"github.com/haproxytech/config-parser/v4/common"
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

func (c *ConfigFile) AddParserData(parser Data) { //nolint:gocognit
	sections := parser.ParserSections
	testOK := parser.TestOK
	TestOKEscaped := parser.TestOKEscaped
	if len(sections) == 0 && !parser.NoSections {
		log.Fatalf("parser %s does not have any section defined", parser.ParserName)
	}
	var lines []string
	var lines2 []string
	for _, s := range sections {
		_, ok := c.Section[s]
		if !ok {
			c.Section[s] = []string{}
		}
		// line = testOK[0]
		if parser.ParserMultiple {
			lines = testOK
			//nolint:gosimple
			for _, line := range testOK {
				c.Section[s] = append(c.Section[s], line)
			}
			// c.Section[s] = append(c.Section[s], testOK...)
		} else {
			lines = []string{testOK[0]}
			c.Section[s] = append(c.Section[s], testOK[0])
		}
		if parser.ParserMultiple {
			lines2 = TestOKEscaped
			//nolint:gosimple
			for _, line := range TestOKEscaped {
				c.Section[s] = append(c.Section[s], line)
			}
			// c.Section[s] = append(c.Section[s], TestOKEscaped...)
		} else if len(TestOKEscaped) > 0 {
			lines2 = []string{TestOKEscaped[0]}
			c.Section[s] = append(c.Section[s], TestOKEscaped[0])
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

//nolint:gochecknoglobals
var configFile = ConfigFile{}

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasSuffix(dir, "generate") {
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]) + "/..")
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	log.Println(dir)

	configFile.Section = map[string][]string{}

	generateTypes(dir, "")
	generateTypesGeneric(dir)
	generateTypesOther(dir)
	// spoe
	generateTypes(dir, "spoe/")

	filePath := path.Join(dir, "tests", "configs", "haproxy_generated.cfg.go")
	err = renameio.WriteFile(filePath, []byte(configFile.String()), 0644)
	if err != nil {
		log.Println(err)
	}
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		log.Println("File " + filePath + " already exists")
	}
	return !os.IsNotExist(err)
}

func generateTypesOther(dir string) { //nolint:gocognit,gocyclo
	dat, err := ioutil.ReadFile("types/types-other.go")
	if err != nil {
		log.Println(err)
	}
	lines := common.StringSplitIgnoreEmpty(string(dat), '\n')

	parserData := Data{}
	for _, line := range lines {
		if strings.HasPrefix(line, "//sections:") {
			s := strings.Split(line, ":")
			parserData.ParserSections = strings.Split(s[1], ",")
		}
		if strings.HasPrefix(line, "//no-sections:true") {
			parserData.NoSections = true
		}
		if strings.HasPrefix(line, "//name:") {
			data := common.StringSplitIgnoreEmpty(line, ':')
			items := common.StringSplitIgnoreEmpty(data[1], ' ')
			parserData.ParserName = data[1]
			if len(items) > 1 {
				parserData.ParserName = items[0]
				parserData.ParserSecondName = items[1]
			}
		}
		if strings.HasPrefix(line, "//struct-name:") {
			data := common.StringSplitIgnoreEmpty(line, ':')
			parserData.StructName = data[1]
		}
		if strings.HasPrefix(line, "//dir:") {
			data := common.StringSplitIgnoreEmpty(line, ':')
			parserData.Dir = data[1]
		}
		if strings.HasPrefix(line, "//is-multiple:true") {
			parserData.ParserMultiple = true
		}
		if strings.HasPrefix(line, "//no-init:true") {
			parserData.NoInit = true
		}
		if strings.HasPrefix(line, "//no-name:true") {
			parserData.NoName = true
		}
		if strings.HasPrefix(line, "//no-parse:true") {
			parserData.NoParse = true
		}
		if strings.HasPrefix(line, "//is-interface:true") {
			parserData.IsInterface = true
		}
		if strings.HasPrefix(line, "//parser-type:") {
			data := common.StringSplitIgnoreEmpty(line, ':')
			parserData.ParserTypeOverride = data[1]
		}
		if strings.HasPrefix(line, "//no-get:true") {
			parserData.NoGet = true
		}
		if strings.HasPrefix(line, `//test:"ok"`) {
			data := strings.SplitN(line, ":", 3)
			parserData.TestOKEscaped = append(parserData.TestOKEscaped, data[2])
		}
		if strings.HasPrefix(line, "//test:ok") {
			data := strings.SplitN(line, ":", 3)
			parserData.TestOK = append(parserData.TestOK, data[2])
		}
		if strings.HasPrefix(line, `//test:"fail"`) {
			data := strings.SplitN(line, ":", 3)
			parserData.TestFailEscaped = append(parserData.TestFailEscaped, data[2])
		}
		if strings.HasPrefix(line, "//test:fail") {
			data := strings.SplitN(line, ":", 3)
			parserData.TestFail = append(parserData.TestFail, data[2])
		}
		if strings.HasPrefix(line, "//test:alias") {
			data := strings.SplitN(line, ":", 5)
			aliasTestData := AliasTestData{
				Alias: data[2],
				Test:  data[4],
			}
			switch data[3] {
			case "ok":
				parserData.TestAliasOK = append(parserData.TestAliasOK, aliasTestData)
			case "fail":
				parserData.TestAliasFail = append(parserData.TestAliasFail, aliasTestData)
			default:
				log.Fatalf("not able to process line %s", line)
			}
		}
		if strings.HasPrefix(line, "//test:skip") {
			parserData.TestSkip = true
		}
		if strings.HasPrefix(line, "//has-alias:true") {
			parserData.HasAlias = true
		}

		if !strings.HasPrefix(line, "type ") {
			continue
		}

		if parserData.ParserName == "" {
			parserData = Data{}
			continue
		}
		data := common.StringSplitIgnoreEmpty(line, ' ')
		if parserData.StructName == "" {
			parserData.StructName = data[1]
		}
		parserData.ParserType = data[1]
		if parserData.ParserTypeOverride != "" {
			parserData.ParserType = parserData.ParserTypeOverride
		}

		filename := parserData.ParserName
		if parserData.ParserSecondName != "" {
			filename = fmt.Sprintf("%s %s", filename, parserData.ParserSecondName)
		}

		// Main parser file.
		filePath := path.Join(dir, "parsers", parserData.Dir, cleanFileName(filename)+".go")
		log.Println(filePath)

		if !fileExists(filePath) {
			f, err := os.Create(filePath)
			CheckErr(err)
			defer f.Close()

			parserData.ModeOther = true
			err = typeOthersAPITemplate.Execute(f, parserData)
			CheckErr(err)
		}

		// Generated parser file.
		filePath = path.Join(dir, "parsers", parserData.Dir, cleanFileName(filename)+"_generated.go")
		log.Println(filePath)
		f, err := os.Create(filePath)
		CheckErr(err)
		defer f.Close()

		parserData.ModeOther = true
		err = typeTemplate.Execute(f, parserData)
		CheckErr(err)

		if !parserData.TestSkip {
			parserData.TestFail = append(parserData.TestFail, "---")
			parserData.TestFail = append(parserData.TestFail, "--- ---")
			dataDir := ""

			filePath = path.Join(dir, dataDir, "tests", cleanFileName(filename)+"_generated_test.go")
			log.Println(filePath)
			f, err = os.Create(filePath)
			CheckErr(err)
			defer f.Close()

			err = testTemplate.Execute(f, parserData)
			CheckErr(err)
		}

		configFile.AddParserData(parserData)
		parserData = Data{}
	}
}

func generateTypesGeneric(dir string) { //nolint:gocognit
	dat, err := ioutil.ReadFile("types/types-generic.go")
	if err != nil {
		log.Println(err)
	}
	lines := common.StringSplitIgnoreEmpty(string(dat), '\n')

	parsers := map[string]*Data{}
	parserData := &Data{}
	for _, line := range lines {
		if strings.HasPrefix(line, "//sections:") {
			s := strings.Split(line, ":")
			parserData.ParserSections = strings.Split(s[1], ",")
		}
		if strings.HasPrefix(line, "//no-sections:true") {
			parserData.NoSections = true
		}
		if strings.HasPrefix(line, "//name:") {
			data := common.StringSplitIgnoreEmpty(line, ':')
			items := common.StringSplitIgnoreEmpty(data[1], ' ')
			parserData.ParserName = data[1]
			if len(items) > 1 {
				parserData.ParserName = items[0]
				parserData.ParserSecondName = items[1]
			}
		}
		if strings.HasPrefix(line, "//no-init:true") {
			parserData.NoInit = true
		}
		if strings.HasPrefix(line, "//no-name:true") {
			parserData.NoName = true
		}
		if strings.HasPrefix(line, "//no-parse:true") {
			parserData.NoParse = true
		}
		if strings.HasPrefix(line, `//test:"ok"`) {
			data := strings.SplitN(line, ":", 3)
			parserData.TestOKEscaped = append(parserData.TestOKEscaped, data[2])
		}
		if strings.HasPrefix(line, "//test:ok") {
			data := strings.SplitN(line, ":", 3)
			parserData.TestOK = append(parserData.TestOK, data[2])
		}
		if strings.HasPrefix(line, `//test:"fail"`) {
			data := strings.SplitN(line, ":", 3)
			parserData.TestFailEscaped = append(parserData.TestFailEscaped, data[2])
		}
		if strings.HasPrefix(line, "//test:fail") {
			data := strings.SplitN(line, ":", 3)
			parserData.TestFail = append(parserData.TestFail, data[2])
		}
		if strings.HasPrefix(line, "//test:alias") {
			data := strings.SplitN(line, ":", 5)
			aliasTestData := AliasTestData{
				Alias: data[2],
				Test:  data[4],
			}
			switch data[3] {
			case "ok":
				parserData.TestAliasOK = append(parserData.TestAliasOK, aliasTestData)
			case "fail":
				parserData.TestAliasFail = append(parserData.TestAliasFail, aliasTestData)
			default:
				log.Fatalf("not able to process line %s", line)
			}
		}
		if strings.HasPrefix(line, "//gen:") {
			data := common.StringSplitIgnoreEmpty(line, ':')
			parserData = &Data{}
			parserData.StructName = data[1]
			parsers[data[1]] = parserData
		}
		if strings.HasPrefix(line, "//has-alias:true") {
			parserData.HasAlias = true
		}

		if !strings.HasPrefix(line, "type ") {
			continue
		}

		if parserData.ParserName == "" {
			parserData = &Data{}
			continue
		}
		data := common.StringSplitIgnoreEmpty(line, ' ')
		parserType := data[1]

		for _, parserData := range parsers {
			parserData.ParserType = parserType

			filename := parserData.ParserName
			if parserData.ParserSecondName != "" {
				filename = fmt.Sprintf("%s %s", filename, parserData.ParserSecondName)
			}

			filePath := path.Join(dir, "parsers", cleanFileName(filename)+"_generated.go")
			log.Println(filePath)
			f, err := os.Create(filePath)
			CheckErr(err)
			defer f.Close()

			err = typeTemplate.Execute(f, parserData)
			CheckErr(err)

			parserData.TestFail = append(parserData.TestFail, "---")
			parserData.TestFail = append(parserData.TestFail, "--- ---")

			filePath = path.Join(dir, "tests", cleanFileName(filename)+"_generated_test.go")
			log.Println(filePath)
			f, err = os.Create(filePath)
			CheckErr(err)
			defer f.Close()

			err = testTemplate.Execute(f, parserData)
			CheckErr(err)
		}
		// configFile.AddParserData(parserData)
		parsers = map[string]*Data{}
		parserData = &Data{}
	}
}

func generateTypes(dir string, dataDir string) { //nolint:gocognit
	dat, err := ioutil.ReadFile(dataDir + "types/types.go")
	if err != nil {
		log.Println(err)
	}
	lines := common.StringSplitIgnoreEmpty(string(dat), '\n')

	parserData := Data{}

	for _, line := range lines {
		parserData.DataDir = dataDir
		if strings.HasPrefix(line, "//deprecated:") {
			parserData.Deprecated = true
		}
		if strings.HasPrefix(line, "//sections:") {
			s := strings.Split(line, ":")
			parserData.ParserSections = strings.Split(s[1], ",")
		}
		if strings.HasPrefix(line, "//no-sections:true") {
			parserData.NoSections = true
		}
		if strings.HasPrefix(line, "//name:") {
			data := common.StringSplitIgnoreEmpty(line, ':')
			items := common.StringSplitIgnoreEmpty(data[1], ' ')
			parserData.ParserName = data[1]
			if len(items) > 1 {
				parserData.ParserName = items[0]
				parserData.ParserSecondName = items[1]
			}
		}
		if strings.HasPrefix(line, "//is-multiple:true") {
			parserData.ParserMultiple = true
		}
		if strings.HasPrefix(line, "//no-init:true") {
			parserData.NoInit = true
		}
		if strings.HasPrefix(line, "//no-name:true") {
			parserData.NoName = true
		}
		if strings.HasPrefix(line, "//no-parse:true") {
			parserData.NoParse = true
		}
		if strings.HasPrefix(line, `//test:"ok"`) && !parserData.Deprecated {
			data := strings.SplitN(line, ":", 3)
			parserData.TestOKEscaped = append(parserData.TestOKEscaped, data[2])
		}
		if strings.HasPrefix(line, "//test:ok") && !parserData.Deprecated {
			data := strings.SplitN(line, ":", 3)
			parserData.TestOK = append(parserData.TestOK, data[2])
		}
		if strings.HasPrefix(line, `//test:"fail"`) && !parserData.Deprecated {
			data := strings.SplitN(line, ":", 3)
			parserData.TestFailEscaped = append(parserData.TestFailEscaped, data[2])
		}
		if strings.HasPrefix(line, "//test:fail") && !parserData.Deprecated {
			data := strings.SplitN(line, ":", 3)
			parserData.TestFail = append(parserData.TestFail, data[2])
		}
		if strings.HasPrefix(line, "//test:alias") {
			data := strings.SplitN(line, ":", 5)
			aliasTestData := AliasTestData{
				Alias: data[2],
				Test:  data[4],
			}
			switch data[3] {
			case "ok":
				parserData.TestAliasOK = append(parserData.TestAliasOK, aliasTestData)
			case "fail":
				parserData.TestAliasFail = append(parserData.TestAliasFail, aliasTestData)
			default:
				log.Fatalf("not able to process line %s", line)
			}
		}
		if strings.HasPrefix(line, "//has-alias:true") {
			parserData.HasAlias = true
		}

		if !strings.HasPrefix(line, "type ") {
			continue
		}

		if parserData.ParserName == "" {
			parserData = Data{}
			continue
		}
		data := common.StringSplitIgnoreEmpty(line, ' ')
		parserData.StructName = data[1]
		parserData.ParserType = data[1]

		filename := parserData.ParserName
		if parserData.ParserSecondName != "" {
			filename = fmt.Sprintf("%s %s", filename, parserData.ParserSecondName)
		}

		filePath := path.Join(dir, dataDir, "parsers", cleanFileName(filename)+"_generated.go")
		log.Println(filePath)
		f, err := os.Create(filePath)
		CheckErr(err)
		defer f.Close()

		err = typeTemplate.Execute(f, parserData)
		CheckErr(err)

		// parserData.TestFail = append(parserData.TestFail, "") parsers should not get empty line!
		parserData.TestFail = append(parserData.TestFail, "---")
		parserData.TestFail = append(parserData.TestFail, "--- ---")

		filePath = path.Join(dir, dataDir, "tests", cleanFileName(filename)+"_generated_test.go")
		log.Println(filePath)
		f, err = os.Create(filePath)
		CheckErr(err)
		defer f.Close()

		err = testTemplate.Execute(f, parserData)
		CheckErr(err)

		configFile.AddParserData(parserData)
		parserData = Data{}
	}
}

func cleanFileName(filename string) string {
	return strings.ReplaceAll(filename, " ", "-")
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const license = `// Code generated by go generate; DO NOT EDIT.
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
`

//nolint:gochecknoglobals
var typeOthersAPITemplate = template.Must(template.New("").Parse(license +
	`{{- if .ModeOther}}
package {{ .Dir }}
{{- else }}
package parsers
{{- end }}

import (
	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/{{ .DataDir }}types"
)

type {{ .StructName }} struct {
	Name string
	// Mode string
	data        []types.{{ .ParserType }}
	preComments []string // comments that appear before the the actual line
}

{{- if not .NoInit }}

func (p *{{ .StructName }}) Init() {
{{- if .ParserMultiple }}
        p.Name = "{{ .ParserName }}"
        p.data = []types.{{ .ParserType }}{}
{{- else }}
        p.data = nil
{{- end }}
		p.preComments = nil
        // Following line forces compilation to fail:
        Function not implemented!
}
{{- end }}

func (h *{{ .StructName }} ) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
        return "", &errors.ParseError{Parser: parseErrorLines("{{ .ParserName }}"), Line: line}

        // Following line forces compilation to fail:
        Function not implemented!
}

func (h *{{ .StructName }} ) Result() ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.ErrFetch
	}

	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = common.ReturnResultLine{
			Data:   "{{ .ParserName }}" + req.String(),
			Comment: req.GetComment(),
		}
	}
	return result, nil
}

func parseErrorLines(s string) string {
	var r string
	parts := strings.Split(s, "-")
	if len(parts) == 1 {
		r = strings.Title(parts[0])
	} else {
		r = strings.Title(parts[0]) + strings.Title(parts[1])
	}
	return r + "Lines"
}
`))

//nolint:gochecknoglobals
var typeTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
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

{{- if .ModeOther}}
package {{ .Dir }}
{{- else }}
package parsers
{{- end }}

import (
	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/{{ .DataDir }}types"
)

{{- if not .NoInit }}

func (p *{{ .StructName }}) Init() {
{{- if .ParserMultiple }}
	p.data = []types.{{ .ParserType }}{}
{{- else }}
    p.data = nil
{{- end }}
    p.preComments = []string{}
}
{{- end }}

{{- if not .NoName }}

func (p *{{ .StructName }}) GetParserName() string {
{{- if .HasAlias}}
    if p.Alias != "" {
	    return p.Alias
	}
{{- end }}
{{- if .ModeOther}}
    return p.Name
{{- else }}
{{- if eq .ParserSecondName "" }}
	return "{{ .ParserName }}"
{{- else }}
	return "{{ .ParserName }} {{ .ParserSecondName }}"
{{- end }}
{{- end }}
}
{{- end }}

{{- if not .NoGet }}

func (p *{{ .StructName }}) Get(createIfNotExist bool) (common.ParserData, error) {
{{- if .ParserMultiple }}
	if len(p.data) == 0 && !createIfNotExist {
		return nil, errors.ErrFetch
	}
{{- else }}
	if p.data == nil {
		if createIfNotExist {
			p.data = &types.{{ .ParserType }}{}
			return p.data, nil
		}
		return nil, errors.ErrFetch
	}
{{- end }}
	return p.data, nil
}
{{- end }}

func (p *{{ .StructName }}) GetOne(index int) (common.ParserData, error) {
{{- if .ParserMultiple }}
	if index < 0 || index >= len(p.data) {
		return nil, errors.ErrFetch
	}
	return p.data[index], nil
{{- else }}
	if index > 0 {
		return nil, errors.ErrFetch
	}
	if p.data == nil {
		return nil, errors.ErrFetch
	}
	return p.data, nil
{{- end }}
}

func (p *{{ .StructName }}) Delete(index int) error {
{{- if .ParserMultiple }}
	if index < 0 || index >= len(p.data) {
		return errors.ErrFetch
	}
	copy(p.data[index:], p.data[index+1:])
{{- if .IsInterface }}
	p.data[len(p.data)-1] = nil
{{- else }}
	p.data[len(p.data)-1] = types.{{ .ParserType }}{}
{{- end }}
	p.data = p.data[:len(p.data)-1]
	return nil
{{- else }}
	p.Init()
	return nil
{{- end }}
}

func (p *{{ .StructName }}) Insert(data common.ParserData, index int) error {
{{- if .ParserMultiple }}
	if data == nil {
		return errors.ErrInvalidData
	}
	switch newValue := data.(type) {
	case []types.{{ .ParserType }}:
		p.data = newValue
	case *types.{{ .ParserType }}:
		if index > -1 {
			if index > len(p.data) {
				return errors.ErrIndexOutOfRange
			}
{{- if .IsInterface }}
			p.data = append(p.data, nil)
{{- else }}
			p.data = append(p.data, types.{{ .ParserType }}{})
{{- end }}
			copy(p.data[index+1:], p.data[index:])
			p.data[index] = *newValue
		} else {
			p.data = append(p.data, *newValue)
		}
	case types.{{ .ParserType }}:
		if index > -1 {
			if index > len(p.data) {
				return errors.ErrIndexOutOfRange
			}
{{- if .IsInterface }}
			p.data = append(p.data, nil)
{{- else }}
			p.data = append(p.data, types.{{ .ParserType }}{})
{{- end }}
			copy(p.data[index+1:], p.data[index:])
			p.data[index] = newValue
		} else {
			p.data = append(p.data, newValue)
		}
	default:
		return errors.ErrInvalidData
	}
	return nil
{{- else }}
	return p.Set(data, index)
{{- end }}
}

func (p *{{ .StructName }}) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
{{- if .ParserMultiple }}
	switch newValue := data.(type) {
	case []types.{{ .ParserType }}:
		p.data = newValue
	case *types.{{ .ParserType }}:
		if index > -1 && index < len(p.data) {
			p.data[index] = *newValue
		} else if index == -1 {
			p.data = append(p.data, *newValue)
		} else {
			return errors.ErrIndexOutOfRange
		}
	case types.{{ .ParserType }}:
		if index > -1 && index < len(p.data) {
			p.data[index] = newValue
		} else if index == -1 {
			p.data = append(p.data, newValue)
		} else {
			return errors.ErrIndexOutOfRange
		}
	default:
		return errors.ErrInvalidData
	}
{{- else }}
	switch newValue := data.(type) {
	case *types.{{ .ParserType }}:
		p.data = newValue
	case types.{{ .ParserType }}:
		p.data = &newValue
	default:
		return errors.ErrInvalidData
	}
{{- end }}
	return nil
}

func (p *{{ .StructName }}) PreParse(line string, parts, previousParts []string, preComments []string, comment string) (changeState string, err error) {
	changeState, err = p.Parse(line, parts, previousParts, comment)
	if err == nil && preComments != nil {
		p.preComments = append(p.preComments, preComments...)
	}
	return changeState, err
}

{{- if and .ParserMultiple (not .NoParse) }}

func (p *{{ .StructName }}) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
{{- if eq .ParserSecondName "" }}
	if parts[0] == "{{ .ParserName }}" {
{{- else }}
	if len(parts) > 1 && parts[0] == "{{ .ParserName }}"  && parts[1] == "{{ .ParserSecondName }}" {
{{- end }}
		data, err := p.parse(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "{{ .StructName }}", Line: line}
		}
		p.data = append(p.data, *data)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "{{ .StructName }}", Line: line}
}
{{- end }}

func (p *{{ .StructName }}) ResultAll() ([]common.ReturnResultLine, []string, error) {
	res, err := p.Result()
	return res, p.preComments, err
}
`))

//nolint:gochecknoglobals
var testTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
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

package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/haproxytech/config-parser/v4/{{ .DataDir }}parsers{{- if .ModeOther}}/{{ .Dir }}{{- end }}"
)
{{ $StructName := .StructName }}
func Test{{ $StructName }}{{ .Dir }}(t *testing.T) {
	tests := map[string]bool{
	{{- range $index, $val := .TestOK}}
		"{{- $val -}}": true,
	{{- end }}
	{{- range $index, $val := .TestOKEscaped}}
		` + "`" + `{{- $val -}}` + "`" + `: true,
	{{- end }}
	{{- range $index, $val := .TestFail}}
		"{{- $val -}}": false,
	{{- end }}
	{{- range $index, $val := .TestFailEscaped}}
		` + "`" + `{{- $val -}}` + "`" + `: true,
	{{- end }}
	}
	parser := {{- if .ModeOther}} &{{ .Dir }}{{- else }} &parsers{{- end }}.{{ $StructName }}{}
	for command, shouldPass := range tests {
		t.Run(command, func(t *testing.T) {
		line :=strings.TrimSpace(command)
		lines := strings.SplitN(line,"\n", -1)
		var err error
		parser.Init()
		if len(lines)> 1{
			for _,line = range(lines){
			  line = strings.TrimSpace(line)
				if err=ProcessLine(line, parser);err!=nil{
					break
				}
			}
		}else{
			err = ProcessLine(line, parser)
		}
			if shouldPass {
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				result, err := parser.Result()
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				var returnLine string
				if result[0].Comment == "" {
					returnLine = result[0].Data
				} else {
					returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
				}
				if command != returnLine {
					t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, command))
				}
			} else {
				if err == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
				}
				_, parseErr := parser.Result()
				if parseErr == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
				}
			}
		})
	}
}
{{- if .HasAlias}}

func TestAlias{{ $StructName }}{{ .Dir }}(t *testing.T) {
	tests := map[string]bool{ {{ $AliasName := .StructName }}
	{{- range $index, $val := .TestAliasOK}} {{ $AliasName = $val.Alias }}
		"{{- $val.Test -}}": true,
	{{- end }}
	{{- range $index, $val := .TestAliasFail}}
		"{{- $val.Test -}}": false,
	{{- end }}
	}
	parser := {{- if .ModeOther}} &{{ .Dir }}{{- else }} &parsers{{- end }}.{{ $StructName }}{Alias:"{{ $AliasName }}"}
	for command, shouldPass := range tests {
		t.Run(command, func(t *testing.T) {
		line :=strings.TrimSpace(command)
		lines := strings.SplitN(line,"\n", -1)
		var err error
		parser.Init()
		if len(lines)> 1{
			for _,line = range(lines){
			  line = strings.TrimSpace(line)
				if err=ProcessLine(line, parser);err!=nil{
					break
				}
			}
		}else{
			err = ProcessLine(line, parser)
		}
			if shouldPass {
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				result, err := parser.Result()
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				var returnLine string
				if result[0].Comment == "" {
					returnLine = result[0].Data
				} else {
					returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
				}
				if command != returnLine {
					t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, command))
				}
			} else {
				if err == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
				}
				_, parseErr := parser.Result()
				if parseErr == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
				}
			}
		})
	}
}
{{- end }}
`))
