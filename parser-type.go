package parser

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
)

type ParserType interface {
	Init()
	Clear()
	Parse(line string, parts, previousParts []string, comment string) (changeState string, err error)
	GetParserName() string
	Get(createIfNotExist bool) (common.ParserData, error)
	Set(data common.ParserData) error
	SetStr(data string) error
	Result(AddComments bool) ([]common.ReturnResultLine, error)
}

type ParserTypes struct {
	parsers []ParserType
}

func (p *ParserTypes) Get(attribute string, createIfNotExist ...bool) (common.ParserData, error) {
	createNew := false
	if len(createIfNotExist) > 0 && createIfNotExist[0] {
		createNew = true
	}
	for _, parser := range p.parsers {
		if parser.GetParserName() == attribute {
			return parser.Get(createNew)
		}
	}
	return nil, fmt.Errorf("attribute not found, no available parser for it")
}

func (p *ParserTypes) Clear(attribute string) {
	for _, parser := range p.parsers {
		if parser.GetParserName() == attribute {
			parser.Clear()
			break //we should only have one parser for attribute
		}
	}
}

func (p *ParserTypes) Set(attribute string, data common.ParserData) error {
	for index, parser := range p.parsers {
		if parser.GetParserName() == attribute {
			return p.parsers[index].Set(data)
		}
	}
	return fmt.Errorf("attribute not available")
}
