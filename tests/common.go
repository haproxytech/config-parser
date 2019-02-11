package tests

import (
	configparser "github.com/haproxytech/config-parser"
	"github.com/haproxytech/config-parser/common"
)

func ProcessLine(line string, parser configparser.ParserType) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(line, ' ', '\t')
	parser.Init()
	_, err := parser.Parse("  "+line, parts, []string{}, comment)
	return err
}
