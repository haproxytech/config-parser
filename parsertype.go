package parser

import "fmt"

type ParserType interface {
	Init()
	Parse(line, wholeLine, previousLine string) (changeState string, err error)
	Valid() bool
	GetParserName() string
	String() []string
}

type ParserTypes struct {
	parsers []ParserType
	maxSize int
}

func (p *ParserTypes) Get(atrtibute string) (ParserType, error) {
	for _, parser := range p.parsers {
		if parser.GetParserName() == atrtibute && parser.Valid() {
			return parser, nil
		}
	}
	return nil, fmt.Errorf("atrtibute not found")
}
