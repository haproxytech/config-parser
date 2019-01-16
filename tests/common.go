package tests

import (
	configparser "github.com/haproxytech/config-parser"
	"github.com/haproxytech/config-parser/common"
)

func ProcessLine(line string, parser configparser.ParserType) error {
	parts := common.StringSplitIgnoreEmpty(line, ' ')
	parser.Init()
	_, err := parser.Parse("  "+line, parts, []string{})
	return err
}
