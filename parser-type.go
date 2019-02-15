package parser

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type ParserType interface {
	Init()
	Parse(line string, parts, previousParts []string, comment string) (changeState string, err error)
	GetParserName() string
	Get(createIfNotExist bool) (common.ParserData, error)
	GetOne(index int) (common.ParserData, error)
	Delete(index int) error
	Insert(data common.ParserData, index int) error
	Set(data common.ParserData, index int) error
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
	return nil, errors.ParserMissingErr
}

//HasParser checks if we have a parser for attribute
func (p *ParserTypes) HasParser(attribute string) bool {
	for _, parser := range p.parsers {
		if parser.GetParserName() == attribute {
			return true
		}
	}
	return false
}

//Set sets data in parser, if you can have multiple items, index is a must
func (p *ParserTypes) Set(attribute string, data common.ParserData, index ...int) error {
	setIndex := -1
	if len(index) > 0 && index[0] > -1 {
		setIndex = index[0]
	}
	for i, parser := range p.parsers {
		if parser.GetParserName() == attribute {
			return p.parsers[i].Set(data, setIndex)
		}
	}
	return fmt.Errorf("attribute not available")
}

func (p *ParserTypes) Insert(attribute string, data common.ParserData, index ...int) error {
	setIndex := -1
	if len(index) > 0 && index[0] > -1 {
		setIndex = index[0]
	}
	for i, parser := range p.parsers {
		if parser.GetParserName() == attribute {
			return p.parsers[i].Insert(data, setIndex)
		}
	}
	return fmt.Errorf("attribute not available")
}

func (p *ParserTypes) Delete(attribute string, index ...int) error {
	setIndex := -1
	if len(index) > 0 && index[0] > -1 {
		setIndex = index[0]
	}
	for i, parser := range p.parsers {
		if parser.GetParserName() == attribute {
			return p.parsers[i].Delete(setIndex)
		}
	}
	return fmt.Errorf("attribute not available")
}
