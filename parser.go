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
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"sync"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/parsers/extra"
	"github.com/haproxytech/config-parser/types"
)

type Section string

const (
	Comments  Section = "#"
	Defaults  Section = "defaults"
	Global    Section = "global"
	Resolvers Section = "resolvers"
	UserList  Section = "userlist"
	Peers     Section = "peers"
	Mailers   Section = "mailers"
	Frontends Section = "frontend"
	Backends  Section = "backend"
	Listen    Section = "listen"
	Cache     Section = "cache"
	Program   Section = "program"
)

const (
	CommentsSectionName = "data"
	GlobalSectionName   = "data"
	DefaultSectionName  = "data"
)

//Parser reads and writes configuration on given file
type Parser struct {
	Parsers map[Section]map[string]*Parsers
	mutex   *sync.Mutex
}

func (p *Parser) lock() {
	p.mutex.Lock()
}

func (p *Parser) unLock() {
	p.mutex.Unlock()
}

//Get get attribute from defaults section
func (p *Parser) Get(sectionType Section, sectionName string, attribute string, createIfNotExist ...bool) (common.ParserData, error) {
	p.lock()
	defer p.unLock()
	st, ok := p.Parsers[sectionType]
	if !ok {
		return nil, errors.ErrSectionMissing
	}
	section, ok := st[sectionName]
	if !ok {
		return nil, errors.ErrSectionMissing
	}
	createNew := false
	if len(createIfNotExist) > 0 && createIfNotExist[0] {
		createNew = true
	}
	return section.Get(attribute, createNew)
}

//GetOne get attribute from defaults section
func (p *Parser) GetOne(sectionType Section, sectionName string, attribute string, index ...int) (common.ParserData, error) {
	p.lock()
	defer p.unLock()
	setIndex := -1
	if len(index) > 0 && index[0] > -1 {
		setIndex = index[0]
	}
	st, ok := p.Parsers[sectionType]
	if !ok {
		return nil, errors.ErrSectionMissing
	}
	section, ok := st[sectionName]
	if !ok {
		return nil, errors.ErrSectionMissing
	}
	return section.GetOne(attribute, setIndex)
}

//SectionsGet lists all sections of certain type
func (p *Parser) SectionsGet(sectionType Section) ([]string, error) {
	p.lock()
	defer p.unLock()
	st, ok := p.Parsers[sectionType]
	if !ok {
		return nil, errors.ErrSectionMissing
	}
	result := make([]string, len(st))
	index := 0
	for sectionName := range st {
		result[index] = sectionName
		index++
	}
	return result, nil
}

//SectionsDelete deletes one section of sectionType
func (p *Parser) SectionsDelete(sectionType Section, sectionName string) error {
	p.lock()
	defer p.unLock()
	_, ok := p.Parsers[sectionType]
	if !ok {
		return errors.ErrSectionMissing
	}
	delete(p.Parsers[sectionType], sectionName)
	return nil
}

//SectionsCreate creates one section of sectionType
func (p *Parser) SectionsCreate(sectionType Section, sectionName string) error {
	p.lock()
	defer p.unLock()
	st, ok := p.Parsers[sectionType]
	if !ok {
		return errors.ErrSectionMissing
	}
	_, ok = st[sectionName]
	if ok {
		return errors.ErrSectionAlreadyExists
	}

	parsers := ConfiguredParsers{
		State:    "",
		Active:   *p.Parsers[Comments][CommentsSectionName],
		Comments: p.Parsers[Comments][CommentsSectionName],
		Defaults: p.Parsers[Defaults][DefaultSectionName],
		Global:   p.Parsers[Global][GlobalSectionName],
	}

	previousLine := []string{}
	parts := []string{string(sectionType), sectionName}
	comment := ""
	p.ProcessLine(fmt.Sprintf("%s %s", sectionType, sectionName), parts, previousLine, comment, parsers)
	return nil
}

