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
	"github.com/haproxytech/config-parser/v4/parsers"
	"github.com/haproxytech/config-parser/v4/parsers/extra"
	"github.com/haproxytech/config-parser/v4/parsers/filters"
	"github.com/haproxytech/config-parser/v4/parsers/http"
	"github.com/haproxytech/config-parser/v4/parsers/simple"
	"github.com/haproxytech/config-parser/v4/parsers/stats"
	"github.com/haproxytech/config-parser/v4/parsers/tcp"
)

func addParser(parser map[string]ParserInterface, sequence *[]Section, p ParserInterface) {
	p.Init()
	parser[p.GetParserName()] = p
	*sequence = append(*sequence, Section(p.GetParserName()))
}

func (p *configParser) createParsers(parser map[string]ParserInterface, sequence []Section) *Parsers {
	addParser(parser, &sequence, &extra.Section{Name: "defaults"})
	addParser(parser, &sequence, &extra.Section{Name: "global"})
	addParser(parser, &sequence, &extra.Section{Name: "frontend"})
	addParser(parser, &sequence, &extra.Section{Name: "backend"})
	addParser(parser, &sequence, &extra.Section{Name: "listen"})
	addParser(parser, &sequence, &extra.Section{Name: "resolvers"})
	addParser(parser, &sequence, &extra.Section{Name: "userlist"})
	addParser(parser, &sequence, &extra.Section{Name: "peers"})
	addParser(parser, &sequence, &extra.Section{Name: "mailers"})
	addParser(parser, &sequence, &extra.Section{Name: "cache"})
	addParser(parser, &sequence, &extra.Section{Name: "program"})
	addParser(parser, &sequence, &extra.Section{Name: "http-errors"})
	addParser(parser, &sequence, &extra.Section{Name: "ring"})
	if !p.Options.DisableUnProcessed {
		addParser(parser, &sequence, &extra.UnProcessed{})
	}

	for _, parser := range parser {
		parser.Init()
	}

	return &Parsers{Parsers: parser, ParserSequence: sequence}
}

func (p *configParser) getStartParser() *Parsers {
	parser := map[string]ParserInterface{}
	sequence := []Section{}
	if p.Options.UseMd5Hash {
		addParser(parser, &sequence, &extra.ConfigHash{})
	}
	addParser(parser, &sequence, &extra.ConfigVersion{})
	addParser(parser, &sequence, &extra.Comments{})
	return p.createParsers(parser, sequence)
}

