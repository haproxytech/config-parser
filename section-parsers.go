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
	"strings"

	"github.com/haproxytech/config-parser/v2/parsers"
	"github.com/haproxytech/config-parser/v2/parsers/extra"
	"github.com/haproxytech/config-parser/v2/parsers/filters"
	"github.com/haproxytech/config-parser/v2/parsers/http"
	"github.com/haproxytech/config-parser/v2/parsers/simple"
	"github.com/haproxytech/config-parser/v2/parsers/stats"
	"github.com/haproxytech/config-parser/v2/parsers/tcp"
)

func getParsersSequenceForSection(sectionName string, parsers map[string]ParserInterface) []ParserInterface {
	pseq := []ParserInterface{}
	switch {
	case sectionName == "":
		for _, pname := range parserSequenceStart {
			parser := parsers[string(pname)]
			pseq = append(pseq, parser)
		}
	case sectionName == string(Global):
		for _, pname := range parserSequenceGlobal {
			parser := parsers[string(pname)]
			pseq = append(pseq, parser)
		}
	case sectionName == string(Defaults):
		for _, pname := range parserSequenceDefault {
			parser := parsers[string(pname)]
			pseq = append(pseq, parser)
		}
	case strings.HasPrefix(sectionName, string(Frontends)):
		for _, pname := range parserSequenceFrontend {
			parser := parsers[string(pname)]
			pseq = append(pseq, parser)
		}
	case strings.HasPrefix(sectionName, string(Backends)):
		for _, pname := range parserSequenceBackend {
			parser := parsers[string(pname)]
			pseq = append(pseq, parser)
		}
	case strings.HasPrefix(sectionName, string(Mailers)):
		for _, pname := range parserSequenceMailers {
			parser := parsers[string(pname)]
			pseq = append(pseq, parser)
		}
	case strings.HasPrefix(sectionName, string(UserList)):
		for _, pname := range parserSequenceUserlist {
			parser := parsers[string(pname)]
			pseq = append(pseq, parser)
		}
	case strings.HasPrefix(sectionName, string(Cache)):
		for _, pname := range parserSequenceCache {
			parser := parsers[string(pname)]
			pseq = append(pseq, parser)
		}
	case strings.HasPrefix(sectionName, string(Peers)):
		for _, pname := range parserSequencePeers {
			parser := parsers[string(pname)]
			pseq = append(pseq, parser)
		}
	case strings.HasPrefix(sectionName, string(Resolvers)):
		for _, pname := range parserSequenceResolver {
			parser := parsers[string(pname)]
			pseq = append(pseq, parser)
		}
	case strings.HasPrefix(sectionName, string(Listen)):
		for _, pname := range parserSequenceSections {
			parser := parsers[string(pname)]
			pseq = append(pseq, parser)
		}
	}
	return pseq
}

var parserSequenceSections = []Section{"defaults", "global", "frontend", "backend", "listen", "resolvers", "userlist", "peers", "mailers", "cache", "program", "http-errors", "ring", "unprocessed"}

func createParsers(parser map[string]ParserInterface) *Parsers {
	parser["defaults"] = &extra.Section{Name: "defaults"}
	parser["global"] = &extra.Section{Name: "global"}
	parser["frontend"] = &extra.Section{Name: "frontend"}
	parser["backend"] = &extra.Section{Name: "backend"}
	parser["listen"] = &extra.Section{Name: "listen"}
	parser["resolvers"] = &extra.Section{Name: "resolvers"}
	parser["userlist"] = &extra.Section{Name: "userlist"}
	parser["peers"] = &extra.Section{Name: "peers"}
	parser["mailers"] = &extra.Section{Name: "mailers"}
	parser["cache"] = &extra.Section{Name: "cache"}
	parser["program"] = &extra.Section{Name: "program"}
	parser["http-errors"] = &extra.Section{Name: "http-errors"}
	parser["ring"] = &extra.Section{Name: "ring"}
	parser["unprocessed"] = &extra.UnProcessed{}

	for _, parser := range parser {
		parser.Init()
	}

	return &Parsers{Parsers: parser}
}

var parserSequenceStart = append([]Section{"configversion", "comments"}, parserSequenceSections...)

