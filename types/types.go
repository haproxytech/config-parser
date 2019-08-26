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

package types

import "github.com/haproxytech/config-parser/params"

//sections:frontend,backend
//name:acl
//is-multiple:true
//test:ok:acl url_stats path_beg /stats
//test:ok:acl url_static path_beg -i /static /images /javascript /stylesheets
//test:ok:acl url_static path_end -i .jpg .gif .png .css .js
//test:ok:acl be_app_ok nbsrv(be_app) gt 0
//test:ok:acl be_static_ok nbsrv(be_static) gt 0
//test:ok:acl key req.hdr(X-Add-ACL-Key) -m found
//test:ok:acl add path /addacl
//test:ok:acl del path /delacl
//test:ok:acl myhost hdr(Host) -f myhost.lst
//test:ok:acl clear dst_port 80
//test:ok:acl secure dst_port 8080
//test:ok:acl login_page url_beg /login
//test:ok:acl logout url_beg /logout
//test:ok:acl uid_given url_reg /login?userid=[^&]+
//test:ok:acl cookie_set hdr_sub(cookie) SEEN=1
//test:fail:acl cookie
//test:fail:acl
type ACL struct {
	Name      string
	Criterion string
	Value     string
	Comment   string
}

//sections:frontend
//name:bind
//is-multiple:true
//test:ok:bind :80,:443
//test:ok:bind 10.0.0.1:10080,10.0.0.1:10443
//test:ok:bind /var/run/ssl-frontend.sock user root mode 600 accept-proxy
//test:ok:bind :80
//test:ok:bind :443 ssl crt /etc/haproxy/site.pem
//test:ok:bind :443 ssl crt /etc/haproxy/site.pem alpn h2,http/1.1
//test:ok:bind :::443 v4v6 ssl crt /etc/haproxy/site.pem alpn h2,http/1.1
//test:ok:bind ipv6@:80
//test:ok:bind ipv4@public_ssl:443 ssl crt /etc/haproxy/site.pem
//test:ok:bind unix@ssl-frontend.sock user root mode 600 accept-proxy
//test:fail:bind
type Bind struct {
	Path    string //can be address:port or socket path
	Params  []params.BindOption
	Comment string
}

//sections:defaults,backend
//name:balance
//is-multiple:false
//test:ok:balance roundrobin
//test:ok:balance uri depth 8
//test:ok:balance uri
//test:fails:balance something
//test:fail:balance
type Balance struct {
	Algorithm string
	Arguments []string
	Comment   string
}

//sections:global
//name:cpu-map
//is-multiple:true
//test:ok:cpu-map 1-4 0-3
//test:ok:cpu-map 1/all 0-3
//test:ok:cpu-map auto:1-4 0-3
//test:ok:cpu-map auto:1-4 0-1 2-3
//test:fail:cpu-map
type CPUMap struct {
	Process string
	CPUSet  string
	Comment string
}

//sections:defaults,backend
//name:default-server
//is-multiple:true
//test:ok:default-server inter 1000 weight 13
//test:ok:default-server fall 1 rise 2 inter 3s port 4444
//test:fail:default-server
type DefaultServer struct {
	Params  []params.ServerOption
	Comment string
}

//sections:defaults,frontend,backend
//name:errorfile
//no-init:true
//is-multiple:true
//test:ok:errorfile 400 /etc/haproxy/errorfiles/400badreq.http
//test:ok:errorfile 408 /dev/null # work around Chrome pre-connect bug
//test:ok:errorfile 403 /etc/haproxy/errorfiles/403forbid.http
//test:ok:errorfile 503 /etc/haproxy/errorfiles/503sorry.http
//test:fail:errorfile
type ErrorFile struct {
	Code    string
	File    string
	Comment string
}

//sections:userlists
//name:group
//is-multiple:true
//test:ok:group G1 users tiger,scott
//test:ok:group G1
//test:fail:group
type Group struct {
	Name    string
	Users   []string
	Comment string
}

