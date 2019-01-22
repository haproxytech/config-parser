package filters

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type Filter interface {
	Parse(parts []string, comment string) error
	Result() common.ReturnResultLine
}

type Filters struct {
	Filters []Filter
}

func (h *Filters) Init() {
	h.Filters = []Filter{}
}

func (h *Filters) GetParserName() string {
	return "filter"
}

func (f *Filters) ParseFilter(filter Filter, parts []string, comment string) error {
	err := filter.Parse(parts, "")
	if err != nil {
		return &errors.ParseError{Parser: "FilterLines", Line: ""}
	}
	f.Filters = append(f.Filters, filter)
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

func (h *Filters) Valid() bool {
	if len(h.Filters) > 0 {
		return true
	}
	return false
}

func (h *Filters) Result(AddComments bool) []common.ReturnResultLine {
	result := make([]common.ReturnResultLine, len(h.Filters))
	for index, req := range h.Filters {
		result[index] = req.Result()
	}
	return result
}
