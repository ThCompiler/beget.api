package api

import (
	"encoding/json"
	"github.com/ThCompiler/go.beget.api/core"
	"net/url"
)

const (
	inputFormatField = "input_format"
	inputDataField   = "input_data"
)

// BasicMethod is the basis of other api methods.
// It implements the general core operation from [github.com/ThCompiler/go.beget.api/pkg/beget/core.APIMethod] interface.
// It implements GetURL and Error methods and general operations of storing error or of creating method information.
// Other methods embed the basic one and use CallError and CallMethod to build specific information about itself.
// The [github.com/ThCompiler/go.beget.api/pkg/beget/api/dns.CallGetData] can be used as an example.
type BasicMethod struct {
	err error
	uRL *url.URL
}

// GetURL returns suffix url of api method. The url is expected to contain the required query parameters
func (b *BasicMethod) GetURL() *url.URL {
	return b.uRL
}

// Error returns any errors when generating information about the request to the method
func (b *BasicMethod) Error() error {
	return b.err
}

// CallError stores any errors that occurred as a result of the method build
func CallError(err error) *BasicMethod {
	return &BasicMethod{
		err: err,
	}
}

// CallMethod creates an [url.URL] and saves it as the result of the method build.
// As parameters, it expects the suffix of a specific method endpoint (methodPath)
// and the request body (requestBody) if it is necessary for this method (otherwise nil).
// It's expected that requestBody can be marshaled.
// If requestBody isn't nil, it will be marshaled and added to query params.
// Also, a field describing the type of input data will also be added with the "json" value.
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
	vals.Add(inputFormatField, string(core.JSON))
	vals.Add(inputDataField, string(request))
	uRL.RawQuery = vals.Encode()

	return &BasicMethod{
		uRL: uRL,
		err: nil,
	}
}
