package frontend

import (
	"fmt"

	bindoptions "github.com/haproxytech/config-parser/bind-options"
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type Bind struct {
	Path    string //can be address:port or socket path
	Params  []bindoptions.BindOption
	Comment string
}

type Binds struct {
	Binds []Bind
}

func (h *Binds) Init() {
	h.Binds = []Bind{}
}

func (h *Binds) GetParserName() string {
	return "bind"
}

func (h *Binds) parseBindLine(line string, parts []string, comment string) (Bind, error) {
	if len(parts) >= 1 {
		data := Bind{
			Path:    parts[1],
			Params:  bindoptions.Parse(parts[2:]),
			Comment: comment,
		}
		return data, nil
	}
	return Bind{}, &errors.ParseError{Parser: "BindLines", Line: line}
}

func (h *Binds) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "bind" {
		request, err := h.parseBindLine(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "BindLines", Line: line}
		}
		h.Binds = append(h.Binds, request)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "BindLines", Line: line}
}

func (h *Binds) Valid() bool {
	if len(h.Binds) > 0 {
		return true
	}
	return false
}

func (h *Binds) Result(AddComments bool) []common.ReturnResultLine {
	result := make([]common.ReturnResultLine, len(h.Binds))
	for index, req := range h.Binds {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("bind %s %s", req.Path, bindoptions.String(req.Params)),
			Comment: req.Comment,
		}
	}
	return result
}
