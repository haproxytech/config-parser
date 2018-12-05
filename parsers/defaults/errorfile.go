package defaults

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/helpers"
)

type ErrorFile struct {
	Code string
	File string
}

type ErrorFileLines struct {
	ErrorFileLines []ErrorFile
	AllowedCode    map[string]bool
}

func (l *ErrorFileLines) Init() {
	l.ErrorFileLines = []ErrorFile{}
	l.AllowedCode = map[string]bool{}
	helpers.AddToBoolMap(l.AllowedCode,
		"200", "400", "403", "405", "408", "425", "429",
		"500", "502", "503", "504")
}

func (l *ErrorFileLines) GetParserName() string {
	return "errorfile"
}

func (l *ErrorFileLines) parseErrorFileLine(line string) (ErrorFile, error) {
	parts := helpers.StringSplitIgnoreEmpty(line, ' ')
	if len(parts) < 3 {
		return ErrorFile{}, &errors.ParseError{Parser: "ErrorFileLines", Line: line}
	}
	errorfile := ErrorFile{
		File: parts[2],
	}
	code := parts[1]
	if _, ok := l.AllowedCode[code]; !ok {
		return errorfile, nil
	}
	errorfile.Code = code
	return errorfile, nil
}

func (l *ErrorFileLines) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(line, "errorfile ") {
		if data, err := l.parseErrorFileLine(line); err == nil {
			l.ErrorFileLines = append(l.ErrorFileLines, data)
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "ErrorFileLines", Line: line}
}

func (l *ErrorFileLines) Valid() bool {
	if len(l.ErrorFileLines) > 0 {
		return true
	}
	return false
}

func (l *ErrorFileLines) String() []string {
	result := make([]string, len(l.ErrorFileLines))
	for index, data := range l.ErrorFileLines {
		result[index] = fmt.Sprintf("  errorfile %s %s", data.Code, data.File)
	}
	return result
}
