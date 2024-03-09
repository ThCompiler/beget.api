// Package api implements API method-specific functionality.
// Package contains:
//   - [github.com/ThCompiler/go.beget.api/api/result] package,
//     that implements structures representing the result of API method responses.
//   - [github.com/ThCompiler/go.beget.api/api/dns] package,
//     that implements the functionality of [Dns methods].
//
// [Dns methods]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns
package api

import (
	"encoding/json"
	"net/url"

	"github.com/ThCompiler/go.beget.api/core"
)

const (
	inputFormatField = "input_format"
	inputDataField   = "input_data"
)

// BasicMethod is the basis of other api methods.
//
// It implements the general core operation from [github.com/ThCompiler/go.beget.api/core.APIMethod] interface.
// It implements GetURL and Error methods and general operations of storing error or of creating method information.
// Other methods embed the basic one and use CallError and CallMethod to build specific information about itself.
//
// The [github.com/ThCompiler/go.beget.api/api/dns.CallGetData] can be used as an example.
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
func CallMethod(methodPath string, requestBody any, urlVales url.Values) *BasicMethod {
	uRL := &url.URL{}
	uRL.Path = methodPath

	{
		vales := uRL.Query()

		for key, values := range urlVales {
			for _, val := range values {
				vales.Add(key, val)
			}
		}

		uRL.RawQuery = vales.Encode()
	}

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

	values := uRL.Query()
	values.Add(inputFormatField, string(core.JSON))
	values.Add(inputDataField, string(request))
	uRL.RawQuery = values.Encode()

	return &BasicMethod{
		uRL: uRL,
		err: nil,
	}
}