//Set sets attribute from defaults section, can be nil to disable/remove
func (p *Parser) Set(sectionType Section, sectionName string, attribute string, data common.ParserData, index ...int) error {
	p.lock()
	defer p.unLock()
	setIndex := -1
	if len(index) > 0 && index[0] > -1 {
		setIndex = index[0]
	}
	st, ok := p.Parsers[sectionType]
	if !ok {
		return errors.ErrSectionMissing
	}
	section, ok := st[sectionName]
	if !ok {
		return errors.ErrSectionMissing
	}
	return section.Set(attribute, data, setIndex)
}

//Delete remove attribute on defined index, in case of single attributes, index is ignored
func (p *Parser) Delete(sectionType Section, sectionName string, attribute string, index ...int) error {
	p.lock()
	defer p.unLock()
	setIndex := -1
	if len(index) > 0 && index[0] > -1 {
		setIndex = index[0]
	}
	st, ok := p.Parsers[sectionType]
	if !ok {
		return errors.ErrSectionMissing
	}
	section, ok := st[sectionName]
	if !ok {
		return errors.ErrSectionMissing
	}
	return section.Delete(attribute, setIndex)
}

//Insert put attribute on defined index, in case of single attributes, index is ignored
func (p *Parser) Insert(sectionType Section, sectionName string, attribute string, data common.ParserData, index ...int) error {
	p.lock()
	defer p.unLock()
	setIndex := -1
	if len(index) > 0 && index[0] > -1 {
		setIndex = index[0]
	}
	st, ok := p.Parsers[sectionType]
	if !ok {
		return errors.ErrSectionMissing
	}
	section, ok := st[sectionName]
	if !ok {
		return errors.ErrSectionMissing
	}
	return section.Insert(attribute, data, setIndex)
}

//HasParser checks if we have a parser for attribute
func (p *Parser) HasParser(sectionType Section, attribute string) bool {
	p.lock()
	defer p.unLock()
	st, ok := p.Parsers[sectionType]
	if !ok {
		return false
	}
	sectionName := ""
	for name := range st {
		sectionName = name
		break
	}
	section, ok := st[sectionName]
	if !ok {
		return false
	}
	return section.HasParser(attribute)
}

func (p *Parser) writeParsers(sectionName string, parsers []ParserInterface, result *strings.Builder, useIndentation bool) {
	sectionNameWritten := false
	if sectionName == "" {
		sectionNameWritten = true
	}
	for _, parser := range parsers {
		lines, err := parser.Result()
		if err != nil {
			continue
		}
		if !sectionNameWritten {
			result.WriteString("\n")
			result.WriteString(sectionName)
			result.WriteString(" \n")
			sectionNameWritten = true
		}
		for _, line := range lines {
			if useIndentation {
				result.WriteString("  ")
			}
			result.WriteString(line.Data)
			if line.Comment != "" {
				result.WriteString(" # ")
				result.WriteString(line.Comment)
			}
			result.WriteString("\n")
		}
	}
}

func (p *Parser) getSortedList(data map[string]*Parsers) []string {
	result := make([]string, len(data))
	index := 0
	for parserSectionName := range data {
		result[index] = parserSectionName
		index++
	}
	sort.Strings(result)
	return result
}

//String returns configuration in writable form
func (p *Parser) String() string {
	p.lock()
	defer p.unLock()
	var result strings.Builder

	p.writeParsers("", p.Parsers[Comments][CommentsSectionName].parsers, &result, false)
	p.writeParsers("global", p.Parsers[Global][GlobalSectionName].parsers, &result, true)
	p.writeParsers("defaults", p.Parsers[Defaults][DefaultSectionName].parsers, &result, true)

	sections := []Section{UserList, Peers, Mailers, Resolvers, Cache, Frontends, Backends, Listen, Program}

	for _, section := range sections {
		sortedSections := p.getSortedList(p.Parsers[section])
		for _, sectionName := range sortedSections {
			p.writeParsers(fmt.Sprintf("%s %s", section, sectionName), p.Parsers[section][sectionName].parsers, &result, true)
		}
	}
	return result.String()
}

func (p *Parser) Save(filename string) error {
	d1 := []byte(p.String())
	err := ioutil.WriteFile(filename, d1, 0644)
	if err != nil {
		return err
	}
	return nil
}

