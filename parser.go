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
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
	"sync"

	"github.com/google/renameio"
	"github.com/haproxytech/config-parser/v2/common"
	"github.com/haproxytech/config-parser/v2/errors"
	"github.com/haproxytech/config-parser/v2/parsers/extra"
	"github.com/haproxytech/config-parser/v2/types"
)

type Section string

const (
	Comments   Section = "#"
	Defaults   Section = "defaults"
	Global     Section = "global"
	Resolvers  Section = "resolvers"
	UserList   Section = "userlist"
	Peers      Section = "peers"
	Mailers    Section = "mailers"
	Frontends  Section = "frontend"
	Backends   Section = "backend"
	Listen     Section = "listen"
	Cache      Section = "cache"
	Program    Section = "program"
	HTTPErrors Section = "http-errors"
	Ring       Section = "ring"
	//spoe sections
	SPOEAgent   Section = "spoe-agent"
	SPOEGroup   Section = "spoe-group"
	SPOEMessage Section = "spoe-message"
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
		Active:   p.Parsers[Comments][CommentsSectionName],
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

func (p *Parser) writeSection(sectionName string, comments []string, result *strings.Builder) {
	result.WriteString("\n")
	for _, line := range comments {
		result.WriteString("# ")
		result.WriteString(line)
		result.WriteString("\n")
	}
	result.WriteString(sectionName)
	result.WriteString(" \n")
}

func (p *Parser) writeParsers(sectionName string, parsersData *Parsers, result *strings.Builder, useIndentation bool) {
	parsers := parsersData.Parsers
	sectionNameWritten := false
	switch sectionName {
	case "":
		sectionNameWritten = true
	case "global", "defaults":
		break
	default:
		p.writeSection(sectionName, parsersData.PreComments, result)
		sectionNameWritten = true
	}
	for _, parser := range getParsersSequenceForSection(sectionName, parsers) {
		lines, comments, err := parser.ResultAll()
		if err != nil {
			continue
		}
		if !sectionNameWritten {
			p.writeSection(sectionName, parsersData.PreComments, result)
			sectionNameWritten = true
		}
		for _, line := range comments {
			if useIndentation {
				result.WriteString("  ")
			}
			result.WriteString("# ")
			result.WriteString(line)
			result.WriteString("\n")
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
	for _, line := range parsersData.PostComments {
		if useIndentation {
			result.WriteString("  ")
		}
		result.WriteString("# ")
		result.WriteString(line)
		result.WriteString("\n")
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

	p.writeParsers("", p.Parsers[Comments][CommentsSectionName], &result, false)
	p.writeParsers("global", p.Parsers[Global][GlobalSectionName], &result, true)
	p.writeParsers("defaults", p.Parsers[Defaults][DefaultSectionName], &result, true)

	sections := []Section{UserList, Peers, Mailers, Resolvers, Cache, Ring, HTTPErrors, Frontends, Backends, Listen, Program}

	for _, section := range sections {
		sortedSections := p.getSortedList(p.Parsers[section])
		for _, sectionName := range sortedSections {
			p.writeParsers(fmt.Sprintf("%s %s", section, sectionName), p.Parsers[section][sectionName], &result, true)
		}
	}
	return result.String()
}

func (p *Parser) Save(filename string) error {
	d1 := []byte(p.String())
	err := renameio.WriteFile(filename, d1, 0644)
	if err != nil {
		return err
	}
	return nil
}

//ProcessLine parses line plus determines if we need to change state
func (p *Parser) ProcessLine(line string, parts, previousParts []string, comment string, config ConfiguredParsers) ConfiguredParsers {
	if config.State != "" {
		if parts[0] == "" && comment != "" && comment != "##_config-snippet_### BEGIN" && comment != "##_config-snippet_### END" {
			if line[0] == ' ' {
				config.ActiveComments = append(config.ActiveComments, comment)
			} else {
				config.ActiveSectionComments = append(config.ActiveSectionComments, comment)
			}
			return config
		}
	}
	for _, section := range config.ActiveSection {
		parser := config.Active.Parsers[string(section)]
		if parser == nil {
			continue
		}
		if newState, err := parser.PreParse(line, parts, previousParts, config.ActiveComments, comment); err == nil {
			if newState != "" {
				// log.Printf("change state from %s to %s\n", config.State, newState)
				if config.ActiveComments != nil {
					config.Active.PostComments = config.ActiveComments
				}
				config.State = newState
				if config.State == "" {
					config.Active = config.Comments
					config.ActiveSection = parserSequenceStart
				}
				if config.State == "defaults" {
					config.Active = config.Defaults
					config.ActiveSection = parserSequenceDefault
				}
				if config.State == "global" {
					config.Active = config.Global
					config.ActiveSection = parserSequenceGlobal
				}
				if config.State == "frontend" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Frontend = getFrontendParser()
					p.Parsers[Frontends][data.Name] = config.Frontend
					config.Active = config.Frontend
					config.ActiveSection = parserSequenceFrontend
				}
				if config.State == "backend" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Backend = getBackendParser()
					p.Parsers[Backends][data.Name] = config.Backend
					config.Active = config.Backend
					config.ActiveSection = parserSequenceBackend
				}
				if config.State == "listen" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Listen = getListenParser()
					p.Parsers[Listen][data.Name] = config.Listen
					config.Active = config.Listen
					config.ActiveSection = parserSequenceSections
				}
				if config.State == "resolvers" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Resolver = getResolverParser()
					p.Parsers[Resolvers][data.Name] = config.Resolver
					config.Active = config.Resolver
					config.ActiveSection = parserSequenceResolver
				}
				if config.State == "userlist" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Userlist = getUserlistParser()
					p.Parsers[UserList][data.Name] = config.Userlist
					config.Active = config.Userlist
					config.ActiveSection = parserSequenceUserlist
				}
				if config.State == "peers" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Peers = getPeersParser()
					p.Parsers[Peers][data.Name] = config.Peers
					config.Active = config.Peers
					config.ActiveSection = parserSequencePeers
				}
				if config.State == "mailers" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Mailers = getMailersParser()
					p.Parsers[Mailers][data.Name] = config.Mailers
					config.Active = config.Mailers
					config.ActiveSection = parserSequenceMailers
				}
				if config.State == "cache" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Cache = getCacheParser()
					p.Parsers[Cache][data.Name] = config.Cache
					config.Active = config.Cache
					config.ActiveSection = parserSequenceCache
				}
				if config.State == "program" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Program = getProgramParser()
					p.Parsers[Program][data.Name] = config.Program
					config.Active = config.Program
					config.ActiveSection = parserSequenceProgram
				}
				if config.State == "http-errors" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.HTTPErrors = getHTTPErrorsParser()
					p.Parsers[HTTPErrors][data.Name] = config.HTTPErrors
					config.Active = config.HTTPErrors
				}
				if config.State == "ring" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Ring = getRingParser()
					p.Parsers[Ring][data.Name] = config.Ring
					config.Active = config.Ring
				}
				if config.State == "snippet_beg" {
					config.Previous = config.Active
					config.Active = &Parsers{Parsers: map[string]ParserInterface{"configsnippet": parser}}
				}
				if config.State == "snippet_end" {
					config.Active = config.Previous
				}
				if config.ActiveSectionComments != nil {
					config.Active.PreComments = config.ActiveSectionComments
					config.ActiveSectionComments = nil
				}
			}
			config.ActiveComments = nil
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

func (p *Parser) Process(reader io.Reader) error {
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
	p.Parsers[HTTPErrors] = map[string]*Parsers{}
	p.Parsers[Ring] = map[string]*Parsers{}

	parsers := ConfiguredParsers{
		State:          "",
		ActiveComments: nil,
		Active:         p.Parsers[Comments][CommentsSectionName],
		Comments:       p.Parsers[Comments][CommentsSectionName],
		Defaults:       p.Parsers[Defaults][DefaultSectionName],
		Global:         p.Parsers[Global][GlobalSectionName],
	}

	bufferedReader := bufio.NewReader(reader)

	var line string
	var err error
	previousLine := []string{}
	parsers.ActiveSection = parserSequenceStart
	for {
		line, err = bufferedReader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.Trim(line, "\n")

		if line == "" {
			if parsers.State == "" {
				parsers.State = "#"
			}
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
	if parsers.ActiveComments != nil {
		parsers.Active.PostComments = parsers.ActiveComments
	}
	if parsers.ActiveSectionComments != nil {
		parsers.Active.PostComments = append(parsers.Active.PostComments, parsers.ActiveSectionComments...)
	}
	return nil
}

func (p *Parser) ParseData(dat string) error {
	return p.Process(strings.NewReader(dat))
}
