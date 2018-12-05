# ![HAProxy](../assets/images/haproxy-weblogo-210x49.png "HAProxy")

## HAProxy configuration parser

### Supported options

| Name | 100% supported | Default | Global |
| - |:-:|:-:|:-:|
| [cpu-map](#nbproc) |:white_circle:|:white_circle:|:large_blue_circle:|
| [daemon](#nbproc) |:white_circle:|:white_circle:|:large_blue_circle:|
| [log](#log) |:white_circle:|:large_blue_circle:|:large_blue_circle:|
| [master-worker](#nbproc) |:white_circle:|:white_circle:|:large_blue_circle:|
| [maxconn](#maxconn) |:large_blue_circle:|:large_blue_circle:|:large_blue_circle:|
| [nbproc](#nbproc) |:white_circle:|:white_circle:|:large_blue_circle:|
| [option redispatch](#options) |:white_circle:|:large_blue_circle:|:white_circle:|
| [option dontlognull](#options) |:white_circle:|:large_blue_circle:|:white_circle:|
| [option http-server-close](#options) |:white_circle:|:large_blue_circle:|:white_circle:|
| [option http-keep-alive](#options) |:white_circle:|:large_blue_circle:|:white_circle:|
| [ssl-default-bind-options](#ssl) |:large_blue_circle:|:white_circle:|:large_blue_circle:|
| [ssl-default-bind-ciphers](#ssl) |:large_blue_circle:|:white_circle:|:large_blue_circle:|
| [stats socket](#stats) |:large_blue_circle:|:white_circle:|:large_blue_circle:|
| [stats timeout](#stats) |:large_blue_circle:|:white_circle:|:large_blue_circle:|
| [timeout http-request](#timeouts) |:large_blue_circle:|:large_blue_circle:|:white_circle:|
| [timeout connect](#timeouts) |:large_blue_circle:|:large_blue_circle:|:white_circle:|
| [timeout client](#timeouts) |:large_blue_circle:|:large_blue_circle:|:white_circle:|
| [timeout queue](#timeouts) |:large_blue_circle:|:large_blue_circle:|:white_circle:|
| [timeout server](#timeouts) |:large_blue_circle:|:large_blue_circle:|:white_circle:|
| [timeout tunnel](#timeouts) |:large_blue_circle:|:large_blue_circle:|:white_circle:|
| [timeout http-keep-alive](#timeouts) |:large_blue_circle:|:large_blue_circle:|:white_circle:|
| [tune.ssl.default-dh-param](#ssl) |:large_blue_circle:|:white_circle:|:large_blue_circle:|


### Options

#### Log

- `global`
- `<address> [len <length>] <facility> [<level> [<minlevel>]]`
  - `facility`: "kern", "user", "mail", "daemon", "auth", "syslog", "lpr", "news", 
                "uucp", "cron", "auth2", "ftp", "ntp", "audit", "alert", "cron2",
				"local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7"
  - `level` & `minlevel`: "emerg", "alert", "crit", "err", "warning", "notice", "info", "debug"

#### maxconn

- `maxconn`: number

#### nbproc

- `nbproc`: number
- `cpu-map`: number number
- `master-worker`
- `daemon`

#### Options

- `option redispatch`
- `option dontlognull`
- `option http-server-close`
- `option http-keep-alive`

#### SSL
- `tune.ssl.default-dh-param`
- `ssl-default-bind-options`
- `ssl-default-bind-ciphers`

#### Stats
- `stats socket` `address:port`|`<path>`  `[params]`
- `stats timeout`: timeout to keep socket connection open

#### Timeouts

- `timeout http-request`
- `timeout connect`
- `timeout client`
- `timeout queue`
- `timeout server`
- `timeout tunnel`
- `timeout http-keep-alive`
