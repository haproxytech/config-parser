package tests

import (
	configparser "github.com/haproxytech/config-parser"
)

func ProcessLine(line string, parser configparser.ParserType) error {
	parser.Init()
	_, err := parser.Parse(line, "  "+line, "")
	return err
}
