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

import "github.com/haproxytech/config-parser/v2/common"

//name:section
//dir:extra
//no-init:true
type Section struct {
	Name    string
	Comment string
}

//name:config-version
//dir:extra
//no-init:true
//no-get:true
type ConfigVersion struct {
	Value int64
}

//name:comments
//dir:extra
//is-multiple:true
//no-init:true
//no-parse:true
type Comments struct {
	Value string
}

//name:unprocessed
//dir:extra
//is-multiple:true
//no-init:true
//no-parse:true
//test:skip
type UnProcessed struct {
	Value string
}

//name:simple-option
//struct-name:Option
//dir:simple
//no-init:true
type SimpleOption struct {
	NoOption bool
	Comment  string
}

//name:simple-timeout
//struct-name:Timeout
//dir:simple
//no-init:true
type SimpleTimeout struct {
	Value   string
	Comment string
}

//name:simple-word
//struct-name:Word
//dir:simple
//parser-type:StringC
type SimpleWord struct{}

//name:simple-number
//struct-name:Number
//dir:simple
//parser-type:Int64C
type SimpleNumber struct{}

//name:simple-string
//struct-name:String
//dir:simple
//parser-type:StringC
type SimpleString struct{}

//name:simple-time
//struct-name:Time
//dir:simple
//parser-type:StringC
type SimpleTime struct{}

type Filter interface {
	Parse(parts []string, comment string) error
	Result() common.ReturnResultLine
}

//name:filter
//dir:filters
//is-multiple:true
//parser-type:Filter
//is-interface:true
//no-init:true
//no-parse:true
type Filters struct{}

type HTTPAction interface {
	Parse(parts []string, comment string) error
	String() string
	GetComment() string
}

//name:http-request
//struct-name:Requests
//dir:http
//is-multiple:true
//parser-type:HTTPAction
//is-interface:true
//no-init:true
//no-parse:true
//test:fail:http-request
//test:fail:http-request capture req.cook_cnt(FirstVisit),bool strlen 10
//test:ok:http-request capture req.cook_cnt(FirstVisit),bool len 10
//test:ok:http-request deny deny_status 0 unless { src 127.0.0.1 }
type HTTPRequests struct{}

//name:http-response
//struct-name:Responses
//dir:http
//is-multiple:true
//parser-type:HTTPAction
//is-interface:true
//no-init:true
//no-parse:true
//test:fail:http-response
//test:ok:http-response capture res.hdr(Server) id 0
type HTTPResponses struct{}

type TCPType interface {
	Parse(parts []string, comment string) error
	String() string
	GetComment() string
}

type TCPAction interface {
	Parse(command []string) error
	String() string
}

//name:tcp-request
//struct-name:Requests
//dir:tcp
//is-multiple:true
//parser-type:TCPType
//is-interface:true
//no-init:true
//no-parse:true
//test:fail:tcp-request
type TCPRequests struct{}

//name:tcp-response
//struct-name:Responses
//dir:tcp
//is-multiple:true
//parser-type:TCPType
//is-interface:true
//no-init:true
//no-parse:true
//test:fail:tcp-response
type TCPResponses struct{}

//name:redirect
//dir:http
//is-multiple:true
//parser-type:HTTPAction
//is-interface:true
//no-init:true
//no-parse:true
//test:fail:redirect
//test:ok:redirect prefix http://www.bar.com code 301 if { hdr(host) -i foo.com }
type Redirect struct{}

type StatsSettings interface {
	Parse(parts []string, comment string) error
	String() string
	GetComment() string
}

//name:stats
//struct-name:Stats
//dir:stats
//is-multiple:true
//parser-type:StatsSettings
//is-interface:true
//no-init:true
//no-parse:true
//test:fail:stats
//test:ok:stats admin if LOCALHOST
//test:ok:stats auth admin1:AdMiN123
//test:fail:stats auth admin1:
//test:fail:stats auth
//test:ok:stats enable
//test:ok:stats hide-version
//test:ok:stats show-legends
//test:fail:stats NON-EXISTS
//test:ok:stats maxconn 10
//test:fail:stats maxconn WORD
//test:ok:stats realm HAProxy\\ Statistics
//test:ok:stats refresh 10s
//test:fail:stats refresh
//test:ok:stats scope .
//test:fail:stats scope
//test:ok:stats show-desc Master node for Europe, Asia, Africa
//test:ok:stats show-node
//test:ok:stats show-node Europe-1
//test:ok:stats uri /admin?stats
//test:fail:stats uri
//test:ok:stats bind-process all
//test:ok:stats bind-process odd
//test:ok:stats bind-process even
//test:ok:stats bind-process 1 2 3 4
//test:ok:stats bind-process 1-4
//test:fail:stats bind-process none
//test:fail:stats bind-process 1+4
//test:fail:stats bind-process none-none
//test:fail:stats bind-process 1-4 1-3
//test:ok:stats http-request realm HAProxy\\ Statistics
//test:ok:stats http-request realm HAProxy\\ Statistics if something
//test:ok:stats http-request auth if something
//test:ok:stats http-request deny unless something
//test:ok:stats http-request allow
//test:fail:stats http-request
//test:fail:stats http-request none
type Stats struct{}
