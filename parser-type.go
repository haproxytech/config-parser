/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package parser

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

//nolint:golint
type ParserInterface interface {
	Init()
	Parse(line string, parts, previousParts []string, comment string) (changeState string, err error)
	GetParserName() string
	Get(createIfNotExist bool) (common.ParserData, error)
	GetOne(index int) (common.ParserData, error)
	Delete(index int) error
	Insert(data common.ParserData, index int) error
	Set(data common.ParserData, index int) error
	Result() ([]common.ReturnResultLine, error)
}

type Parsers struct {
	parsers []ParserInterface
}

func (p *Parsers) Get(attribute string, createIfNotExist ...bool) (common.ParserData, error) {
	createNew := false
	if len(createIfNotExist) > 0 && createIfNotExist[0] {
		createNew = true
	}
	for _, parser := range p.parsers {
		if parser.GetParserName() == attribute {
			return parser.Get(createNew)
		}
	}
	return nil, errors.ErrParserMissing
}

func (p *Parsers) GetOne(attribute string, index ...int) (common.ParserData, error) {
	setIndex := -1
	if len(index) > 0 && index[0] > -1 {
		setIndex = index[0]
	}
	for _, parser := range p.parsers {
		if parser.GetParserName() == attribute {
			return parser.GetOne(setIndex)
		}
	}
	return nil, errors.ErrParserMissing
}

//HasParser checks if we have a parser for attribute
func (p *Parsers) HasParser(attribute string) bool {
	for _, parser := range p.parsers {
		if parser.GetParserName() == attribute {
			return true
		}
	}
	return false
}

//Set sets data in parser, if you can have multiple items, index is a must
func (p *Parsers) Set(attribute string, data common.ParserData, index ...int) error {
	setIndex := -1
	if len(index) > 0 && index[0] > -1 {
		setIndex = index[0]
	}
	for i, parser := range p.parsers {
		if parser.GetParserName() == attribute {
			return p.parsers[i].Set(data, setIndex)
		}
	}
	return errors.ErrAttributeNotFound
}

func (p *Parsers) Insert(attribute string, data common.ParserData, index ...int) error {
	setIndex := -1
	if len(index) > 0 && index[0] > -1 {
		setIndex = index[0]
	}
	for i, parser := range p.parsers {
		if parser.GetParserName() == attribute {
			return p.parsers[i].Insert(data, setIndex)
		}
	}
	return errors.ErrAttributeNotFound
}

func (p *Parsers) Delete(attribute string, index ...int) error {
	setIndex := -1
	if len(index) > 0 && index[0] > -1 {
		setIndex = index[0]
	}
	for i, parser := range p.parsers {
		if parser.GetParserName() == attribute {
			return p.parsers[i].Delete(setIndex)
		}
	}
	return errors.ErrAttributeNotFound
}
