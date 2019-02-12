package parser

import (
	"github.com/haproxytech/config-parser/parsers"
	"github.com/haproxytech/config-parser/parsers/extra"
	"github.com/haproxytech/config-parser/parsers/filters"
	"github.com/haproxytech/config-parser/parsers/http"
	"github.com/haproxytech/config-parser/parsers/option"
	"github.com/haproxytech/config-parser/parsers/simple"
	"github.com/haproxytech/config-parser/parsers/tcp"
)

func createParsers(parser []ParserType) *ParserTypes {
	p := ParserTypes{
		parsers: append(parser, []ParserType{
			&extra.SectionName{Name: "defaults"},
			&extra.SectionName{Name: "global"},
			&extra.SectionName{Name: "frontend"},
			&extra.SectionName{Name: "backend"},
			&extra.SectionName{Name: "listen"},
			&extra.SectionName{Name: "resolvers"},
			&extra.SectionName{Name: "userlist"},
			&extra.SectionName{Name: "peers"},
			&extra.SectionName{Name: "mailers"},
			&extra.SectionName{Name: "cache"},
			&parsers.UnProcessed{},
		}...),
	}
	for _, parser := range p.parsers {
		parser.Init()
	}
	return &p
}

func getStartParser() *ParserTypes {
	return createParsers([]ParserType{
		&extra.ConfigVersion{},
		&extra.Comments{},
	})
}

func getDefaultParser() *ParserTypes {
	return createParsers([]ParserType{
		&parsers.Mode{},
		&parsers.Balance{},
		&parsers.MaxConn{},
		&parsers.Log{},

		&simple.SimpleOption{Name: "httpclose"},
		&simple.SimpleOption{Name: "http-use-htx"},
		&simple.SimpleOption{Name: "redispatch"},
		&simple.SimpleOption{Name: "dontlognull"},
		&simple.SimpleOption{Name: "http-server-close"},
		&simple.SimpleOption{Name: "http-keep-alive"},
		&simple.SimpleOption{Name: "httplog"},
		&simple.SimpleOption{Name: "clitcpka"},

		&simple.SimpleTimeout{Name: "http-request"},
		&simple.SimpleTimeout{Name: "connect"},
		&simple.SimpleTimeout{Name: "client"},
		&simple.SimpleTimeout{Name: "queue"},
		&simple.SimpleTimeout{Name: "server"},
		&simple.SimpleTimeout{Name: "tunnel"},
		&simple.SimpleTimeout{Name: "http-keep-alive"},

		&parsers.DefaultServer{},
		&parsers.ErrorFile{},
		&parsers.DefaultBackend{},
	})
}

func getGlobalParser() *ParserTypes {
	return createParsers([]ParserType{
		&parsers.Daemon{},
		//&simple.SimpleFlag{Name: "master-worker"},
		&parsers.MasterWorker{},
		&parsers.NbProc{},
		&parsers.NbThread{},
		&parsers.CpuMap{},
		&parsers.Mode{},
		&parsers.MaxConn{},
		&simple.SimpleString{Name: "pidfile"},
		&parsers.Socket{},
		&parsers.StatsTimeout{},
		&simple.SimpleNumber{Name: "tune.ssl.default-dh-param"},
		&simple.SimpleStringMultiple{Name: "ssl-default-bind-options"},
		&simple.SimpleString{Name: "ssl-default-bind-ciphers"},
		&parsers.Log{},
	})
}

func getFrontendParser() *ParserTypes {
	return createParsers([]ParserType{
		&parsers.Mode{},
		&parsers.MaxConn{},
		&parsers.Bind{},
		&simple.SimpleString{Name: "log-tag"},
		&parsers.Log{},

		&simple.SimpleOption{Name: "httpclose"},
		&simple.SimpleOption{Name: "http-use-htx"},
		&simple.SimpleOption{Name: "tcplog"},
		&simple.SimpleOption{Name: "dontlognull"},
		&simple.SimpleOption{Name: "contstats"},
		&simple.SimpleOption{Name: "log-separate-errors"},
		&simple.SimpleOption{Name: "clitcpka"},

		&option.OptionHTTPLog{},

		&simple.SimpleTimeout{Name: "http-request"},
		&simple.SimpleTimeout{Name: "client"},
		&simple.SimpleTimeout{Name: "http-keep-alive"},

		&filters.Filters{},
		&http.HTTPRequests{},
		&http.HTTPResponses{},
		&tcp.TCPRequests{},

		&simple.SimpleString{Name: "monitor-uri"},

		&parsers.UseBackend{},
		&parsers.DefaultBackend{},
	})
}

func getBackendParser() *ParserTypes {
	return createParsers([]ParserType{
		&parsers.Mode{},
		&parsers.Balance{},

		&simple.SimpleOption{Name: "httpclose"},
		&simple.SimpleOption{Name: "http-use-htx"},
		&simple.SimpleOption{Name: "forwardfor"},
		&simple.SimpleOption{Name: "http-keep-alive"},
		&simple.SimpleOption{Name: "contstats"},
		&simple.SimpleOption{Name: "httplog"},
		&simple.SimpleString{Name: "log-tag"},
		&parsers.OptionHttpchk{},

		&parsers.Log{},

		&simple.SimpleTimeout{Name: "http-request"},
		&simple.SimpleTimeout{Name: "client"},
		&simple.SimpleTimeout{Name: "http-keep-alive"},
		&simple.SimpleTimeout{Name: "check"},
		&simple.SimpleTimeout{Name: "tunnel"},
		&simple.SimpleTimeout{Name: "server"},

		&parsers.DefaultServer{},
		&parsers.Stick{},
		&filters.Filters{},
		&http.HTTPRequests{},
		&http.HTTPResponses{},
		&tcp.TCPRequests{},
		&tcp.TCPResponses{},
		&simple.SimpleString{Name: "cookie"},
		&parsers.UseServer{},
		&parsers.Server{},
	})
}

func getListenParser() *ParserTypes {
	return createParsers([]ParserType{})
}

func getResolverParser() *ParserTypes {
	return createParsers([]ParserType{
		&parsers.Nameserver{},
		&simple.SimpleTimeTwoWords{Keywords: []string{"hold", "obsolete"}},
		&simple.SimpleTimeTwoWords{Keywords: []string{"hold", "valid"}},
		&simple.SimpleTimeout{Name: "retry"},
		&simple.SimpleString{Name: "accepted_payload_size"},
	})
}

func getUserlistParser() *ParserTypes {
	return createParsers([]ParserType{
		&parsers.Group{},
		&parsers.User{},
	})
}

func getPeersParser() *ParserTypes {
	return createParsers([]ParserType{
		&parsers.Peer{},
	})
}

func getMailersParser() *ParserTypes {
	return createParsers([]ParserType{
		&simple.SimpleTimeTwoWords{Keywords: []string{"timeout", "mail"}},
		&parsers.Mailer{},
	})
}
func getCacheParser() *ParserTypes {
	return createParsers([]ParserType{
		&simple.SimpleNumber{Name: "total-max-size"},
		&simple.SimpleNumber{Name: "max-object-size"},
		&simple.SimpleNumber{Name: "max-age"},
	})
}
