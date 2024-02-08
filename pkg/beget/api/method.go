package api

import (
	"encoding/json"
	"net/url"

	infoapi "github.com/ThCompiler/go.beget.api/pkg/beget/core/info/api"
)

const (
	inputFormatField = "input_format"
	inputDataField   = "input_data"
)

type BasicMethod struct {
	err error
	uRL *url.URL
}

func (b *BasicMethod) GetURL() *url.URL {
	return b.uRL
}

func (b *BasicMethod) Error() error {
	return b.err
}

func CallError(err error) *BasicMethod {
	return &BasicMethod{
		err: err,
	}
}

func CallMethod(methodPath string, requestBody any) *BasicMethod {
	uRL := &url.URL{}
	uRL.Path = methodPath

	if requestBody == nil {
		return &BasicMethod{
			uRL: uRL,
			err: nil,
		}
	}

	request, err := json.Marshal(&requestBody)
	if err != nil {
		return CallError(err)
	}

	vals := uRL.Query()
	vals.Add(inputFormatField, string(infoapi.JSON))
	vals.Add(inputDataField, string(request))
	uRL.RawQuery = vals.Encode()

	return &BasicMethod{
		uRL: uRL,
		err: nil,
	}
}
