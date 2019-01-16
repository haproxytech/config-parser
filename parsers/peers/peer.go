package peers

import (
	"fmt"
	"strconv"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type Peer struct {
	Name string
	IP   string
	Port int64
}

type Peers struct {
	Peers []Peer
}

func (l *Peers) Init() {
	l.Peers = []Peer{}
}

func (l *Peers) GetParserName() string {
	return "peer"
}

func (l *Peers) parsePeerLine(line string, parts []string) (Peer, error) {
	if len(parts) >= 2 {
		adr := common.StringSplitIgnoreEmpty(parts[2], ':')
		if len(adr) >= 2 {
			if port, err := strconv.ParseInt(adr[1], 10, 64); err == nil {
				return Peer{
					Name: parts[1],
					IP:   adr[0],
					Port: port,
				}, nil
			}
		}
	}
	return Peer{}, &errors.ParseError{Parser: "PeerLines", Line: line}
}

func (l *Peers) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if parts[0] == "peer" {
		nameserver, err := l.parsePeerLine(line, parts)
		if err != nil {
			return "", &errors.ParseError{Parser: "PeerLines", Line: line}
		}
		l.Peers = append(l.Peers, nameserver)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "PeerLines", Line: line}
}

func (l *Peers) Valid() bool {
	if len(l.Peers) > 0 {
		return true
	}
	return false
}

func (l *Peers) Result(AddComments bool) []string {
	result := make([]string, len(l.Peers))
	for index, peer := range l.Peers {
		result[index] = fmt.Sprintf("  peer %s %s:%d", peer.Name, peer.IP, peer.Port)
	}
	return result
}
