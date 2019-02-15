package filters

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Filters struct {
	Name string
	data []types.Filter
}

func (h *Filters) Init() {
	h.data = []types.Filter{}
	h.Name = "filter"
}

func (f *Filters) ParseFilter(filter types.Filter, parts []string, comment string) error {
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
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = req.Result()
	}
	return result, nil
}
