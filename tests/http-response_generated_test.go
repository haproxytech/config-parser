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

	"github.com/deyunluo/config-parser/v4/parsers/http"
)

func TestResponseshttp(t *testing.T) {
	tests := map[string]bool{
		`http-response capture res.hdr(Server) id 0`:                          true,
		"http-response set-map(map.lst) %[src] %[res.hdr(X-Value)] if value":  true,
		"http-response set-map(map.lst) %[src] %[res.hdr(X-Value)]":           true,
		"http-response add-acl(map.lst) [src]":                                true,
		"http-response add-header X-value value":                              true,
		"http-response del-acl(map.lst) [src]":                                true,
		"http-response allow":                                                 true,
		"http-response del-header X-value":                                    true,
		"http-response del-map(map.lst) %[src] if ! value":                    true,
		"http-response del-map(map.lst) %[src]":                               true,
		"http-response deny":                                                  true,
		"http-response lua.foo":                                               true,
		"http-response lua.foo if FALSE":                                      true,
		"http-response lua.foo param":                                         true,
		"http-response lua.foo param param2":                                  true,
		"http-response redirect prefix https://mysite.com":                    true,
		"http-response replace-header User-agent curl foo":                    true,
		"http-response replace-value X-Forwarded-For ^192.168.(.*)$ 172.16.1": true,
		"http-response sc-inc-gpc0(1)":                                        true,
		"http-response sc-inc-gpc0(1) if FALSE":                               true,
		"http-response sc-inc-gpc1(1)":                                        true,
		"http-response sc-inc-gpc1(1) if FALSE":                               true,
		"http-response sc-set-gpt0(1) hdr(Host),lower":                        true,
		"http-response sc-set-gpt0(1) 10":                                     true,
		"http-response sc-set-gpt0(1) hdr(Host),lower if FALSE":               true,
		"http-response send-spoe-group engine group":                          true,
		"http-response set-header X-value value":                              true,
		"http-response set-log-level silent":                                  true,
		"http-response set-mark 20":                                           true,
		"http-response set-mark 0x1Ab":                                        true,
		"http-response set-nice 0":                                            true,
		"http-response set-nice 0 if FALSE":                                   true,
		"http-response set-status 503":                                        true,
		"http-response set-tos 0 if FALSE":                                    true,
		"http-response set-tos 0":                                             true,
		"http-response set-var(req.my_var) res.fhdr(user-agent),lower":        true,
		"http-response silent-drop":                                           true,
		"http-response silent-drop if FALSE":                                  true,
		"http-response unset-var(req.my_var)":                                 true,
		"http-response unset-var(req.my_var) if FALSE":                        true,
		"http-response track-sc0 src if FALSE":                                true,
		"http-response track-sc0 src table tr if FALSE":                       true,
		"http-response track-sc0 src":                                         true,
		"http-response track-sc1 src if FALSE":                                true,
		"http-response track-sc1 src table tr if FALSE":                       true,
		"http-response track-sc1 src":                                         true,
		"http-response track-sc2 src if FALSE":                                true,
		"http-response track-sc2 src table tr if FALSE":                       true,
		"http-response track-sc2 src":                                         true,
		"http-response strict-mode on":                                        true,
		"http-response strict-mode on if FALSE":                               true,
		"http-response":                                                       false,
		"http-response set-map(map.lst) %[src]":                               false,
		"http-response add-acl(map.lst)":                                      false,
		"http-response add-header X-value":                                    false,
		"http-response del-acl(map.lst)":                                      false,
		"http-response del-header":                                            false,
		"http-response del-map(map.lst)":                                      false,
		"http-response lua.":                                                  false,
		"http-response lua. if FALSE":                                         false,
		"http-response lua. param":                                            false,
		"http-response redirect prefix":                                       false,
		"http-response replace-header User-agent curl":                        false,
		"http-response replace-value X-Forwarded-For ^192.168.(.*)$":          false,
		"http-response sc-inc-gpc0":                                           false,
		"http-response sc-inc-gpc1":                                           false,
		"http-response sc-set-gpt0(1)":                                        false,
		"http-response sc-set-gpt0":                                           false,
		"http-response sc-set-gpt0(1) if FALSE":                               false,
		"http-response send-spoe-group engine":                                false,
		"http-response set-header X-value":                                    false,
		"http-response set-log-level":                                         false,
		"http-response set-mark":                                              false,
		"http-response set-nice":                                              false,
		"http-response set-status":                                            false,
		"http-response set-tos":                                               false,
		"http-response set-var(req.my_var)":                                   false,
		"http-response unset-var(req.)":                                       false,
		"http-response unset-var(req)":                                        false,
		"http-response track-sc0":                                             false,
		"http-response track-sc1":                                             false,
		"http-response track-sc2":                                             false,
		"http-response strict-mode":                                           false,
		"http-response strict-mode if FALSE":                                  false,
		"---":                                                                 false,
		"--- ---":                                                             false,
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
