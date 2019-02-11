package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Daemon struct {
	data *types.Enabled
}

func (d *Daemon) Init() {
	d.data = nil
}

func (d *Daemon) GetParserName() string {
	return "daemon"
}

func (d *Daemon) Get(createIfNotExist bool) (common.ParserData, error) {
	if d.data == nil {
		if createIfNotExist {
			d.data = &types.Enabled{}
			return d.data, nil
		}
		return nil, errors.FetchError
	}
	return d.data, nil
}

func (p *Daemon) GetOne(index int) (common.ParserData, error) {
	if index != 0 {
		return nil, errors.FetchError
	}
	if p.data == nil {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (d *Daemon) Set(data common.ParserData, index int) error {
	if data == nil {
		d.data = nil
		return nil
	}
	switch newValue := data.(type) {
	case *types.Enabled:
		d.data = newValue
	case types.Enabled:
		d.data = &newValue
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (d *Daemon) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "daemon" {
		d.data = &types.Enabled{
			Comment: comment,
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "Daemon", Line: line}
}

func (d *Daemon) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if d.data == nil {
		return nil, errors.FetchError
	}
	return []common.ReturnResultLine{
		common.ReturnResultLine{
			Data:    "daemon",
			Comment: d.data.Comment,
		},
	}, nil
}
