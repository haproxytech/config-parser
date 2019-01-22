package mailers

import (
	"fmt"
	"strconv"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type Mailer struct {
	Name    string
	IP      string
	Port    int64
	Comment string
}

type Mailers struct {
	Mailers []Mailer
}

func (l *Mailers) Init() {
	l.Mailers = []Mailer{}
}

func (l *Mailers) GetParserName() string {
	return "peer"
}

func (l *Mailers) parseMailerLine(line string, parts []string, comment string) (Mailer, error) {
	if len(parts) >= 2 {
		adr := common.StringSplitIgnoreEmpty(parts[2], ':')
		if len(adr) >= 2 {
			if port, err := strconv.ParseInt(adr[1], 10, 64); err == nil {
				return Mailer{
					Name:    parts[1],
					IP:      adr[0],
					Port:    port,
					Comment: comment,
				}, nil
			}
		}
	}
	return Mailer{}, &errors.ParseError{Parser: "MailerLines", Line: line}
}

func (l *Mailers) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "mailer" {
		nameserver, err := l.parseMailerLine(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "MailerLines", Line: line}
		}
		l.Mailers = append(l.Mailers, nameserver)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "MailerLines", Line: line}
}

func (l *Mailers) Valid() bool {
	if len(l.Mailers) > 0 {
		return true
	}
	return false
}

func (l *Mailers) Result(AddComments bool) []common.ReturnResultLine {
	result := make([]common.ReturnResultLine, len(l.Mailers))
	for index, peer := range l.Mailers {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("mailer %s %s:%d", peer.Name, peer.IP, peer.Port),
			Comment: peer.Comment,
		}
	}
	return result
}
