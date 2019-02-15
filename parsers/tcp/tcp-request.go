package tcp

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/parsers/tcp/actions"
	"github.com/haproxytech/config-parser/types"
)

type TCPRequests struct {
	Name string
	Mode string //frontent, backend
	data []types.TCPAction
}

func (h *TCPRequests) Init() {
	h.Name = "tcp-request"
	h.data = []types.TCPAction{}
}

func (f *TCPRequests) ParseTCPRequest(request types.TCPAction, parts []string, comment string) error {
	err := request.Parse(parts, comment)
	if err != nil {
		return &errors.ParseError{Parser: "HTTPRequestLines", Line: ""}
	}
	f.data = append(f.data, request)
	return nil
}

func (h *TCPRequests) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) >= 2 && parts[0] == "tcp-request" {
		var err error
		switch parts[1] {
		case "connection":
			if h.Mode == "backend" {
				return "", &errors.ParseError{Parser: "HTTPRequestLines", Line: line}
			}
			err = h.ParseTCPRequest(&actions.Connection{}, parts, comment)
		case "session":
			if h.Mode == "backend" {
				return "", &errors.ParseError{Parser: "HTTPRequestLines", Line: line}
			}
			err = h.ParseTCPRequest(&actions.Session{}, parts, comment)
		case "content":
			err = h.ParseTCPRequest(&actions.Content{}, parts, comment)
		case "inspect-delay":
			err = h.ParseTCPRequest(&actions.InspectDelay{}, parts, comment)
		default:
			return "", &errors.ParseError{Parser: "HTTPRequestLines", Line: line}
		}
		if err != nil {
			return "", err
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "HTTPRequestLines", Line: line}
}

func (h *TCPRequests) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    "tcp-request " + req.String(),
			Comment: req.GetComment(),
		}
	}
	return result, nil
}
