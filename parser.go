package parser

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/parsers/extra"
	"github.com/haproxytech/config-parser/types"
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

//GetDefaultsAttr get attribute from defaults section
func (p *Parser) GetDefaultsAttr(attribute string) (common.ParserData, error) {
	return p.Default.Get(attribute)
}

//GetGlobalAttr get attribute from global section
func (p *Parser) GetGlobalAttr(attribute string) (common.ParserData, error) {
	return p.Global.Get(attribute)
}

//GetUserlistAttr get attribute from listen sections
func (p *Parser) GetUserlistAttr(section string, attribute string) (common.ParserData, error) {
	return p.get(p.UserLists, section, attribute)
}

//GetUserlistAttr get attribute from peers sections
func (p *Parser) GetPeersAttr(section string, attribute string) (common.ParserData, error) {
	return p.get(p.Peers, section, attribute)
}

//GetMailersAttr get attribute from mailer sections
func (p *Parser) GetMailersAttr(section string, attribute string) (common.ParserData, error) {
	return p.get(p.Mailers, section, attribute)
}

//GetMailersAttr get attribute from mailer sections
func (p *Parser) GetResolversAttr(section string, attribute string) (common.ParserData, error) {
	return p.get(p.Resolvers, section, attribute)
}

//GetFrontendAttr get attribute from frontend sections
func (p *Parser) GetFrontendAttr(frontendName string, attribute string) (common.ParserData, error) {
	return p.get(p.Frontends, frontendName, attribute)
}

//GetBackendAttr get attribute from backend sections
func (p *Parser) GetBackendAttr(backendName string, attribute string) (common.ParserData, error) {
	return p.get(p.Backends, backendName, attribute)
}

//GetListenAttr get attribute from listen sections
func (p *Parser) GetListenAttr(section string, attribute string) (common.ParserData, error) {
	return p.get(p.Listen, section, attribute)
}

func (p *Parser) writeParsers(parsers []ParserType, result *strings.Builder) {
	for _, parser := range parsers {
		lines, err := parser.Result(true)
		if err != nil {
			continue
		}
		for _, line := range lines {
			result.WriteString("  ")
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

	p.writeParsers(p.Comments.parsers, &result)

	result.WriteString("\ndefaults ")
	result.WriteString("\n")
	p.writeParsers(p.Default.parsers, &result)

	result.WriteString("\nglobal ")
	result.WriteString("\n")
	p.writeParsers(p.Global.parsers, &result)

	for parserSectionName, section := range p.UserLists {
		result.WriteString("\nuserlist ")
		result.WriteString(parserSectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}

	for parserSectionName, section := range p.Peers {
		result.WriteString("\npeers ")
		result.WriteString(parserSectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}

	for parserSectionName, section := range p.Mailers {
		result.WriteString("\nmailers ")
		result.WriteString(parserSectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}

	for parserSectionName, section := range p.Resolvers {
		result.WriteString("\nresolvers ")
		result.WriteString(parserSectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}

	for parserSectionName, section := range p.Frontends {
		result.WriteString("\nfrontend ")
		result.WriteString(parserSectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}

	for parserSectionName, section := range p.Backends {
		result.WriteString("\nbackend ")
		result.WriteString(parserSectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}
	for parserSectionName, section := range p.Listen {
		result.WriteString("\nlisten ")
		result.WriteString(parserSectionName)
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
					p.Frontends[data.Name] = config.Frontend
					config.Active = *config.Frontend
				}
				if config.State == "backend" {
					parserSectionName := parser.(*extra.SectionName)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Backend = getBackendParser()
					p.Backends[data.Name] = config.Backend
					config.Active = *config.Backend
				}
				if config.State == "listen" {
					parserSectionName := parser.(*extra.SectionName)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Listen = getListenParser()
					p.Listen[data.Name] = config.Listen
					config.Active = *config.Listen
				}
				if config.State == "resolvers" {
					parserSectionName := parser.(*extra.SectionName)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Resolver = getResolverParser()
					p.Resolvers[data.Name] = config.Resolver
					config.Active = *config.Resolver
				}
				if config.State == "userlist" {
					parserSectionName := parser.(*extra.SectionName)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Userlist = getUserlistParser()
					p.UserLists[data.Name] = config.Userlist
					config.Active = *config.Userlist
				}
				if config.State == "peers" {
					parserSectionName := parser.(*extra.SectionName)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Peers = getPeersParser()
					p.Peers[data.Name] = config.Peers
					config.Active = *config.Peers
				}
				if config.State == "mailers" {
					parserSectionName := parser.(*extra.SectionName)
					rawData, _ := parserSectionName.Get(false)
					data := rawData.(*types.Section)
					config.Mailers = getMailersParser()
					p.Mailers[data.Name] = config.Mailers
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
		parts, comment := common.StringSplitWithCommentIgnoreEmpty(line, ' ')
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
