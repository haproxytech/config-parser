package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type NameserverLines struct {
	data []types.Nameserver
}

func (l *NameserverLines) Init() {
	l.data = []types.Nameserver{}
}

func (l *NameserverLines) GetParserName() string {
	return "nameserver"
}

func (l *NameserverLines) Clear() {
	l.Init()
}

func (l *NameserverLines) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(l.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return l.data, nil
}

func (l *NameserverLines) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case []types.Nameserver:
		l.data = newValue
	case types.Nameserver:
		l.data = append(l.data, newValue)
	}
	return fmt.Errorf("casting error")
}

func (l *NameserverLines) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := l.Get(false)
	l.Clear()
	_, err := l.Parse(data, parts, []string{}, comment)
	if err != nil {
		l.Set(oldData)
	}
	return err
}

func (l *NameserverLines) parseNameserverLine(line string, parts []string, comment string) (*types.Nameserver, error) {
	if len(parts) >= 3 {
		return &types.Nameserver{
			Name:    parts[1],
			Address: parts[2],
			Comment: comment,
		}, nil
	}
	return nil, &errors.ParseError{Parser: "NameserverLines", Line: line}
}

func (l *NameserverLines) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "nameserver" {
		nameserver, err := l.parseNameserverLine(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "NameserverLines", Line: line}
		}
		l.data = append(l.data, *nameserver)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "NameserverLines", Line: line}
}

func (l *NameserverLines) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(l.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(l.data))
	for index, nameserver := range l.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("nameserver %s %s", nameserver.Name, nameserver.Address),
			Comment: nameserver.Comment,
		}
	}
	return result, nil
}
