package http

import (
	"strings"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/parsers/http/actions"
	"github.com/haproxytech/config-parser/types"
)

type HTTPRequests struct {
	Name string
	data []types.HTTPAction
}

func (h *HTTPRequests) Init() {
	h.Name = "http-request"
	h.data = []types.HTTPAction{}
}

func (f *HTTPRequests) ParseHTTPRequest(request types.HTTPAction, parts []string, comment string) error {
	err := request.Parse(parts, comment)
	if err != nil {
		return &errors.ParseError{Parser: "HTTPRequestLines", Line: ""}
	}
	f.data = append(f.data, request)
	return nil
}

func (h *HTTPRequests) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) >= 2 && parts[0] == "http-request" {
		var err error
		switch parts[1] {
		case "add-header":
			err = h.ParseHTTPRequest(&actions.AddHeader{}, parts, comment)
		case "allow":
			err = h.ParseHTTPRequest(&actions.Allow{}, parts, comment)
		case "auth":
			err = h.ParseHTTPRequest(&actions.Auth{}, parts, comment)
		case "del-header":
			err = h.ParseHTTPRequest(&actions.DelHeader{}, parts, comment)
		case "deny":
			err = h.ParseHTTPRequest(&actions.Deny{}, parts, comment)
		case "redirect":
			err = h.ParseHTTPRequest(&actions.Redirect{}, parts, comment)
		case "replace-header":
			err = h.ParseHTTPRequest(&actions.ReplaceHeader{}, parts, comment)
		case "replace-value":
			err = h.ParseHTTPRequest(&actions.ReplaceValue{}, parts, comment)
		case "send-spoe-group":
			err = h.ParseHTTPRequest(&actions.SendSpoeGroup{}, parts, comment)
		case "set-header":
			err = h.ParseHTTPRequest(&actions.SetHeader{}, parts, comment)
		case "set-log-level":
			err = h.ParseHTTPRequest(&actions.SetLogLevel{}, parts, comment)
		case "set-var":
			err = h.ParseHTTPRequest(&actions.SetVar{}, parts, comment)
		case "tarpit":
			err = h.ParseHTTPRequest(&actions.Tarpit{}, parts, comment)
		default:
			if strings.HasPrefix(parts[1], "add-acl(") {
				err = h.ParseHTTPRequest(&actions.AddAcl{}, parts, comment)
			} else if strings.HasPrefix(parts[1], "del-acl(") {
				err = h.ParseHTTPRequest(&actions.DelAcl{}, parts, comment)
			} else if strings.HasPrefix(parts[1], "set-var(") {
				err = h.ParseHTTPRequest(&actions.SetVar{}, parts, comment)
			} else {
				return "", &errors.ParseError{Parser: "HTTPRequestLines", Line: line}
			}
		}
		if err != nil {
			return "", err
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "HTTPRequestLines", Line: line}
}

func (h *HTTPRequests) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(h.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(h.data))
	for index, req := range h.data {
		result[index] = common.ReturnResultLine{
			Data:    "http-request " + req.String(),
			Comment: req.GetComment(),
		}
	}
	return result, nil
}
