package tcp

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/parsers/tcp/actions"
	"github.com/haproxytech/config-parser/types"
)

type TCPResponses struct {
	Name string
	Mode string //frontent, backend
	data []types.TCPAction
}

func (h *TCPResponses) Init() {
	h.Name = "tcp-response"
	h.data = []types.TCPAction{}
}

func (f *TCPResponses) ParseTCPRequest(request types.TCPAction, parts []string, comment string) error {
	err := request.Parse(parts, comment)
	if err != nil {
		return &errors.ParseError{Parser: "TCPResponses", Line: ""}
	}
	f.data = append(f.data, request)
	return nil
}

func (h *TCPResponses) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) >= 2 && parts[0] == "tcp-response" {
		var err error
		switch parts[1] {
		case "content":
			err = h.ParseTCPRequest(&actions.Content{}, parts, comment)
		case "inspect-delay":
			err = h.ParseTCPRequest(&actions.InspectDelay{}, parts, comment)
		default:
			return "", &errors.ParseError{Parser: "TCPResponses", Line: line}
		}
		if err != nil {
			return "", err
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "TCPResponses", Line: line}
}

func (h *TCPResponses) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    "tcp-response " + req.String(),
			Comment: req.GetComment(),
		}
	}
	return result, nil
}
