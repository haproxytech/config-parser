package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/params"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type DefaultServers struct {
	data []types.DefaultServer
}

func (h *DefaultServers) Init() {
	h.data = []types.DefaultServer{}
}

func (h *DefaultServers) GetParserName() string {
	return "default-server"
}

func (h *DefaultServers) Clear() {
	h.Init()
}

func (h *DefaultServers) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(h.data) == 0 && !createIfNotExist {
		return nil, &errors.FetchError{}
	}
	return h.data, nil
}

func (h *DefaultServers) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case []types.DefaultServer:
		h.data = newValue
	case types.DefaultServer:
		h.data = append(h.data, newValue)
	case *types.DefaultServer:
		h.data = append(h.data, *newValue)
	}
	return fmt.Errorf("casting error")
}

func (h *DefaultServers) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := h.Get(false)
	h.Clear()
	_, err := h.Parse(data, parts, []string{}, comment)
	if err != nil {
		h.Set(oldData)
	}
	return err
}

func (h *DefaultServers) parseLine(line string, parts []string, comment string) (*types.DefaultServer, error) {
	if len(parts) >= 1 {
		data := &types.DefaultServer{
			Params:  params.ParseServerOptions(parts),
			Comment: comment,
		}
		return data, nil
	}
	return nil, &errors.ParseError{Parser: "DefaultServers", Line: line}
}

func (h *DefaultServers) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "default-server" {
		data, err := h.parseLine(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "DefaultServers", Line: line}
		}
		h.data = append(h.data, *data)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "DefaultServers", Line: line}
}

func (h *DefaultServers) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, &errors.FetchError{}
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("default-server %s", params.ServerOptionsString(req.Params)),
			Comment: req.Comment,
		}
	}
	return result, nil
}
