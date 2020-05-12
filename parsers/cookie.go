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

package parsers

import (
	"strconv"
	"strings"

	"github.com/haproxytech/config-parser/v2/common"
	"github.com/haproxytech/config-parser/v2/errors"
	"github.com/haproxytech/config-parser/v2/types"
)

type Cookie struct {
	data *types.Cookie
}

func (p *Cookie) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "cookie" {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "Cookie", Line: line, Message: "Parse error"}
		}
		data := &types.Cookie{
			Domain:  []string{},
			Name:    parts[1],
			Comment: comment,
		}

		for i := 2; i < len(parts); i++ {
			el := parts[i]
			switch el {
			case "dynamic":
				data.Dynamic = true
			case "httponly":
				data.Httponly = true
			case "indirect":
				data.Indirect = true
			case "nocache":
				data.Nocache = true
			case "postonly":
				data.Postonly = true
			case "preserve":
				data.Preserve = true
			case "secure":
				data.Secure = true
			case "domain":
				if (i + 1) < len(parts) {
					i++
					data.Domain = append(data.Domain, parts[i])
				}
			case "maxidle":
				if (i + 1) < len(parts) {
					i++
					if data.Maxidle, err = strconv.ParseInt(parts[i], 10, 64); err != nil {
						return "", &errors.ParseError{Parser: "maxidle", Line: line, Message: err.Error()}
					}
				}
			case "maxlife":
				if (i + 1) < len(parts) {
					i++
					if data.Maxlife, err = strconv.ParseInt(parts[i], 10, 64); err != nil {
						return "", &errors.ParseError{Parser: "maxlife", Line: line, Message: err.Error()}
					}
				}
			default:
				if el == "rewrite" || el == "insert" || el == "prefix" {
					data.Type = el
				}
			}
		}

		p.data = data
		return "", nil
	}
	return "", &errors.ParseError{Parser: "Cookie", Line: line}
}

func (p *Cookie) Result() ([]common.ReturnResultLine, error) {
	if p.data == nil {
		return nil, errors.ErrFetch
	}

	var result strings.Builder
	result.WriteString("cookie")

	if p.data.Name != "" {
		result.WriteString(" ")
		result.WriteString(p.data.Name)
	}

	if len(p.data.Domain) > 0 {
		for _, domain := range p.data.Domain {
			result.WriteString(" domain ")
			result.WriteString(domain)
		}
	}
	if p.data.Dynamic {
		result.WriteString(" dynamic")
	}
	if p.data.Httponly {
		result.WriteString(" httponly")
	}
	if p.data.Indirect {
		result.WriteString(" indirect")
	}
	if p.data.Maxidle > 0 {
		result.WriteString(" maxidle ")
		result.WriteString(strconv.Itoa(int(p.data.Maxidle)))
	}
	if p.data.Maxlife > 0 {
		result.WriteString(" maxlife ")
		result.WriteString(strconv.Itoa(int(p.data.Maxlife)))
	}
	if p.data.Nocache {
		result.WriteString(" nocache")
	}
	if p.data.Postonly {
		result.WriteString(" postonly")
	}
	if p.data.Preserve {
		result.WriteString(" preserve")
	}
	if p.data.Type != "" {
		result.WriteString(" ")
		result.WriteString(p.data.Type)
	}
	if p.data.Secure {
		result.WriteString(" secure")
	}

	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    result.String(),
			Comment: p.data.Comment,
		},
	}, nil
}
