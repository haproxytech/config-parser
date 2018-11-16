package parser

import (
	"github.com/haproxytech/config-parser/parsers"
	"github.com/haproxytech/config-parser/parsers/extra"
	"github.com/haproxytech/config-parser/parsers/simple"
	"github.com/haproxytech/config-parser/parsers/stats"
)

func getStartParser() ParserTypes {
	p := ParserTypes{
		parsers: []ParserType{
			&extra.Comments{},
			&extra.SectionName{Name: "defaults"},
			&extra.SectionName{Name: "global"},
			&extra.SectionName{Name: "frontend"},
			&extra.SectionName{Name: "backend"},
			&extra.SectionName{Name: "listen"},
			&extra.SectionName{Name: "resolvers"},
			&extra.UnProcessed{},
		},
	}
	for _, parser := range p.parsers {
		parser.Init()
	}
	return p
}

func getDefaultParser() ParserTypes {
	p := ParserTypes{
		parsers: []ParserType{
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
			&extra.SectionName{Name: "listen"},
			&extra.SectionName{Name: "resolvers"},
			&extra.UnProcessed{},
		},
	}
	for _, parser := range p.parsers {
		parser.Init()
	}
	return p
}

func getGlobalParser() ParserTypes {
	p := ParserTypes{
		parsers: []ParserType{
			&parsers.Daemon{},
			&simple.SimpleNumber{Name: "nbproc"},
			&simple.SimpleString{Name: "pidfile"},
			&parsers.MaxConn{},
			&stats.SocketLines{},
			&stats.Timeout{},
			&simple.SimpleNumber{Name: "tune.ssl.default-dh-param"},
			&simple.SimpleStringMultiple{Name: "ssl-default-bind-options"},
			&simple.SimpleString{Name: "ssl-default-bind-ciphers"},
			&extra.SectionName{Name: "defaults"},
			&extra.SectionName{Name: "frontend"},
			&extra.SectionName{Name: "backend"},
			&extra.SectionName{Name: "listen"},
			&extra.SectionName{Name: "resolvers"},
			&extra.UnProcessed{},
		},
	}
	for _, parser := range p.parsers {
		parser.Init()
	}
	return p
}

func getFrontendParser() ParserTypes {
	p := ParserTypes{
		parsers: []ParserType{
			&extra.SectionName{Name: "frontend"},
			&extra.SectionName{Name: "backend"},
			&extra.SectionName{Name: "global"},
			&extra.SectionName{Name: "defaults"},
			&extra.SectionName{Name: "listen"},
			&extra.SectionName{Name: "resolvers"},
			&extra.UnProcessed{},
		},
	}
	for _, parser := range p.parsers {
		parser.Init()
	}
	return p
}

func getBackendParser() ParserTypes {
	p := ParserTypes{
		parsers: []ParserType{
			&extra.SectionName{Name: "frontend"},
			&extra.SectionName{Name: "backend"},
			&extra.SectionName{Name: "global"},
			&extra.SectionName{Name: "defaults"},
			&extra.SectionName{Name: "listen"},
			&extra.SectionName{Name: "resolvers"},
			&extra.UnProcessed{},
		},
	}
	for _, parser := range p.parsers {
		parser.Init()
	}
	return p
}

func getListenParser() ParserTypes {
	p := ParserTypes{
		parsers: []ParserType{
			&extra.SectionName{Name: "frontend"},
			&extra.SectionName{Name: "backend"},
			&extra.SectionName{Name: "global"},
			&extra.SectionName{Name: "defaults"},
			&extra.SectionName{Name: "listen"},
			&extra.SectionName{Name: "resolvers"},
			&extra.UnProcessed{},
		},
	}
	for _, parser := range p.parsers {
		parser.Init()
	}
	return p
}

func getResolverParser() ParserTypes {
	p := ParserTypes{
		parsers: []ParserType{
			&parsers.NameserverLines{},
			&simple.SimpleTimeTwoWords{Name: "hold obsolete"},
			&simple.SimpleTimeTwoWords{Name: "hold valid"},
			&simple.SimpleTimeout{Name: "retry"},
			&simple.SimpleString{Name: "accepted_payload_size"},
			&extra.SectionName{Name: "frontend"},
			&extra.SectionName{Name: "backend"},
			&extra.SectionName{Name: "global"},
			&extra.SectionName{Name: "defaults"},
			&extra.SectionName{Name: "listen"},
			&extra.SectionName{Name: "resolvers"},
			&extra.UnProcessed{},
		},
	}
	for _, parser := range p.parsers {
		parser.Init()
	}
	return p
}
