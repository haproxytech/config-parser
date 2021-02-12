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
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
	"sync"

	"github.com/google/renameio"
	"github.com/haproxytech/config-parser/v3/common"
	"github.com/haproxytech/config-parser/v3/errors"
	"github.com/haproxytech/config-parser/v3/parsers/extra"
	"github.com/haproxytech/config-parser/v3/types"
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
	// spoe sections
	SPOEAgent   Section = "spoe-agent"
	SPOEGroup   Section = "spoe-group"
	SPOEMessage Section = "spoe-message"
)

const (
	CommentsSectionName = "data"
	GlobalSectionName   = "data"
	DefaultSectionName  = "data"
)

type Options struct {
	UseV2HTTPCheck bool
	UseMd5Hash     bool
}

// Parser reads and writes configuration on given file
type Parser struct {
	Parsers map[Section]map[string]*Parsers
	Options Options
	mutex   *sync.Mutex
}

func (p *Parser) lock() {
	p.mutex.Lock()
}

func (p *Parser) unLock() {
	p.mutex.Unlock()
}

// Get get attribute from defaults section
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

// GetOne get attribute from defaults section
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

// SectionsGet lists all sections of certain type
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

// SectionsDelete deletes one section of sectionType
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

// SectionsCreate creates one section of sectionType
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

// Set sets attribute from defaults section, can be nil to disable/remove
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

// Delete remove attribute on defined index, in case of single attributes, index is ignored
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

// Insert put attribute on defined index, in case of single attributes, index is ignored
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

// HasParser checks if we have a parser for attribute
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

func (p *Parser) writeSection(sectionName string, comments []string, result io.StringWriter) {
	_, _ = result.WriteString("\n")
	for _, line := range comments {
		_, _ = result.WriteString("# ")
		_, _ = result.WriteString(line)
		_, _ = result.WriteString("\n")
	}
	_, _ = result.WriteString(sectionName)
	_, _ = result.WriteString(" \n")
}

func (p *Parser) writeParsers(sectionName string, parsersData *Parsers, result io.StringWriter, useIndentation bool) {
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
	for _, parserName := range parsersData.ParserSequence {
		parser := parsersData.Parsers[string(parserName)]
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
				_, _ = result.WriteString("  ")
			}
			_, _ = result.WriteString("# ")
			_, _ = result.WriteString(line)
			_, _ = result.WriteString("\n")
		}
		for _, line := range lines {
			if useIndentation {
				_, _ = result.WriteString("  ")
			}
			_, _ = result.WriteString(line.Data)
			if line.Comment != "" {
				_, _ = result.WriteString(" # ")
				_, _ = result.WriteString(line.Comment)
			}
			_, _ = result.WriteString("\n")
		}
	}
	for _, line := range parsersData.PostComments {
		if useIndentation {
			_, _ = result.WriteString("  ")
		}
		_, _ = result.WriteString("# ")
		_, _ = result.WriteString(line)
		_, _ = result.WriteString("\n")
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

// String returns configuration in writable form
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
	if p.Options.UseMd5Hash {
		data, err := p.StringWithHash()
		if err != nil {
			return err
		}
		return p.save([]byte(data), filename)
	}
	return p.save([]byte(p.String()), filename)
}

func (p *Parser) save(data []byte, filename string) error {
	err := renameio.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (p *Parser) StringWithHash() (string, error) {
	var result strings.Builder
	content := p.String()
	//nolint:gosec
	hash := md5.Sum([]byte(content))
	result.WriteString(fmt.Sprintf("# _md5hash=%x\n", hash))
	result.WriteString(content)
	if err := p.Set(Comments, CommentsSectionName, "# _md5hash", &types.ConfigHash{Value: fmt.Sprintf("%x", hash)}); err != nil {
		return "", err
	}

	return result.String(), nil
}

// ProcessLine parses line plus determines if we need to change state
func (p *Parser) ProcessLine(line string, parts, previousParts []string, comment string, config ConfiguredParsers) ConfiguredParsers { //nolint:gocognit,gocyclo
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
	for _, section := range config.Active.ParserSequence {
		parser := config.Active.Parsers[string(section)]
		if newState, err := parser.PreParse(line, parts, previousParts, config.ActiveComments, comment); err == nil {
			if newState != "" {
				// log.Printf("change state from %s to %s\n", state, newState)
				if config.ActiveComments != nil {
					config.Active.PostComments = config.ActiveComments
				}
				config.State = newState
				if config.State == "" {
					config.Active = config.Comments
				}
				if config.State == "defaults" {
					config.Active = config.Defaults
				}
				if config.State == "global" {
					config.Active = config.Global
				}
				if config.State == "frontend" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Frontend = p.getFrontendParser()
					p.Parsers[Frontends][data.Name] = config.Frontend
					config.Active = config.Frontend
				}
				if config.State == "backend" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Backend = p.getBackendParser()
					p.Parsers[Backends][data.Name] = config.Backend
					config.Active = config.Backend
				}
				if config.State == "listen" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Listen = p.getListenParser()
					p.Parsers[Listen][data.Name] = config.Listen
					config.Active = config.Listen
				}
				if config.State == "resolvers" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Resolver = p.getResolverParser()
					p.Parsers[Resolvers][data.Name] = config.Resolver
					config.Active = config.Resolver
				}
				if config.State == "userlist" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Userlist = p.getUserlistParser()
					p.Parsers[UserList][data.Name] = config.Userlist
					config.Active = config.Userlist
				}
				if config.State == "peers" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Peers = p.getPeersParser()
					p.Parsers[Peers][data.Name] = config.Peers
					config.Active = config.Peers
				}
				if config.State == "mailers" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Mailers = p.getMailersParser()
					p.Parsers[Mailers][data.Name] = config.Mailers
					config.Active = config.Mailers
				}
				if config.State == "cache" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Cache = p.getCacheParser()
					p.Parsers[Cache][data.Name] = config.Cache
					config.Active = config.Cache
				}
				if config.State == "program" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Program = p.getProgramParser()
					p.Parsers[Program][data.Name] = config.Program
					config.Active = config.Program
				}
				if config.State == "http-errors" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.HTTPErrors = p.getHTTPErrorsParser()
					p.Parsers[HTTPErrors][data.Name] = config.HTTPErrors
					config.Active = config.HTTPErrors
				}
				if config.State == "ring" {
					parserSectionName := parser.(*extra.Section)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Ring = p.getRingParser()
					p.Parsers[Ring][data.Name] = config.Ring
					config.Active = config.Ring
				}
				if config.State == "snippet_beg" {
					config.Previous = config.Active
					config.Active = &Parsers{
						Parsers:        map[string]ParserInterface{"config-snippet": parser},
						ParserSequence: []Section{"config-snippet"},
					}
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
		CommentsSectionName: p.getStartParser(),
	}

	p.Parsers[Defaults] = map[string]*Parsers{
		DefaultSectionName: p.getDefaultParser(),
	}

	p.Parsers[Global] = map[string]*Parsers{
		GlobalSectionName: p.getGlobalParser(),
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
		parts, comment := common.StringSplitWithCommentIgnoreEmpty(line)
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
