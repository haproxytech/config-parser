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
	"fmt"
	"strconv"
	"strings"

	"github.com/haproxytech/config-parser/v2/common"
	"github.com/haproxytech/config-parser/v2/errors"
	"github.com/haproxytech/config-parser/v2/types"
)

type HTTPCheck struct {
	data []types.HTTPCheck
}

func (h *HTTPCheck) parse(line string, parts []string, comment string) (*types.HTTPCheck, error) {
	if len(parts) < 2 {
		return nil, &errors.ParseError{Parser: "HttpCheck", Line: line, Message: "http-check type not provided"}
	}

	hcType := parts[1]

	hc := &types.HTTPCheck{
		Comment: comment,
		Type:    hcType,
	}

	if hcType == "comment" {
		if len(parts) < 3 {
			return nil, &errors.ParseError{Parser: "HttpCheck", Line: line, Message: "http-check comment: comment string not provided"}
		}
		hc.CheckComment = parts[2]
	}

	if hcType == "connect" {
		for i := 2; i < len(parts); i++ {
			el := parts[i]
			switch el {
			case "default":
				hc.Connect.Default = true
			case "port":
				if (i + 1) < len(parts) {
					hc.Connect.Port = parts[i+1]
					i++
				}
			case "addr":
				if (i + 1) < len(parts) {
					hc.Connect.Addr = parts[i+1]
					i++
				}
			case "send-proxy":
				hc.Connect.SendProxy = true
			case "via-socks4":
				hc.Connect.ViaSOCKS4 = true
			case "ssl":
				hc.Connect.SSL = true
			case "sni":
				if (i + 1) < len(parts) {
					hc.Connect.SNI = parts[i+1]
					i++
				}
			case "alpn":
				if (i + 1) < len(parts) {
					hc.Connect.ALPN = parts[i+1]
					i++
				}
			case "linger":
				hc.Connect.Linger = true
			case "proto":
				if (i + 1) < len(parts) {
					hc.Connect.Proto = parts[i+1]
					i++
				}
			case "comment":
				if (i + 1) < len(parts) {
					hc.CheckComment = parts[i+1]
					i++
				}
			}
		}
	}

	if hcType == "expect" {
		var i int
	LoopExpect:
		for i = 2; i < len(parts); i++ {
			el := parts[i]
			switch el {
			case "min-recv":
				if (i + 1) < len(parts) {
					minRecv, err := strconv.ParseInt(parts[i+1], 10, 64)
					if err != nil {
						return nil, &errors.ParseError{Parser: "HttpCheck", Line: line, Message: err.Error()}
					}
					hc.Expect.MinRecv = &minRecv
					i++
				}
			case "comment":
				if (i + 1) < len(parts) {
					hc.CheckComment = parts[i+1]
					i++
				}
			case "ok-status":
				if (i + 1) < len(parts) {
					hc.Expect.OKStatus = parts[i+1]
					i++
				}
			case "error-status":
				if (i + 1) < len(parts) {
					hc.Expect.ErrorStatus = parts[i+1]
					i++
				}
			case "tout-status":
				if (i + 1) < len(parts) {
					hc.Expect.TimeoutStatus = parts[i+1]
					i++
				}
			case "on-success":
				if (i + 1) < len(parts) {
					hc.Expect.OnSuccess = parts[i+1]
					i++
				}
			case "on-error":
				if (i + 1) < len(parts) {
					hc.Expect.OnError = parts[i+1]
					i++
				}
			case "status-code":
				if (i + 1) < len(parts) {
					hc.Expect.StatusCode = parts[i+1]
					i++
				}
			case "!":
				hc.ExclamationMark = true
				hc.Expect.ExclamationMark = true
			default:
				break LoopExpect
			}
		}

		// if we broke out of the loop, whatever is leftover should be
		// `<match> <pattern>`. Prevent panics with bounds checks for safety.
		if i >= len(parts) {
			return nil, &errors.ParseError{Parser: "HttpCheck", Line: line, Message: "http-check expect: match not provided"}
		}
		hc.Match = parts[i]
		hc.Expect.Match = parts[i]

		if i+1 >= len(parts) {
			return nil, &errors.ParseError{Parser: "HttpCheck", Line: line, Message: "http-check expect: pattern not provided"}
		}
		// Since pattern is always the last option provided, we can safely join
		// the remainder as part of the pattern.
		pattern := strings.Join(parts[i+1:], " ")
		hc.Pattern = pattern
		hc.Expect.Pattern = pattern
	}

	if hcType == "send" {
		for i := 2; i < len(parts); i++ {
			el := parts[i]
			switch el {
			case "meth":
				if (i + 1) < len(parts) {
					hc.Send.Method = parts[i+1]
					i++
				}
			case "uri":
				if (i + 1) < len(parts) {
					hc.Send.URI = parts[i+1]
					i++
				}
			case "uri-lf":
				if (i + 1) < len(parts) {
					hc.Send.URILogFormat = parts[i+1]
					i++
				}
			case "ver":
				if (i + 1) < len(parts) {
					hc.Send.Version = parts[i+1]
					i++
				}
			// NOTE: HAProxy config supports header values containing spaces,
			// which config-parser normally would support with `\ `.
			// However, because parts is split by spaces and hdr can be provided
			// multiple times with other directives surrounding it, it's
			// impossible to read ahead to put the pieces together.
			// As such, header values with spaces are not supported.
			case "hdr":
				if (i+1) < len(parts) && (i+2) < len(parts) {
					hc.Send.Header = append(hc.Send.Header, types.HTTPCheckSendHeader{Name: parts[i+1], Format: parts[i+2]})
					i++
				}
			case "body":
				if (i + 1) < len(parts) {
					hc.Send.Body = parts[i+1]
					i++
				}
			case "body-lf":
				if (i + 1) < len(parts) {
					hc.Send.BodyLogFormat = parts[i+1]
					i++
				}
			case "comment":
				if (i + 1) < len(parts) {
					hc.CheckComment = parts[i+1]
					i++
				}
			}
		}
	}

	return hc, nil
}

