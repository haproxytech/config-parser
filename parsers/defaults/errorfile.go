package defaults

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type ErrorFileLines struct {
	data        []types.ErrorFile
	allowedCode map[string]bool
}

func (l *ErrorFileLines) Init() {
	l.data = []types.ErrorFile{}
	l.allowedCode = map[string]bool{}
	common.AddToBoolMap(l.allowedCode,
		"200", "400", "403", "405", "408", "425", "429",
		"500", "502", "503", "504")
}

func (l *ErrorFileLines) GetParserName() string {
	return "errorfile"
}

func (p *ErrorFileLines) Clear() {
	p.Init()
}

func (p *ErrorFileLines) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(p.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *ErrorFileLines) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case []types.ErrorFile:
		p.data = newValue
	case types.ErrorFile:
		p.data = append(p.data, newValue)
	}
	return fmt.Errorf("casting error")
}

func (p *ErrorFileLines) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := p.Get(false)
	p.Clear()
	_, err := p.Parse(data, parts, []string{}, comment)
	if err != nil {
		p.Set(oldData)
	}
	return err
}

func (l *ErrorFileLines) parseErrorFileLine(line string, comment string) (*types.ErrorFile, error) {
	parts := common.StringSplitIgnoreEmpty(line, ' ')
	if len(parts) < 3 {
		return nil, &errors.ParseError{Parser: "ErrorFileLines", Line: line}
	}
	errorfile := &types.ErrorFile{
		File:    parts[2],
		Comment: comment,
	}
	code := parts[1]
	if _, ok := l.allowedCode[code]; !ok {
		return errorfile, nil
	}
	errorfile.Code = code
	return errorfile, nil
}

func (l *ErrorFileLines) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "errorfile" {
		if data, err := l.parseErrorFileLine(line, comment); err == nil {
			l.data = append(l.data, *data)
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "ErrorFileLines", Line: line}
}

func (l *ErrorFileLines) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(l.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(l.data))
	for index, data := range l.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("errorfile %s %s", data.Code, data.File),
			Comment: data.Comment,
		}
	}
	return result, nil
}
