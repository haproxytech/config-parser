package main

import (
	"io/ioutil"
	"log"
	"strings"

	"config-parser/parsers"
	"config-parser/parsers/extra"
	"config-parser/parsers/simple"
)

type ParserType interface {
	Init()
	Parse(line, wholeLine, previousLine string) (changeState string, err error)
	Valid() bool
	String() []string
}

type Parser struct {
	Data map[string][]ParserType
}

func (p *Parser) String() string {
	var result strings.Builder

	parsersList := []string{"#", "defaults", "global"}
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
	return result.String()
}

func (p *Parser) LoadData(filename string) (map[string][]ParserType, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	parsersNone := []ParserType{
		&extra.Comments{},
		&extra.SectionName{Name: "defaults"},
		&extra.SectionName{Name: "global"},
		&extra.UnProcessed{}}
	for _, parser := range parsersNone {
		parser.Init()
	}
	parsersDefaults := []ParserType{
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
		&extra.UnProcessed{}}
	for _, parser := range parsersDefaults {
		parser.Init()
	}
	parsersGlobal := []ParserType{
		&parsers.Daemon{},
		&parsers.MaxConn{},
		&parsers.StatsSocket{},
		&simple.SimpleNumber{Name: "tune.ssl.default-dh-param"},
		&simple.SimpleStringMultiple{Name: "ssl-default-bind-options"},
		&simple.SimpleString{Name: "ssl-default-bind-ciphers"},
		&extra.SectionName{Name: "defaults"},
		&extra.UnProcessed{}}
	for _, parser := range parsersGlobal {
		parser.Init()
	}

	lines := strings.Split(string(dat), "\n")
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
						log.Printf("change state from %s to %s\n", state, newState)
						state = newState
					}
					break
				}
			}
			previousLine = ""
		case "global":
			if part[0] != ' ' {
				state = line
				previousLine = ""
				log.Println("change state to", state)
				continue
			}
			for _, parser := range parsersGlobal {
				if newState, err := parser.Parse(line, part, previousLine); err == nil {
					//should we have an option to remove it when found?
					if newState != "" {
						log.Printf("change state from %s to %s\n", state, newState)
						state = newState
					}
					break
				}
			}
		case "defaults":
			if part[0] != ' ' {
				state = line
				previousLine = ""
				log.Println("change state to", state)
				continue
			}
			for _, parser := range parsersDefaults {
				if newState, err := parser.Parse(line, part, previousLine); err == nil {
					//should we have an option to remove it when found?
					if newState != "" {
						log.Printf("change state from %s to %s\n", state, newState)
						state = newState
					}
					break
				}
			}
		}
	}
	p.Data = map[string][]ParserType{
		"#":        parsersNone,
		"global":   parsersGlobal,
		"defaults": parsersDefaults,
	}
	return p.Data, nil
}
