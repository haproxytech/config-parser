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

const defaults_balanceroundrobin = `
defaults test
  balance roundrobin
`
const defaults_balanceuri = `
defaults test
  balance uri
`
const defaults_balanceuriwhole = `
defaults test
  balance uri whole
`
const defaults_balanceurilen12 = `
defaults test
  balance uri len 12
`
const defaults_balanceuridepth8 = `
defaults test
  balance uri depth 8
`
const defaults_balanceuridepth8whole = `
defaults test
  balance uri depth 8 whole
`
const defaults_balanceuridepth8len12whole = `
defaults test
  balance uri depth 8 len 12 whole
`
const defaults_balanceurlparam = `
defaults test
  balance url_param
`
const defaults_balanceurlparamsessionid = `
defaults test
  balance url_param session_id
`
const defaults_balanceurlparamcheckpost10 = `
defaults test
  balance url_param check_post 10
`
const defaults_balanceurlparamcheckpost10maxwai = `
defaults test
  balance url_param check_post 10 max_wait 20
`
const defaults_balanceurlparamsessionidcheckpos = `
defaults test
  balance url_param session_id check_post 10 max_wait 20
`
const defaults_balancehdrhdrName = `
defaults test
  balance hdr(hdrName)
`
const defaults_balancehdrhdrNameusedomainonly = `
defaults test
  balance hdr(hdrName) use_domain_only
`
const defaults_balancerandom = `
defaults test
  balance random
`
const defaults_balancerandom15 = `
defaults test
  balance random(15)
`
const defaults_balancerdpcookie = `
defaults test
  balance rdp-cookie
`
const defaults_balancerdpcookiesomething = `
defaults test
  balance rdp-cookie(something)
`
const defaults_balancehashreqcookieclientid = `
defaults test
  balance hash req.cookie(clientid)
`
const defaults_balancehashreqhdripxforwardedfor = `
defaults test
  balance hash req.hdr_ip(x-forwarded-for,-1),ipmask(24)
`
const defaults_balancehashreqhdripxforwardedfor_ = `
defaults test
  balance hash req.hdr_ip(x-forwarded-for ,-1),ipmask(24)
`
const defaults_persistrdpcookie = `
defaults test
  persist rdp-cookie
`
const defaults_persistrdpcookiecookies = `
defaults test
  persist rdp-cookie(cookies)
`
const defaults_cookietest = `
defaults test
  cookie test
`
const defaults_cookiemyCookiedomaindom1indirect = `
defaults test
  cookie myCookie domain dom1 indirect postonly
`
const defaults_cookiemyCookiedomaindom1domaindo = `
defaults test
  cookie myCookie domain dom1 domain dom2 indirect postonly
`
const defaults_cookiemyCookieindirectmaxidle10m = `
defaults test
  cookie myCookie indirect maxidle 10 maxlife 5 postonly
`
const defaults_cookiemyCookieindirectmaxidle10 = `
defaults test
  cookie myCookie indirect maxidle 10
`
const defaults_cookiemyCookieindirectmaxlife10 = `
defaults test
  cookie myCookie indirect maxlife 10
`
const defaults_cookiemyCookiedomaindom1domaindo_ = `
defaults test
  cookie myCookie domain dom1 domain dom2 httponly indirect maxidle 10 maxlife 5 nocache postonly preserve rewrite secure
`
const defaults_cookiemyCookieattrSameSiteStrict = `
defaults test
  cookie myCookie attr \"SameSite=Strict\" attr \"mykey=myvalue\" insert
`
const defaults_defaultserveraddr127001 = `
defaults test
  default-server addr 127.0.0.1
`
const defaults_defaultserveraddr1 = `
defaults test
  default-server addr ::1
`
const defaults_defaultserveragentcheck = `
defaults test
  default-server agent-check
`
const defaults_defaultserveragentsendname = `
defaults test
  default-server agent-send name
`
const defaults_defaultserveragentinter1000ms = `
defaults test
  default-server agent-inter 1000ms
`
const defaults_defaultserveragentaddr127001 = `
defaults test
  default-server agent-addr 127.0.0.1
`
const defaults_defaultserveragentaddrsitecom = `
defaults test
  default-server agent-addr site.com
`
const defaults_defaultserveragentport1 = `
defaults test
  default-server agent-port 1
`
const defaults_defaultserveragentport65535 = `
defaults test
  default-server agent-port 65535
`
const defaults_defaultserverallow0rtt = `
defaults test
  default-server allow-0rtt
`
const defaults_defaultserveralpnh2 = `
defaults test
  default-server alpn h2
`
const defaults_defaultserveralpnhttp11 = `
defaults test
  default-server alpn http/1.1
`
const defaults_defaultserveralpnh2http11 = `
defaults test
  default-server alpn h2,http/1.1
`
const defaults_defaultserverbackup = `
defaults test
  default-server backup
`
const defaults_defaultservercafilecertcrt = `
defaults test
  default-server ca-file cert.crt
`
const defaults_defaultservercheck = `
defaults test
  default-server check
`
const defaults_defaultserverchecksendproxy = `
defaults test
  default-server check-send-proxy
`
const defaults_defaultservercheckalpnhttp10 = `
defaults test
  default-server check-alpn http/1.0
`
const defaults_defaultservercheckalpnhttp11http = `
defaults test
  default-server check-alpn http/1.1,http/1.0
`
const defaults_defaultservercheckprotoh2 = `
defaults test
  default-server check-proto h2
`
const defaults_defaultservercheckssl = `
defaults test
  default-server check-ssl
`
const defaults_defaultservercheckviasocks4 = `
defaults test
  default-server check-via-socks4
`
const defaults_defaultserverciphersECDHERSAAES1 = `
defaults test
  default-server ciphers ECDHE-RSA-AES128-GCM-SHA256
`
const defaults_defaultserverciphersECDHEECDSACH = `
defaults test
  default-server ciphers ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA:ECDHE-RSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-RSA-AES256-SHA256:DHE-RSA-AES256-SHA:ECDHE-ECDSA-DES-CBC3-SHA:ECDHE-RSA-DES-CBC3-SHA:EDH-RSA-DES-CBC3-SHA:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:DES-CBC3-SHA:!DSS
`
const defaults_defaultserverciphersuitesECDHEEC = `
defaults test
  default-server ciphersuites ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA:ECDHE-RSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-RSA-AES256-SHA256:DHE-RSA-AES256-SHA:ECDHE-ECDSA-DES-CBC3-SHA:ECDHE-RSA-DES-CBC3-SHA:EDH-RSA-DES-CBC3-SHA:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:DES-CBC3-SHA:!DSS
`
const defaults_defaultservercookievalue = `
defaults test
  default-server cookie value
`
const defaults_defaultservercrlfilefilepem = `
defaults test
  default-server crl-file file.pem
`
const defaults_defaultservercrtcertpem = `
defaults test
  default-server crt cert.pem
`
const defaults_defaultserverdisabled = `
defaults test
  default-server disabled
`
const defaults_defaultserverenabled = `
defaults test
  default-server enabled
`
const defaults_defaultservererrorlimit50 = `
defaults test
  default-server error-limit 50
`
const defaults_defaultserverfall30 = `
defaults test
  default-server fall 30
`
const defaults_defaultserverfall1rise2inter3spo = `
defaults test
  default-server fall 1 rise 2 inter 3s port 4444
`
const defaults_defaultserverforcesslv3 = `
defaults test
  default-server force-sslv3
`
const defaults_defaultserverforcetlsv10 = `
defaults test
  default-server force-tlsv10
`
const defaults_defaultserverforcetlsv11 = `
defaults test
  default-server force-tlsv11
`
const defaults_defaultserverforcetlsv12 = `
defaults test
  default-server force-tlsv12
`
const defaults_defaultserverforcetlsv13 = `
defaults test
  default-server force-tlsv13
`
const defaults_defaultserverinitaddrlastlibcnon = `
defaults test
  default-server init-addr last,libc,none
`
const defaults_defaultserverinitaddrlastlibcnon_ = `
defaults test
  default-server init-addr last,libc,none,127.0.0.1
`
const defaults_defaultserverinter1500ms = `
defaults test
  default-server inter 1500ms
`
const defaults_defaultserverinter1000weight13 = `
defaults test
  default-server inter 1000 weight 13
`
const defaults_defaultserverfastinter2500ms = `
defaults test
  default-server fastinter 2500ms
`
const defaults_defaultserverfastinterunknown = `
defaults test
  default-server fastinter unknown
`
const defaults_defaultserverdowninter3500ms = `
defaults test
  default-server downinter 3500ms
`
const defaults_defaultserverlogprotolegacy = `
defaults test
  default-server log-proto legacy
`
const defaults_defaultserverlogprotooctetcount = `
defaults test
  default-server log-proto octet-count
`
const defaults_defaultservermaxconn1 = `
defaults test
  default-server maxconn 1
`
const defaults_defaultservermaxconn50 = `
defaults test
  default-server maxconn 50
`
const defaults_defaultservermaxqueue0 = `
defaults test
  default-server maxqueue 0
`
const defaults_defaultservermaxqueue1000 = `
defaults test
  default-server maxqueue 1000
`
const defaults_defaultservermaxreuse1 = `
defaults test
  default-server max-reuse -1
`
const defaults_defaultservermaxreuse0 = `
defaults test
  default-server max-reuse 0
`
const defaults_defaultservermaxreuse1_ = `
defaults test
  default-server max-reuse 1
`
const defaults_defaultserverminconn1 = `
defaults test
  default-server minconn 1
`
const defaults_defaultserverminconn50 = `
defaults test
  default-server minconn 50
`
const defaults_defaultservernamespacetest = `
defaults test
  default-server namespace test
`
const defaults_defaultservernoagentcheck = `
defaults test
  default-server no-agent-check
`
const defaults_defaultservernobackup = `
defaults test
  default-server no-backup
`
const defaults_defaultservernocheck = `
defaults test
  default-server no-check
`
const defaults_defaultservernocheckssl = `
defaults test
  default-server no-check-ssl
`
const defaults_defaultservernosendproxyv2 = `
defaults test
  default-server no-send-proxy-v2
`
const defaults_defaultservernosendproxyv2ssl = `
defaults test
  default-server no-send-proxy-v2-ssl
`
const defaults_defaultservernosendproxyv2sslcn = `
defaults test
  default-server no-send-proxy-v2-ssl-cn
`
const defaults_defaultservernossl = `
defaults test
  default-server no-ssl
`
const defaults_defaultservernosslreuse = `
defaults test
  default-server no-ssl-reuse
`
const defaults_defaultservernosslv3 = `
defaults test
  default-server no-sslv3
`
const defaults_defaultservernotlstickets = `
defaults test
  default-server no-tls-tickets
`
const defaults_defaultservernotlsv10 = `
defaults test
  default-server no-tlsv10
`
const defaults_defaultservernotlsv11 = `
defaults test
  default-server no-tlsv11
`
const defaults_defaultservernotlsv12 = `
defaults test
  default-server no-tlsv12
`
const defaults_defaultservernotlsv13 = `
defaults test
  default-server no-tlsv13
`
const defaults_defaultservernoverifyhost = `
defaults test
  default-server no-verifyhost
`
const defaults_defaultservernotfo = `
defaults test
  default-server no-tfo
`
const defaults_defaultservernonstick = `
defaults test
  default-server non-stick
`
const defaults_defaultservernpnhttp11http10 = `
defaults test
  default-server npn http/1.1,http/1.0
`
const defaults_defaultserverobservelayer4 = `
defaults test
  default-server observe layer4
`
const defaults_defaultserverobservelayer7 = `
defaults test
  default-server observe layer7
`
const defaults_defaultserveronerrorfastinter = `
defaults test
  default-server on-error fastinter
`
const defaults_defaultserveronerrorfailcheck = `
defaults test
  default-server on-error fail-check
`
const defaults_defaultserveronerrorsuddendeath = `
defaults test
  default-server on-error sudden-death
`
const defaults_defaultserveronerrormarkdown = `
defaults test
  default-server on-error mark-down
`
const defaults_defaultserveronmarkeddownshutdow = `
defaults test
  default-server on-marked-down shutdown-sessions
`
const defaults_defaultserveronmarkedupshutdownb = `
defaults test
  default-server on-marked-up shutdown-backup-session
`
const defaults_defaultserverpoolmaxconn1 = `
defaults test
  default-server pool-max-conn -1
`
const defaults_defaultserverpoolmaxconn0 = `
defaults test
  default-server pool-max-conn 0
`
const defaults_defaultserverpoolmaxconn100 = `
defaults test
  default-server pool-max-conn 100
`
const defaults_defaultserverpoolpurgedelay0 = `
defaults test
  default-server pool-purge-delay 0
`
const defaults_defaultserverpoolpurgedelay5 = `
defaults test
  default-server pool-purge-delay 5
`
const defaults_defaultserverpoolpurgedelay500 = `
defaults test
  default-server pool-purge-delay 500
`
const defaults_defaultserverport27015 = `
defaults test
  default-server port 27015
`
const defaults_defaultserverport27016 = `
defaults test
  default-server port 27016
`
const defaults_defaultserverprotoh2 = `
defaults test
  default-server proto h2
`
const defaults_defaultserverredirhttpimage1mydo = `
defaults test
  default-server redir http://image1.mydomain.com
`
const defaults_defaultserverredirhttpsimage1myd = `
defaults test
  default-server redir https://image1.mydomain.com
`
const defaults_defaultserverrise2 = `
defaults test
  default-server rise 2
`
const defaults_defaultserverrise200 = `
defaults test
  default-server rise 200
`
const defaults_defaultserverresolveoptsallowdup = `
defaults test
  default-server resolve-opts allow-dup-ip
`
const defaults_defaultserverresolveoptsignorewe = `
defaults test
  default-server resolve-opts ignore-weight
`
const defaults_defaultserverresolveoptsallowdup_ = `
defaults test
  default-server resolve-opts allow-dup-ip,ignore-weight
`
const defaults_defaultserverresolveoptspreventd = `
defaults test
  default-server resolve-opts prevent-dup-ip,ignore-weight
`
const defaults_defaultserverresolvepreferipv4 = `
defaults test
  default-server resolve-prefer ipv4
`
const defaults_defaultserverresolvepreferipv6 = `
defaults test
  default-server resolve-prefer ipv6
`
const defaults_defaultserverresolvenet100008 = `
defaults test
  default-server resolve-net 10.0.0.0/8
`
const defaults_defaultserverresolvenet100008100 = `
defaults test
  default-server resolve-net 10.0.0.0/8,10.0.0.0/16
`
const defaults_defaultserverresolversmydns = `
defaults test
  default-server resolvers mydns
`
const defaults_defaultserversendproxy = `
defaults test
  default-server send-proxy
`
const defaults_defaultserversendproxyv2 = `
defaults test
  default-server send-proxy-v2
`
const defaults_defaultserverproxyv2optionsssl = `
defaults test
  default-server proxy-v2-options ssl
`
const defaults_defaultserverproxyv2optionssslce = `
defaults test
  default-server proxy-v2-options ssl,cert-cn
`
const defaults_defaultserverproxyv2optionssslce_ = `
defaults test
  default-server proxy-v2-options ssl,cert-cn,ssl-cipher,cert-sig,cert-key,authority,crc32c,unique-id
`
const defaults_defaultserversendproxyv2ssl = `
defaults test
  default-server send-proxy-v2-ssl
`
const defaults_defaultserversendproxyv2sslcn = `
defaults test
  default-server send-proxy-v2-ssl-cn
`
const defaults_defaultserverslowstart2000ms = `
defaults test
  default-server slowstart 2000ms
`
const defaults_defaultserversniTODO = `
defaults test
  default-server sni TODO
`
const defaults_defaultserversourceTODO = `
defaults test
  default-server source TODO
`
const defaults_defaultserverssl = `
defaults test
  default-server ssl
`
const defaults_defaultserversslmaxverSSLv3 = `
defaults test
  default-server ssl-max-ver SSLv3
`
const defaults_defaultserversslmaxverTLSv10 = `
defaults test
  default-server ssl-max-ver TLSv1.0
`
const defaults_defaultserversslmaxverTLSv11 = `
defaults test
  default-server ssl-max-ver TLSv1.1
`
const defaults_defaultserversslmaxverTLSv12 = `
defaults test
  default-server ssl-max-ver TLSv1.2
`
const defaults_defaultserversslmaxverTLSv13 = `
defaults test
  default-server ssl-max-ver TLSv1.3
`
const defaults_defaultserversslminverSSLv3 = `
defaults test
  default-server ssl-min-ver SSLv3
`
const defaults_defaultserversslminverTLSv10 = `
defaults test
  default-server ssl-min-ver TLSv1.0
`
const defaults_defaultserversslminverTLSv11 = `
defaults test
  default-server ssl-min-ver TLSv1.1
`
const defaults_defaultserversslminverTLSv12 = `
defaults test
  default-server ssl-min-ver TLSv1.2
`
const defaults_defaultserversslminverTLSv13 = `
defaults test
  default-server ssl-min-ver TLSv1.3
`
const defaults_defaultserversslreuse = `
defaults test
  default-server ssl-reuse
`
const defaults_defaultserverstick = `
defaults test
  default-server stick
`
const defaults_defaultserversocks412700181 = `
defaults test
  default-server socks4 127.0.0.1:81
`
const defaults_defaultservertcput20ms = `
defaults test
  default-server tcp-ut 20ms
`
const defaults_defaultservertfo = `
defaults test
  default-server tfo
`
const defaults_defaultservertrackTODO = `
defaults test
  default-server track TODO
`
const defaults_defaultservertlstickets = `
defaults test
  default-server tls-tickets
`
const defaults_defaultserververifynone = `
defaults test
  default-server verify none
`
const defaults_defaultserververifyrequired = `
defaults test
  default-server verify required
`
const defaults_defaultserververifyhostsitecom = `
defaults test
  default-server verifyhost site.com
`
const defaults_defaultserverweight1 = `
defaults test
  default-server weight 1
`
const defaults_defaultserverweight128 = `
defaults test
  default-server weight 128
`
const defaults_defaultserverweight256 = `
defaults test
  default-server weight 256
`
const defaults_defaultserverpoollowconn384 = `
defaults test
  default-server pool-low-conn 384
`
const defaults_defaultserverwsh1 = `
defaults test
  default-server ws h1
`
const defaults_defaultserverwsh2 = `
defaults test
  default-server ws h2
`
const defaults_defaultserverwsauto = `
defaults test
  default-server ws auto
`
const defaults_emailalertfromadminexamplecom = `
defaults test
  email-alert from admin@example.com
`
const defaults_emailalerttoazxy = `
defaults test
  email-alert to a@z,x@y
`
const defaults_emailalertlevelwarning = `
defaults test
  email-alert level warning
`
const defaults_emailalertmailerslocalmailers = `
defaults test
  email-alert mailers local-mailers
`
const defaults_emailalertmyhostnamesrv01example = `
defaults test
  email-alert myhostname srv01.example.com
`
const defaults_emailalerttosupportexamplecom = `
defaults test
  email-alert to support@example.com
`
const defaults_emailalerttoabcd = `
defaults test
  email-alert to "a@b, c@d"
`
const defaults_errorfile400etchaproxyerrorfiles = `
defaults test
  errorfile 400 /etc/haproxy/errorfiles/400badreq.http
`
const defaults_errorfile408devnullworkaroundChr = `
defaults test
  errorfile 408 /dev/null # work around Chrome pre-connect bug
`
const defaults_errorfile403etchaproxyerrorfiles = `
defaults test
  errorfile 403 /etc/haproxy/errorfiles/403forbid.http
`
const defaults_errorfile503etchaproxyerrorfiles = `
defaults test
  errorfile 503 /etc/haproxy/errorfiles/503sorry.http
`
const defaults_errorloc302400httpwwwmyawesomesi = `
defaults test
  errorloc302 400 http://www.myawesomesite.com/error_page
`
const defaults_errorloc302404httpwwwmyawesomesi = `
defaults test
  errorloc302 404 http://www.myawesomesite.com/not_found
`
const defaults_errorloc302501errorpage = `
defaults test
  errorloc302 501 /error_page
`
const defaults_errorloc303400httpwwwmyawesomesi = `
defaults test
  errorloc303 400 http://www.myawesomesite.com/error_page
`
const defaults_errorloc303404httpwwwmyawesomesi = `
defaults test
  errorloc303 404 http://www.myawesomesite.com/not_found
`
const defaults_errorloc303501errorpage = `
defaults test
  errorloc303 501 /error_page
`
const defaults_errorfileserrorssection400 = `
defaults test
  errorfiles errors_section 400
`
const defaults_errorfileserrorssection400401500 = `
defaults test
  errorfiles errors_section 400 401 500
`
const defaults_errorfileserrorssection = `
defaults test
  errorfiles errors_section
`
const defaults_hashtypemapbased = `
defaults test
  hash-type map-based
`
const defaults_hashtypemapbasedavalanche = `
defaults test
  hash-type map-based avalanche
`
const defaults_hashtypeconsistent = `
defaults test
  hash-type consistent
`
const defaults_hashtypeconsistentavalanche = `
defaults test
  hash-type consistent avalanche
`
const defaults_hashtypeavalanche = `
defaults test
  hash-type avalanche
`
const defaults_hashtypemapbasedsdbm = `
defaults test
  hash-type map-based sdbm
`
const defaults_hashtypemapbaseddjb2 = `
defaults test
  hash-type map-based djb2
`
const defaults_hashtypemapbasedwt6 = `
defaults test
  hash-type map-based wt6
`
const defaults_hashtypemapbasedcrc32 = `
defaults test
  hash-type map-based crc32
`
const defaults_hashtypeconsistentsdbm = `
defaults test
  hash-type consistent sdbm
`
const defaults_hashtypeconsistentdjb2 = `
defaults test
  hash-type consistent djb2
`
const defaults_hashtypeconsistentwt6 = `
defaults test
  hash-type consistent wt6
`
const defaults_hashtypeconsistentcrc32 = `
defaults test
  hash-type consistent crc32
`
const defaults_hashtypemapbasedsdbmavalanche = `
defaults test
  hash-type map-based sdbm avalanche
`
const defaults_hashtypemapbaseddjb2avalanche = `
defaults test
  hash-type map-based djb2 avalanche
`
const defaults_hashtypemapbasedwt6avalanche = `
defaults test
  hash-type map-based wt6 avalanche
`
const defaults_hashtypemapbasedcrc32avalanche = `
defaults test
  hash-type map-based crc32 avalanche
`
const defaults_hashtypeconsistentsdbmavalanche = `
defaults test
  hash-type consistent sdbm avalanche
`
const defaults_hashtypeconsistentdjb2avalanche = `
defaults test
  hash-type consistent djb2 avalanche
`
const defaults_hashtypeconsistentwt6avalanche = `
defaults test
  hash-type consistent wt6 avalanche
`
const defaults_hashtypeconsistentcrc32avalanche = `
defaults test
  hash-type consistent crc32 avalanche
`
const defaults_httpreusenever = `
defaults test
  http-reuse never
`
const defaults_httpreusesafe = `
defaults test
  http-reuse safe
`
const defaults_httpreuseaggressive = `
defaults test
  http-reuse aggressive
`
const defaults_httpreusealways = `
defaults test
  http-reuse always
`
const defaults_logglobal = `
defaults test
  log global
`
const defaults_nolog = `
defaults test
  no log
`
const defaults_logstdoutformatshortdaemonsendlo = `
defaults test
  log stdout format short daemon # send log to systemd
`
const defaults_logstdoutformatrawdaemonsendever = `
defaults test
  log stdout format raw daemon # send everything to stdout
`
const defaults_logstderrformatrawdaemonnoticese = `
defaults test
  log stderr format raw daemon notice # send important events to stderr
`
const defaults_log127001514local0noticeonlysend = `
defaults test
  log 127.0.0.1:514 local0 notice # only send important events
`
const defaults_log127001514local0noticenoticesa = `
defaults test
  log 127.0.0.1:514 local0 notice notice # same but limit output level
`
const defaults_log1270011515len8192formatrfc542 = `
defaults test
  log 127.0.0.1:1515 len 8192 format rfc5424 local2 info
`
const defaults_log1270011515sample12local0 = `
defaults test
  log 127.0.0.1:1515 sample 1:2 local0
`
const defaults_log1270011515len8192formatrfc542_ = `
defaults test
  log 127.0.0.1:1515 len 8192 format rfc5424 sample 1,2-5:6 local2 info
`
const defaults_log1270011515formatrfc5424sample = `
defaults test
  log 127.0.0.1:1515 format rfc5424 sample 1,2-5:6 local2 info
`
const defaults_log1270011515formatrfc5424sample_ = `
defaults test
  log 127.0.0.1:1515 format rfc5424 sample 1-5:6 local2
`
const defaults_log1270011515sample16local2 = `
defaults test
  log 127.0.0.1:1515 sample 1:6 local2
`
const defaults_optionhttpchkOPTIONSHTTP11rnHost = `
defaults test
  option httpchk OPTIONS * HTTP/1.1\\r\\nHost:\\ www
`
const defaults_optionhttpchkuri = `
defaults test
  option httpchk <uri>
`
const defaults_optionhttpchkmethoduri = `
defaults test
  option httpchk <method> <uri>
`
const defaults_optionhttpchkmethoduriversion = `
defaults test
  option httpchk <method> <uri> <version>
`
const defaults_uniqueidformatXocicpfifpTsrtpid = `
defaults test
  unique-id-format %{+X}o_%ci:%cp_%fi:%fp_%Ts_%rt:%pid
`
const defaults_uniqueidformatXocpfifpTsrtpid = `
defaults test
  unique-id-format %{+X}o_%cp_%fi:%fp_%Ts_%rt:%pid
`
const defaults_uniqueidformatXofifpTsrtpid = `
defaults test
  unique-id-format %{+X}o_%fi:%fp_%Ts_%rt:%pid
`
const defaults_uniqueidheaderXUniqueID = `
defaults test
  unique-id-header X-Unique-ID
`
const defaults_loadserverstatefromfileglobal = `
defaults test
  load-server-state-from-file global
`
const defaults_loadserverstatefromfilelocal = `
defaults test
  load-server-state-from-file local
`
const defaults_loadserverstatefromfilenone = `
defaults test
  load-server-state-from-file none
`
const defaults_monitorurihaproxytest = `
defaults test
  monitor-uri /haproxy_test
`
const defaults_httpsendnameheader = `
defaults test
  http-send-name-header
`
const defaults_httpsendnameheaderXMyAwesomeHead = `
defaults test
  http-send-name-header X-My-Awesome-Header
`
const defaults_optionhttprestrictreqhdrnamespre = `
defaults test
  option http-restrict-req-hdr-names preserve
`
const defaults_optionhttprestrictreqhdrnamesdel = `
defaults test
  option http-restrict-req-hdr-names delete
`
const defaults_optionhttprestrictreqhdrnamesrej = `
defaults test
  option http-restrict-req-hdr-names reject
`
const defaults_source1921681200 = `
defaults test
  source 192.168.1.200
`
const defaults_source1921681200usesrcclientip = `
defaults test
  source 192.168.1.200 usesrc clientip
`
const defaults_source192168120080usesrcclientip = `
defaults test
  source 192.168.1.200:80 usesrc clientip
`
const defaults_source1921681200usesrcclient = `
defaults test
  source 192.168.1.200 usesrc client
`
const defaults_source192168120080usesrcclient = `
defaults test
  source 192.168.1.200:80 usesrc client
`
const defaults_source0000usesrcclientip = `
defaults test
  source 0.0.0.0 usesrc clientip
`
const defaults_source0000usesrchdripxforwardedf = `
defaults test
  source 0.0.0.0 usesrc hdr_ip(x-forwarded-for,-1)
`
const defaults_source1921681200interfacename = `
defaults test
  source 192.168.1.200 interface name
`
const defaults_source1921681200usesrc1921681201 = `
defaults test
  source 192.168.1.200 usesrc 192.168.1.201
`
const defaults_source1921681200usesrchdriphdr = `
defaults test
  source 192.168.1.200 usesrc hdr_ip(hdr)
`
const defaults_source1921681200usesrchdriphdroc = `
defaults test
  source 192.168.1.200 usesrc hdr_ip(hdr,occ)
`
const defaults_optionoriginalto = `
defaults test
  option originalto
`
const defaults_optionoriginaltoexcept127001 = `
defaults test
  option originalto except 127.0.0.1
`
const defaults_optionoriginaltoheaderXClientDst = `
defaults test
  option originalto header X-Client-Dst
`
const defaults_optionoriginaltoexcept127001head = `
defaults test
  option originalto except 127.0.0.1 header X-Client-Dst
`
const defaults_optionoriginaltocomment = `
defaults test
  option originalto # comment
`
const defaults_optionoriginaltoexcept127001comm = `
defaults test
  option originalto except 127.0.0.1 # comment
`
const defaults_httperrorstatus400 = `
defaults test
  http-error status 400
`
const defaults_httperrorstatus400defaulterrorfi = `
defaults test
  http-error status 400 default-errorfiles
`
const defaults_httperrorstatus400errorfilemyfan = `
defaults test
  http-error status 400 errorfile /my/fancy/errorfile
`
const defaults_httperrorstatus400errorfilesmyer = `
defaults test
  http-error status 400 errorfiles myerror
`
const defaults_httperrorstatus200contenttypetex = `
defaults test
  http-error status 200 content-type "text/plain" string "My content"
`
const defaults_httperrorstatus400contenttypetex = `
defaults test
  http-error status 400 content-type "text/plain" lf-string "Hello, you are: %[src]"
`
const defaults_httperrorstatus400contenttypetex_ = `
defaults test
  http-error status 400 content-type "text/plain" file /my/fancy/response/file
`
const defaults_httperrorstatus400contenttypetex__ = `
defaults test
  http-error status 400 content-type "text/plain" lf-file /my/fancy/lof/format/response/file
`
const defaults_httperrorstatus400contenttypetex___ = `
defaults test
  http-error status 400 content-type "text/plain" string "My content" hdr X-value value
`
const defaults_httperrorstatus400contenttypetex____ = `
defaults test
  http-error status 400 content-type "text/plain" string "My content" hdr X-value x-value hdr Y-value y-value
`
const defaults_httpcheckcommenttestcomment = `
defaults test
  http-check comment testcomment
`
const defaults_httpcheckconnect = `
defaults test
  http-check connect
`
const defaults_httpcheckconnectdefault = `
defaults test
  http-check connect default
`
const defaults_httpcheckconnectport8080 = `
defaults test
  http-check connect port 8080
`
const defaults_httpcheckconnectaddr8888 = `
defaults test
  http-check connect addr 8.8.8.8
`
const defaults_httpcheckconnectsendproxy = `
defaults test
  http-check connect send-proxy
`
const defaults_httpcheckconnectviasocks4 = `
defaults test
  http-check connect via-socks4
`
const defaults_httpcheckconnectssl = `
defaults test
  http-check connect ssl
`
const defaults_httpcheckconnectsnihaproxy1wteu = `
defaults test
  http-check connect sni haproxy.1wt.eu
`
const defaults_httpcheckconnectalpnh2http11 = `
defaults test
  http-check connect alpn h2,http/1.1
`
const defaults_httpcheckconnectprotoh2 = `
defaults test
  http-check connect proto h2
`
const defaults_httpcheckconnectlinger = `
defaults test
  http-check connect linger
`
const defaults_httpcheckconnectcommenttestcomme = `
defaults test
  http-check connect comment testcomment
`
const defaults_httpcheckconnectport443addr8888s = `
defaults test
  http-check connect port 443 addr 8.8.8.8 send-proxy via-socks4 ssl sni haproxy.1wt.eu alpn h2,http/1.1 linger proto h2 comment testcomment
`
const defaults_httpcheckdisableon404 = `
defaults test
  http-check disable-on-404
`
const defaults_httpcheckexpectstatus200 = `
defaults test
  http-check expect status 200
`
const defaults_httpcheckexpectminrecv50status20 = `
defaults test
  http-check expect min-recv 50 status 200
`
const defaults_httpcheckexpectcommenttestcommen = `
defaults test
  http-check expect comment testcomment status 200
`
const defaults_httpcheckexpectokstatusL7OKstatu = `
defaults test
  http-check expect ok-status L7OK status 200
`
const defaults_httpcheckexpecterrorstatusL7RSPs = `
defaults test
  http-check expect error-status L7RSP status 200
`
const defaults_httpcheckexpecttoutstatusL7TOUTs = `
defaults test
  http-check expect tout-status L7TOUT status 200
`
const defaults_httpcheckexpectonsuccessmylogfor = `
defaults test
  http-check expect on-success \"my-log-format\" status 200
`
const defaults_httpcheckexpectonerrormylogforma = `
defaults test
  http-check expect on-error \"my-log-format\" status 200
`
const defaults_httpcheckexpectstatuscode500stat = `
defaults test
  http-check expect status-code \"500\" status 200
`
const defaults_httpcheckexpectstringSQLError = `
defaults test
  http-check expect ! string SQL\\ Error
`
const defaults_httpcheckexpectrstatus5 = `
defaults test
  http-check expect ! rstatus ^5
`
const defaults_httpcheckexpectrstringtag09afhtm = `
defaults test
  http-check expect rstring <!--tag:[0-9a-f]*--></html>
`
const defaults_httpchecksendmethGET = `
defaults test
  http-check send meth GET
`
const defaults_httpchecksendurihealth = `
defaults test
  http-check send uri /health
`
const defaults_httpchecksendverHTTP11 = `
defaults test
  http-check send ver \"HTTP/1.1\"
`
const defaults_httpchecksendcommenttestcomment = `
defaults test
  http-check send comment testcomment
`
const defaults_httpchecksendmethGETurihealthver = `
defaults test
  http-check send meth GET uri /health ver \"HTTP/1.1\" hdr Host example.com hdr Accept-Encoding gzip body '{\"key\":\"value\"}'
`
const defaults_httpchecksendurilfmylogformatbod = `
defaults test
  http-check send uri-lf my-log-format body-lf 'my-log-format'
`
const defaults_httpchecksendstate = `
defaults test
  http-check send-state
`
const defaults_httpchecksetvarcheckportint1234 = `
defaults test
  http-check set-var(check.port) int(1234)
`
const defaults_httpchecksetvarfmtcheckportint12 = `
defaults test
  http-check set-var-fmt(check.port) int(1234)
`
const defaults_httpcheckunsetvartxnfrom = `
defaults test
  http-check unset-var(txn.from)
`
const defaults_tcpcheckcommenttestcomment = `
defaults test
  tcp-check comment testcomment
`
const defaults_tcpcheckconnect = `
defaults test
  tcp-check connect
`
const defaults_tcpcheckconnectport443ssl = `
defaults test
  tcp-check connect port 443 ssl
`
const defaults_tcpcheckconnectport110linger = `
defaults test
  tcp-check connect port 110 linger
`
const defaults_tcpcheckconnectport143 = `
defaults test
  tcp-check connect port 143
`
const defaults_tcpcheckexpectstringPONG = `
defaults test
  tcp-check expect string +PONG
`
const defaults_tcpcheckexpectstringrolemaster = `
defaults test
  tcp-check expect string role:master
`
const defaults_tcpcheckexpectstringOK = `
defaults test
  tcp-check expect string +OK
`
const defaults_tcpchecksendlftestfmt = `
defaults test
  tcp-check send-lf testfmt
`
const defaults_tcpchecksendlftestfmtcommenttest = `
defaults test
  tcp-check send-lf testfmt comment testcomment
`
const defaults_tcpchecksendbinarytesthexstring = `
defaults test
  tcp-check send-binary testhexstring
`
const defaults_tcpchecksendbinarytesthexstringc = `
defaults test
  tcp-check send-binary testhexstring comment testcomment
`
const defaults_tcpchecksendbinarylftesthexfmt = `
defaults test
  tcp-check send-binary-lf testhexfmt
`
const defaults_tcpchecksendbinarylftesthexfmtco = `
defaults test
  tcp-check send-binary-lf testhexfmt comment testcomment
`
const defaults_tcpchecksetvarcheckportint1234 = `
defaults test
  tcp-check set-var(check.port) int(1234)
`
const defaults_tcpcheckexpectstringOKPOP3ready = `
defaults test
  tcp-check expect string +OK\ POP3\ ready
`
const defaults_tcpcheckexpectstringOKIMAP4ready = `
defaults test
  tcp-check expect string *\ OK\ IMAP4\ ready
`
const defaults_tcpchecksendPINGrn = `
defaults test
  tcp-check send PING\r\n
`
const defaults_tcpchecksendPINGrncommenttestcom = `
defaults test
  tcp-check send PING\r\n comment testcomment
`
const defaults_tcpchecksendQUITrn = `
defaults test
  tcp-check send QUIT\r\n
`
const defaults_tcpchecksendQUITrncommenttestcom = `
defaults test
  tcp-check send QUIT\r\n comment testcomment
`
const defaults_tcpchecksendinforeplicationrn = `
defaults test
  tcp-check send info\ replication\r\n
`
const defaults_tcpchecksetvarfmtchecknameH = `
defaults test
  tcp-check set-var-fmt(check.name) "%H"
`
const defaults_tcpchecksetvarfmttxnfromaddrsrcs = `
defaults test
  tcp-check set-var-fmt(txn.from) "addr=%[src]:%[src_port]"
`
const defaults_tcpcheckunsetvartxnfrom = `
defaults test
  tcp-check unset-var(txn.from)
`
const defaults_statsauthadmin1AdMiN123 = `
defaults test
  stats auth admin1:AdMiN123
`
const defaults_statsenable = `
defaults test
  stats enable
`
const defaults_statshideversion = `
defaults test
  stats hide-version
`
const defaults_statsshowlegends = `
defaults test
  stats show-legends
`
const defaults_statsshowmodules = `
defaults test
  stats show-modules
`
const defaults_statsmaxconn10 = `
defaults test
  stats maxconn 10
`
const defaults_statsrealmHAProxyStatistics = `
defaults test
  stats realm HAProxy\\ Statistics
`
const defaults_statsrefresh10s = `
defaults test
  stats refresh 10s
`
const defaults_statsscope = `
defaults test
  stats scope .
`
const defaults_statsshowdescMasternodeforEurope = `
defaults test
  stats show-desc Master node for Europe, Asia, Africa
`
const defaults_statsshownode = `
defaults test
  stats show-node
`
const defaults_statsshownodeEurope1 = `
defaults test
  stats show-node Europe-1
`
const defaults_statsuriadminstats = `
defaults test
  stats uri /admin?stats
`
const defaults_statsbindprocessall = `
defaults test
  stats bind-process all
`
const defaults_statsbindprocessodd = `
defaults test
  stats bind-process odd
`
const defaults_statsbindprocesseven = `
defaults test
  stats bind-process even
`
const defaults_statsbindprocess1234 = `
defaults test
  stats bind-process 1 2 3 4
`
const defaults_statsbindprocess14 = `
defaults test
  stats bind-process 1-4
`
