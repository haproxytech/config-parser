package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/params"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Servers struct {
	data []types.Server
}

func (h *Servers) Init() {
	h.data = []types.Server{}
}

func (h *Servers) GetParserName() string {
	return "server"
}

func (h *Servers) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(h.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return h.data, nil
}

func (h *Servers) Set(data common.ParserData) error {
	if data == nil {
		h.Init()
		return nil
	}
	switch newValue := data.(type) {
	case []types.Server:
		h.data = newValue
	case *types.Server:
		h.data = append(h.data, *newValue)
	case types.Server:
		h.data = append(h.data, newValue)
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (h *Servers) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := h.Get(false)
	h.Init()
	_, err := h.Parse(data, parts, []string{}, comment)
	if err != nil {
		h.Set(oldData)
	}
	return err
}

func (h *Servers) parseLine(line string, parts []string, comment string) (*types.Server, error) {
	if len(parts) >= 3 {
		data := &types.Server{
			Name:    parts[1],
			Address: parts[2],
			Params:  params.ParseServerOptions(parts[3:]),
			Comment: comment,
		}
		return data, nil
	}
	return nil, &errors.ParseError{Parser: "Server", Line: line}
}

func (h *Servers) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "server" {
		data, err := h.parseLine(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "Server", Line: line}
		}
		h.data = append(h.data, *data)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "Server", Line: line}
}

func (h *Servers) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("server %s %s %s", req.Name, req.Address, params.ServerOptionsString(req.Params)),
			Comment: req.Comment,
		}
	}
	return result, nil
}
