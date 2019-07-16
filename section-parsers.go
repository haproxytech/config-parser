/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
			&extra.Section{Name: "program"},
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
		&parsers.OptionHTTPLog{},

		&simple.SimpleWord{Name: "log-tag"},

		&simple.SimpleString{Name: "log-format"},
		&simple.SimpleString{Name: "log-format-sd"},
		&simple.SimpleString{Name: "cookie"},

		&simple.SimpleOption{Name: "tcplog"},
		&simple.SimpleOption{Name: "httpclose"},
		&simple.SimpleOption{Name: "http-use-htx"},
		&parsers.OptionRedispatch{},
		&simple.SimpleOption{Name: "dontlognull"},
		&simple.SimpleOption{Name: "log-separate-errors"},
		&simple.SimpleOption{Name: "http-server-close"},
		&simple.SimpleOption{Name: "http-keep-alive"},
		&simple.SimpleOption{Name: "http-pretend-keepalive"},
		&simple.SimpleOption{Name: "clitcpka"},
		&simple.SimpleOption{Name: "contstats"},
		&simple.SimpleOption{Name: "ssl-hello-chk"},
		&parsers.OptionSmtpchk{},
		&simple.SimpleOption{Name: "ldap-check"},
		&parsers.OptionMysqlCheck{},
		&simple.SimpleOption{Name: "pgsql-check"},
		&simple.SimpleOption{Name: "tcp-check"},
		&simple.SimpleOption{Name: "redis-check"},
		&parsers.OptionHttpchk{},

		&simple.SimpleOption{Name: "external-check"},
		&parsers.OptionForwardFor{},

		&simple.SimpleTimeout{Name: "http-request"},
		&simple.SimpleTimeout{Name: "check"},
		&simple.SimpleTimeout{Name: "connect"},
		&simple.SimpleTimeout{Name: "client"},
		&simple.SimpleTimeout{Name: "queue"},
		&simple.SimpleTimeout{Name: "server"},
		&simple.SimpleTimeout{Name: "tunnel"},
		&simple.SimpleTimeout{Name: "http-keep-alive"},

		&simple.SimpleNumber{Name: "retries"},

		&parsers.ExternalCheckPath{},
		&parsers.ExternalCheckCommand{},
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
		&parsers.ExternalCheck{},
		&parsers.NbProc{},
		&parsers.NbThread{},
		&parsers.CpuMap{},
		&parsers.Mode{},
		&parsers.MaxConn{},
		&simple.SimpleString{Name: "pidfile"},
		&parsers.Socket{},
		&parsers.StatsTimeout{},
		&simple.SimpleNumber{Name: "tune.ssl.default-dh-param"},
		&simple.SimpleString{Name: "ssl-default-bind-options"},
		&simple.SimpleWord{Name: "ssl-default-bind-ciphers"},
		&parsers.Log{},
	})
}

func getFrontendParser() *ParserTypes {
	return createParsers([]ParserType{
		&parsers.Mode{},
		&parsers.MaxConn{},
		&parsers.Bind{},
		&parsers.Acl{},
		&simple.SimpleWord{Name: "log-tag"},
		&simple.SimpleString{Name: "log-format"},
		&simple.SimpleString{Name: "log-format-sd"},

		&parsers.Log{},

		&simple.SimpleOption{Name: "httpclose"},
		&simple.SimpleOption{Name: "forceclose"},
		&simple.SimpleOption{Name: "http-server-close"},
		&simple.SimpleOption{Name: "http-keep-alive"},
		&simple.SimpleOption{Name: "http-use-htx"},
		&parsers.OptionForwardFor{},
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
		&http.Redirect{},

		&simple.SimpleWord{Name: "monitor-uri"},

		&parsers.UseBackend{},
		&parsers.DefaultBackend{},
		&parsers.StickTable{},
		&http.HTTPResponses{},
	})
}

func getBackendParser() *ParserTypes {
	return createParsers([]ParserType{
		&parsers.Mode{},
		&parsers.Balance{},
		&parsers.Acl{},

		&simple.SimpleOption{Name: "httpclose"},
		&simple.SimpleOption{Name: "forceclose"},
		&simple.SimpleOption{Name: "http-server-close"},
		&simple.SimpleOption{Name: "http-keep-alive"},
		&simple.SimpleOption{Name: "http-pretend-keepalive"},
		&simple.SimpleOption{Name: "http-use-htx"},
		&parsers.OptionForwardFor{},
		&simple.SimpleOption{Name: "ssl-hello-chk"},
		&parsers.OptionSmtpchk{},
		&simple.SimpleOption{Name: "ldap-check"},
		&parsers.OptionMysqlCheck{},
		&simple.SimpleOption{Name: "pgsql-check"},
		&simple.SimpleOption{Name: "tcp-check"},
		&simple.SimpleOption{Name: "redis-check"},
		&parsers.OptionRedispatch{},
		&simple.SimpleOption{Name: "external-check"},

		&simple.SimpleString{Name: "log-tag"},

		&parsers.OptionHttpchk{},
		&parsers.ExternalCheckPath{},
		&parsers.ExternalCheckCommand{},

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
		&http.HTTPRequests{},
		&http.Redirect{},
		&simple.SimpleString{Name: "cookie"},
		&parsers.UseServer{},
		&parsers.StickTable{},
		&parsers.Server{},
		&simple.SimpleNumber{Name: "retries"},
		&tcp.TCPResponses{},
		&http.HTTPResponses{},
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
		&simple.SimpleWord{Name: "accepted_payload_size"},
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

func getProgramParser() *ParserTypes {
	return createParsers([]ParserType{
		&simple.SimpleString{Name: "command"},
	})
}