func getStartParser() *Parsers {
	return createParsers(map[string]ParserInterface{
		"configversion": &extra.ConfigVersion{},
		"comments":      &extra.Comments{},
	})
}

var parserSequenceDefault = append([]Section{"mode", "hashtype", "balance", "maxconn",
	"log", "option httplog", "stats", "option log-tag", "log-format",
	"log-format-sd", "cookie", "bindprocess", "option tcplog",
	"option httpclose", "option http-use-htx", "option redispatch", "option dontlognull",
	"option log-separate-errors", "option http-buffer-request", "option http-server-close",
	"option http-keep-alive", "option http-pretend-keepalive", "option clitcpka", "option contstats",
	"option ssl-hello-chk", "option smtpchk", "option ldap-check", "option mysqlcheck",
	"option abortonclose", "option pgsqlcheck", "option tcp-check", "option redis-check",
	"option splice-auto", "option splice-request", "option splice-response",
	"option logasap", "option log-health-checks", "option allbackups", "option external-check",
	"option forwardfor", "option httpchk", "externalcheckpath", "externalcheckcommand", "httpreuse", "timeout http-request",
	"timeout check", "timeout connect", "timeout client",
	"timeout client-fin", "timeout queue", "timeout server",
	"timeout server-fin", "timeout tunnel", "timeout http-keep-alive",
	"retries", "defaultserver",
	"errorfile", "defaultbackend", "uniqueidformat", "uniqueidheader",
	"configsnippet"}, parserSequenceSections...)

func getDefaultParser() *Parsers {
	return createParsers(map[string]ParserInterface{
		"mode":     &parsers.Mode{},
		"monitor-uri": &parsers.MonitorURI{},
		"hashtype": &parsers.HashType{},
		"balance":  &parsers.Balance{},

		"maxconn":       &parsers.MaxConn{},
		"log":           &parsers.Log{},
		"optionhttplog": &parsers.OptionHTTPLog{},
		"stats":         &stats.Stats{Mode: "defaults"},

		"log-tag": &simple.Word{Name: "log-tag"},

		"log-format":    &simple.String{Name: "log-format"},
		"log-format-sd": &simple.String{Name: "log-format-sd"},
		"cookie":        &parsers.Cookie{},
		"bindprocess":   &parsers.BindProcess{},

		"option tcplog":                 &simple.Option{Name: "tcplog"},
		"option httpclose":              &simple.Option{Name: "httpclose"},
		"option http-use-htx":           &simple.Option{Name: "http-use-htx"},
		"option redispatch":             &parsers.OptionRedispatch{},
		"option dontlognull":            &simple.Option{Name: "dontlognull"},
		"option log-separate-errors":    &simple.Option{Name: "log-separate-errors"},
		"option http-buffer-request":    &simple.Option{Name: "http-buffer-request"},
		"option http-server-close":      &simple.Option{Name: "http-server-close"},
		"option http-keep-alive":        &simple.Option{Name: "http-keep-alive"},
		"option http-pretend-keepalive": &simple.Option{Name: "http-pretend-keepalive"},
		"option clitcpka":               &simple.Option{Name: "clitcpka"},
		"option contstats":              &simple.Option{Name: "contstats"},
		"option ssl-hello-chk":          &simple.Option{Name: "ssl-hello-chk"},
		"option smtpchk":                &parsers.OptionSmtpchk{},
		"option ldap-check":             &simple.Option{Name: "ldap-check"},
		"option mysqlcheck":             &parsers.OptionMysqlCheck{},
		"option abortonclose":           &simple.Option{Name: "abortonclose"},
		"option pgsqlcheck":             &parsers.OptionPgsqlCheck{},
		"option tcp-check":              &simple.Option{Name: "tcp-check"},
		"option redis-check":            &simple.Option{Name: "redis-check"},
		"option splice-auto":            &simple.Option{Name: "splice-auto"},
		"option splice-request":         &simple.Option{Name: "splice-request"},
		"option splice-response":        &simple.Option{Name: "splice-response"},
		"option logasap":                &simple.Option{Name: "logasap"},
		"option log-health-checks":      &simple.Option{Name: "log-health-checks"},
		"option allbackups":             &simple.Option{Name: "allbackups"},
		"option external-check":         &simple.Option{Name: "external-check"},
		"option forwardfor":             &parsers.OptionForwardFor{},

		"option httpchk":       &parsers.OptionHttpchk{},
		"externalcheckpath":    &parsers.ExternalCheckPath{},
		"externalcheckcommand": &parsers.ExternalCheckCommand{},

		"httpreuse":               &parsers.HTTPReuse{},
		"timeout-http-request":    &simple.Timeout{Name: "http-request"},
		"timeout-check":           &simple.Timeout{Name: "check"},
		"timeout-connect":         &simple.Timeout{Name: "connect"},
		"timeout-client":          &simple.Timeout{Name: "client"},
		"timeout-client-fin":      &simple.Timeout{Name: "client-fin"},
		"timeout-queue":           &simple.Timeout{Name: "queue"},
		"timeout-server":          &simple.Timeout{Name: "server"},
		"timeout-server-fin":      &simple.Timeout{Name: "server-fin"},
		"timeout-tunnel":          &simple.Timeout{Name: "tunnel"},
		"timeout-http-keep-alive": &simple.Timeout{Name: "http-keep-alive"},

		"retries": &simple.Number{Name: "retries"},

		"defaultserver":  &parsers.DefaultServer{},
		"errorfile":      &parsers.ErrorFile{},
		"defaultbackend": &parsers.DefaultBackend{},
		"uniqueidformat": &parsers.UniqueIDFormat{},
		"uniqueidheader": &parsers.UniqueIDHeader{},
		"configsnippet":  &parsers.ConfigSnippet{},
	})
}

