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
package configs

const configFull = `# _version=10
# HAProxy Technologies
# https://www.haproxy.com/

# some random global comment
global
  daemon
  master-worker
  nbproc 5
  nbthread 6
  cpu-map 1 3
  cpu-map 2 1
  cpu-map 3 2
  maxconn 5000
  pidfile /var/run/haproxy.pid
  stats socket /var/run/haproxy-runtime-api.1.sock level admin mode 777 expose-fd listeners process 1
  stats socket /var/run/haproxy-runtime-api.2.sock level admin mode 777 expose-fd listeners process 2
  stats socket $PWD/haproxy-runtime-api.3.sock level admin mode 777 expose-fd listeners process 3
  stats timeout 120s
  tune.ssl.default-dh-param 2048
  ssl-default-bind-options no-sslv3 no-tls-tickets
  ssl-default-bind-ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-DSS-AES128-GCM-SHA256:kEDH+AESGCM:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-DSS-AES128-SHA256:DHE-RSA-AES256-SHA256:DHE-DSS-AES256-SHA:DHE-RSA-AES256-SHA:!aNULL:!eNULL:!EXPORT:!DES:!RC4:!3DES:!MD5:!PSK
  log 127.0.0.1:514 local0 notice
  # random comment before snippet
  ###_config-snippet_### BEGIN
  tune.ssl.default-dh-param 2048
  tune.bufsize 32768
  ###_config-snippet_### END
  # random comment after snippet

# some random defaults comment
defaults
  maxconn 2000
  log global
  log 127.0.0.1:514 local0 notice
  log-format '%ci:%cp [%tr] %ft %b/%s %TR/%Tw/%Tc/%Tr/%Ta %ST %B %CC %CS %tsc %ac/%fc/%bc/%sc/%rc %sq/%bq %hr %hs "%HM %[var(txn.base)] %HV"'
  option redispatch
  option dontlognull
  option http-server-close
  option http-keep-alive
  timeout http-request 5s
  timeout check 15s
  timeout connect 5s
  timeout client 50s
  timeout queue 5s
  timeout server 50s
  timeout tunnel 1h
  timeout http-keep-alive 1m
  load-server-state-from-file global

# some random userlist L1
userlist L1
  group G1 users tiger,scott
  group G2 users xdb,scott
  user nopwd groups G2
  user tiger password $6$k6y3o.eP$JlKBx9za9667qe4(...)xHSwRv6J.C0/D7cV91
  user scott insecure-password elgato
  user xdb insecure-password hello

# some random userlist L2
userlist L2
  group G1
  group G2
  # some random user comment #1
  # some random user comment #2
  user tiger password $6$k6y3o.eP$JlKBx(...)xHSwRv6J.C0/D7cV91 groups G1
  user scott insecure-password elgato groups G1,G2
  user xdb insecure-password hello groups G2

peers mypeers
  peer haproxy1 192.168.0.1:1024
  peer haproxy2 192.168.0.2:1024
  peer haproxy3 10.2.0.1:1024

peers mypeers2
  peer haproxy4 192.168.0.1:1024
  peer haproxy5 192.168.0.2:1024
  peer haproxy6 10.2.0.1:1024

mailers mymailers
  timeout mail 20s
  mailer smtp1 192.168.0.1:587
  mailer smtp2 192.168.0.2:587

mailers mymailers2
  timeout mail 10s
  mailer smtp1 192.168.0.3:587
  mailer smtp2 192.168.0.4:587

resolvers mirko
  nameserver ns_1_new_name_for_0 0.0.0.0:8080
  nameserver ns_2_new_name 0.0.0.0:8081
  hold obsolete 5s
  hold valid 5s
  timeout retry 5s
  accepted_payload_size 323

cache foobar
  total-max-size 4
  max-age 240

frontend healthz
  mode http
  monitor-uri /healthz
  no log

frontend http
  mode http
  bind 0.0.0.0:80 name bind_1
  bind :::80 v4v6 name bind_2
  default_backend default_backend

frontend https
  mode http
  bind 0.0.0.0:443 name bind_1
  bind :::443 v4v6 name bind_2
  http-request set-var(txn.Base) base
  http-request set-header X-Forwarded-Proto https if { ssl_fc }
  default_backend default_backend

frontend xyz
  mode http
  acl network_allowed src 20.30.40.50 8.9.9.0/27
  acl ratelimit_is_abuse src_http_req_rate(Abuse) ge 10
  acl ratelimit_inc_cnt_abuse src_inc_gpc0(Abuse) gt 0
  acl ratelimit_cnt_abuse src_get_gpc0(Abuse) gt 0
  option forwardfor
  http-request deny if !network_allowed
  default_backend default_backend

frontend xyz2
  mode http
  maxconn 2000
  bind 0.0.0.0:9981 name http2
  acl key req.hdr(X-Add-ACL-Key) -m found
  acl add path /addacl
  acl del path /delacl
  acl myhost hdr(Host) -f myhost.lst
  http-request add-acl(myhost.lst) %[req.hdr(X-Add-ACL-Key)] if key add # add-acl
  http-request del-acl(myhost.lst) %[req.hdr(X-Add-ACL-Key)] if key del
  http-request add-header Via 1.1\ %[env(HOSTNAME)] # AH1
  http-request add-header Via2 1.1\ %[env(HOSTNAME)] if key # AH2
  http-request del-header Via2 if something else # AH3
  http-request allow if something really cool # allow1
  http-request redirect scheme https code 302 if !{ ssl_fc }
  http-request set-log-level level if !{ ssl_fc }
  http-request set-log-level level2
  use_backend default-ingress-default-backend-8080 if { req.hdr(host) -i foo.bar } { path_beg / } # alctl: farm switching rule #deny #deny #deny
  use_backend default-http-svc-8080 if { req.hdr(host) -i foo.bar } { path_beg /app } # alctl: farm switching rule #deny #deny #deny
  default_backend default_backend2
  http-response set-status some_status if !{ ssl_fc }

frontend xyz3
  acl network_allowed src 20.30.40.50 8.9.9.0/27
  acl network_allowed src 20.30.40.51 8.9.9.0/27
  acl other acl src 20.30.40.52 8.9.9.0/27
  http-request allow if src 192.168.0.0/16
  http-request set-header X-SSL %[ssl_fc]
  http-request set-var(req.my_var) req.fhdr(user-agent),lower
  http-request set-header X-SSL2 %[ssl_fc] if something
  http-response allow if src 192.168.0.1/16
  http-response set-header X-SSL1 %[ssl_fc]
  http-response set-var(req.my_var1) req.fhdr(user-agent),lower

frontend xyz4
  mode http
  http-request allow if something
  http-request allow

frontend xyz5
  mode http
  maxconn 2000
  bind 192.168.1.1:80 name webserv
  bind 192.168.1.1:8080 name webserv2
  log-tag bla
  log global
  option httpclose
  option dontlognull
  option contstats
  option log-separate-errors
  option clitcpka
  option httplog
  timeout http-request 2s
  timeout client 4s
  timeout http-keep-alive 3s
  filter trace name BEFORE-HTTP-COMP random-parsing hexdump
  filter compression
  filter trace name AFTER-HTTP-COMP random-forwarding
  tcp-request connection accept if TRUE
  tcp-request connection reject if FALSE
  tcp-request content accept if TRUE
  tcp-request content reject if FALSE
  tcp-request inspect-delay 30s
  http-request allow if src 192.168.0.0/16
  http-request set-header X-SSL %[ssl_fc]
  http-request set-var(req.my_var) req.fhdr(user-agent),lower
  http-request auth unless auth_ok
  use_backend test_2 if TRUE
  default_backend test_2
  http-response allow if src 192.168.0.0/16
  http-response set-header X-SSL %[ssl_fc]
  http-response set-var(req.my_var) req.fhdr(user-agent),lower

backend default_backend
  mode http
  http-request deny deny_status 400 # deny

backend default_backend2
  mode http
  balance uri
  option httpchk OPTIONS * HTTP/1.1\r\nHost:\ www
  server SRV_PYkL1 127.0.0.1:5851 maxconn 1000 weight 1 check # alctl: server SRV_PYkL1 configuration.
  server SRV_VqMNT 127.0.0.1:5852 maxconn 1000 weight 1 check # alctl: server SRV_VqMNT configuration.
  server SRV_LkIZ9 127.0.0.1:5853 maxconn 1000 weight 1 check # alctl: server SRV_LkIZ9 configuration.
  server THE_NEW_GUY 127.0.0.5:9345 # Newly added
  # server SRV_LkIZw 127.0.0.1:5853 maxconn 1000 weight 1 check disabled #alctl: server SRV_LkIZ9 configuration.

backend test
  mode http
  balance roundrobin
  option http-keep-alive
  option forwardfor header X-Forwarded-For
  log-tag bla
  option httpchk HEAD /
  log global
  timeout check 2s
  timeout tunnel 5s
  timeout server 3s
  default-server fall 2
  default-server rise 4
  default-server inter 5s
  default-server port 8888
  stick store-request src table test
  stick match src table test
  stick on src table test
  stick store-response src
  stick store-response src_port table test_port
  stick store-response src table test if TRUE
  cookie BLA
  use-server webserv if TRUE
  use-server webserv2 unless TRUE
  stick-table type string len 1000 size 1m expire 5m store gpc0,conn_rate(30s)
  server webserv 192.168.1.1:9200 maxconn 1000 ssl weight 10 cookie BLAH
  server webserv2 192.168.1.1:9300 maxconn 1000 ssl weight 10 cookie BLAH
  tcp-response content accept if TRUE
  tcp-response content reject if FALSE
  option httplog
  option contstats
  option contstats

listen stats
  bind *:1024 process 1
  mode http
  stats enable
  stats realm HAProxy\ Statistics
  stats uri /
  no log
  option forceclose
`
