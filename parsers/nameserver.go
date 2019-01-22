package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type Nameserver struct {
	Name    string
	Address string
	Comment string
}

type NameserverLines struct {
	NameserverLines []Nameserver
}

func (l *NameserverLines) Init() {
	l.NameserverLines = []Nameserver{}
}

func (l *NameserverLines) GetParserName() string {
	return "nameserver"
}

func (l *NameserverLines) parseNameserverLine(line string, parts []string, comment string) (Nameserver, error) {
	if len(parts) >= 3 {
		return Nameserver{
			Name:    parts[1],
			Address: parts[2],
			Comment: comment,
		}, nil
	}
	return Nameserver{}, &errors.ParseError{Parser: "NameserverLines", Line: line}
}

func (l *NameserverLines) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "nameserver" {
		nameserver, err := l.parseNameserverLine(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "NameserverLines", Line: line}
		}
		l.NameserverLines = append(l.NameserverLines, nameserver)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "NameserverLines", Line: line}
}

func (l *NameserverLines) Valid() bool {
	if len(l.NameserverLines) > 0 {
		return true
	}
	return false
}

func (l *NameserverLines) Result(AddComments bool) []common.ReturnResultLine {
	result := make([]common.ReturnResultLine, len(l.NameserverLines))
	for index, nameserver := range l.NameserverLines {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("nameserver %s %s", nameserver.Name, nameserver.Address),
			Comment: nameserver.Comment,
		}
	}
	return result
}
