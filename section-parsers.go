package parser

import (
	"github.com/haproxytech/config-parser/parsers"
	"github.com/haproxytech/config-parser/parsers/defaults"
	"github.com/haproxytech/config-parser/parsers/extra"
	"github.com/haproxytech/config-parser/parsers/filters"
	"github.com/haproxytech/config-parser/parsers/frontend"
	"github.com/haproxytech/config-parser/parsers/global"
	"github.com/haproxytech/config-parser/parsers/httprequest"
	"github.com/haproxytech/config-parser/parsers/mailers"
	"github.com/haproxytech/config-parser/parsers/peers"
	"github.com/haproxytech/config-parser/parsers/simple"
	"github.com/haproxytech/config-parser/parsers/stats"
	"github.com/haproxytech/config-parser/parsers/userlist"
)

func createParsers(parsers []ParserType) *ParserTypes {
	p := ParserTypes{
		parsers: append(parsers, []ParserType{
			&extra.SectionName{Name: "defaults"},
			&extra.SectionName{Name: "global"},
			&extra.SectionName{Name: "frontend"},
			&extra.SectionName{Name: "backend"},
			&extra.SectionName{Name: "listen"},
			&extra.SectionName{Name: "resolvers"},
			&extra.SectionName{Name: "userlist"},
			&extra.SectionName{Name: "peers"},
			&extra.SectionName{Name: "mailers"},
			&extra.UnProcessed{},
		}...),
	}
	for _, parser := range p.parsers {
		parser.Init()
	}
	return &p
}

func getStartParser() *ParserTypes {
	return createParsers([]ParserType{
		&extra.Comments{},
	})
}

func getDefaultParser() *ParserTypes {
	return createParsers([]ParserType{
		&parsers.Mode{},
		&parsers.Balance{},
		&parsers.MaxConn{},
		&parsers.LogLines{},

		&simple.SimpleOption{Name: "http-use-htx"},
		&simple.SimpleOption{Name: "redispatch"},
		&simple.SimpleOption{Name: "dontlognull"},
		&simple.SimpleOption{Name: "http-server-close"},
		&simple.SimpleOption{Name: "http-keep-alive"},
		&simple.SimpleOption{Name: "httplog"},

		&simple.SimpleTimeout{Name: "http-request"},
		&simple.SimpleTimeout{Name: "connect"},
		&simple.SimpleTimeout{Name: "client"},
		&simple.SimpleTimeout{Name: "queue"},
		&simple.SimpleTimeout{Name: "server"},
		&simple.SimpleTimeout{Name: "tunnel"},
		&simple.SimpleTimeout{Name: "http-keep-alive"},

		&defaults.ErrorFileLines{},
		&parsers.DefaultBackend{},
	})
}

func getGlobalParser() *ParserTypes {
	return createParsers([]ParserType{
		&parsers.Daemon{},
		//&simple.SimpleFlag{Name: "master-worker"},
		&parsers.MasterWorker{},
		//&simple.SimpleNumber{Name: "nbproc"},
		&global.NbProc{},
		&global.NbThread{},
		&global.CpuMapLines{},
		&simple.SimpleString{Name: "pidfile"},
		&parsers.Mode{},
		&parsers.MaxConn{},
		&stats.SocketLines{},
		&stats.Timeout{},
		&simple.SimpleNumber{Name: "tune.ssl.default-dh-param"},
		&simple.SimpleStringMultiple{Name: "ssl-default-bind-options"},
		&simple.SimpleString{Name: "ssl-default-bind-ciphers"},
		&parsers.LogLines{},
	})
}

func getFrontendParser() *ParserTypes {
	return createParsers([]ParserType{
		&parsers.Mode{},
		&parsers.MaxConn{},
		&frontend.Binds{},
		&simple.SimpleString{Name: "log-tag"},
		&parsers.LogLines{},
		&simple.SimpleOption{Name: "http-use-htx"},
		&simple.SimpleOption{Name: "httplog"},
		&simple.SimpleOption{Name: "dontlognull"},
		&simple.SimpleOption{Name: "contstats"},
		&simple.SimpleOption{Name: "log-separate-errors"},
		&filters.Filters{},
		&httprequest.HTTPRequestAddAcls{},
		&httprequest.HTTPRequestDelAcls{},
		&httprequest.HTTPRequestAddHeaders{},
		&httprequest.HTTPRequestAllows{},
		&httprequest.HTTPRequestRedirects{},
		&httprequest.HTTPRequestDenials{},
		&httprequest.HTTPRequests{},

		&frontend.UseBackends{},
		&parsers.DefaultBackend{},
	})
}

func getBackendParser() *ParserTypes {
	return createParsers([]ParserType{
		&parsers.Mode{},
		&parsers.Balance{},
		&simple.SimpleOption{Name: "http-use-htx"},
		&simple.SimpleOption{Name: "forwardfor"},

		&filters.Filters{},
		&httprequest.HTTPRequestAddAcls{},
		&httprequest.HTTPRequestDelAcls{},
		&httprequest.HTTPRequestAddHeaders{},
		&httprequest.HTTPRequestAllows{},
		&httprequest.HTTPRequestRedirects{},
		&httprequest.HTTPRequestDenials{},
		&httprequest.HTTPRequests{},
	})
}

func getListenParser() *ParserTypes {
	return createParsers([]ParserType{})
}

func getResolverParser() *ParserTypes {
	return createParsers([]ParserType{
		&parsers.NameserverLines{},
		&simple.SimpleTimeTwoWords{Keywords: []string{"hold", "obsolete"}},
		&simple.SimpleTimeTwoWords{Keywords: []string{"hold", "valid"}},
		&simple.SimpleTimeout{Name: "retry"},
		&simple.SimpleString{Name: "accepted_payload_size"},
	})
}

func getUserlistParser() *ParserTypes {
	return createParsers([]ParserType{
		&userlist.GroupLines{},
		&userlist.UserLines{},
	})
}

func getPeersParser() *ParserTypes {
	return createParsers([]ParserType{
		&peers.Peers{},
	})
}

func getMailersParser() *ParserTypes {
	return createParsers([]ParserType{
		&simple.SimpleTimeTwoWords{Keywords: []string{"timeout", "mail"}},
		&mailers.Mailers{},
	})
}
