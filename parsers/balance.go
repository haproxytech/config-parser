package parsers

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type Balance struct {
	Algorithm string
	Arguments []string
}

func (b *Balance) Init() {
	b.Arguments = []string{}
}

func (b *Balance) GetParserName() string {
	return "balance"
}

func (b *Balance) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if parts[0] == "balance" {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "Balance", Line: line, Message: "Parse error"}
		}
		switch parts[1] {
		case "roundrobin", "static-rr", "leastconn", "first", "source", "random":
			b.Algorithm = parts[1]
			return "", nil
		case "uri", "url_param":
			b.Algorithm = parts[1]
			if len(parts) > 2 {
				b.Arguments = parts[2:]
				return "", nil
			}
			return "", &errors.ParseError{Parser: "Balance", Line: line}
		}
		if strings.HasPrefix(parts[1], "hdr(") && strings.HasSuffix(parts[1], ")") {
			b.Algorithm = parts[1]
			return "", nil
		}
		if strings.HasPrefix(parts[1], "rdp-cookie(") && strings.HasSuffix(parts[1], ")") {
			b.Algorithm = parts[1]
			return "", nil
		}
		return "", &errors.ParseError{Parser: "Balance", Line: line}
	}
	return "", &errors.ParseError{Parser: "Balance", Line: line}
}

func (b *Balance) Valid() bool {
	return b.Algorithm != ""
}

func (b *Balance) Result(AddComments bool) []string {
	params := ""
	if len(b.Arguments) > 0 {
		params = fmt.Sprintf(" %s", strings.Join(b.Arguments, " "))
	}
	return []string{fmt.Sprintf("  balance %s%s", b.Algorithm, params)}
}