//sections:defaults,backend
//name:hash-type
//test:ok:hash-type map-based
//test:ok:hash-type map-based avalanche
//test:ok:hash-type consistent
//test:ok:hash-type consistent avalanche
//test:ok:hash-type avalanche
//test:ok:hash-type map-based sdbm
//test:ok:hash-type map-based djb2
//test:ok:hash-type map-based wt6
//test:ok:hash-type map-based crc32
//test:ok:hash-type consistent sdbm
//test:ok:hash-type consistent djb2
//test:ok:hash-type consistent wt6
//test:ok:hash-type consistent crc32
//test:ok:hash-type map-based sdbm avalanche
//test:ok:hash-type map-based djb2 avalanche
//test:ok:hash-type map-based wt6 avalanche
//test:ok:hash-type map-based crc32 avalanche
//test:ok:hash-type consistent sdbm avalanche
//test:ok:hash-type consistent djb2 avalanche
//test:ok:hash-type consistent wt6 avalanche
//test:ok:hash-type consistent crc32 avalanche
//test:fail:hash-type
type HashType struct {
	Method   string
	Function string
	Modifier string
	Comment  string
}

//sections:defaults,frontend,backend
//name:log
//is-multiple:true
//no-init:true
//no-parse:true
//test:ok:log global
//test:ok:log stdout format short daemon # send log to systemd
//test:ok:log stdout format raw daemon # send everything to stdout
//test:ok:log stderr format raw daemon notice # send important events to stderr
//test:ok:log 127.0.0.1:514 local0 notice # only send important events
//test:ok:log 127.0.0.1:514 local0 notice notice # same but limit output level
//test:ok:log 127.0.0.1:1515 len 8192 format rfc5424 local2 info
//test:fail:log
type Log struct {
	Global   bool
	NoLog    bool
	Address  string
	Length   int64
	Format   string
	Facility string
	Level    string
	MinLevel string
	Comment  string
}

//sections:mailers
//name:mailer
//is-multiple:true
//test:ok:mailer smtp1 192.168.0.1:587
//test:ok:mailer smtp1 192.168.0.1:587 # just some comment
//test:fail:mailer
type Mailer struct {
	Name    string
	IP      string
	Port    int64
	Comment string
}

//sections:frontend,backend
//name:option forwardfor
//no-parse:true
//test:ok:option forwardfor
//test:ok:option forwardfor except A
//test:ok:option forwardfor except A header B
//test:ok:option forwardfor except A header B if-none
//test:ok:option forwardfor # comment
//test:ok:option forwardfor except A # comment
//test:fail:option forwardfor except
//test:fail:option forwardfor except A header
//test:fail:option forwardfor header
type OptionForwardFor struct {
	Except  string
	Header  string
	IfNone  bool
	Comment string
}

//sections:defaults, backend
//name:option httpchk
//no-parse:true
//test:ok:option httpchk OPTIONS * HTTP/1.1\\r\\nHost:\\ www
//test:ok:option httpchk <uri>
//test:ok:option httpchk <method> <uri>
//test:ok:option httpchk <method> <uri> <version>
type OptionHttpchk struct {
	Method  string
	URI     string
	Version string
	Comment string
}

//sections:frontend
//name:option httplog
//no-parse:true
//test:ok:option httplog
//test:ok:no option httplog
//test:ok:option httplog clf
//test:ok:option httplog # comment
//test:ok:option httplog clf # comment
type OptionHTTPLog struct {
	NoOption bool
	Clf      bool
	Comment  string
}

//sections:backend
//name:option mysql-check
//no-parse:true
//test:ok:option mysql-check
//test:ok:option mysql-check user john
//test:ok:option mysql-check user john post-41
//test:ok:option mysql-check # comment
//test:fail:option mysql-check user
//test:fail:option mysql-check user # comment
type OptionMysqlCheck struct {
	NoOption bool
	User     string
	Post41   bool
	Comment  string
}

//sections:backend
//name:option redispatch
//no-parse:true
//test:ok:option redispatch
//test:ok:no option redispatch
//test:ok:option redispatch 1
//test:ok:option redispatch # comment
//test:ok:option redispatch -1 # comment
type OptionRedispatch struct {
	NoOption bool
	Interval *int64
	Comment  string
}

//sections:backend
//name:option smtpchk
//no-parse:true
//test:ok:option smtpchk
//test:ok:no option smtpchk
//test:ok:option smtpchk HELO mydomain.org
//test:ok:option smtpchk EHLO mydomain.org
//test:ok:option smtpchk # comment
//test:ok:option smtpchk HELO mydomain.org # comment
type OptionSmtpchk struct {
	NoOption bool
	Hello    string
	Domain   string
	Comment  string
}

