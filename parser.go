package parser

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/haproxytech/config-parser/helpers"
	"github.com/haproxytech/config-parser/parsers"
	"github.com/haproxytech/config-parser/parsers/extra"
	"github.com/haproxytech/config-parser/parsers/simple"
)

type ParserType interface {
	Init()
	Parse(line, wholeLine, previousLine string) (changeState string, err error)
	Valid() bool
	GetParserName() string
	String() []string
}

type ParserTypes []ParserType

func (p *ParserTypes) Get(atrtibute string) (ParserType, error) {
	for _, parser := range *p {
		if parser.GetParserName() == atrtibute && parser.Valid() {
			return parser, nil
		}
	}
	return nil, fmt.Errorf("atrtibute not found")
}

//Parser reads and writes configuration on given file
type Parser struct {
	Data      map[string]ParserTypes
	Frontends map[string]ParserTypes
	Backends  map[string]ParserTypes
}

func (p *Parser) get(data map[string]ParserTypes, key string, atrtibute string) (ParserType, error) {
	for _, parser := range data[key] {
		if parser.GetParserName() == atrtibute && parser.Valid() {
			return parser, nil
		}
	}
	return nil, fmt.Errorf("atrtibute not found")
}

//GetAttr retrieves data from global and default part of configuration
func (p *Parser) GetAttr(section string, atrtibute string) (ParserType, error) {
	return p.get(p.Data, section, atrtibute)
}

//GetFrontendAttr ...
func (p *Parser) GetFrontendAttr(section string, atrtibute string) (ParserType, error) {
	return p.get(p.Frontends, section, atrtibute)
}

//GetBackendAttr ...
func (p *Parser) GetBackendAttr(section string, atrtibute string) (ParserType, error) {
	return p.get(p.Backends, section, atrtibute)
}

