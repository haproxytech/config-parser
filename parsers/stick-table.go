package parsers

import (
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type StickTable struct {
	data []types.StickTable
}

func (h *StickTable) parse(line string, parts []string, comment string) (*types.StickTable, error) {
	if len(parts) >= 3 && parts[0] == "stick-table" && parts[1] == "type" {
		index := 2
		data := &types.StickTable{
			Type:    parts[index],
			Comment: comment,
		}
		index++
		for index < len(parts) {
			switch parts[index] {
			case "len":
				index++
				if index == len(parts) {
					return nil, &errors.ParseError{Parser: "StickTable", Line: line}
				}
				data.Length = parts[index]
			case "size":
				index++
				if index == len(parts) {
					return nil, &errors.ParseError{Parser: "StickTable", Line: line}
				}
				data.Size = parts[index]
			case "expire":
				index++
				if index == len(parts) {
					return nil, &errors.ParseError{Parser: "StickTable", Line: line}
				}
				data.Expire = parts[index]
			case "nopurge":
				data.NoPurge = true
			case "peers":
				index++
				if index == len(parts) {
					return nil, &errors.ParseError{Parser: "StickTable", Line: line}
				}
				data.Peers = parts[index]
			case "store":
				index++
				if index == len(parts) {
					return nil, &errors.ParseError{Parser: "StickTable", Line: line}
				}
				data.Store = parts[index]
			default:
				return nil, &errors.ParseError{Parser: "StickTable", Line: line}
			}
			index++
		}
		return data, nil
	}
	return nil, &errors.ParseError{Parser: "StickTable", Line: line}
}

func (h *StickTable) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		var data strings.Builder
		data.WriteString("stick-table type ")
		data.WriteString(req.Type)
		if req.Length != "" {
			data.WriteString(" len ")
			data.WriteString(req.Length)
		}
		if req.Size != "" {
			data.WriteString(" size ")
			data.WriteString(req.Size)
		}
		if req.Expire != "" {
			data.WriteString(" expire ")
			data.WriteString(req.Expire)
		}
		if req.NoPurge {
			data.WriteString(" nopurge")
		}
		if req.Peers != "" {
			data.WriteString(" peers ")
			data.WriteString(req.Peers)
		}
		if req.Store != "" {
			data.WriteString(" store ")
			data.WriteString(req.Store)
		}
		result[index] = common.ReturnResultLine{
			Data:    data.String(),
			Comment: req.Comment,
		}
	}
	return result, nil
}
