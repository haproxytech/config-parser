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

//nolint:dupl
package actions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haproxytech/config-parser/v3/common"
)

type Return struct { // http-request return [status <code>] [content-type <type>] [ {default-errorfile | <content-format> <content>} ] [ hdr <name> <fmt> ]* [{if | unless} <condition>]
	Comment       string
	Status        *int64
	ContentType   string
	ContentFormat string
	Content       string
	Hdrs          []*Hdr
	Cond          string
	CondTest      string
}

type Hdr struct {
	Name string
	Fmt  string
}

func IsPayload(in string) bool {
	_, ok := map[string]struct{}{
		"string":    {},
		"lf-string": {},
		"file":      {},
		"lf-file":   {},
	}[in]
	return ok
}

func AllowedErrorCode(code int64) bool {
	_, ok := map[int64]struct{}{
		200: {},
		400: {},
		403: {},
		404: {},
		405: {},
		408: {},
		410: {},
		413: {},
		425: {},
		429: {},
		500: {},
		503: {},
		504: {},
	}[code]
	return ok
}

func allowedStatusCode(code int64) bool {
	return code >= 200 && code <= 509
}

func (f *Return) Parse(parts []string, comment string) error {
	f.Hdrs = []*Hdr{}
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 4 {
		command, condition := common.SplitRequest(parts[2:])
		if len(command) > 1 {
			for i := 0; i < len(command); i++ {
				switch command[i] {
				case "status":
					i++
					code, err := strconv.ParseInt(command[i], 10, 64)
					if err != nil {
						return fmt.Errorf("failed to parse status code")
					}
					f.Status = &code
				case "content-type":
					i++
					contentType, err := unQuote(command[i])
					if err != nil {
						return fmt.Errorf("failed to unquote content-type %s", command[i])
					}
					f.ContentType = contentType
				case "errorfile", "errorfiles", "file", "lf-file", "string", "lf-string":
					f.ContentFormat = command[i]
					i++
					content, err := unQuote(command[i])
					if err != nil {
						return fmt.Errorf("failed to unquote content %s", command[i])
					}
					f.Content = content
				case "default-errorfiles":
					f.ContentFormat = command[i]
				case "hdr":
					hdr := Hdr{}
					i++
					hdr.Name = command[i]
					i++
					hdr.Fmt = command[i]
					f.Hdrs = append(f.Hdrs, &hdr)
				}
			}
			if len(condition) > 1 {
				f.Cond = condition[0]
				f.CondTest = strings.Join(condition[1:], " ")
			}

			return nil
		}
	} else if len(parts) == 2 {
		return nil
	}
	return fmt.Errorf("not enough params")
}

func unQuote(in string) (string, error) {
	if len(in) < 2 {
		return in, nil
	}
	if in[0] != '\'' && in[0] != '`' && in[0] != '"' {
		return in, nil
	}
	return strconv.Unquote(in)
}

func (f *Return) String() string {
	var result strings.Builder
	result.WriteString("return")
	if f.Status != nil {
		if IsPayload(f.ContentFormat) {
			if allowedStatusCode(*f.Status) {
				result.WriteString(" status ")
				result.WriteString(strconv.FormatInt(*f.Status, 10))
			}
		} else {
			if AllowedErrorCode(*f.Status) {
				result.WriteString(" status ")
				result.WriteString(strconv.FormatInt(*f.Status, 10))
			}
		}
	}
	if f.ContentType != "" {
		result.WriteString(" content-type ")
		result.WriteString(strconv.Quote(f.ContentType))
	}
	if f.ContentFormat != "" {
		result.WriteString(" ")
		result.WriteString(f.ContentFormat)
		if f.Content != "" && f.ContentFormat != "default-errorfiles" {
			result.WriteString(" ")
			if f.ContentFormat == "string" || f.ContentFormat == "lf-string" {
				result.WriteString(strconv.Quote(f.Content))
			} else {
				result.WriteString(f.Content)
			}
		}
	}
	if IsPayload(f.ContentFormat) {
		for _, hdr := range f.Hdrs {
			result.WriteString(" hdr ")
			result.WriteString(hdr.Name)
			result.WriteString(" ")
			result.WriteString(hdr.Fmt)
		}
	}
	if f.Cond != "" {
		result.WriteString(" ")
		result.WriteString(f.Cond)
		result.WriteString(" ")
		result.WriteString(f.CondTest)
	}
	return result.String()
}

func (f *Return) GetComment() string {
	return f.Comment
}
