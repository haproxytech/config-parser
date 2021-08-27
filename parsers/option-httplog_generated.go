// Code generated by go generate; DO NOT EDIT.
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
package parsers

import (
	"github.com/deyunluo/config-parser/v4/common"
	"github.com/deyunluo/config-parser/v4/errors"
	"github.com/deyunluo/config-parser/v4/types"
)

func (p *OptionHTTPLog) Init() {
	p.data = nil
	p.preComments = []string{}
}

func (p *OptionHTTPLog) GetParserName() string {
	return "option httplog"
}

func (p *OptionHTTPLog) Get(createIfNotExist bool) (common.ParserData, error) {
	if p.data == nil {
		if createIfNotExist {
			p.data = &types.OptionHTTPLog{}
			return p.data, nil
		}
		return nil, errors.ErrFetch
	}
	return p.data, nil
}

func (p *OptionHTTPLog) GetOne(index int) (common.ParserData, error) {
	if index > 0 {
		return nil, errors.ErrFetch
	}
	if p.data == nil {
		return nil, errors.ErrFetch
	}
	return p.data, nil
}

func (p *OptionHTTPLog) Delete(index int) error {
	p.Init()
	return nil
}

func (p *OptionHTTPLog) Insert(data common.ParserData, index int) error {
	return p.Set(data, index)
}

func (p *OptionHTTPLog) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.OptionHTTPLog:
		p.data = newValue
	case types.OptionHTTPLog:
		p.data = &newValue
	default:
		return errors.ErrInvalidData
	}
	return nil
}

func (p *OptionHTTPLog) PreParse(line string, parts, previousParts []string, preComments []string, comment string) (changeState string, err error) {
	changeState, err = p.Parse(line, parts, previousParts, comment)
	if err == nil && preComments != nil {
		p.preComments = append(p.preComments, preComments...)
	}
	return changeState, err
}

func (p *OptionHTTPLog) ResultAll() ([]common.ReturnResultLine, []string, error) {
	res, err := p.Result()
	return res, p.preComments, err
}
