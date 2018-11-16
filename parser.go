package parser

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/haproxytech/config-parser/helpers"
	"github.com/haproxytech/config-parser/parsers/extra"
)

//Parser reads and writes configuration on given file
type Parser struct {
	Comments  ParserTypes
	Default   ParserTypes
	Global    ParserTypes
	Frontends map[string]ParserTypes
	Backends  map[string]ParserTypes
	Listen    map[string]ParserTypes
	Resolvers map[string]ParserTypes
}

func (p *Parser) get(data map[string]ParserTypes, key string, atrtibute string) (ParserType, error) {
	for _, parser := range data[key].parsers {
		if parser.GetParserName() == atrtibute && parser.Valid() {
			return parser, nil
		}
	}
	return nil, fmt.Errorf("atrtibute not found")
}

//GetDefaultsAttr get atribute from defaults section
func (p *Parser) GetDefaultsAttr(atrtibute string) (ParserType, error) {
	return p.Default.Get(atrtibute)
}

//GetGlobalAttr get atribute from global section
func (p *Parser) GetGlobalAttr(atrtibute string) (ParserType, error) {
	return p.Global.Get(atrtibute)
}

//GetFrontendAttr get atribute from frontend sections
func (p *Parser) GetFrontendAttr(frontendName string, atrtibute string) (ParserType, error) {
	return p.get(p.Frontends, frontendName, atrtibute)
}

//GetBackendAttr get atribute from backend sections
func (p *Parser) GetBackendAttr(backendName string, atrtibute string) (ParserType, error) {
	return p.get(p.Backends, backendName, atrtibute)
}

//GetListenAttr get atribute from listen sections
func (p *Parser) GetListenAttr(section string, atrtibute string) (ParserType, error) {
	return p.get(p.Listen, section, atrtibute)
}

//String returns configuration in writable form
func (p *Parser) String() string {
	var result strings.Builder

	for _, parser := range p.Comments.parsers {
		if !parser.Valid() {
			continue
		}
		for _, line := range parser.String() {
			result.WriteString(line)
			result.WriteString("\n")
		}
	}

	result.WriteString("\ndefaults ")
	result.WriteString("\n")
	for _, parser := range p.Default.parsers {
		if !parser.Valid() {
			continue
		}
		for _, line := range parser.String() {
			result.WriteString(line)
			result.WriteString("\n")
		}
	}
	result.WriteString("\nglobal ")
	result.WriteString("\n")
	for _, parser := range p.Global.parsers {
		if !parser.Valid() {
			continue
		}
		for _, line := range parser.String() {
			result.WriteString(line)
			result.WriteString("\n")
		}
	}

	for sectionName, section := range p.Resolvers {
		result.WriteString("\nresolvers ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		for _, parser := range section.parsers {
			if !parser.Valid() {
				continue
			}
			for _, line := range parser.String() {
				result.WriteString(line)
				result.WriteString("\n")
			}
		}
	}

	for sectionName, section := range p.Frontends {
		result.WriteString("\nfrontend ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		for _, parser := range section.parsers {
			if !parser.Valid() {
				continue
			}
			for _, line := range parser.String() {
				result.WriteString(line)
				result.WriteString("\n")
			}
		}
	}

	for sectionName, section := range p.Backends {
		result.WriteString("\nbackend ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		for _, parser := range section.parsers {
			if !parser.Valid() {
				continue
			}
			for _, line := range parser.String() {
				result.WriteString(line)
				result.WriteString("\n")
			}
		}
	}
	for sectionName, section := range p.Listen {
		result.WriteString("\nlisten ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		for _, parser := range section.parsers {
			if !parser.Valid() {
				continue
			}
			for _, line := range parser.String() {
				result.WriteString(line)
				result.WriteString("\n")
			}
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
func (p *Parser) ProcessLine(state string, activeParser ParserTypes, line, part, previousLine string, parserFrontend, parserBackend, parserListen, parserResolver ParserTypes) (newState string, newParserFrontend, newParserBackend, newParserListen, newParserResolver ParserTypes) {
	for _, parser := range activeParser.parsers {
		if newState, err := parser.Parse(line, part, previousLine); err == nil {
			//should we have an option to remove it when found?
			if newState != "" {
				//log.Printf("change state from %s to %s\n", state, newState)
				state = newState
				if state == "frontend" {
					sectionName := parser.(*extra.SectionName)
					parserFrontend = getFrontendParser()
					p.Frontends[sectionName.SectionName] = parserFrontend
				}
				if state == "backend" {
					sectionName := parser.(*extra.SectionName)
					parserBackend = getBackendParser()
					p.Backends[sectionName.SectionName] = parserBackend
				}
				if state == "listen" {
					sectionName := parser.(*extra.SectionName)
					parserListen = getListenParser()
					p.Listen[sectionName.SectionName] = parserListen
				}
				if state == "resolvers" {
					sectionName := parser.(*extra.SectionName)
					parserResolver = getResolverParser()
					p.Resolvers[sectionName.SectionName] = parserResolver
				}
			}
			break
		}
	}
	return state, parserFrontend, parserBackend, parserListen, parserResolver
}

func (p *Parser) LoadData(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	p.Comments = getStartParser()
	p.Default = getDefaultParser()
	p.Global = getGlobalParser()

	p.Frontends = map[string]ParserTypes{}
	p.Backends = map[string]ParserTypes{}
	p.Listen = map[string]ParserTypes{}
	p.Resolvers = map[string]ParserTypes{}

	var parserFrontend ParserTypes
	var parserBackend ParserTypes
	var parserListen ParserTypes
	var parserResolver ParserTypes

	lines := helpers.StringSplitIgnoreEmpty(string(dat), '\n')
	state := ""
	previousLine := ""
	for _, part := range lines {
		if part == "" {
			continue
		}
		line := strings.Trim(part, " ")
		switch state {
		case "":
			state, parserFrontend, parserBackend, parserListen, parserResolver =
				p.ProcessLine(state, p.Comments, line, part, previousLine, parserFrontend, parserBackend, parserListen, parserListen)
			previousLine = line
		case "defaults":
			state, parserFrontend, parserBackend, parserListen, parserResolver =
				p.ProcessLine(state, p.Default, line, part, previousLine, parserFrontend, parserBackend, parserListen, parserResolver)
			previousLine = line
		case "global":
			state, parserFrontend, parserBackend, parserListen, parserResolver =
				p.ProcessLine(state, p.Global, line, part, previousLine, parserFrontend, parserBackend, parserListen, parserResolver)
			previousLine = line
		case "frontend":
			state, parserFrontend, parserBackend, parserListen, parserResolver =
				p.ProcessLine(state, parserFrontend, line, part, previousLine, parserFrontend, parserBackend, parserListen, parserResolver)
			previousLine = line
		case "backend":
			state, parserFrontend, parserBackend, parserListen, parserResolver =
				p.ProcessLine(state, parserBackend, line, part, previousLine, parserFrontend, parserBackend, parserListen, parserResolver)
			previousLine = line
		case "listen":
			state, parserFrontend, parserBackend, parserListen, parserResolver =
				p.ProcessLine(state, parserListen, line, part, previousLine, parserFrontend, parserBackend, parserListen, parserResolver)
			previousLine = line
		case "resolvers":
			state, parserFrontend, parserBackend, parserListen, parserResolver =
				p.ProcessLine(state, parserResolver, line, part, previousLine, parserFrontend, parserBackend, parserListen, parserResolver)
			previousLine = line
		}
	}
	return nil
}