//String returns configuration in writable form
func (p *Parser) String() string {
	var result strings.Builder

	parsersList := []string{"#", "global", "defaults"}
	for _, parserName := range parsersList {
		parsers := p.Data[parserName]
		result.WriteString("\n")
		result.WriteString(parserName)
		result.WriteString("\n")
		for _, parser := range parsers {
			if !parser.Valid() {
				continue
			}
			for _, line := range parser.String() {
				result.WriteString(line)
				result.WriteString("\n")
			}
		}
	}
	for parserName, parsers := range p.Frontends {
		result.WriteString("\nfrontend ")
		result.WriteString(parserName)
		result.WriteString("\n")
		for _, parser := range parsers {
			if !parser.Valid() {
				continue
			}
			for _, line := range parser.String() {
				result.WriteString(line)
				result.WriteString("\n")
			}
		}
	}
	for parserName, parsers := range p.Backends {
		result.WriteString("\nbackend ")
		result.WriteString(parserName)
		result.WriteString("\n")
		for _, parser := range parsers {
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

func getFrontendParser() ParserTypes {
	return ParserTypes{
		&extra.SectionName{Name: "frontend"},
		&extra.SectionName{Name: "backend"},
		&extra.SectionName{Name: "global"},
		&extra.SectionName{Name: "defaults"},
		&extra.UnProcessed{}}
}

func getBackendParser() ParserTypes {
	return ParserTypes{
		&extra.SectionName{Name: "frontend"},
		&extra.SectionName{Name: "backend"},
		&extra.SectionName{Name: "global"},
		&extra.SectionName{Name: "defaults"},
		&extra.UnProcessed{}}
}

func (p *Parser) Save(filename string) error {
	d1 := []byte(p.String())
	err := ioutil.WriteFile(filename, d1, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (p *Parser) LoadData(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	parsersNone := ParserTypes{
		&extra.Comments{},
		&extra.SectionName{Name: "defaults"},
		&extra.SectionName{Name: "global"},
		&extra.SectionName{Name: "frontend"},
		&extra.SectionName{Name: "backend"},
		&extra.UnProcessed{}}
	for _, parser := range parsersNone {
		parser.Init()
	}
	parsersDefaults := ParserTypes{
		&parsers.MaxConn{},
		&parsers.LogLines{},

		&simple.SimpleOption{Name: "redispatch"},
		&simple.SimpleOption{Name: "dontlognull"},
		&simple.SimpleOption{Name: "http-server-close"},
		&simple.SimpleOption{Name: "http-keep-alive"},

		&simple.SimpleTimeout{Name: "http-request"},
		&simple.SimpleTimeout{Name: "connect"},
		&simple.SimpleTimeout{Name: "client"},
		&simple.SimpleTimeout{Name: "queue"},
		&simple.SimpleTimeout{Name: "server"},
		&simple.SimpleTimeout{Name: "tunnel"},
		&simple.SimpleTimeout{Name: "http-keep-alive"},

		&extra.SectionName{Name: "global"},
		&extra.SectionName{Name: "frontend"},
		&extra.SectionName{Name: "backend"},
		&extra.UnProcessed{}}
	for _, parser := range parsersDefaults {
		parser.Init()
	}
	parsersGlobal := ParserTypes{
		&parsers.Daemon{},
		&simple.SimpleNumber{Name: "nbproc"},
		&simple.SimpleString{Name: "pidfile"},
		&parsers.MaxConn{},
		&parsers.StatsSocket{},
		&parsers.StatsTimeout{},
		&simple.SimpleNumber{Name: "tune.ssl.default-dh-param"},
		&simple.SimpleStringMultiple{Name: "ssl-default-bind-options"},
		&simple.SimpleString{Name: "ssl-default-bind-ciphers"},
		&extra.SectionName{Name: "defaults"},
		&extra.SectionName{Name: "frontend"},
		&extra.SectionName{Name: "backend"},
		&extra.UnProcessed{}}
	for _, parser := range parsersGlobal {
		parser.Init()
	}
	frontends := map[string]ParserTypes{}
	backends := map[string]ParserTypes{}
	//active_frontend := ""
	//active_backend := ""
	var parserFrontend ParserTypes
	var parserBackend ParserTypes

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
			//search for segments
			for _, parser := range parsersNone {
				if newState, err := parser.Parse(line, part, previousLine); err == nil {
					//should we have an option to remove it when found?
					if newState != "" {
						//log.Printf("change state from %s to %s\n", state, newState)
						state = newState
						if state == "frontend" {
							sectionName := parser.(*extra.SectionName)
							parserFrontend = getFrontendParser()
							frontends[sectionName.SectionName] = parserFrontend
						}
						if state == "backend" {
							sectionName := parser.(*extra.SectionName)
							parserBackend = getBackendParser()
							frontends[sectionName.SectionName] = parserBackend
						}
					}
					break
				}
			}
			previousLine = ""
		case "global":
			for _, parser := range parsersGlobal {
				if newState, err := parser.Parse(line, part, previousLine); err == nil {
					//should we have an option to remove it when found?
					if newState != "" {
						//log.Printf("change state from %s to %s\n", state, newState)
						state = newState
						if state == "frontend" {
							sectionName := parser.(*extra.SectionName)
							parserFrontend = getFrontendParser()
							frontends[sectionName.SectionName] = parserFrontend
						}
						if state == "backend" {
							sectionName := parser.(*extra.SectionName)
							parserBackend = getBackendParser()
							frontends[sectionName.SectionName] = parserBackend
						}
					}
					break
				}
			}
		case "defaults":
			for _, parser := range parsersDefaults {
				if newState, err := parser.Parse(line, part, previousLine); err == nil {
					//should we have an option to remove it when found?
					if newState != "" {
						//log.Printf("change state from %s to %s\n", state, newState)
						state = newState
						if state == "frontend" {
							sectionName := parser.(*extra.SectionName)
							parserFrontend = getFrontendParser()
							frontends[sectionName.SectionName] = parserFrontend
						}
						if state == "backend" {
							sectionName := parser.(*extra.SectionName)
							parserBackend = getBackendParser()
							frontends[sectionName.SectionName] = parserBackend
						}
					}
					break
				}
			}
		case "frontend":
			for _, parser := range parserFrontend {
				if newState, err := parser.Parse(line, part, previousLine); err == nil {
					//should we have an option to remove it when found?
					if newState != "" {
						//log.Printf("change state from %s to %s\n", state, newState)
						state = newState
						if state == "frontend" {
							sectionName := parser.(*extra.SectionName)
							parserFrontend = getFrontendParser()
							frontends[sectionName.SectionName] = parserFrontend
						}
						if state == "backend" {
							sectionName := parser.(*extra.SectionName)
							parserBackend = getBackendParser()
							backends[sectionName.SectionName] = parserBackend
						}
					}
					break
				}
			}
		case "backend":
			for _, parser := range parserBackend {
				if newState, err := parser.Parse(line, part, previousLine); err == nil {
					//should we have an option to remove it when found?
					if newState != "" {
						//log.Printf("change state from %s to %s\n", state, newState)
						state = newState
						if state == "frontend" {
							sectionName := parser.(*extra.SectionName)
							parserFrontend = getFrontendParser()
							frontends[sectionName.SectionName] = parserFrontend
						}
						if state == "backend" {
							sectionName := parser.(*extra.SectionName)
							parserBackend = getBackendParser()
							backends[sectionName.SectionName] = parserBackend
						}
					}
					break
				}
			}
		}
	}
	p.Data = map[string]ParserTypes{
		"#":        parsersNone,
		"global":   parsersGlobal,
		"defaults": parsersDefaults,
	}
	p.Frontends = frontends
	p.Backends = backends
	return nil
}
