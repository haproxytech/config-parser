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

import "github.com/haproxytech/config-parser/common"

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
type UnProcessed struct {
	Value string
}

//name:simple-option
//dir:simple
//no-init:true
type SimpleOption struct {
	NoOption bool
	Comment  string
}

//name:simple-timeout
//dir:simple
//no-init:true
type SimpleTimeout struct {
	Value   string
	Comment string
}

//name:simple-word
//dir:simple
//parser-type:StringC
type SimpleWord struct{}

//name:simple-number
//dir:simple
//parser-type:Int64C
type SimpleNumber struct{}

//name:simple-string
//dir:simple
//parser-type:StringC
type SimpleString struct{}

//name:simple-time-two-words
//dir:simple
//parser-type:StringC
//no-init:true
type SimpleTimeTwoWords struct{}

//name:simple-time
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
//dir:http
//is-multiple:true
//parser-type:HTTPAction
//is-interface:true
//no-init:true
//no-parse:true
type HTTPRequests struct{}

//name:http-response
//dir:http
//is-multiple:true
//parser-type:HTTPAction
//is-interface:true
//no-init:true
//no-parse:true
type HTTPResponses struct{}

type TCPAction interface {
	Parse(parts []string, comment string) error
	String() string
	GetComment() string
}

//name:tcp-request
//dir:tcp
//is-multiple:true
//parser-type:TCPAction
//is-interface:true
//no-init:true
//no-parse:true
type TCPRequests struct{}

//name:tcp-response
//dir:tcp
//is-multiple:true
//parser-type:TCPAction
//is-interface:true
//no-init:true
//no-parse:true
type TCPResponses struct{}

//name:redirect
//dir:http
//is-multiple:true
//parser-type:HTTPAction
//is-interface:true
//no-init:true
//no-parse:true
type Redirect struct{}
