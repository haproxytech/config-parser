// Code generated by go generate; DO NOT EDIT.
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

package integration_test

const global_cpumap1403 = `
global
  cpu-map 1-4 0-3
`
const global_cpumap1all03 = `
global
  cpu-map 1/all 0-3
`
const global_cpumapauto1403 = `
global
  cpu-map auto:1-4 0-3
`
const global_cpumapauto140123 = `
global
  cpu-map auto:1-4 0-1 2-3
`
const global_statssocket1270018080 = `
global
  stats socket 127.0.0.1:8080
`
const global_statssocket1270018080modeadmin = `
global
  stats socket 127.0.0.1:8080 mode admin
`
const global_statssocketsomepathtosocket = `
global
  stats socket /some/path/to/socket
`
const global_statssocketsomepathtosocketmodea = `
global
  stats socket /some/path/to/socket mode admin
`
const global_luaprependpathusrsharehaproxylua = `
global
  lua-prepend-path /usr/share/haproxy-lua/?/init.lua
`
const global_luaprependpathusrsharehaproxylua_ = `
global
  lua-prepend-path /usr/share/haproxy-lua/?/init.lua cpath
`
const global_lualoadetchaproxyluafoolua = `
global
  lua-load /etc/haproxy/lua/foo.lua
`
const global_sslenginerdrand = `
global
  ssl-engine rdrand
`
const global_sslenginerdrandALL = `
global
  ssl-engine rdrand ALL
`
const global_sslenginerdrandRSADSA = `
global
  ssl-engine rdrand RSA,DSA
`
const global_sslmodeasync = `
global
  ssl-mode-async
`
const global_h1caseadjustcontenttypeContentTy = `
global
  h1-case-adjust content-type Content-Type
`
const global_unixbindprefixpre = `
global
  unix-bind prefix pre
`
const global_unixbindprefixpremodetest = `
global
  unix-bind prefix pre mode test
`
const global_unixbindprefixpremodetestusergga = `
global
  unix-bind prefix pre mode test user ggalinec
`
const global_unixbindprefixpremodetestusergga_ = `
global
  unix-bind prefix pre mode test user ggalinec uid 12345
`
const global_unixbindprefixpremodetestusergga__ = `
global
  unix-bind prefix pre mode test user ggalinec uid 12345 group haproxy
`
const global_unixbindprefixpremodetestusergga___ = `
global
  unix-bind prefix pre mode test user ggalinec uid 12345 group haproxy gid 6789
`
const global_threadgroupname110 = `
global
  thread-group name 1-10
`
const global_threadgroupname10 = `
global
  thread-group name 10
`
const global_setvarproccurrentstatestrprimary = `
global
  set-var proc.current_state str(primary)
`
const global_setvarprocprioint100 = `
global
  set-var proc.prio int(100)
`
const global_setvarprocthresholdint200subproc = `
global
  set-var proc.threshold int(200),sub(proc.prio)
`
const global_setvarfmtproccurrentstateprimary = `
global
  set-var-fmt proc.current_state "primary"
`
const global_setvarfmtprocbootidpidt = `
global
  set-var-fmt proc.bootid "%pid|%t"
`
const global_numacpumapping = `
global
  numa-cpu-mapping
`
const global_nonumacpumapping = `
global
  no numa-cpu-mapping
`
const global_defaultpathcurrent = `
global
  default-path current
`
const global_defaultpathconfig = `
global
  default-path config
`
const global_defaultpathparent = `
global
  default-path parent
`
const global_defaultpathoriginsomepath = `
global
  default-path origin /some/path
`
const global_defaultpathcurrentcomment = `
global
  default-path current # comment
`
const global_defaultpathoriginsomepathcomment = `
global
  default-path origin /some/path # comment
`
const global_tunequicsocketownerlistener = `
global
  tune.quic.socket-owner listener
`
const global_tunequicsocketownerconnection = `
global
  tune.quic.socket-owner connection
`
const global_httpclientresolverspreferipv4 = `
global
  httpclient.resolvers.prefer ipv4
`
const global_httpclientresolverspreferipv6 = `
global
  httpclient.resolvers.prefer ipv6
`
const global_httpclientsslverifynone = `
global
  httpclient.ssl.verify none
`
const global_httpclientsslverifyrequired = `
global
  httpclient.ssl.verify required
`
const global_httpclientsslverify = `
global
  httpclient.ssl.verify
`
const global_httperrcodes400402444446480490 = `
global
  http-err-codes 400,402-444,446-480,490
`
const global_httperrcodes400499450500 = `
global
  http-err-codes 400-499 -450 +500
`
const global_httperrcodes400408comment = `
global
  http-err-codes 400-408 # comment
`
const global_httpfailcodes400402444446480490 = `
global
  http-fail-codes 400,402-444,446-480,490
`
const global_httpfailcodes400499450500 = `
global
  http-fail-codes 400-499 -450 +500
`
const global_httpfailcodes400408comment = `
global
  http-fail-codes 400-408 # comment
`