var parserSequenceGlobal = append([]Section{"daemon", "localpeer", "chroot", "user",
	"group", "master-worker", "externalcheck", "nosplice", "nbproc", "nbthread",
	"cpumap", "mode", "maxconn", "pidfile", "socket", "statstimeout",
	"tune.bufsize", "tune.maxrewrite", "tune.ssl.default-dh-param",
	"ssl-default-bind-options", "ssl-default-bind-ciphers",
	"ssl-default-bind-ciphersuites", "ssl-default-server-options",
	"ssl-default-server-ciphers", "ssl-default-server-ciphersuites",
	"ssl-server-verify", "log", "logsendhostname", "luaload", "configsnippet"}, parserSequenceSections...)

func getGlobalParser() *Parsers {
	return createParsers(map[string]ParserInterface{
		"daemon":    &parsers.Daemon{},
		"localpeer": &simple.String{Name: "localpeer"},
		"chroot":    &simple.Word{Name: "chroot"},
		"user":      &simple.Word{Name: "user"},
		"group":     &simple.Word{Name: "group"},
		//&simple.SimpleFlag{Name: "master-worker"},
		"master-worker":                   &parsers.MasterWorker{},
		"externalcheck":                   &parsers.ExternalCheck{},
		"nosplice":                        &parsers.NoSplice{},
		"nbproc":                          &parsers.NbProc{},
		"nbthread":                        &parsers.NbThread{},
		"cpumap":                          &parsers.CPUMap{},
		"mode":                            &parsers.Mode{},
		"maxconn":                         &parsers.MaxConn{},
		"pidfile":                         &simple.String{Name: "pidfile"},
		"socket":                          &parsers.Socket{},
		"statstimeout":                    &parsers.StatsTimeout{},
		"tune.bufsize":                    &simple.Number{Name: "tune.bufsize"},
		"tune.maxrewrite":                 &simple.Number{Name: "tune.maxrewrite"},
		"tune.ssl.default-dh-param":       &simple.Number{Name: "tune.ssl.default-dh-param"},
		"ssl-default-bind-options":        &simple.String{Name: "ssl-default-bind-options"},
		"ssl-default-bind-ciphers":        &simple.Word{Name: "ssl-default-bind-ciphers"},
		"ssl-default-bind-ciphersuites":   &simple.Word{Name: "ssl-default-bind-ciphersuites"},
		"ssl-default-server-options":      &simple.String{Name: "ssl-default-server-options"},
		"ssl-default-server-ciphers":      &simple.Word{Name: "ssl-default-server-ciphers"},
		"ssl-default-server-ciphersuites": &simple.Word{Name: "ssl-default-server-ciphersuites"},
		"ssl-server-verify":               &simple.Word{Name: "ssl-server-verify"},
		"log":                             &parsers.Log{},
		"logsendhostname":                 &parsers.LogSendHostName{},
		"luaload":                         &parsers.LuaLoad{},
		"configsnippet":                   &parsers.ConfigSnippet{},
	})
}

