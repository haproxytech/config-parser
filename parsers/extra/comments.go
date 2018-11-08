package extra

import (
	"strings"

	"github.com/haproxytech/config-parser/errors"
)

type Comments struct {
	comments []string
}

func (c *Comments) Init() {
	c.comments = []string{}
}

func (c *Comments) GetParserName() string {
	return "#"
}

func (c *Comments) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	if strings.HasPrefix(wholeLine, "#") {
		c.comments = append(c.comments, line)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "Comments", Line: line}
}

func (c *Comments) Valid() bool {
	if len(c.comments) > 0 {
		return true
	}
	return false
}

func (c *Comments) String() []string {
	return c.comments
}
