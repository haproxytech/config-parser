package peers

import (
	"fmt"
	"strconv"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Peers struct {
	data []types.Peer
}

func (l *Peers) Init() {
	l.data = []types.Peer{}
}

func (l *Peers) GetParserName() string {
	return "peer"
}

func (l *Peers) Clear() {
	l.Init()
}

func (l *Peers) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(l.data) == 0 && !createIfNotExist {
		return nil, &errors.FetchError{}
	}
	return l.data, nil
}

func (l *Peers) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case []types.Peer:
		l.data = newValue
	case types.Peer:
		l.data = append(l.data, newValue)
	case *types.Peer:
		l.data = append(l.data, *newValue)
	}
	return fmt.Errorf("casting error")
}

func (l *Peers) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := l.Get(false)
	l.Clear()
	_, err := l.Parse(data, parts, []string{}, comment)
	if err != nil {
		l.Set(oldData)
	}
	return err
}

func (l *Peers) parsePeerLine(line string, parts []string, comment string) (*types.Peer, error) {
	if len(parts) >= 2 {
		adr := common.StringSplitIgnoreEmpty(parts[2], ':')
		if len(adr) >= 2 {
			if port, err := strconv.ParseInt(adr[1], 10, 64); err == nil {
				return &types.Peer{
					Name:    parts[1],
					IP:      adr[0],
					Port:    port,
					Comment: comment,
				}, nil
			}
		}
	}
	return nil, &errors.ParseError{Parser: "PeerLines", Line: line}
}

func (l *Peers) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "peer" {
		peer, err := l.parsePeerLine(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "PeerLines", Line: line}
		}
		l.data = append(l.data, *peer)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "PeerLines", Line: line}
}

func (l *Peers) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(l.data) == 0 {
		return nil, &errors.FetchError{}
	}
	result := make([]common.ReturnResultLine, len(l.data))
	for index, peer := range l.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("peer %s %s:%d", peer.Name, peer.IP, peer.Port),
			Comment: peer.Comment,
		}
	}
	return result, nil
}