var parserSequenceFrontend = append([]Section{"mode", "maxconn", "bind", "acl",
	"bindprocess", "log-tag", "log-format", "log-format-sd", "log", "httpclose",
	"forceclose", "http-buffer-request", "http-server-close", "http-keep-alive",
	"http-use-htx", "optionforwardfor", "tcplog", "dontlognull", "contstats",
	"log-separate-errors", "clitcpka", "splice-auto", "splice-request",
	"splice-response", "logasap", "optionhttplog", "timeout-http-request",
	"timeout-client", "timeout-client-fin", "timeout-timeout-http-keep-alive",
	"filters", "requests", "stats-frontend", "requests-frontend", "redirect",
	"monitor-uri", "configsnippet", "usebackend", "defaultbackend", "sticktable",
	"responses-frontend", "uniqueidformat", "uniqueidheader"}, parserSequenceSections...)

func getFrontendParser() *Parsers {
	return createParsers(map[string]ParserInterface{
		"mode":          &parsers.Mode{},
		"monitor-uri": &parsers.MonitorURI{},
		"monitor-fail": &parsers.MonitorFail{},
		"maxconn":       &parsers.MaxConn{},
		"bind":          &parsers.Bind{},
		"acl":           &parsers.ACL{},
		"bindprocess":   &parsers.BindProcess{},
		"log-tag":       &simple.Word{Name: "log-tag"},
		"log-format":    &simple.String{Name: "log-format"},
		"log-format-sd": &simple.String{Name: "log-format-sd"},

		"log": &parsers.Log{},

		"httpclose":           &simple.Option{Name: "httpclose"},
		"forceclose":          &simple.Option{Name: "forceclose"},
		"http-buffer-request": &simple.Option{Name: "http-buffer-request"},
		"http-server-close":   &simple.Option{Name: "http-server-close"},
		"http-keep-alive":     &simple.Option{Name: "http-keep-alive"},
		"http-use-htx":        &simple.Option{Name: "http-use-htx"},
		"optionforwardfor":    &parsers.OptionForwardFor{},
		"tcplog":              &simple.Option{Name: "tcplog"},
		"dontlognull":         &simple.Option{Name: "dontlognull"},
		"contstats":           &simple.Option{Name: "contstats"},
		"log-separate-errors": &simple.Option{Name: "log-separate-errors"},
		"clitcpka":            &simple.Option{Name: "clitcpka"},
		"splice-auto":         &simple.Option{Name: "splice-auto"},
		"splice-request":      &simple.Option{Name: "splice-request"},
		"splice-response":     &simple.Option{Name: "splice-response"},
		"logasap":             &simple.Option{Name: "logasap"},
		"optionhttplog":       &parsers.OptionHTTPLog{},

		"timeout-http-request":            &simple.Timeout{Name: "http-request"},
		"timeout-client":                  &simple.Timeout{Name: "client"},
		"timeout-client-fin":              &simple.Timeout{Name: "client-fin"},
		"timeout-timeout-http-keep-alive": &simple.Timeout{Name: "http-keep-alive"},

		"filters":           &filters.Filters{},
		"requests":          &tcp.Requests{},
		"stats-frontend":    &stats.Stats{Mode: "frontend"},
		"requests-frontend": &http.Requests{Mode: "frontend"},
		"redirect":          &http.Redirect{},

		"monitor-uri": &simple.Word{Name: "monitor-uri"},

		"configsnippet":      &parsers.ConfigSnippet{},
		"usebackend":         &parsers.UseBackend{},
		"defaultbackend":     &parsers.DefaultBackend{},
		"sticktable":         &parsers.StickTable{},
		"responses-frontend": &http.Responses{Mode: "frontend"},
		"uniqueidformat":     &parsers.UniqueIDFormat{},
		"uniqueidheader":     &parsers.UniqueIDHeader{},
	})
}

