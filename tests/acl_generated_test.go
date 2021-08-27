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

	"github.com/deyunluo/config-parser/v4/parsers"
)

func TestACL(t *testing.T) {
	tests := map[string]bool{
		"acl url_stats path_beg /stats":                                       true,
		"acl url_static path_beg -i /static /images /javascript /stylesheets": true,
		"acl url_static path_end -i .jpg .gif .png .css .js":                  true,
		"acl be_app_ok nbsrv(be_app) gt 0":                                    true,
		"acl be_static_ok nbsrv(be_static) gt 0":                              true,
		"acl key req.hdr(X-Add-ACL-Key) -m found":                             true,
		"acl add path /addacl":                                                true,
		"acl del path /delacl":                                                true,
		"acl myhost hdr(Host) -f myhost.lst":                                  true,
		"acl clear dst_port 80":                                               true,
		"acl secure dst_port 8080":                                            true,
		"acl login_page url_beg /login":                                       true,
		"acl logout url_beg /logout":                                          true,
		"acl uid_given url_reg /login?userid=[^&]+":                           true,
		"acl cookie_set hdr_sub(cookie) SEEN=1":                               true,
		"acl cookie":                                                          false,
		"acl":                                                                 false,
		"---":                                                                 false,
		"--- ---":                                                             false,
	}
	parser := &parsers.ACL{}
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