//sections:backend
//name:external-check path
//no-parse:true
//test:ok:external-check path /usr/bin:/bin
type ExternalCheckPath struct {
	Path    string
	Comment string
}

//sections:backend
//name:external-check command
//no-parse:true
//test:ok:external-check command /bin/true
type ExternalCheckCommand struct {
	Command string
	Comment string
}

//sections:peers
//name:peer
//is-multiple:true
//test:ok:peer name 127.0.0.1:8080
//test:fail:peer name 127.0.0.1
//test:fail:peer name :8080
//test:fail:peer
type Peer struct {
	Name    string
	IP      string
	Port    int64
	Comment string
}

//sections:backend
//name:server
//is-multiple:true
//test:ok:server name 127.0.0.1:8080
//test:ok:server name 127.0.0.1
//test:ok:server name 127.0.0.1 backup
//test:fail:server
type Server struct {
	Name    string
	Address string
	Params  []params.ServerOption
	Comment string
}

//sections:frontend,backend
//name:stick-table
//test:ok:stick-table type ip size 1m expire 5m store gpc0,conn_rate(30s)
//test:ok:stick-table type ip size 1m expire 5m store gpc0,conn_rate(30s) # comment
//test:ok:stick-table type string len 1000 size 1m expire 5m store gpc0,conn_rate(30s)
//test:ok:stick-table type string len 1000 size 1m expire 5m nopurge peers aaaaa store gpc0,conn_rate(30s)
//test:fail:stick-table type string len 1000 size 1m expire 5m something peers aaaaa store gpc0,conn_rate(30s)
//test:fail:stick-table type
//test:fail:stick-table
type StickTable struct {
	Type   string
	Length string
	Size   string

	Expire  string
	NoPurge bool
	Peers   string
	Store   string
	Comment string
}

//sections:backend
//name:stats socket
//is-multiple:true
//test:ok:stats socket 127.0.0.1:8080
//test:ok:stats socket 127.0.0.1:8080 mode admin
//test:ok:stats socket /some/path/to/socket
//test:ok:stats socket /some/path/to/socket mode admin
//atest:fail:stats socket /some/path/to/socket mode
//test:fail:stats socket
type Socket struct {
	Path    string //can be address:port
	Params  []params.BindOption
	Comment string
}

//sections:backend
//name:stick
//is-multiple:true
//no-parse:true
//test:ok:stick on src table pop if !localhost
//test:ok:stick match src table pop if !localhost
//test:ok:stick store-request src table pop if !localhost
//test:fail:stick
type Stick struct {
	Type     string
	Pattern  string
	Table    string
	Cond     string
	CondTest string
	Comment  string
}

//sections:resolvers
//name:nameserver
//is-multiple:true
//test:ok:nameserver dns1 10.0.0.1:53
//test:ok:nameserver dns1 10.0.0.1:53 # comment
//test:fail:nameserver
type Nameserver struct {
	Name    string
	Address string
	Comment string
}

//sections:frontend
//name:use_backend
//is-multiple:true
//test:ok:use_backend test if TRUE
//test:ok:use_backend test if TRUE # deny
//test:ok:use_backend test # deny
//test:fail:use_backend
type UseBackend struct {
	Name     string
	Cond     string
	CondTest string
	Comment  string
}

//sections:userlists
//name:user
//is-multiple:true
//test:ok:user tiger password $6$k6y3o.eP$JlKBx(...)xHSwRv6J.C0/D7cV91 groups G1
//test:ok:user panda insecure-password elgato groups G1,G2
//test:ok:user bear insecure-password hello groups G2
//test:fail:user
type User struct {
	Name       string
	Password   string
	IsInsecure bool
	Groups     []string
	Comment    string
}

//sections:backend
//name:use-server
//is-multiple:true
//no-parse:true
//test:ok:use-server www if { req_ssl_sni -i www.example.com }
//test:ok:use-server www if { req_ssl_sni -i www.example.com } # comment
//test:fail:use-server
type UseServer struct {
	Name     string
	Cond     string
	CondTest string
	Comment  string
}
