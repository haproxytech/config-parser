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

package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/haproxytech/config-parser/v4/parsers/http"
)

func TestCheckshttp(t *testing.T) {
	tests := map[string]bool{
		"http-check comment testcomment":         true,
		"http-check connect":                     true,
		"http-check connect default":             true,
		"http-check connect port 8080":           true,
		"http-check connect addr 8.8.8.8":        true,
		"http-check connect send-proxy":          true,
		"http-check connect via-socks4":          true,
		"http-check connect ssl":                 true,
		"http-check connect sni haproxy.1wt.eu":  true,
		"http-check connect alpn h2,http/1.1":    true,
		"http-check connect proto h2":            true,
		"http-check connect linger":              true,
		"http-check connect comment testcomment": true,
		"http-check connect port 443 addr 8.8.8.8 send-proxy via-socks4 ssl sni haproxy.1wt.eu alpn h2,http/1.1 linger proto h2 comment testcomment": true,
		"http-check disable-on-404":                                 true,
		"http-check expect status 200":                              true,
		"http-check expect min-recv 50 status 200":                  true,
		"http-check expect comment testcomment status 200":          true,
		"http-check expect ok-status L7OK status 200":               true,
		"http-check expect error-status L7RSP status 200":           true,
		"http-check expect tout-status L7TOUT status 200":           true,
		"http-check expect on-success \"my-log-format\" status 200": true,
		"http-check expect on-error \"my-log-format\" status 200":   true,
		"http-check expect status-code \"500\" status 200":          true,
		"http-check expect ! string SQL\\ Error":                    true,
		"http-check expect ! rstatus ^5":                            true,
		"http-check expect rstring <!--tag:[0-9a-f]*--></html>":     true,
		"http-check send meth GET":                                  true,
		"http-check send uri /health":                               true,
		"http-check send ver \"HTTP/1.1\"":                          true,
		"http-check send comment testcomment":                       true,
		"http-check send meth GET uri /health ver \"HTTP/1.1\" hdr Host example.com hdr Accept-Encoding gzip body '{\"key\":\"value\"}'": true,
		"http-check send uri-lf my-log-format body-lf 'my-log-format'":                                                                   true,
		"http-check send-state":                         true,
		`http-check set-var(check.port) int(1234)`:      true,
		`http-check unset-var(txn.from)`:                true,
		"http-check":                                    false,
		"http-check comment":                            false,
		"http-check expect":                             false,
		"http-check expect status":                      false,
		"http-check expect comment testcomment":         false,
		"http-check set-var(check.port)":                false,
		"http-check set-var(check.port) int(1234) if x": false,
		"http-check unset-var(txn.from) if x":           false,
		"---":                                           false,
		"--- ---":                                       false,
	}
	parser := &http.Checks{}
	for command, shouldPass := range tests {
		t.Run(command, func(t *testing.T) {
			line := strings.TrimSpace(command)
			lines := strings.SplitN(line, "\n", -1)
			var err error
			parser.Init()
			if len(lines) > 1 {
				for _, line = range lines {
					line = strings.TrimSpace(line)
					if err = ProcessLine(line, parser); err != nil {
						break
					}
				}
			} else {
				err = ProcessLine(line, parser)
			}
			if shouldPass {
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				result, err := parser.Result()
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				var returnLine string
				if result[0].Comment == "" {
					returnLine = result[0].Data
				} else {
					returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
				}
				if command != returnLine {
					t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, command))
				}
			} else {
				if err == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
				}
				_, parseErr := parser.Result()
				if parseErr == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
				}
			}
		})
	}
}