var parserSequenceBackend = append([]Section{"mode", "hashtype", "balance", "acl",
	"bindprocess", "option httpclose", "option forceclose", "option http-buffer-request",
	"option http-server-close", "option http-keep-alive", "option http-pretend-keepalive",
	"option http-use-htx", "option forwardfor", "option ssl-hello-chk", "option smtpchk",
	"option ldap-check", "option mysqlcheck", "option abortonclose", "option pgsqlcheck",
	"option tcp-check", "option redis-check", "option redispatch", "option external-check", "option splice-auto",
	"option splice-request", "option splice-response", "option log-health-checks", "option log-tag",
	"option allbackups", "option httpchk", "externalcheckpath", "externalcheckcommand",
	"log", "timeout http-request", "timeout queue", "timeout http-keep-alive",
	"timeout check", "timeout tunnel", "timeout server", "timeout server-fin",
	"timeout connect", "defaultserver", "stick", "filters", "tcp-request", "backend",
	"httpreuse", "http-request", "http-redirect", "cookie", "useserver", "sticktable",
	"configsnippet", "server", "retries", "tcp-response", "http-response"}, parserSequenceSections...)

func getBackendParser() *Parsers {
	return createParsers(map[string]ParserInterface{
		"mode":        &parsers.Mode{},
		"hashtype":    &parsers.HashType{},
		"balance":     &parsers.Balance{},
		"acl":         &parsers.ACL{},
		"bindprocess": &parsers.BindProcess{},

		"option httpclose":              &simple.Option{Name: "httpclose"},
		"option forceclose":             &simple.Option{Name: "forceclose"},
		"option http-buffer-request":    &simple.Option{Name: "http-buffer-request"},
		"option http-server-close":      &simple.Option{Name: "http-server-close"},
		"option http-keep-alive":        &simple.Option{Name: "http-keep-alive"},
		"option http-pretend-keepalive": &simple.Option{Name: "http-pretend-keepalive"},
		"option http-use-htx":           &simple.Option{Name: "http-use-htx"},
		"option forwardfor":             &parsers.OptionForwardFor{},
		"option ssl-hello-chk":          &simple.Option{Name: "ssl-hello-chk"},
		"option smtpchk":                &parsers.OptionSmtpchk{},
		"option ldap-check":             &simple.Option{Name: "ldap-check"},
		"option mysqlcheck":             &parsers.OptionMysqlCheck{},
		"option abortonclose":           &simple.Option{Name: "abortonclose"},
		"option pgsqlcheck":             &parsers.OptionPgsqlCheck{},
		"option tcp-check":              &simple.Option{Name: "tcp-check"},
		"option redis-check":            &simple.Option{Name: "redis-check"},
		"option redispatch":             &parsers.OptionRedispatch{},
		"option external-check":         &simple.Option{Name: "external-check"},
		"option splice-auto":            &simple.Option{Name: "splice-auto"},
		"option splice-request":         &simple.Option{Name: "splice-request"},
		"option splice-response":        &simple.Option{Name: "splice-response"},
		"option log-health-checks":      &simple.Option{Name: "log-health-checks"},
		"option log-tag":                &simple.String{Name: "log-tag"},
		"option allbackups":             &simple.Option{Name: "allbackups"},

		"option httpchk":       &parsers.OptionHttpchk{},
		"externalcheckpath":    &parsers.ExternalCheckPath{},
		"externalcheckcommand": &parsers.ExternalCheckCommand{},

		"log": &parsers.Log{},

		"timeout-http-request":    &simple.Timeout{Name: "http-request"},
		"timeout-queue":           &simple.Timeout{Name: "queue"},
		"timeout-http-keep-alive": &simple.Timeout{Name: "http-keep-alive"},
		"timeout-check":           &simple.Timeout{Name: "check"},
		"timeout-tunnel":          &simple.Timeout{Name: "tunnel"},
		"timeout-server":          &simple.Timeout{Name: "server"},
		"timeout-server-fin":      &simple.Timeout{Name: "server-fin"},
		"timeout-connect":         &simple.Timeout{Name: "connect"},

		"defaultserver":    &parsers.DefaultServer{},
		"stick":            &parsers.Stick{},
		"filters":          &filters.Filters{},
		"requests":         &tcp.Requests{},
		"backend":          &stats.Stats{Mode: "backend"},
		"httpreuse":        &parsers.HTTPReuse{},
		"request-backend":  &http.Requests{Mode: "backend"},
		"redirect":         &http.Redirect{},
		"cookie":           &parsers.Cookie{},
		"useserver":        &parsers.UseServer{},
		"sticktable":       &parsers.StickTable{},
		"configsnippet":    &parsers.ConfigSnippet{},
		"server":           &parsers.Server{},
		"retries":          &simple.Number{Name: "retries"},
		"responses":        &tcp.Responses{},
		"response-backend": &http.Responses{Mode: "backend"},
	})
}

