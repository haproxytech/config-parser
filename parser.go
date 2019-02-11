package parser

import (
	"fmt"
	"io/ioutil"
	"strings"

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
)

const (
	CommentsSectionName = "data"
	GlobalSectionName   = "data"
	DefaultSectionName  = "data"
)

//Parser reads and writes configuration on given file
type Parser struct {
	Parsers map[Section]map[string]*ParserTypes
}

func (p *Parser) get(data map[string]*ParserTypes, key string, attribute string) (common.ParserData, error) {
	for _, parser := range data[key].parsers {
		if parser.GetParserName() == attribute {
			return parser.Get(false)
		}
	}
	return nil, fmt.Errorf("attribute not found")
}

func (p *Parser) getOrCreate(data map[string]*ParserTypes, key string, attribute string, createIfNotExist bool) (common.ParserData, error) {
	for _, parser := range data[key].parsers {
		if parser.GetParserName() == attribute {
			return parser.Get(createIfNotExist)
		}
	}
	return nil, fmt.Errorf("attribute not found")
}

//Get get attribute from defaults section
func (p *Parser) Get(sectionType Section, sectionName string, attribute string, createIfNotExist ...bool) (common.ParserData, error) {
	st, ok := p.Parsers[sectionType]
	if !ok {
		return nil, errors.SectionMissingErr
	}
	section, ok := st[sectionName]
	if !ok {
		return nil, fmt.Errorf("Section [%s] not found", sectionName)
	}
	createNew := false
	if len(createIfNotExist) > 0 && createIfNotExist[0] {
		createNew = true
	}
	return section.Get(attribute, createNew)
}

//SectionsGet lists all sections of certain type
func (p *Parser) SectionsGet(sectionType Section) ([]string, error) {
	st, ok := p.Parsers[sectionType]
	if !ok {
		return nil, errors.SectionMissingErr
	}
	result := make([]string, len(st))
	index := 0
	for sectionName, _ := range st {
		result[index] = sectionName
		index++
	}
	return result, nil
}

//Set sets attribute from defaults section, can be nil to disable/remove
func (p *Parser) Set(sectionType Section, sectionName string, attribute string, data common.ParserData) error {
	st, ok := p.Parsers[sectionType]
	if !ok {
		return errors.SectionMissingErr
	}
	section, ok := st[sectionName]
	if !ok {
		return fmt.Errorf("Section [%s] not found", sectionName)
	}
	return section.Set(attribute, data)
}

//HasParser checks if we have a parser for attribute
func (p *Parser) HasParser(sectionType Section, attribute string) bool {
	st, ok := p.Parsers[sectionType]
	if !ok {
		return false
	}
	sectionName := ""
	for name, _ := range st {
		sectionName = name
		break
	}
	section, ok := st[sectionName]
	if !ok {
		return false
	}
	return section.HasParser(attribute)
}