func (p *configParser) getDefaultParser() *Parsers {
	parser := map[string]ParserInterface{}
	sequence := []Section{}
	addParser(parser, &sequence, &parsers.Mode{})
	addParser(parser, &sequence, &parsers.MonitorURI{})
	addParser(parser, &sequence, &parsers.HashType{})
	addParser(parser, &sequence, &parsers.Balance{})
	addParser(parser, &sequence, &parsers.MaxConn{})
	addParser(parser, &sequence, &simple.Number{Name: "backlog"})
	addParser(parser, &sequence, &parsers.Log{})
	addParser(parser, &sequence, &parsers.OptionHTTPLog{})
	addParser(parser, &sequence, &stats.Stats{Mode: "defaults"})
	addParser(parser, &sequence, &simple.Word{Name: "log-tag"})
	addParser(parser, &sequence, &simple.String{Name: "log-format"})
	addParser(parser, &sequence, &simple.String{Name: "log-format-sd"})
	addParser(parser, &sequence, &parsers.Cookie{})
	addParser(parser, &sequence, &simple.Word{Name: "dynamic-cookie-key"})
	addParser(parser, &sequence, &parsers.BindProcess{})
	addParser(parser, &sequence, &simple.Option{Name: "tcplog"})
	addParser(parser, &sequence, &simple.Option{Name: "httpclose"})
	addParser(parser, &sequence, &simple.Option{Name: "http-use-htx"})
	addParser(parser, &sequence, &parsers.OptionRedispatch{})
	addParser(parser, &sequence, &simple.Option{Name: "dontlognull"})
	addParser(parser, &sequence, &simple.Option{Name: "log-separate-errors"})
	addParser(parser, &sequence, &simple.Option{Name: "http-buffer-request"})
	addParser(parser, &sequence, &simple.Option{Name: "http-server-close"})
	addParser(parser, &sequence, &simple.Option{Name: "http-keep-alive"})
	addParser(parser, &sequence, &simple.Option{Name: "http-pretend-keepalive"})
	addParser(parser, &sequence, &simple.Option{Name: "http-no-delay"})
	addParser(parser, &sequence, &simple.Option{Name: "http-proxy"})
	addParser(parser, &sequence, &simple.Option{Name: "tcpka"})
	addParser(parser, &sequence, &simple.Option{Name: "clitcpka"})
	addParser(parser, &sequence, &simple.Option{Name: "srvtcpka"})
	addParser(parser, &sequence, &simple.Option{Name: "contstats"})
	addParser(parser, &sequence, &simple.Option{Name: "ssl-hello-chk"})
	addParser(parser, &sequence, &parsers.OptionSmtpchk{})
	addParser(parser, &sequence, &simple.Option{Name: "ldap-check"})
	addParser(parser, &sequence, &parsers.OptionMysqlCheck{})
	addParser(parser, &sequence, &simple.Option{Name: "abortonclose"})
	addParser(parser, &sequence, &parsers.OptionPgsqlCheck{})
	addParser(parser, &sequence, &simple.Option{Name: "redis-check"})
	addParser(parser, &sequence, &simple.Option{Name: "splice-auto"})
	addParser(parser, &sequence, &simple.Option{Name: "splice-request"})
	addParser(parser, &sequence, &simple.Option{Name: "splice-response"})
	addParser(parser, &sequence, &simple.Option{Name: "logasap"})
	addParser(parser, &sequence, &simple.Option{Name: "log-health-checks"})
	addParser(parser, &sequence, &simple.Option{Name: "allbackups"})
	addParser(parser, &sequence, &simple.Option{Name: "external-check"})
	addParser(parser, &sequence, &parsers.OptionForwardFor{})
	addParser(parser, &sequence, &simple.Option{Name: "accept-invalid-http-request"})
	addParser(parser, &sequence, &simple.Option{Name: "accept-invalid-http-response"})
	addParser(parser, &sequence, &simple.Option{Name: "h1-case-adjust-bogus-client"})
	addParser(parser, &sequence, &simple.Option{Name: "h1-case-adjust-bogus-server"})
	addParser(parser, &sequence, &simple.Option{Name: "disable-h2-upgrade"})
	addParser(parser, &sequence, &simple.Option{Name: "tcp-check"})
	addParser(parser, &sequence, &tcp.Checks{})
	addParser(parser, &sequence, &parsers.OptionHttpchk{})
	if p.Options.UseV2HTTPCheck {
		addParser(parser, &sequence, &parsers.HTTPCheckV2{})
	} else {
		addParser(parser, &sequence, &http.Checks{Mode: "defaults"})
	}
	addParser(parser, &sequence, &parsers.ExternalCheckPath{})
	addParser(parser, &sequence, &parsers.ExternalCheckCommand{})
	addParser(parser, &sequence, &parsers.HTTPReuse{})
	addParser(parser, &sequence, &simple.Timeout{Name: "http-request"})
	addParser(parser, &sequence, &simple.Timeout{Name: "check"})
	addParser(parser, &sequence, &simple.Timeout{Name: "connect"})
	addParser(parser, &sequence, &simple.Timeout{Name: "client"})
	addParser(parser, &sequence, &simple.Timeout{Name: "client-fin"})
	addParser(parser, &sequence, &simple.Timeout{Name: "queue"})
	addParser(parser, &sequence, &simple.Timeout{Name: "server"})
	addParser(parser, &sequence, &simple.Timeout{Name: "server-fin"})
	addParser(parser, &sequence, &simple.Timeout{Name: "tunnel"})
	addParser(parser, &sequence, &simple.Timeout{Name: "http-keep-alive"})
	addParser(parser, &sequence, &simple.Number{Name: "retries"})
	addParser(parser, &sequence, &parsers.CompressionAlgo{})
	addParser(parser, &sequence, &parsers.CompressionType{})
	addParser(parser, &sequence, &parsers.CompressionOffload{})
	addParser(parser, &sequence, &parsers.DefaultServer{})
	addParser(parser, &sequence, &parsers.LoadServerStateFromFile{})
	addParser(parser, &sequence, &parsers.ErrorFile{})
	addParser(parser, &sequence, &parsers.DefaultBackend{})
	addParser(parser, &sequence, &parsers.UniqueIDFormat{})
	addParser(parser, &sequence, &parsers.UniqueIDHeader{})
	addParser(parser, &sequence, &parsers.ConfigSnippet{})
	return p.createParsers(parser, sequence)
}