func (h *HTTPCheck) Result() ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.ErrFetch
	}

	result := make([]common.ReturnResultLine, len(h.data))
	for index, c := range h.data {
		var sb strings.Builder
		sb.WriteString("http-check")
		if c.Type != "" {
			sb.WriteString(" ")
			sb.WriteString(c.Type)
		}

		if c.Type == "comment" {
			if c.CheckComment != "" {
				sb.WriteString(" ")
				sb.WriteString(c.CheckComment)
			} else {
				return nil, fmt.Errorf("http-check comment: comment string required but not provided")
			}
		}

		if c.Type == "connect" {
			if c.Connect.Default {
				sb.WriteString(" default")
			}
			if c.Connect.Port != "" {
				sb.WriteString(" port ")
				sb.WriteString(c.Connect.Port)
			}
			if c.Connect.Addr != "" {
				sb.WriteString(" addr ")
				sb.WriteString(c.Connect.Addr)
			}
			if c.Connect.SendProxy {
				sb.WriteString(" send-proxy")
			}
			if c.Connect.ViaSOCKS4 {
				sb.WriteString(" via-socks4")
			}
			if c.Connect.SSL {
				sb.WriteString(" ssl")
			}
			if c.Connect.SNI != "" {
				sb.WriteString(" sni ")
				sb.WriteString(c.Connect.SNI)
			}
			if c.Connect.ALPN != "" {
				sb.WriteString(" alpn ")
				sb.WriteString(c.Connect.ALPN)
			}
			if c.Connect.Linger {
				sb.WriteString(" linger")
			}
			if c.Connect.Proto != "" {
				sb.WriteString(" proto ")
				sb.WriteString(c.Connect.Proto)
			}
			if c.CheckComment != "" {
				sb.WriteString(" comment ")
				sb.WriteString(c.CheckComment)
			}
		}

		if c.Type == "expect" {
			if c.Expect.MinRecv != nil {
				sb.WriteString(" min-recv ")
				sb.WriteString(strconv.Itoa(int(*c.Expect.MinRecv)))
			}
			if c.CheckComment != "" {
				sb.WriteString(" comment ")
				sb.WriteString(c.CheckComment)
			}
			if c.Expect.OKStatus != "" {
				sb.WriteString(" ok-status ")
				sb.WriteString(c.Expect.OKStatus)
			}
			if c.Expect.ErrorStatus != "" {
				sb.WriteString(" error-status ")
				sb.WriteString(c.Expect.ErrorStatus)
			}
			if c.Expect.TimeoutStatus != "" {
				sb.WriteString(" tout-status ")
				sb.WriteString(c.Expect.TimeoutStatus)
			}
			if c.Expect.OnSuccess != "" {
				sb.WriteString(" on-success ")
				sb.WriteString(c.Expect.OnSuccess)
			}
			if c.Expect.OnError != "" {
				sb.WriteString(" on-error ")
				sb.WriteString(c.Expect.OnError)
			}
			if c.Expect.StatusCode != "" {
				sb.WriteString(" status-code ")
				sb.WriteString(c.Expect.StatusCode)
			}
			if c.ExclamationMark {
				c.Expect.ExclamationMark = c.ExclamationMark
			}
			if c.Expect.ExclamationMark {
				sb.WriteString(" !")
			}
			if c.Match != "" {
				c.Expect.Match = c.Match
			}
			if c.Expect.Match != "" {
				sb.WriteString(" ")
				sb.WriteString(c.Expect.Match)
			}
			if c.Pattern != "" {
				c.Expect.Pattern = c.Pattern
			}
			if c.Expect.Pattern != "" {
				sb.WriteString(" ")
				sb.WriteString(c.Expect.Pattern)
			}
		}

		if c.Type == "send" {
			if c.Send.Method != "" {
				sb.WriteString(" meth ")
				sb.WriteString(c.Send.Method)
			}
			if c.Send.URI != "" {
				sb.WriteString(" uri ")
				sb.WriteString(c.Send.URI)
			}
			if c.Send.URILogFormat != "" {
				sb.WriteString(" uri-lf ")
				sb.WriteString(c.Send.URILogFormat)
			}
			if c.Send.Version != "" {
				sb.WriteString(" ver ")
				sb.WriteString(c.Send.Version)
			}
			for _, h := range c.Send.Header {
				sb.WriteString(" hdr ")
				sb.WriteString(h.Name)
				sb.WriteString(" ")
				sb.WriteString(h.Format)
			}
			if c.Send.Body != "" {
				sb.WriteString(" body ")
				sb.WriteString(c.Send.Body)
			}
			if c.Send.BodyLogFormat != "" {
				sb.WriteString(" body-lf ")
				sb.WriteString(c.Send.BodyLogFormat)
			}
			if c.CheckComment != "" {
				sb.WriteString(" comment ")
				sb.WriteString(c.CheckComment)
			}
		}

		result[index] = common.ReturnResultLine{
			Data:    sb.String(),
			Comment: c.Comment,
		}
	}

	return result, nil
}
