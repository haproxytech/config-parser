package mailers

import (
	"fmt"
	"strconv"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Mailers struct {
	data []types.Mailer
}

func (l *Mailers) Init() {
	l.data = []types.Mailer{}
}

func (l *Mailers) GetParserName() string {
	return "peer"
}

func (l *Mailers) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(l.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return l.data, nil
}

func (l *Mailers) Set(data common.ParserData) error {
	if data == nil {
		l.Init()
		return nil
	}
	switch newValue := data.(type) {
	case []types.Mailer:
		l.data = newValue
	case *types.Mailer:
		l.data = append(l.data, *newValue)
	case types.Mailer:
		l.data = append(l.data, newValue)
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (l *Mailers) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := l.Get(false)
	l.Init()
	_, err := l.Parse(data, parts, []string{}, comment)
	if err != nil {
		l.Set(oldData)
	}
	return err
}

func (l *Mailers) parseMailerLine(line string, parts []string, comment string) (*types.Mailer, error) {
	if len(parts) >= 2 {
		adr := common.StringSplitIgnoreEmpty(parts[2], ':')
		if len(adr) >= 2 {
			if port, err := strconv.ParseInt(adr[1], 10, 64); err == nil {
				return &types.Mailer{
					Name:    parts[1],
					IP:      adr[0],
					Port:    port,
					Comment: comment,
				}, nil
			}
		}
	}
	return nil, &errors.ParseError{Parser: "MailerLines", Line: line}
}

func (l *Mailers) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "mailer" {
		mailer, err := l.parseMailerLine(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "MailerLines", Line: line}
		}
		l.data = append(l.data, *mailer)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "MailerLines", Line: line}
}

func (l *Mailers) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(l.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(l.data))
	for index, peer := range l.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("mailer %s %s:%d", peer.Name, peer.IP, peer.Port),
			Comment: peer.Comment,
		}
	}
	return result, nil
}
