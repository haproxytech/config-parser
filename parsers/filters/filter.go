package filters

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type Filter interface {
	Parse(parts []string, comment string) error
	Result() common.ReturnResultLine
}

type Filters struct {
	data []Filter
}

func (h *Filters) Init() {
	h.data = []Filter{}
}

func (h *Filters) GetParserName() string {
	return "filter"
}

func (h *Filters) Clear() {
	h.Init()
}

func (h *Filters) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(h.data) == 0 && !createIfNotExist {
		return nil, &errors.FetchError{}
	}
	return h.data, nil
}

func (h *Filters) Set(data common.ParserData) error {
	switch newValue := data.(type) {
	case []Filter:
		h.data = newValue
	case Filter:
		h.data = append(h.data, newValue)
	case *Filter:
		h.data = append(h.data, *newValue)
	}
	return fmt.Errorf("casting error")
}

func (h *Filters) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := h.Get(false)
	h.Clear()
	_, err := h.Parse(data, parts, []string{}, comment)
	if err != nil {
		h.Set(oldData)
	}
	return err
}

func (f *Filters) ParseFilter(filter Filter, parts []string, comment string) error {
	err := filter.Parse(parts, "")
	if err != nil {
		return &errors.ParseError{Parser: "FilterLines", Line: ""}
	}
	f.data = append(f.data, filter)
	return nil
}

func (h *Filters) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) >= 2 && parts[0] == "filter" {
		var err error
		switch parts[1] {
		case "trace":
			err = h.ParseFilter(&Trace{}, parts, comment)
		case "compression":
			err = h.ParseFilter(&Compression{}, parts, comment)
		case "cache":
			err = h.ParseFilter(&Cache{}, parts, comment)
		default:
			return "", &errors.ParseError{Parser: "FilterLines", Line: line}
		}
		if err != nil {
			return "", err
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "FilterLines", Line: line}
}

func (h *Filters) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, &errors.FetchError{}
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = req.Result()
	}
	return result, nil
}
