package parser

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/parsers/extra"
)

//Parser reads and writes configuration on given file
type Parser struct {
	Comments  *ParserTypes
	Default   *ParserTypes
	Global    *ParserTypes
	Frontends map[string]*ParserTypes
	Backends  map[string]*ParserTypes
	Listen    map[string]*ParserTypes
	Resolvers map[string]*ParserTypes
	UserLists map[string]*ParserTypes
	Peers     map[string]*ParserTypes
	Mailers   map[string]*ParserTypes
}

func (p *Parser) get(data map[string]*ParserTypes, key string, attribute string) (ParserType, error) {
	for _, parser := range data[key].parsers {
		if parser.GetParserName() == attribute && parser.Valid() {
			return parser, nil
		}
	}
	return nil, fmt.Errorf("attribute not found")
}

//GetDefaultsAttr get attribute from defaults section
func (p *Parser) GetDefaultsAttr(attribute string) (ParserType, error) {
	return p.Default.Get(attribute)
}

//GetGlobalAttr get attribute from global section
func (p *Parser) GetGlobalAttr(attribute string) (ParserType, error) {
	return p.Global.Get(attribute)
}

//NewGlobalAttr adds attribute to global section, if exists, replaces it
func (p *Parser) NewGlobalAttr(parser ParserType) {
	p.Global.Set(parser)
}

//GetUserlistAttr get attribute from listen sections
func (p *Parser) GetUserlistAttr(section string, attribute string) (ParserType, error) {
	return p.get(p.UserLists, section, attribute)
}

//GetUserlistAttr get attribute from peers sections
func (p *Parser) GetPeersAttr(section string, attribute string) (ParserType, error) {
	return p.get(p.Peers, section, attribute)
}

//GetMailersAttr get attribute from mailer sections
func (p *Parser) GetMailersAttr(section string, attribute string) (ParserType, error) {
	return p.get(p.Mailers, section, attribute)
}

//GetFrontendAttr get attribute from frontend sections
func (p *Parser) GetFrontendAttr(frontendName string, attribute string) (ParserType, error) {
	return p.get(p.Frontends, frontendName, attribute)
}

//GetBackendAttr get attribute from backend sections
func (p *Parser) GetBackendAttr(backendName string, attribute string) (ParserType, error) {
	return p.get(p.Backends, backendName, attribute)
}

//GetListenAttr get attribute from listen sections
func (p *Parser) GetListenAttr(section string, attribute string) (ParserType, error) {
	return p.get(p.Listen, section, attribute)
}

func (p *Parser) writeParsers(parsers []ParserType, result *strings.Builder) {
	for _, parser := range parsers {
		if !parser.Valid() {
			continue
		}
		for _, line := range parser.Result(true) {
			result.WriteString(line)
			result.WriteString("\n")
		}
	}
}

//String returns configuration in writable form
func (p *Parser) String() string {
	var result strings.Builder

	p.writeParsers(p.Comments.parsers, &result)

	result.WriteString("\ndefaults ")
	result.WriteString("\n")
	p.writeParsers(p.Default.parsers, &result)

	result.WriteString("\nglobal ")
	result.WriteString("\n")
	p.writeParsers(p.Global.parsers, &result)

	for sectionName, section := range p.UserLists {
		result.WriteString("\nuserlist ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}

	for sectionName, section := range p.Peers {
		result.WriteString("\npeers ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}

	for sectionName, section := range p.Mailers {
		result.WriteString("\nmailers ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}

	for sectionName, section := range p.Resolvers {
		result.WriteString("\nresolvers ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}

	for sectionName, section := range p.Frontends {
		result.WriteString("\nfrontend ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}

	for sectionName, section := range p.Backends {
		result.WriteString("\nbackend ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}
	for sectionName, section := range p.Listen {
		result.WriteString("\nlisten ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
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
func (p *Parser) ProcessLine(line string, parts, previousParts []string, config ConfiguredParsers) ConfiguredParsers {
	for _, parser := range config.Active.parsers {
		if newState, err := parser.Parse(line, parts, previousParts); err == nil {
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
					sectionName := parser.(*extra.SectionName)
					config.Frontend = getFrontendParser()
					p.Frontends[sectionName.SectionName] = config.Frontend
					config.Active = *config.Frontend
				}
				if config.State == "backend" {
					sectionName := parser.(*extra.SectionName)
					config.Backend = getBackendParser()
					p.Backends[sectionName.SectionName] = config.Backend
					config.Active = *config.Backend
				}
				if config.State == "listen" {
					sectionName := parser.(*extra.SectionName)
					config.Listen = getListenParser()
					p.Listen[sectionName.SectionName] = config.Listen
					config.Active = *config.Listen
				}
				if config.State == "resolvers" {
					sectionName := parser.(*extra.SectionName)
					config.Resolver = getResolverParser()
					p.Resolvers[sectionName.SectionName] = config.Resolver
					config.Active = *config.Resolver
				}
				if config.State == "userlist" {
					sectionName := parser.(*extra.SectionName)
					config.Userlist = getUserlistParser()
					p.UserLists[sectionName.SectionName] = config.Userlist
					config.Active = *config.Userlist
				}
				if config.State == "peers" {
					sectionName := parser.(*extra.SectionName)
					config.Peers = getPeersParser()
					p.Peers[sectionName.SectionName] = config.Peers
					config.Active = *config.Peers
				}
				if config.State == "mailers" {
					sectionName := parser.(*extra.SectionName)
					config.Mailers = getMailersParser()
					p.Mailers[sectionName.SectionName] = config.Mailers
					config.Active = *config.Mailers
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
	p.Comments = getStartParser()
	p.Default = getDefaultParser()
	p.Global = getGlobalParser()

	p.Frontends = map[string]*ParserTypes{}
	p.Backends = map[string]*ParserTypes{}
	p.Listen = map[string]*ParserTypes{}
	p.Resolvers = map[string]*ParserTypes{}
	p.UserLists = map[string]*ParserTypes{}
	p.Peers = map[string]*ParserTypes{}
	p.Mailers = map[string]*ParserTypes{}

	parsers := ConfiguredParsers{
		State:    "",
		Active:   *p.Comments,
		Comments: p.Comments,
		Defaults: p.Default,
		Global:   p.Global,
	}

	lines := common.StringSplitIgnoreEmpty(string(dat), '\n')
	previousLine := []string{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := common.StringSplitIgnoreEmpty(line, ' ')
		if len(parts) == 0 {
			continue
		}
		parsers = p.ProcessLine(line, parts, previousLine, parsers)
		previousLine = parts
	}
	return nil
}
