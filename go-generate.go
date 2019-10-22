// +build ignore

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

//run this as go run go-generate.go $(pwd)

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/haproxytech/config-parser/v2/common"
)

type Data struct {
	ParserMultiple     bool
	ParserName         string
	ParserSecondName   string
	StructName         string
	ParserType         string
	ParserTypeOverride string
	NoInit             bool
	NoParse            bool
	NoGet              bool
	IsInterface        bool
	Dir                string
	ModeOther          bool
	TestOK             []string
	TestFail           []string
}

func main() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	log.Println(dir)
	generateTypes(dir)
	generateTypesGeneric(dir)
	generateTypesOther(dir)
}

func generateTypesOther(dir string) {
	dat, err := ioutil.ReadFile("types/types-other.go")
	if err != nil {
		log.Println(err)
	}
	lines := common.StringSplitIgnoreEmpty(string(dat), '\n')
	//fmt.Print(lines)

	parserData := Data{}
	for _, line := range lines {
		if strings.HasPrefix(line, "//sections:") {
			//log.Println(line)
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
		if strings.HasPrefix(line, "//test:ok") {
			data := strings.SplitN(line, ":", 3)
			parserData.TestOK = append(parserData.TestOK, data[2])
		}
		if strings.HasPrefix(line, "//test:fail") {
			data := strings.SplitN(line, ":", 3)
			parserData.TestFail = append(parserData.TestFail, data[2])
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

		filePath := path.Join(dir, "parsers", parserData.Dir, cleanFileName(filename)+"_generated.go")
		log.Println(filePath)
		f, err := os.Create(filePath)
		CheckErr(err)
		defer f.Close()

		parserData.ModeOther = true
		err = typeTemplate.Execute(f, parserData)
		CheckErr(err)

		parserData = Data{}
	}
}

func generateTypesGeneric(dir string) {
	dat, err := ioutil.ReadFile("types/types-generic.go")
	if err != nil {
		log.Println(err)
	}
	lines := common.StringSplitIgnoreEmpty(string(dat), '\n')

	parsers := map[string]*Data{}
	parserData := &Data{}
	for _, line := range lines {
		if strings.HasPrefix(line, "//sections:") {
			//log.Println(line)
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
		if strings.HasPrefix(line, "//no-parse:true") {
			parserData.NoParse = true
		}
		if strings.HasPrefix(line, "//test:ok") {
			data := strings.SplitN(line, ":", 3)
			parserData.TestOK = append(parserData.TestOK, data[2])
		}
		if strings.HasPrefix(line, "//test:fail") {
			data := strings.SplitN(line, ":", 3)
			parserData.TestFail = append(parserData.TestFail, data[2])
		}
		if strings.HasPrefix(line, "//gen:") {
			data := common.StringSplitIgnoreEmpty(line, ':')
			parserData = &Data{}
			parserData.StructName = data[1]
			parsers[data[1]] = parserData
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
			//log.Println(parserData)
			//continue
			f, err := os.Create(filePath)
			CheckErr(err)
			defer f.Close()

			err = typeTemplate.Execute(f, parserData)
			CheckErr(err)

			//parserData.TestFail = append(parserData.TestFail, "") parsers should not get empty line!
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
		parsers = map[string]*Data{}
		parserData = &Data{}
	}
}

func generateTypes(dir string) {
	dat, err := ioutil.ReadFile("types/types.go")
	if err != nil {
		log.Println(err)
	}
	lines := common.StringSplitIgnoreEmpty(string(dat), '\n')
	//fmt.Print(lines)

	parserData := Data{}
	for _, line := range lines {
		if strings.HasPrefix(line, "//sections:") {
			//log.Println(line)
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
		if strings.HasPrefix(line, "//no-parse:true") {
			parserData.NoParse = true
		}
		if strings.HasPrefix(line, "//test:ok") {
			data := strings.SplitN(line, ":", 3)
			parserData.TestOK = append(parserData.TestOK, data[2])
		}
		if strings.HasPrefix(line, "//test:fail") {
			data := strings.SplitN(line, ":", 3)
			parserData.TestFail = append(parserData.TestFail, data[2])
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

		filePath := path.Join(dir, "parsers", cleanFileName(filename)+"_generated.go")
		log.Println(filePath)
		f, err := os.Create(filePath)
		CheckErr(err)
		defer f.Close()

		err = typeTemplate.Execute(f, parserData)
		CheckErr(err)

		//parserData.TestFail = append(parserData.TestFail, "") parsers should not get empty line!
		parserData.TestFail = append(parserData.TestFail, "---")
		parserData.TestFail = append(parserData.TestFail, "--- ---")

		filePath = path.Join(dir, "tests", cleanFileName(filename)+"_generated_test.go")
		log.Println(filePath)
		f, err = os.Create(filePath)
		CheckErr(err)
		defer f.Close()

		err = testTemplate.Execute(f, parserData)
		CheckErr(err)

		parserData = Data{}
	}
}

func cleanFileName(filename string) string {
	return strings.Replace(filename, " ", "-", -1)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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
	"github.com/haproxytech/config-parser/v2/common"
	"github.com/haproxytech/config-parser/v2/errors"
	"github.com/haproxytech/config-parser/v2/types"
)

{{- if not .NoInit }}

func (p *{{ .StructName }}) Init() {
{{- if .ParserMultiple }}
	p.data = []types.{{ .ParserType }}{}
{{- else }}
    p.data = nil
{{- end }}
}
{{- end }}

func (p *{{ .StructName }}) GetParserName() string {
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
`))

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

	"github.com/haproxytech/config-parser/v2/parsers"
)

{{ $StructName := .StructName }}
{{- range $index, $val := .TestOK}}
func Test{{ $StructName }}Normal{{$index}}(t *testing.T) {
	parser := &parsers.{{ $StructName }}{}
	line := strings.TrimSpace("{{- $val -}}")
	err := ProcessLine(line, parser)
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
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
{{- end }}

{{- range $index, $val := .TestFail}}
func Test{{ $StructName }}Fail{{$index}}(t *testing.T) {
	parser := &parsers.{{ $StructName }}{}
	line := strings.TrimSpace("{{- $val -}}")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result()
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
{{- end }}
`))