//ProcessLine parses line plus determines if we need to change state
func (p *Parser) ProcessLine(line string, parts, previousParts []string, comment string, config ConfiguredParsers) ConfiguredParsers {
	for _, parser := range config.Active.parsers {
		if newState, err := parser.Parse(line, parts, previousParts, comment); err == nil {
			//should we have an option to remove it when found?
			if newState != "" {
				//log.Printf("change state from %s to %s\n", state, newState)
				config.State = newState
				if config.State == "" {
					config.Active = *config.Comments
				}
				if config.State == "defaults" {
					config.Active = *config.Defaults
				}
				if config.State == "global" {
					config.Active = *config.Global
				}
				if config.State == "frontend" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Frontend = getFrontendParser()
					p.Parsers[Frontends][data.Name] = config.Frontend
					config.Active = *config.Frontend
				}
				if config.State == "backend" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Backend = getBackendParser()
					p.Parsers[Backends][data.Name] = config.Backend
					config.Active = *config.Backend
				}
				if config.State == "listen" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Listen = getListenParser()
					p.Parsers[Listen][data.Name] = config.Listen
					config.Active = *config.Listen
				}
				if config.State == "resolvers" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Resolver = getResolverParser()
					p.Parsers[Resolvers][data.Name] = config.Resolver
					config.Active = *config.Resolver
				}
				if config.State == "userlist" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Userlist = getUserlistParser()
					p.Parsers[UserList][data.Name] = config.Userlist
					config.Active = *config.Userlist
				}
				if config.State == "peers" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Peers = getPeersParser()
					p.Parsers[Peers][data.Name] = config.Peers
					config.Active = *config.Peers
				}
				if config.State == "mailers" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Mailers = getMailersParser()
					p.Parsers[Mailers][data.Name] = config.Mailers
					config.Active = *config.Mailers
				}
				if config.State == "cache" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Cache = getCacheParser()
					p.Parsers[Cache][data.Name] = config.Cache
					config.Active = *config.Cache
				}
				if config.State == "program" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Program = getProgramParser()
					p.Parsers[Program][data.Name] = config.Program
					config.Active = *config.Program
				}
			}
			break
		}
	}
	return config
}

func (p *Parser) LoadData(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return p.ParseData(string(dat))
}

func (p *Parser) ParseData(dat string) error {
	p.mutex = &sync.Mutex{}

	p.Parsers = map[Section]map[string]*Parsers{}
	p.Parsers[Comments] = map[string]*Parsers{
		CommentsSectionName: getStartParser(),
	}
	p.Parsers[Defaults] = map[string]*Parsers{
		DefaultSectionName: getDefaultParser(),
	}
	p.Parsers[Global] = map[string]*Parsers{
		GlobalSectionName: getGlobalParser(),
	}
	p.Parsers[Frontends] = map[string]*Parsers{}
	p.Parsers[Backends] = map[string]*Parsers{}
	p.Parsers[Listen] = map[string]*Parsers{}
	p.Parsers[Resolvers] = map[string]*Parsers{}
	p.Parsers[UserList] = map[string]*Parsers{}
	p.Parsers[Peers] = map[string]*Parsers{}
	p.Parsers[Mailers] = map[string]*Parsers{}
	p.Parsers[Cache] = map[string]*Parsers{}
	p.Parsers[Program] = map[string]*Parsers{}

	parsers := ConfiguredParsers{
		State:    "",
		Active:   *p.Parsers[Comments][CommentsSectionName],
		Comments: p.Parsers[Comments][CommentsSectionName],
		Defaults: p.Parsers[Defaults][DefaultSectionName],
		Global:   p.Parsers[Global][GlobalSectionName],
	}

	lines := common.StringSplitIgnoreEmpty(dat, '\n')
	previousLine := []string{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts, comment := common.StringSplitWithCommentIgnoreEmpty(line, ' ', '\t')
		if len(parts) == 0 && comment != "" {
			parts = []string{""}
		}
		if len(parts) == 0 {
			continue
		}
		parsers = p.ProcessLine(line, parts, previousLine, comment, parsers)
		previousLine = parts
	}
	return nil
}
