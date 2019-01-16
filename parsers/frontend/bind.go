package frontend

import (
	"fmt"

	bindoptions "github.com/haproxytech/config-parser/bind-options"
	"github.com/haproxytech/config-parser/errors"
)

type Bind struct {
	Path   string //can be address:port or socket path
	Params []bindoptions.BindOption
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

func (h *Binds) parseBindLine(line string, parts []string) (Bind, error) {
	if len(parts) >= 1 {
		data := Bind{
			Path:   parts[1],
			Params: bindoptions.Parse(parts[2:]),
		}
		return data, nil
	}
	return Bind{}, &errors.ParseError{Parser: "BindLines", Line: line}
}

func (h *Binds) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if parts[0] == "bind" {
		request, err := h.parseBindLine(line, parts)
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

func (h *Binds) Result(AddComments bool) []string {
	result := make([]string, len(h.Binds))
	for index, req := range h.Binds {
		result[index] = fmt.Sprintf("  bind %s %s", req.Path, bindoptions.String(req.Params))
	}
	return result
}
