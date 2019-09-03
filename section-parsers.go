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

func createParsers(parser []ParserInterface) *Parsers {
	p := Parsers{
		parsers: append(parser, []ParserInterface{
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

func getStartParser() *Parsers {
	return createParsers([]ParserInterface{
		&extra.ConfigVersion{},
		&extra.Comments{},
	})
}

func getDefaultParser() *Parsers {
	return createParsers([]ParserInterface{
		&parsers.Mode{},
		&parsers.HashType{},
		&parsers.Balance{},
		&parsers.MaxConn{},
		&parsers.Log{},
		&parsers.OptionHTTPLog{},

		&simple.Word{Name: "log-tag"},

		&simple.String{Name: "log-format"},
		&simple.String{Name: "log-format-sd"},
		&simple.String{Name: "cookie"},

		&simple.Option{Name: "tcplog"},
		&simple.Option{Name: "httpclose"},
		&simple.Option{Name: "http-use-htx"},
		&parsers.OptionRedispatch{},
		&simple.Option{Name: "dontlognull"},
		&simple.Option{Name: "log-separate-errors"},
		&simple.Option{Name: "http-buffer-request"},
		&simple.Option{Name: "http-server-close"},
		&simple.Option{Name: "http-keep-alive"},
		&simple.Option{Name: "http-pretend-keepalive"},
		&simple.Option{Name: "clitcpka"},
		&simple.Option{Name: "contstats"},
		&simple.Option{Name: "ssl-hello-chk"},
		&parsers.OptionSmtpchk{},
		&simple.Option{Name: "ldap-check"},
		&parsers.OptionMysqlCheck{},
		&simple.Option{Name: "pgsql-check"},
		&simple.Option{Name: "tcp-check"},
		&simple.Option{Name: "redis-check"},
		&parsers.OptionHttpchk{},

		&simple.Option{Name: "external-check"},
		&parsers.OptionForwardFor{},

		&simple.Timeout{Name: "http-request"},
		&simple.Timeout{Name: "check"},
		&simple.Timeout{Name: "connect"},
		&simple.Timeout{Name: "client"},
		&simple.Timeout{Name: "queue"},
		&simple.Timeout{Name: "server"},
		&simple.Timeout{Name: "tunnel"},
		&simple.Timeout{Name: "http-keep-alive"},

		&simple.Number{Name: "retries"},

		&parsers.ExternalCheckPath{},
		&parsers.ExternalCheckCommand{},
		&parsers.DefaultServer{},
		&parsers.ErrorFile{},
		&parsers.DefaultBackend{},
	})
}

func getGlobalParser() *Parsers {
	return createParsers([]ParserInterface{
		&parsers.Daemon{},
		//&simple.SimpleFlag{Name: "master-worker"},
		&parsers.MasterWorker{},
		&parsers.ExternalCheck{},
		&parsers.NbProc{},
		&parsers.NbThread{},
		&parsers.CPUMap{},
		&parsers.Mode{},
		&parsers.MaxConn{},
		&simple.String{Name: "pidfile"},
		&parsers.Socket{},
		&parsers.StatsTimeout{},
		&simple.Number{Name: "tune.ssl.default-dh-param"},
		&simple.String{Name: "ssl-default-bind-options"},
		&simple.Word{Name: "ssl-default-bind-ciphers"},
		&parsers.Log{},
	})
}

func getFrontendParser() *Parsers {
	return createParsers([]ParserInterface{
		&parsers.Mode{},
		&parsers.MaxConn{},
		&parsers.Bind{},
		&parsers.ACL{},
		&simple.Word{Name: "log-tag"},
		&simple.String{Name: "log-format"},
		&simple.String{Name: "log-format-sd"},

		&parsers.Log{},

		&simple.Option{Name: "httpclose"},
		&simple.Option{Name: "forceclose"},
		&simple.Option{Name: "http-buffer-request"},
		&simple.Option{Name: "http-server-close"},
		&simple.Option{Name: "http-keep-alive"},
		&simple.Option{Name: "http-use-htx"},
		&parsers.OptionForwardFor{},
		&simple.Option{Name: "tcplog"},
		&simple.Option{Name: "dontlognull"},
		&simple.Option{Name: "contstats"},
		&simple.Option{Name: "log-separate-errors"},
		&simple.Option{Name: "clitcpka"},

		&parsers.OptionHTTPLog{},

		&simple.Timeout{Name: "http-request"},
		&simple.Timeout{Name: "client"},
		&simple.Timeout{Name: "http-keep-alive"},

		&filters.Filters{},
		&tcp.Requests{},
		&http.Requests{},
		&http.Redirect{},

		&simple.Word{Name: "monitor-uri"},

		&parsers.UseBackend{},
		&parsers.DefaultBackend{},
		&parsers.StickTable{},
		&http.Responses{},
	})
}

func getBackendParser() *Parsers {
	return createParsers([]ParserInterface{
		&parsers.Mode{},
		&parsers.HashType{},
		&parsers.Balance{},
		&parsers.ACL{},

		&simple.Option{Name: "httpclose"},
		&simple.Option{Name: "forceclose"},
		&simple.Option{Name: "http-buffer-request"},
		&simple.Option{Name: "http-server-close"},
		&simple.Option{Name: "http-keep-alive"},
		&simple.Option{Name: "http-pretend-keepalive"},
		&simple.Option{Name: "http-use-htx"},
		&parsers.OptionForwardFor{},
		&simple.Option{Name: "ssl-hello-chk"},
		&parsers.OptionSmtpchk{},
		&simple.Option{Name: "ldap-check"},
		&parsers.OptionMysqlCheck{},
		&simple.Option{Name: "pgsql-check"},
		&simple.Option{Name: "tcp-check"},
		&simple.Option{Name: "redis-check"},
		&parsers.OptionRedispatch{},
		&simple.Option{Name: "external-check"},

		&simple.String{Name: "log-tag"},

		&parsers.OptionHttpchk{},
		&parsers.ExternalCheckPath{},
		&parsers.ExternalCheckCommand{},

		&parsers.Log{},

		&simple.Timeout{Name: "http-request"},
		&simple.Timeout{Name: "queue"},
		&simple.Timeout{Name: "client"},
		&simple.Timeout{Name: "http-keep-alive"},
		&simple.Timeout{Name: "check"},
		&simple.Timeout{Name: "tunnel"},
		&simple.Timeout{Name: "server"},
		&simple.Timeout{Name: "connect"},

		&parsers.DefaultServer{},
		&parsers.Stick{},
		&filters.Filters{},
		&tcp.Requests{},
		&http.Requests{},
		&http.Redirect{},
		&simple.String{Name: "cookie"},
		&parsers.UseServer{},
		&parsers.StickTable{},
		&parsers.Server{},
		&simple.Number{Name: "retries"},
		&tcp.Responses{},
		&http.Responses{},
	})
}

func getListenParser() *Parsers {
	return createParsers([]ParserInterface{})
}

func getResolverParser() *Parsers {
	return createParsers([]ParserInterface{
		&parsers.Nameserver{},
		&simple.TimeTwoWords{Keywords: []string{"hold", "obsolete"}},
		&simple.TimeTwoWords{Keywords: []string{"hold", "valid"}},
		&simple.Timeout{Name: "retry"},
		&simple.Word{Name: "accepted_payload_size"},
	})
}

func getUserlistParser() *Parsers {
	return createParsers([]ParserInterface{
		&parsers.Group{},
		&parsers.User{},
	})
}

func getPeersParser() *Parsers {
	return createParsers([]ParserInterface{
		&parsers.Peer{},
	})
}

func getMailersParser() *Parsers {
	return createParsers([]ParserInterface{
		&simple.TimeTwoWords{Keywords: []string{"timeout", "mail"}},
		&parsers.Mailer{},
	})
}

func getCacheParser() *Parsers {
	return createParsers([]ParserInterface{
		&simple.Number{Name: "total-max-size"},
		&simple.Number{Name: "max-object-size"},
		&simple.Number{Name: "max-age"},
	})
}

func getProgramParser() *Parsers {
	return createParsers([]ParserInterface{
		&simple.String{Name: "command"},
	})
}
