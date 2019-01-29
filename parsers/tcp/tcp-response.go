package tcp

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/parsers/tcp/actions"
)

type TCPResponses struct {
	Mode string //frontent, backend
	data []TCPAction
}

func (h *TCPResponses) Init() {
	h.data = []TCPAction{}
}

func (h *TCPResponses) GetParserName() string {
	return "tcp-response"
}

func (h *TCPResponses) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(h.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return h.data, nil
}

func (h *TCPResponses) Set(data common.ParserData) error {
	if data == nil {
		h.Init()
		return nil
	}
	switch newValue := data.(type) {
	case []TCPAction:
		h.data = newValue
	case *TCPAction:
		h.data = append(h.data, *newValue)
	case TCPAction:
		h.data = append(h.data, newValue)
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (h *TCPResponses) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := h.Get(false)
	h.Init()
	_, err := h.Parse(data, parts, []string{}, comment)
	if err != nil {
		h.Set(oldData)
	}
	return err
}

func (f *TCPResponses) ParseTCPRequest(request TCPAction, parts []string, comment string) error {
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
			Data:    "tpc-response " + req.String(),
			Comment: req.GetComment(),
		}
	}
	return result, nil
}
