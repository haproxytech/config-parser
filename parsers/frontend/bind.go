package frontend

import (
	"fmt"

	bindoptions "github.com/haproxytech/config-parser/bind-options"
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Binds struct {
	data []types.Bind
}

func (h *Binds) Init() {
	h.data = []types.Bind{}
}

func (h *Binds) GetParserName() string {
	return "bind"
}

func (h *Binds) Clear() {
	h.Init()
}

func (h *Binds) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(h.data) == 0 && !createIfNotExist {
		return nil, &errors.FetchError{}
	}
	return h.data, nil
}

func (h *Binds) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case []types.Bind:
		h.data = newValue
	case types.Bind:
		h.data = append(h.data, newValue)
	case *types.Bind:
		h.data = append(h.data, *newValue)
	}
	return fmt.Errorf("casting error")
}

func (h *Binds) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := h.Get(false)
	h.Clear()
	_, err := h.Parse(data, parts, []string{}, comment)
	if err != nil {
		h.Set(oldData)
	}
	return err
}

func (h *Binds) parseBindLine(line string, parts []string, comment string) (*types.Bind, error) {
	if len(parts) >= 1 {
		data := &types.Bind{
			Path:    parts[1],
			Params:  bindoptions.Parse(parts[2:]),
			Comment: comment,
		}
		return data, nil
	}
	return nil, &errors.ParseError{Parser: "BindLines", Line: line}
}

func (h *Binds) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "bind" {
		data, err := h.parseBindLine(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "BindLines", Line: line}
		}
		h.data = append(h.data, *data)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "BindLines", Line: line}
}

func (h *Binds) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, &errors.FetchError{}
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("bind %s %s", req.Path, bindoptions.String(req.Params)),
			Comment: req.Comment,
		}
	}
	return result, nil
}