func (p *Parser) writeParsers(parsers []ParserType, result *strings.Builder, useIndentation bool) {
	for _, parser := range parsers {
		lines, err := parser.Result(true)
		if err != nil {
			continue
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

//String returns configuration in writable form
func (p *Parser) String() string {
	var result strings.Builder

	p.writeParsers(p.Parsers[Comments][CommentsSectionName].parsers, &result, false)

	result.WriteString("\ndefaults ")
	result.WriteString("\n")
	p.writeParsers(p.Parsers[Defaults][DefaultSectionName].parsers, &result, true)

	result.WriteString("\nglobal ")
	result.WriteString("\n")
	p.writeParsers(p.Parsers[Global][GlobalSectionName].parsers, &result, true)

	for parserSectionName, section := range p.Parsers[UserList] {
		result.WriteString("\nuserlist ")
		result.WriteString(parserSectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result, true)
	}

	for parserSectionName, section := range p.Parsers[Peers] {
		result.WriteString("\npeers ")
		result.WriteString(parserSectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result, true)
	}

	for parserSectionName, section := range p.Parsers[Mailers] {
		result.WriteString("\nmailers ")
		result.WriteString(parserSectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result, true)
	}

	for parserSectionName, section := range p.Parsers[Resolvers] {
		result.WriteString("\nresolvers ")
		result.WriteString(parserSectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result, true)
	}

	for parserSectionName, section := range p.Parsers[Cache] {
		result.WriteString("\ncache ")
		result.WriteString(parserSectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result, true)
	}

	for parserSectionName, section := range p.Parsers[Frontends] {
		result.WriteString("\nfrontend ")
		result.WriteString(parserSectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result, true)
	}

	for parserSectionName, section := range p.Parsers[Backends] {
		result.WriteString("\nbackend ")
		result.WriteString(parserSectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result, true)
	}

	for parserSectionName, section := range p.Parsers[Listen] {
		result.WriteString("\nlisten ")
		result.WriteString(parserSectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result, true)
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
					parserSectionName := parser.(*extra.SectionName)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Frontend = getFrontendParser()
					p.Parsers[Frontends][data.Name] = config.Frontend
					config.Active = *config.Frontend
				}
				if config.State == "backend" {
					parserSectionName := parser.(*extra.SectionName)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Backend = getBackendParser()
					p.Parsers[Backends][data.Name] = config.Backend
					config.Active = *config.Backend
				}
				if config.State == "listen" {
					parserSectionName := parser.(*extra.SectionName)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Listen = getListenParser()
					p.Parsers[Listen][data.Name] = config.Listen
					config.Active = *config.Listen
				}
				if config.State == "resolvers" {
					parserSectionName := parser.(*extra.SectionName)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Resolver = getResolverParser()
					p.Parsers[Resolvers][data.Name] = config.Resolver
					config.Active = *config.Resolver
				}
				if config.State == "userlist" {
					parserSectionName := parser.(*extra.SectionName)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Userlist = getUserlistParser()
					p.Parsers[UserList][data.Name] = config.Userlist
					config.Active = *config.Userlist
				}
				if config.State == "peers" {
					parserSectionName := parser.(*extra.SectionName)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Peers = getPeersParser()
					p.Parsers[Peers][data.Name] = config.Peers
					config.Active = *config.Peers
				}
				if config.State == "mailers" {
					parserSectionName := parser.(*extra.SectionName)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Mailers = getMailersParser()
					p.Parsers[Mailers][data.Name] = config.Mailers
					config.Active = *config.Mailers
				}
				if config.State == "cache" {
					parserSectionName := parser.(*extra.SectionName)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Cache = getCacheParser()
					p.Parsers[Cache][data.Name] = config.Cache
					config.Active = *config.Cache
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

	p.Parsers = map[Section]map[string]*ParserTypes{}
	p.Parsers[Comments] = map[string]*ParserTypes{
		CommentsSectionName: getStartParser(),
	}
	p.Parsers[Defaults] = map[string]*ParserTypes{
		DefaultSectionName: getDefaultParser(),
	}
	p.Parsers[Global] = map[string]*ParserTypes{
		GlobalSectionName: getGlobalParser(),
	}
	p.Parsers[Frontends] = map[string]*ParserTypes{}
	p.Parsers[Backends] = map[string]*ParserTypes{}
	p.Parsers[Listen] = map[string]*ParserTypes{}
	p.Parsers[Resolvers] = map[string]*ParserTypes{}
	p.Parsers[UserList] = map[string]*ParserTypes{}
	p.Parsers[Peers] = map[string]*ParserTypes{}
	p.Parsers[Mailers] = map[string]*ParserTypes{}
	p.Parsers[Cache] = map[string]*ParserTypes{}

	parsers := ConfiguredParsers{
		State:    "",
		Active:   *p.Parsers[Comments][CommentsSectionName],
		Comments: p.Parsers[Comments][CommentsSectionName],
		Defaults: p.Parsers[Defaults][DefaultSectionName],
		Global:   p.Parsers[Global][GlobalSectionName],
	}

	lines := common.StringSplitIgnoreEmpty(string(dat), '\n')
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
