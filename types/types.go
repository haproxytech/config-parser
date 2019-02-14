package types

import "github.com/haproxytech/config-parser/params"

//sections:frontend
//name:bind
//is-multiple:true
//test:ok:bind :80,:443
//test:ok:bind 10.0.0.1:10080,10.0.0.1:10443
//test:ok:bind /var/run/ssl-frontend.sock user root mode 600 accept-proxy
//test:ok:bind :80
//test:ok:bind :443 ssl crt /etc/haproxy/site.pem
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
type CpuMap struct {
	Name    string
	Value   string
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

//sections:defaults, backend
//name:option httpchk
//no-parse:true
//test:ok:option httpchk OPTIONS * HTTP/1.1\\r\\nHost:\\ www
type OptionHttpchk struct {
	Method  string
	Uri     string
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

type Section struct {
	Name    string
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
//is-multiple:true
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

type SimpleOption struct {
	NoOption bool
	Comment  string
}

type SimpleTimeout struct {
	Value   string
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
	Name          string
	Pattern       string
	Table         string
	Condition     string
	ConditionType string
	Comment       string
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
//test:fail:use_backend
type UseBackend struct {
	Name          string
	Condition     string
	ConditionKind string
	Comment       string
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
	Name          string
	Condition     string
	ConditionType string
	Comment       string
}
