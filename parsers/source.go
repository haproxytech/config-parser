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

	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/types"
)

type Source struct {
	data        []types.Source
	preComments []string // comments that appear before the the actual line
}

func (s *Source) parse(line string, parts []string, comment string) (*types.Source, error) { //nolint:gocognit
	if len(parts) < 2 {
		return nil, &errors.ParseError{Parser: "Source", Line: line}
	}
	if parts[0] != "source" {
		return nil, &errors.ParseError{Parser: "Source", Line: line}
	}

	data := &types.Source{
		Comment: comment,
	}

	if strings.Contains(parts[1], ":") {
		addressAndPort := strings.Split(parts[1], ":")
		data.Address = addressAndPort[0]
		if port, err := strconv.ParseInt(addressAndPort[1], 10, 64); err == nil {
			data.Port = port
		}
	} else {
		data.Address = parts[1]
	}

	for i := 2; i < len(parts); i++ {
		element := parts[i]
		switch element {
		case "usesrc":
			data.UseSrc = true
			i++
			if i >= len(parts) {
				return nil, &errors.ParseError{Parser: "Source", Line: line}
			}
			switch {
			case strings.HasPrefix(parts[i], "clientip"):
				data.ClientIP = true
			case strings.HasPrefix(parts[i], "client"):
				data.Client = true
			case strings.HasPrefix(parts[i], "hdr_ip"):
				data.HdrIP = true
				param := strings.TrimPrefix(parts[i], "hdr_ip(")
				param = strings.TrimRight(param, ")")
				if strings.Contains(param, ":") {
					HdrAndOcc := strings.Split(param, ",")
					data.Hdr = HdrAndOcc[0]
					data.Occ = HdrAndOcc[1]
				} else {
					data.Hdr = param
				}

			default:
				if strings.Contains(parts[i], ":") {
					addressAndPort := strings.Split(parts[i], ":")
					data.AddressSecond = addressAndPort[0]
					if port, err := strconv.ParseInt(addressAndPort[1], 10, 64); err == nil {
						data.PortSecond = port
					}
				} else {
					data.AddressSecond = parts[i]
				}

			}
		case "interface":
			i++
			if i >= len(parts) {
				return nil, &errors.ParseError{Parser: "Source", Line: line}
			}
			data.Interface = parts[i]
		}
	}

	return data, nil
}

func (s *Source) Result() ([]common.ReturnResultLine, error) {
	if len(s.data) == 0 {
		return nil, errors.ErrFetch
	}
	result := make([]common.ReturnResultLine, len(s.data))
	for index, line := range s.data {
		var sb strings.Builder
		sb.WriteString("source")
		sb.WriteString(" ")
		sb.WriteString(line.Address)
		if line.Port > 0 {
			sb.WriteString(":")
			sb.WriteString(strconv.FormatInt(line.Port, 10))
		}
		if line.UseSrc {
			sb.WriteString(" ")
			sb.WriteString("usesrc")
			sb.WriteString(" ")
			if line.Client {
				sb.WriteString("client")
			}
			if line.ClientIP {
				sb.WriteString("clientip")
			}
			if line.HdrIP {
				sb.WriteString("hdr_ip(")
				sb.WriteString(line.Hdr)
				if line.Occ != "" {
					sb.WriteString(",")
					sb.WriteString(line.Occ)
				}
				sb.WriteString(")")
			}
			if line.AddressSecond != "" {
				sb.WriteString(line.AddressSecond)
				if line.PortSecond > 0 {
					sb.WriteString(":")
					sb.WriteString(strconv.FormatInt(line.PortSecond, 10))
				}
			}
		}
		if line.Interface != "" {
			sb.WriteString(" ")
			sb.WriteString("interface")
			sb.WriteString(" ")
			sb.WriteString(line.Interface)
		}
		result[index] = common.ReturnResultLine{
			Data:    sb.String(),
			Comment: line.Comment,
		}
	}

	return result, nil
}