func getListenParser() *Parsers {
	return createParsers(map[string]ParserInterface{})
}

var parserSequenceResolver = append([]Section{"nameserver", "hold-nx", "hold-obsolete", "hold-other", "hold-refused", "hold-timeout", "hold-valid", "timeout-resolve", "timeout-retry", "accepted_payload_size", "parse-resolv-conf", "resolve_retries"}, parserSequenceSections...)

func getResolverParser() *Parsers {
	return createParsers(map[string]ParserInterface{
		"nameserver": &parsers.Nameserver{},

		"hold-nx":       &simple.TimeTwoWords{Keywords: []string{"hold", "nx"}},
		"hold-obsolete": &simple.TimeTwoWords{Keywords: []string{"hold", "obsolete"}},
		"hold-other":    &simple.TimeTwoWords{Keywords: []string{"hold", "other"}},
		"hold-refused":  &simple.TimeTwoWords{Keywords: []string{"hold", "refused"}},
		"hold-timeout":  &simple.TimeTwoWords{Keywords: []string{"hold", "timeout"}},
		"hold-valid":    &simple.TimeTwoWords{Keywords: []string{"hold", "valid"}},

		"timeout-resolve": &simple.Timeout{Name: "resolve"},
		"timeout-retry":   &simple.Timeout{Name: "retry"},

		"accepted_payload_size": &simple.Word{Name: "accepted_payload_size"},
		"parse-resolv-conf":     &simple.Word{Name: "parse-resolv-conf"},
		"resolve_retries":       &simple.Word{Name: "resolve_retries"},
	})
}

var parserSequenceUserlist = append([]Section{"group", "user"}, parserSequenceSections...)

func getUserlistParser() *Parsers {
	return createParsers(map[string]ParserInterface{
		"group": &parsers.Group{},
		"user":  &parsers.User{},
	})
}

var parserSequencePeers = append([]Section{"peer"}, parserSequenceSections...)

func getPeersParser() *Parsers {
	return createParsers(map[string]ParserInterface{
		"peer": &parsers.Peer{},
	})
}

var parserSequenceMailers = append([]Section{"timeout-mail", "mailer"}, parserSequenceSections...)

func getMailersParser() *Parsers {
	return createParsers(map[string]ParserInterface{
		"timeout-mail": &simple.TimeTwoWords{Keywords: []string{"timeout", "mail"}},
		"mailer":       &parsers.Mailer{},
	})
}

var parserSequenceCache = append([]Section{"total-max-size", "max-object-size", "max-age"}, parserSequenceSections...)

func getCacheParser() *Parsers {
	return createParsers(map[string]ParserInterface{
		"total-max-size":  &simple.Number{Name: "total-max-size"},
		"max-object-size": &simple.Number{Name: "max-object-size"},
		"max-age":         &simple.Number{Name: "max-age"},
	})
}

var parserSequenceProgram = append([]Section{"command", "user", "group", "start-on-reload"}, parserSequenceSections...)

func getProgramParser() *Parsers {
	return createParsers(map[string]ParserInterface{
		"command":         &simple.String{Name: "command"},
		"user":            &simple.String{Name: "user"},
		"group":           &simple.String{Name: "group"},
		"start-on-reload": &simple.Option{Name: "start-on-reload"},
	})
}

func getHTTPErrorsParser() *Parsers {
	return createParsers(map[string]ParserInterface{})
}

func getRingParser() *Parsers {
	return createParsers(map[string]ParserInterface{})
}
