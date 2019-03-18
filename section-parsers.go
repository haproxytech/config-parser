package parser

import (
	"github.com/haproxytech/config-parser/parsers"
	"github.com/haproxytech/config-parser/parsers/extra"
	"github.com/haproxytech/config-parser/parsers/filters"
	"github.com/haproxytech/config-parser/parsers/http"
	"github.com/haproxytech/config-parser/parsers/simple"
	"github.com/haproxytech/config-parser/parsers/tcp"
)

func createParsers(parser []ParserType) *ParserTypes {
	p := ParserTypes{
		parsers: append(parser, []ParserType{
			&extra.Section{Name: "defaults"},
			&extra.Section{Name: "global"},
			&extra.Section{Name: "frontend"},
			&extra.Section{Name: "backend"},
			&extra.Section{Name: "listen"},
			&extra.Section{Name: "resolvers"},
			&extra.Section{Name: "userlist"},
			&extra.Section{Name: "peers"},
			&extra.Section{Name: "mailers"},
			&extra.Section{Name: "cache"},
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
		&simple.SimpleString{Name: "log-format"},
		&simple.SimpleString{Name: "log-format-sd"},

		&parsers.Log{},

		&simple.SimpleOption{Name: "http-tunnel"},
		&simple.SimpleOption{Name: "httpclose"},
		&simple.SimpleOption{Name: "forceclose"},
		&simple.SimpleOption{Name: "http-server-close"},
		&simple.SimpleOption{Name: "http-keep-alive"},
		&simple.SimpleOption{Name: "http-pretend-keepalive"},
		&simple.SimpleOption{Name: "http-use-htx"},
		&simple.SimpleOption{Name: "tcplog"},
		&simple.SimpleOption{Name: "dontlognull"},
		&simple.SimpleOption{Name: "contstats"},
		&simple.SimpleOption{Name: "log-separate-errors"},
		&simple.SimpleOption{Name: "clitcpka"},

		&parsers.OptionHTTPLog{},

		&simple.SimpleTimeout{Name: "http-request"},
		&simple.SimpleTimeout{Name: "client"},
		&simple.SimpleTimeout{Name: "http-keep-alive"},

		&filters.Filters{},
		&tcp.TCPRequests{},
		&http.HTTPRequests{},
		&http.HTTPResponses{},
		&parsers.StickTable{},

		&simple.SimpleString{Name: "monitor-uri"},

		&parsers.UseBackend{},
		&parsers.DefaultBackend{},
	})
}

func getBackendParser() *ParserTypes {
	return createParsers([]ParserType{
		&parsers.Mode{},
		&parsers.Balance{},

		&simple.SimpleOption{Name: "http-tunnel"},
		&simple.SimpleOption{Name: "httpclose"},
		&simple.SimpleOption{Name: "forceclose"},
		&simple.SimpleOption{Name: "http-server-close"},
		&simple.SimpleOption{Name: "http-keep-alive"},
		&simple.SimpleOption{Name: "http-use-htx"},
		&simple.SimpleOption{Name: "forwardfor"},
		&simple.SimpleOption{Name: "contstats"},
		&simple.SimpleOption{Name: "ssl-hello-check"},
		&simple.SimpleOption{Name: "smtpchk"},
		&simple.SimpleOption{Name: "ldap-check"},
		&simple.SimpleOption{Name: "mysql-check"},
		&simple.SimpleOption{Name: "pgsql-check"},
		&simple.SimpleOption{Name: "tcp-check"},
		&simple.SimpleOption{Name: "redis-check"},
		&simple.SimpleOption{Name: "redispatch"},

		&simple.SimpleString{Name: "log-tag"},

		&parsers.OptionHttpchk{},

		&parsers.Log{},

		&simple.SimpleTimeout{Name: "http-request"},
		&simple.SimpleTimeout{Name: "queue"},
		&simple.SimpleTimeout{Name: "client"},
		&simple.SimpleTimeout{Name: "http-keep-alive"},
		&simple.SimpleTimeout{Name: "check"},
		&simple.SimpleTimeout{Name: "tunnel"},
		&simple.SimpleTimeout{Name: "server"},
		&simple.SimpleTimeout{Name: "connect"},

		&parsers.DefaultServer{},
		&parsers.Stick{},
		&filters.Filters{},
		&tcp.TCPRequests{},
		&tcp.TCPResponses{},
		&http.HTTPRequests{},
		&http.HTTPResponses{},
		&parsers.StickTable{},
		&simple.SimpleString{Name: "cookie"},
		&parsers.UseServer{},
		&parsers.Server{},
		&simple.SimpleNumber{Name: "retries"},
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
