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

	"github.com/haproxytech/config-parser/v5/parsers/http"
)

func TestResponseshttp(t *testing.T) {
	tests := map[string]bool{
		`http-response capture res.hdr(Server) id 0`:                                                                                true,
		"http-response set-map(map.lst) %[src] %[res.hdr(X-Value)] if value":                                                        true,
		"http-response set-map(map.lst) %[src] %[res.hdr(X-Value)]":                                                                 true,
		"http-response add-acl(map.lst) [src]":                                                                                      true,
		"http-response add-header X-value value":                                                                                    true,
		"http-response del-acl(map.lst) [src]":                                                                                      true,
		"http-response allow":                                                                                                       true,
		"http-response cache-store cache-name":                                                                                      true,
		"http-response cache-store cache-name if FALSE":                                                                             true,
		"http-response del-header X-value":                                                                                          true,
		"http-response del-map(map.lst) %[src] if ! value":                                                                          true,
		"http-response del-map(map.lst) %[src]":                                                                                     true,
		"http-response deny":                                                                                                        true,
		"http-response deny deny_status 400":                                                                                        true,
		"http-response deny if TRUE":                                                                                                true,
		"http-response deny deny_status 400 if TRUE":                                                                                true,
		"http-response deny deny_status 400 content-type application/json if TRUE":                                                  true,
		"http-response deny deny_status 400 content-type application/json":                                                          true,
		"http-response deny deny_status 400 content-type application/json default-errorfiles":                                       true,
		"http-response deny deny_status 400 content-type application/json errorfile errors":                                         true,
		"http-response deny deny_status 400 content-type application/json string error if TRUE":                                     true,
		"http-response deny deny_status 400 content-type application/json lf-string error hdr host google.com if TRUE":              true,
		"http-response deny deny_status 400 content-type application/json file /var/errors.file":                                    true,
		"http-response deny deny_status 400 content-type application/json lf-file /var/errors.file":                                 true,
		"http-response deny deny_status 400 content-type application/json string error hdr host google.com if TRUE":                 true,
		"http-response deny deny_status 400 content-type application/json string error hdr host google.com hdr x-value bla if TRUE": true,
		"http-response deny deny_status 400 content-type application/json string error hdr host google.com hdr x-value bla":         true,
		"http-response lua.foo":                                                                                                     true,
		"http-response lua.foo if FALSE":                                                                                            true,
		"http-response lua.foo param":                                                                                               true,
		"http-response lua.foo param param2":                                                                                        true,
		"http-response redirect prefix https://mysite.com":                                                                          true,
		"http-response replace-header User-agent curl foo":                                                                          true,
		"http-response replace-value X-Forwarded-For ^192.168.(.*)$ 172.16.1":                                                       true,
		"http-response return status 400 default-errorfiles if { var(txn.myip) -m found }":                                          true,
		"http-response return status 400 errorfile /my/fancy/errorfile if { var(txn.myip) -m found }":                               true,
		"http-response return status 400 errorfiles myerror if { var(txn.myip) -m found }":                                          true,
		"http-response sc-add-gpc(1,2) 1":                                                                                           true,
		"http-response sc-add-gpc(1,2) 1 if is-error":                                                                               true,
		"http-response sc-inc-gpc(1,2)":                                                                                             true,
		"http-response sc-inc-gpc(1,2) if FALSE":                                                                                    true,
		"http-response sc-inc-gpc0(1)":                                                                                              true,
		"http-response sc-inc-gpc0(1) if FALSE":                                                                                     true,
		"http-response sc-inc-gpc1(1)":                                                                                              true,
		"http-response sc-inc-gpc1(1) if FALSE":                                                                                     true,
		"http-response sc-set-gpt0(1) hdr(Host),lower":                                                                              true,
		"http-response sc-set-gpt0(1) 10":                                                                                           true,
		"http-response sc-set-gpt0(1) hdr(Host),lower if FALSE":                                                                     true,
		"http-response send-spoe-group engine group":                                                                                true,
		"http-response set-header X-value value":                                                                                    true,
		"http-response set-log-level silent":                                                                                        true,
		"http-response set-mark 20":                                                                                                 true,
		"http-response set-mark 0x1Ab":                                                                                              true,
		"http-response set-nice 0":                                                                                                  true,
		"http-response set-nice 0 if FALSE":                                                                                         true,
		"http-response set-status 503":                                                                                              true,
		"http-response set-timeout server 20":                                                                                       true,
		"http-response set-timeout tunnel 20":                                                                                       true,
		"http-response set-timeout tunnel 20s if TRUE":                                                                              true,
		"http-response set-timeout server 20s if TRUE":                                                                              true,
		"http-response set-timeout client 20":                                                                                       true,
		"http-response set-timeout client 20s if TRUE":                                                                              true,
		"http-response set-tos 0 if FALSE":                                                                                          true,
		"http-response set-tos 0":                                                                                                   true,
		"http-response set-var(req.my_var) res.fhdr(user-agent),lower":                                                              true,
		"http-response set-var-fmt(req.my_var) res.fhdr(user-agent),lower":                                                          true,
		"http-response silent-drop":                                                                                                 true,
		"http-response silent-drop if FALSE":                                                                                        true,
		"http-response unset-var(req.my_var)":                                                                                       true,
		"http-response unset-var(req.my_var) if FALSE":                                                                              true,
		"http-response track-sc0 src if FALSE":                                                                                      true,
		"http-response track-sc0 src table tr if FALSE":                                                                             true,
		"http-response track-sc0 src":                                                                                               true,
		"http-response track-sc1 src if FALSE":                                                                                      true,
		"http-response track-sc1 src table tr if FALSE":                                                                             true,
		"http-response track-sc1 src":                                                                                               true,
		"http-response track-sc2 src if FALSE":                                                                                      true,
		"http-response track-sc2 src table tr if FALSE":                                                                             true,
		"http-response track-sc2 src":                                                                                               true,
		"http-response track-sc5 src":                                                                                               true,
		"http-response track-sc5 src table a_table":                                                                                 true,
		"http-response track-sc5 src table a_table if some_cond":                                                                    true,
		"http-response track-sc5 src if some_cond":                                                                                  true,
		"http-response strict-mode on":                                                                                              true,
		"http-response strict-mode on if FALSE":                                                                                     true,
		"http-response wait-for-body time 20s":                                                                                      true,
		"http-response wait-for-body time 20s if TRUE":                                                                              true,
		"http-response wait-for-body time 20s at-least 100k":                                                                        true,
		"http-response wait-for-body time 20s at-least 100k if TRUE":                                                                true,
		"http-response set-bandwidth-limit my-limit":                                                                                true,
		"http-response set-bandwidth-limit my-limit limit 1m period 10s":                                                            true,
		"http-response set-bandwidth-limit my-limit period 10s":                                                                     true,
		"http-response set-bandwidth-limit my-limit limit 1m":                                                                       true,
		"http-response set-fc-mark 2000":                                                                                            true,
		"http-response set-fc-tos 200":                                                                                              true,
		`http-response return status 200 content-type "text/plain" string "My content" if { var(txn.myip) -m found }`:               true,
		`http-response return status 200 content-type "text/plain" string "My content" unless { var(txn.myip) -m found }`:                          true,
		`http-response return content-type "text/plain" string "My content" if { var(txn.myip) -m found }`:                                         true,
		`http-response return content-type 'text/plain' string 'My content' if { var(txn.myip) -m found }`:                                         true,
		`http-response return content-type "text/plain" lf-string "Hello, you are: %[src]" if { var(txn.myip) -m found }`:                          true,
		`http-response return content-type "text/plain" file /my/fancy/response/file if { var(txn.myip) -m found }`:                                true,
		`http-response return content-type "text/plain" lf-file /my/fancy/lof/format/response/file if { var(txn.myip) -m found }`:                  true,
		`http-response return content-type "text/plain" string "My content" hdr X-value value if { var(txn.myip) -m found }`:                       true,
		`http-response return content-type "text/plain" string "My content" hdr X-value x-value hdr Y-value y-value if { var(txn.myip) -m found }`: true,
		`http-response return content-type "text/plain" lf-string "Hello, you are: %[src]"`:                                                        true,
		"http-response":                                                            false,
		"http-response set-map(map.lst) %[src]":                                    false,
		"http-response add-acl(map.lst)":                                           false,
		"http-response add-header X-value":                                         false,
		"http-response del-acl(map.lst)":                                           false,
		"http-response cache-store":                                                false,
		"http-response cache-store if FALSE":                                       false,
		"http-response del-header":                                                 false,
		"http-response del-map(map.lst)":                                           false,
		"http-response deny test test":                                             false,
		"http-response lua.":                                                       false,
		"http-response lua. if FALSE":                                              false,
		"http-response lua. param":                                                 false,
		"http-response redirect prefix":                                            false,
		"http-response replace-header User-agent curl":                             false,
		"http-response replace-value X-Forwarded-For ^192.168.(.*)$":               false,
		"http-response return 8 t hdr":                                             false,
		"http-response return hdr":                                                 false,
		"http-response return hdr one":                                             false,
		"http-response return errorfile":                                           false,
		"http-response return 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 file":              false,
		"http-response return 0 hdr":                                               false,
		"http-response return 0 0 hdr 0":                                           false,
		"http-response return e r s n s c m	t e r  s c t e s t e r s c v e hdr ï": false,
		"http-response sc-add-gpc":                                                 false,
		"http-response sc-inc-gpc":                                                 false,
		"http-response sc-inc-gpc0":                                                false,
		"http-response sc-inc-gpc1":                                                false,
		"http-response sc-set-gpt0(1)":                                             false,
		"http-response sc-set-gpt0":                                                false,
		"http-response sc-set-gpt0(1) if FALSE":                                    false,
		"http-response send-spoe-group engine":                                     false,
		"http-response set-header X-value":                                         false,
		"http-response set-log-level":                                              false,
		"http-response set-mark":                                                   false,
		"http-response set-nice":                                                   false,
		"http-response set-status":                                                 false,
		"http-response set-timeout fake-timeout 20s if TRUE":                       false,
		"http-response set-tos":                                                    false,
		"http-response set-var(req.my_var)":                                        false,
		"http-response set-var-fmt(req.my_var)":                                    false,
		"http-response unset-var(req.)":                                            false,
		"http-response unset-var(req)":                                             false,
		"http-response track-sc0":                                                  false,
		"http-response track-sc1":                                                  false,
		"http-response track-sc2":                                                  false,
		"http-response track-sc":                                                   false,
		"http-response track-sc5":                                                  false,
		"http-response track-sc5 src table":                                        false,
		"http-response track-sc5 src if":                                           false,
		"http-response strict-mode":                                                false,
		"http-response strict-mode if FALSE":                                       false,
		"http-response wait-for-body 20s at-least 100k":                            false,
		"http-response wait-for-body time 2000 test":                               false,
		"http-response set-bandwidth-limit my-limit limit":                         false,
		"http-response set-bandwidth-limit my-limit period":                        false,
		"http-response set-bandwidth-limit my-limit 10s":                           false,
		"http-response set-bandwidth-limit my-limit period 10s limit":              false,
		"http-response set-bandwidth-limit my-limit limit period 10s":              false,
		"---":     false,
		"--- ---": false,
	}
	parser := &http.Responses{}
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