func (p *configParser) getGlobalParser() *Parsers {
	parser := map[string]ParserInterface{}
	sequence := []Section{}
	// environment directives are placed before the rest,
	// because HAProxy can use the environment vars in subsequent config
	addParser(parser, &sequence, &simple.ArrayKeyValue{Name: "presetenv"})
	addParser(parser, &sequence, &simple.StringSlice{Name: "resetenv"})
	addParser(parser, &sequence, &simple.ArrayKeyValue{Name: "setenv"})
	addParser(parser, &sequence, &simple.StringSlice{Name: "unsetenv"})
	addParser(parser, &sequence, &parsers.Daemon{})
	addParser(parser, &sequence, &simple.String{Name: "localpeer"})
	addParser(parser, &sequence, &simple.Word{Name: "chroot"})
	addParser(parser, &sequence, &simple.Number{Name: "uid"})
	addParser(parser, &sequence, &simple.Word{Name: "user"})
	addParser(parser, &sequence, &simple.Number{Name: "gid"})
	addParser(parser, &sequence, &simple.Word{Name: "group"})
	addParser(parser, &sequence, &simple.Word{Name: "ca-base"})
	addParser(parser, &sequence, &simple.Word{Name: "crt-base"})
	addParser(parser, &sequence, &parsers.MasterWorker{})
	addParser(parser, &sequence, &parsers.ExternalCheck{})
	addParser(parser, &sequence, &parsers.NoSplice{})
	addParser(parser, &sequence, &parsers.NbProc{})
	addParser(parser, &sequence, &parsers.NbThread{})
	addParser(parser, &sequence, &parsers.CPUMap{})
	addParser(parser, &sequence, &parsers.Mode{})
	addParser(parser, &sequence, &parsers.MaxConn{})
	addParser(parser, &sequence, &simple.Number{Name: "maxconnrate"})
	addParser(parser, &sequence, &simple.Number{Name: "maxcomprate"})
	addParser(parser, &sequence, &simple.Number{Name: "maxcompcpuusage"})
	addParser(parser, &sequence, &simple.Number{Name: "maxpipes"})
	addParser(parser, &sequence, &simple.Number{Name: "maxsessrate"})
	addParser(parser, &sequence, &simple.Number{Name: "maxsslconn"})
	addParser(parser, &sequence, &simple.Number{Name: "maxsslrate"})
	addParser(parser, &sequence, &simple.Number{Name: "maxzlibmem"})
	addParser(parser, &sequence, &simple.String{Name: "pidfile"})
	addParser(parser, &sequence, &parsers.Socket{})
	addParser(parser, &sequence, &parsers.StatsTimeout{})
	addParser(parser, &sequence, &simple.Number{Name: "tune.buffers.limit"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.buffers.reserve"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.bufsize"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.comp.maxlevel"})
	addParser(parser, &sequence, &simple.Enabled{Name: "tune.fail-alloc"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.h2.header-table-size"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.h2.initial-window-size"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.h2.max-concurrent-streams"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.h2.max-frame-size"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.http.cookielen"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.http.logurilen"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.http.maxhdr"})
	addParser(parser, &sequence, &simple.OnOff{Name: "tune.idle-pool.shared"})
	addParser(parser, &sequence, &simple.Time{Name: "tune.idletimer"})
	addParser(parser, &sequence, &simple.OnOff{Name: "tune.listener.multi-queue"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.lua.forced-yield"})
	addParser(parser, &sequence, &simple.Enabled{Name: "tune.lua.maxmem"})
	addParser(parser, &sequence, &simple.Time{Name: "tune.lua.session-timeout"})
	addParser(parser, &sequence, &simple.Time{Name: "tune.lua.task-timeout"})
	addParser(parser, &sequence, &simple.Time{Name: "tune.lua.service-timeout"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.maxaccept"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.maxpollevents"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.maxrewrite"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.pattern.cache-size"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.pipesize"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.pool-high-fd-ratio"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.pool-low-fd-ratio"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.rcvbuf.client"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.rcvbuf.server"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.recv_enough"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.runqueue-depth"})
	addParser(parser, &sequence, &simple.OnOff{Name: "tune.sched.low-latency"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.sndbuf.client"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.sndbuf.server"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.ssl.cachesize"})
	addParser(parser, &sequence, &simple.Enabled{Name: "tune.ssl.force-private-cache"})
	addParser(parser, &sequence, &simple.OnOff{Name: "tune.ssl.keylog"})
	addParser(parser, &sequence, &simple.Time{Name: "tune.ssl.lifetime"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.ssl.maxrecord"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.ssl.default-dh-param"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.ssl.ssl-ctx-cache-size"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.ssl.capture-buffer-size"})
	addParser(parser, &sequence, &simple.Size{Name: "tune.vars.global-max-size"})
	addParser(parser, &sequence, &simple.Size{Name: "tune.vars.proc-max-size"})
	addParser(parser, &sequence, &simple.Size{Name: "tune.vars.reqres-max-size"})
	addParser(parser, &sequence, &simple.Size{Name: "tune.vars.sess-max-size"})
	addParser(parser, &sequence, &simple.Size{Name: "tune.vars.txn-max-size"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.zlib.memlevel"})
	addParser(parser, &sequence, &simple.Number{Name: "tune.zlib.windowsize"})
	addParser(parser, &sequence, &simple.String{Name: "ssl-default-bind-options"})
	addParser(parser, &sequence, &simple.Word{Name: "ssl-default-bind-ciphers"})
	addParser(parser, &sequence, &simple.Word{Name: "ssl-default-bind-ciphersuites"})
	addParser(parser, &sequence, &simple.String{Name: "ssl-default-server-options"})
	addParser(parser, &sequence, &simple.Word{Name: "ssl-default-server-ciphers"})
	addParser(parser, &sequence, &simple.Word{Name: "ssl-default-server-ciphersuites"})
	addParser(parser, &sequence, &simple.Word{Name: "ssl-dh-param-file"})
	addParser(parser, &sequence, &simple.Word{Name: "ssl-server-verify"})
	addParser(parser, &sequence, &simple.Time{Name: "hard-stop-after"})
	addParser(parser, &sequence, &parsers.Log{})
	addParser(parser, &sequence, &parsers.LogSendHostName{})
	addParser(parser, &sequence, &parsers.LuaPrependPath{})
	addParser(parser, &sequence, &parsers.LuaLoad{})
	addParser(parser, &sequence, &simple.Word{Name: "server-state-file"})
	addParser(parser, &sequence, &simple.Word{Name: "server-state-base"})
	addParser(parser, &sequence, &parsers.SslEngine{})
	addParser(parser, &sequence, &parsers.SslModeAsync{})
	addParser(parser, &sequence, &simple.Word{Name: "h1-case-adjust-file"})
	addParser(parser, &sequence, &parsers.H1CaseAdjust{})
	addParser(parser, &sequence, &simple.Enabled{Name: "busy-polling"})
	addParser(parser, &sequence, &simple.Number{Name: "max-spread-checks"})
	addParser(parser, &sequence, &simple.Enabled{Name: "noepoll"})
	addParser(parser, &sequence, &simple.Enabled{Name: "nokqueue"})
	addParser(parser, &sequence, &simple.Enabled{Name: "noevports"})
	addParser(parser, &sequence, &simple.Enabled{Name: "nopoll"})
	addParser(parser, &sequence, &simple.Enabled{Name: "nosplice"})
	addParser(parser, &sequence, &simple.Enabled{Name: "nogetaddrinfo"})
	addParser(parser, &sequence, &simple.Enabled{Name: "noreuseport"})
	addParser(parser, &sequence, &simple.AutoOnOff{Name: "profiling.tasks"})
	addParser(parser, &sequence, &simple.Number{Name: "spread-checks"})
	// the ConfigSnippet must be at the end to parsers load order to ensure
	// the overloading of any option has been declared previously
	addParser(parser, &sequence, &parsers.ConfigSnippet{})
	return p.createParsers(parser, sequence)
}

func (p *configParser) getFrontendParser() *Parsers {
	parser := map[string]ParserInterface{}
	sequence := []Section{}
	addParser(parser, &sequence, &parsers.Mode{})
	addParser(parser, &sequence, &parsers.MaxConn{})
	addParser(parser, &sequence, &simple.Number{Name: "backlog"})
	addParser(parser, &sequence, &parsers.Bind{})
	addParser(parser, &sequence, &parsers.ACL{})
	addParser(parser, &sequence, &parsers.MonitorURI{})
	addParser(parser, &sequence, &parsers.MonitorFail{})
	addParser(parser, &sequence, &parsers.BindProcess{})
	addParser(parser, &sequence, &simple.Word{Name: "log-tag"})
	addParser(parser, &sequence, &simple.String{Name: "log-format"})
	addParser(parser, &sequence, &simple.String{Name: "log-format-sd"})
	addParser(parser, &sequence, &parsers.Log{})
	addParser(parser, &sequence, &simple.Option{Name: "httpclose"})
	addParser(parser, &sequence, &simple.Option{Name: "forceclose"})
	addParser(parser, &sequence, &simple.Option{Name: "http-buffer-request"})
	addParser(parser, &sequence, &simple.Option{Name: "http-server-close"})
	addParser(parser, &sequence, &simple.Option{Name: "http-keep-alive"})
	addParser(parser, &sequence, &simple.Option{Name: "http-use-htx"})
	addParser(parser, &sequence, &simple.Option{Name: "http-no-delay"})
	addParser(parser, &sequence, &simple.Option{Name: "http-proxy"})
	addParser(parser, &sequence, &parsers.OptionForwardFor{})
	addParser(parser, &sequence, &simple.Option{Name: "tcplog"})
	addParser(parser, &sequence, &simple.Option{Name: "dontlognull"})
	addParser(parser, &sequence, &simple.Option{Name: "contstats"})
	addParser(parser, &sequence, &simple.Option{Name: "log-separate-errors"})
	addParser(parser, &sequence, &simple.Option{Name: "tcpka"})
	addParser(parser, &sequence, &simple.Option{Name: "clitcpka"})
	addParser(parser, &sequence, &simple.Option{Name: "splice-auto"})
	addParser(parser, &sequence, &simple.Option{Name: "splice-request"})
	addParser(parser, &sequence, &simple.Option{Name: "splice-response"})
	addParser(parser, &sequence, &simple.Option{Name: "accept-invalid-http-request"})
	addParser(parser, &sequence, &simple.Option{Name: "h1-case-adjust-bogus-client"})
	addParser(parser, &sequence, &simple.Option{Name: "disable-h2-upgrade"})
	addParser(parser, &sequence, &simple.Option{Name: "logasap"})
	addParser(parser, &sequence, &parsers.OptionHTTPLog{})
	addParser(parser, &sequence, &simple.Timeout{Name: "http-request"})
	addParser(parser, &sequence, &simple.Timeout{Name: "client"})
	addParser(parser, &sequence, &simple.Timeout{Name: "client-fin"})
	addParser(parser, &sequence, &simple.Timeout{Name: "http-keep-alive"})
	addParser(parser, &sequence, &filters.Filters{})
	addParser(parser, &sequence, &parsers.CompressionAlgo{})
	addParser(parser, &sequence, &parsers.CompressionType{})
	addParser(parser, &sequence, &parsers.CompressionOffload{})
	addParser(parser, &sequence, &tcp.Requests{})
	addParser(parser, &sequence, &stats.Stats{Mode: "frontend"})
	addParser(parser, &sequence, &http.Requests{Mode: "frontend"})
	addParser(parser, &sequence, &http.Redirect{})
	addParser(parser, &sequence, &parsers.UniqueIDFormat{})
	addParser(parser, &sequence, &parsers.UniqueIDHeader{})
	addParser(parser, &sequence, &parsers.ErrorFile{})
	addParser(parser, &sequence, &parsers.ConfigSnippet{})
	addParser(parser, &sequence, &parsers.UseBackend{})
	addParser(parser, &sequence, &parsers.DefaultBackend{})
	addParser(parser, &sequence, &parsers.StickTable{})
	addParser(parser, &sequence, &tcp.Responses{})
	addParser(parser, &sequence, &http.Responses{Mode: "frontend"})
	addParser(parser, &sequence, &parsers.DeclareCapture{})
	return p.createParsers(parser, sequence)
}

func (p *configParser) getBackendParser() *Parsers {
	parser := map[string]ParserInterface{}
	sequence := []Section{}
	addParser(parser, &sequence, &parsers.Mode{})
	addParser(parser, &sequence, &parsers.HashType{})
	addParser(parser, &sequence, &parsers.Balance{})
	addParser(parser, &sequence, &parsers.ACL{})
	addParser(parser, &sequence, &parsers.ForcePersist{})
	addParser(parser, &sequence, &parsers.BindProcess{})
	addParser(parser, &sequence, &simple.Option{Name: "httpclose"})
	addParser(parser, &sequence, &simple.Option{Name: "forceclose"})
	addParser(parser, &sequence, &simple.Option{Name: "http-buffer-request"})
	addParser(parser, &sequence, &simple.Option{Name: "http-server-close"})
	addParser(parser, &sequence, &simple.Option{Name: "http-keep-alive"})
	addParser(parser, &sequence, &simple.Option{Name: "http-pretend-keepalive"})
	addParser(parser, &sequence, &simple.Option{Name: "http-use-htx"})
	addParser(parser, &sequence, &simple.Option{Name: "http-no-delay"})
	addParser(parser, &sequence, &simple.Option{Name: "http-proxy"})
	addParser(parser, &sequence, &parsers.OptionForwardFor{})
	addParser(parser, &sequence, &simple.Option{Name: "ssl-hello-chk"})
	addParser(parser, &sequence, &parsers.OptionSmtpchk{})
	addParser(parser, &sequence, &simple.Option{Name: "ldap-check"})
	addParser(parser, &sequence, &parsers.OptionMysqlCheck{})
	addParser(parser, &sequence, &simple.Option{Name: "abortonclose"})
	addParser(parser, &sequence, &parsers.OptionPgsqlCheck{})
	addParser(parser, &sequence, &simple.Option{Name: "redis-check"})
	addParser(parser, &sequence, &parsers.OptionRedispatch{})
	addParser(parser, &sequence, &simple.Option{Name: "external-check"})
	addParser(parser, &sequence, &simple.Option{Name: "splice-auto"})
	addParser(parser, &sequence, &simple.Option{Name: "splice-request"})
	addParser(parser, &sequence, &simple.Option{Name: "splice-response"})
	addParser(parser, &sequence, &simple.Option{Name: "log-health-checks"})
	addParser(parser, &sequence, &simple.String{Name: "log-tag"})
	addParser(parser, &sequence, &simple.Option{Name: "tcpka"})
	addParser(parser, &sequence, &simple.Option{Name: "srvtcpka"})
	addParser(parser, &sequence, &simple.Option{Name: "allbackups"})
	addParser(parser, &sequence, &simple.Option{Name: "accept-invalid-http-response"})
	addParser(parser, &sequence, &simple.Option{Name: "h1-case-adjust-bogus-server"})
	addParser(parser, &sequence, &simple.Option{Name: "tcp-check"})
	addParser(parser, &sequence, &tcp.Checks{})
	addParser(parser, &sequence, &parsers.OptionHttpchk{})
	if p.Options.UseV2HTTPCheck {
		addParser(parser, &sequence, &parsers.HTTPCheckV2{})
	} else {
		addParser(parser, &sequence, &http.Checks{Mode: "backend"})
	}
	addParser(parser, &sequence, &parsers.ExternalCheckPath{})
	addParser(parser, &sequence, &parsers.ExternalCheckCommand{})
	addParser(parser, &sequence, &parsers.Log{})
	addParser(parser, &sequence, &simple.Timeout{Name: "http-request"})
	addParser(parser, &sequence, &simple.Timeout{Name: "queue"})
	addParser(parser, &sequence, &simple.Timeout{Name: "http-keep-alive"})
	addParser(parser, &sequence, &simple.Timeout{Name: "check"})
	addParser(parser, &sequence, &simple.Timeout{Name: "tunnel"})
	addParser(parser, &sequence, &simple.Timeout{Name: "server"})
	addParser(parser, &sequence, &simple.Timeout{Name: "server-fin"})
	addParser(parser, &sequence, &simple.Timeout{Name: "connect"})
	addParser(parser, &sequence, &parsers.DefaultServer{})
	addParser(parser, &sequence, &parsers.Stick{})
	addParser(parser, &sequence, &filters.Filters{})
	addParser(parser, &sequence, &parsers.CompressionAlgo{})
	addParser(parser, &sequence, &parsers.CompressionType{})
	addParser(parser, &sequence, &parsers.CompressionOffload{})
	addParser(parser, &sequence, &tcp.Requests{})
	addParser(parser, &sequence, &stats.Stats{Mode: "backend"})
	addParser(parser, &sequence, &parsers.HTTPReuse{})
	addParser(parser, &sequence, &http.Requests{Mode: "backend"})
	addParser(parser, &sequence, &http.Redirect{})
	addParser(parser, &sequence, &parsers.Cookie{})
	addParser(parser, &sequence, &simple.Word{Name: "dynamic-cookie-key"})
	addParser(parser, &sequence, &parsers.UseServer{})
	addParser(parser, &sequence, &parsers.StickTable{})
	addParser(parser, &sequence, &parsers.ConfigSnippet{})
	addParser(parser, &sequence, &parsers.ErrorFile{})
	addParser(parser, &sequence, &parsers.Server{})
	addParser(parser, &sequence, &simple.Number{Name: "retries"})
	addParser(parser, &sequence, &tcp.Responses{})
	addParser(parser, &sequence, &http.Responses{Mode: "backend"})
	addParser(parser, &sequence, &parsers.ServerTemplate{})
	addParser(parser, &sequence, &parsers.LoadServerStateFromFile{})
	return p.createParsers(parser, sequence)
}

func (p *configParser) getListenParser() *Parsers {
	parser := map[string]ParserInterface{}
	sequence := []Section{}
	if p.Options.UseListenSectionParsers {
		addParser(parser, &sequence, &parsers.Mode{})
		addParser(parser, &sequence, &parsers.HashType{})
		addParser(parser, &sequence, &parsers.Balance{})
		addParser(parser, &sequence, &parsers.MaxConn{})
		addParser(parser, &sequence, &simple.Number{Name: "backlog"})
		addParser(parser, &sequence, &parsers.Bind{})
		addParser(parser, &sequence, &parsers.ACL{})
		addParser(parser, &sequence, &parsers.ForcePersist{})
		addParser(parser, &sequence, &parsers.MonitorURI{})
		addParser(parser, &sequence, &parsers.MonitorFail{})
		addParser(parser, &sequence, &parsers.BindProcess{})
		addParser(parser, &sequence, &simple.Word{Name: "log-tag"})
		addParser(parser, &sequence, &simple.String{Name: "log-format"})
		addParser(parser, &sequence, &simple.String{Name: "log-format-sd"})
		addParser(parser, &sequence, &parsers.Log{})
		addParser(parser, &sequence, &simple.Option{Name: "httpclose"})
		addParser(parser, &sequence, &simple.Option{Name: "forceclose"})
		addParser(parser, &sequence, &simple.Option{Name: "http-buffer-request"})
		addParser(parser, &sequence, &simple.Option{Name: "http-server-close"})
		addParser(parser, &sequence, &simple.Option{Name: "http-keep-alive"})
		addParser(parser, &sequence, &simple.Option{Name: "http-pretend-keepalive"})
		addParser(parser, &sequence, &simple.Option{Name: "http-use-htx"})
		addParser(parser, &sequence, &simple.Option{Name: "http-no-delay"})
		addParser(parser, &sequence, &simple.Option{Name: "http-proxy"})
		addParser(parser, &sequence, &parsers.OptionForwardFor{})
		addParser(parser, &sequence, &simple.Option{Name: "ssl-hello-chk"})
		addParser(parser, &sequence, &parsers.OptionSmtpchk{})
		addParser(parser, &sequence, &simple.Option{Name: "ldap-check"})
		addParser(parser, &sequence, &parsers.OptionMysqlCheck{})
		addParser(parser, &sequence, &simple.Option{Name: "abortonclose"})
		addParser(parser, &sequence, &parsers.OptionPgsqlCheck{})
		addParser(parser, &sequence, &simple.Option{Name: "redis-check"})
		addParser(parser, &sequence, &parsers.OptionRedispatch{})
		addParser(parser, &sequence, &simple.Option{Name: "external-check"})
		addParser(parser, &sequence, &simple.Option{Name: "tcplog"})
		addParser(parser, &sequence, &simple.Option{Name: "dontlognull"})
		addParser(parser, &sequence, &simple.Option{Name: "contstats"})
		addParser(parser, &sequence, &simple.Option{Name: "log-separate-errors"})
		addParser(parser, &sequence, &simple.Option{Name: "tcpka"})
		addParser(parser, &sequence, &simple.Option{Name: "clitcpka"})
		addParser(parser, &sequence, &simple.Option{Name: "splice-auto"})
		addParser(parser, &sequence, &simple.Option{Name: "splice-request"})
		addParser(parser, &sequence, &simple.Option{Name: "splice-response"})
		addParser(parser, &sequence, &simple.Option{Name: "log-health-checks"})
		addParser(parser, &sequence, &simple.String{Name: "log-tag"})
		addParser(parser, &sequence, &simple.Option{Name: "tcpka"})
		addParser(parser, &sequence, &simple.Option{Name: "srvtcpka"})
		addParser(parser, &sequence, &simple.Option{Name: "allbackups"})
		addParser(parser, &sequence, &simple.Option{Name: "accept-invalid-http-request"})
		addParser(parser, &sequence, &simple.Option{Name: "h1-case-adjust-bogus-client"})
		addParser(parser, &sequence, &simple.Option{Name: "disable-h2-upgrade"})
		addParser(parser, &sequence, &simple.Option{Name: "logasap"})
		addParser(parser, &sequence, &parsers.OptionHTTPLog{})
		addParser(parser, &sequence, &simple.Option{Name: "accept-invalid-http-response"})
		addParser(parser, &sequence, &simple.Option{Name: "h1-case-adjust-bogus-server"})
		addParser(parser, &sequence, &simple.Option{Name: "tcp-check"})
		addParser(parser, &sequence, &tcp.Checks{})
		addParser(parser, &sequence, &parsers.OptionHttpchk{})
		if p.Options.UseV2HTTPCheck {
			addParser(parser, &sequence, &parsers.HTTPCheckV2{})
		} else {
			addParser(parser, &sequence, &http.Checks{Mode: "listen"})
		}
		addParser(parser, &sequence, &parsers.ExternalCheckPath{})
		addParser(parser, &sequence, &parsers.ExternalCheckCommand{})
		addParser(parser, &sequence, &simple.Timeout{Name: "http-request"})
		addParser(parser, &sequence, &simple.Timeout{Name: "client"})
		addParser(parser, &sequence, &simple.Timeout{Name: "client-fin"})
		addParser(parser, &sequence, &simple.Timeout{Name: "queue"})
		addParser(parser, &sequence, &simple.Timeout{Name: "http-keep-alive"})
		addParser(parser, &sequence, &simple.Timeout{Name: "check"})
		addParser(parser, &sequence, &simple.Timeout{Name: "tunnel"})
		addParser(parser, &sequence, &simple.Timeout{Name: "server"})
		addParser(parser, &sequence, &simple.Timeout{Name: "server-fin"})
		addParser(parser, &sequence, &simple.Timeout{Name: "connect"})
		addParser(parser, &sequence, &parsers.DefaultServer{})
		addParser(parser, &sequence, &parsers.Stick{})
		addParser(parser, &sequence, &filters.Filters{})
		addParser(parser, &sequence, &parsers.CompressionAlgo{})
		addParser(parser, &sequence, &parsers.CompressionType{})
		addParser(parser, &sequence, &parsers.CompressionOffload{})
		addParser(parser, &sequence, &tcp.Requests{})
		addParser(parser, &sequence, &stats.Stats{Mode: "listen"})
		addParser(parser, &sequence, &parsers.HTTPReuse{})
		addParser(parser, &sequence, &http.Requests{Mode: "listen"})
		addParser(parser, &sequence, &http.Redirect{})
		addParser(parser, &sequence, &parsers.Cookie{})
		addParser(parser, &sequence, &simple.Word{Name: "dynamic-cookie-key"})
		addParser(parser, &sequence, &parsers.UseServer{})
		addParser(parser, &sequence, &parsers.UniqueIDFormat{})
		addParser(parser, &sequence, &parsers.UniqueIDHeader{})
		addParser(parser, &sequence, &parsers.ErrorFile{})
		addParser(parser, &sequence, &parsers.ConfigSnippet{})
		addParser(parser, &sequence, &parsers.UseBackend{})
		addParser(parser, &sequence, &parsers.DefaultBackend{})
		addParser(parser, &sequence, &parsers.StickTable{})
		addParser(parser, &sequence, &parsers.ConfigSnippet{})
		addParser(parser, &sequence, &parsers.ErrorFile{})
		addParser(parser, &sequence, &parsers.Server{})
		addParser(parser, &sequence, &simple.Number{Name: "retries"})
		addParser(parser, &sequence, &tcp.Responses{})
		addParser(parser, &sequence, &http.Responses{Mode: "listen"})
		addParser(parser, &sequence, &parsers.DeclareCapture{})
	}
	return p.createParsers(parser, sequence)
}

func (p *configParser) getResolverParser() *Parsers {
	parser := map[string]ParserInterface{}
	sequence := []Section{}
	addParser(parser, &sequence, &parsers.Nameserver{})
	addParser(parser, &sequence, &simple.TimeTwoWords{Keywords: []string{"hold", "nx"}})
	addParser(parser, &sequence, &simple.TimeTwoWords{Keywords: []string{"hold", "obsolete"}})
	addParser(parser, &sequence, &simple.TimeTwoWords{Keywords: []string{"hold", "other"}})
	addParser(parser, &sequence, &simple.TimeTwoWords{Keywords: []string{"hold", "refused"}})
	addParser(parser, &sequence, &simple.TimeTwoWords{Keywords: []string{"hold", "timeout"}})
	addParser(parser, &sequence, &simple.TimeTwoWords{Keywords: []string{"hold", "valid"}})
	addParser(parser, &sequence, &simple.Timeout{Name: "resolve"})
	addParser(parser, &sequence, &simple.Timeout{Name: "retry"})
	addParser(parser, &sequence, &simple.Word{Name: "accepted_payload_size"})
	addParser(parser, &sequence, &simple.Word{Name: "parse-resolv-conf"})
	addParser(parser, &sequence, &simple.Word{Name: "resolve_retries"})
	return p.createParsers(parser, sequence)
}

func (p *configParser) getUserlistParser() *Parsers {
	parser := map[string]ParserInterface{}
	sequence := []Section{}
	addParser(parser, &sequence, &parsers.Group{})
	addParser(parser, &sequence, &parsers.User{})
	return p.createParsers(parser, sequence)
}

func (p *configParser) getPeersParser() *Parsers {
	parser := map[string]ParserInterface{}
	sequence := []Section{}
	addParser(parser, &sequence, &parsers.Peer{})
	return p.createParsers(parser, sequence)
}

func (p *configParser) getMailersParser() *Parsers {
	parser := map[string]ParserInterface{}
	sequence := []Section{}
	addParser(parser, &sequence, &simple.TimeTwoWords{Keywords: []string{"timeout", "mail"}})
	addParser(parser, &sequence, &parsers.Mailer{})
	return p.createParsers(parser, sequence)
}

func (p *configParser) getCacheParser() *Parsers {
	parser := map[string]ParserInterface{}
	sequence := []Section{}
	addParser(parser, &sequence, &simple.Number{Name: "total-max-size"})
	addParser(parser, &sequence, &simple.Number{Name: "max-object-size"})
	addParser(parser, &sequence, &simple.Number{Name: "max-age"})
	addParser(parser, &sequence, &simple.Number{Name: "max-secondary-entries"})
	addParser(parser, &sequence, &parsers.ProcessVary{})
	return p.createParsers(parser, sequence)
}

func (p *configParser) getProgramParser() *Parsers {
	parser := map[string]ParserInterface{}
	sequence := []Section{}
	addParser(parser, &sequence, &simple.String{Name: "command"})
	addParser(parser, &sequence, &simple.String{Name: "user"})
	addParser(parser, &sequence, &simple.String{Name: "group"})
	addParser(parser, &sequence, &simple.Option{Name: "start-on-reload"})
	return p.createParsers(parser, sequence)
}

func (p *configParser) getHTTPErrorsParser() *Parsers {
	return p.createParsers(map[string]ParserInterface{}, []Section{})
}

func (p *configParser) getRingParser() *Parsers {
	return p.createParsers(map[string]ParserInterface{}, []Section{})
}
