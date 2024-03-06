package core

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const (
	loginField        = "login"
	passwordField     = "passwd"
	outputFormatField = "output_format"
)

// Client is a client of Beget.com to access by API
type Client struct {
	Login    string // Login of Beget.com user
	Password string // Password for API access of Beget.com user
}

// PrepareRequestWithClient is a function to prepare request to Beget.API system.
// As input parameters, it expects to receive information about the user and a built api method to call and http.Client.
// If http.client is nil, it will be created as an empty http.Client.
// As a result, either a ready-to-execute request to the Beget.API is returned, or one of the errors:
//   - Errors creating api method.
//   - Errors resolving BEGET host
//   - Errors creating HTTP request
//
// Currently, Golang cannot infer the Result generic parameter from the [APIMethod] interface,
// so when calling the method, you need to explicitly specify the generic parameter.
//
// For example, with method [github.com/ThCompiler/go.beget.api/api/dns.CallGetData] ([GetData]):
//
//	 import "github.com/ThCompiler/go.beget.api/api/result"
//
//		client := Client{ login: "user", password: "password" }
//		req, _ := PrepareRequest[result.GetData](client, dns.CallGetData("some.domain.com"))
//
// [GetData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func PrepareRequestWithClient[Result any](c Client, method APIMethod[Result], client *http.Client) (*BegetRequest[Result], error) {
	if method.Error() != nil {
		return nil, errors.Wrap(method.Error(), ErrFromAPIMethod.Error())
	}

	apiURL := method.GetURL()

	host := Host
	if GetMode() == Test {
		host = testHost
	}

	hostURL, err := url.Parse(host)
	if err != nil {
		return nil, errors.Wrap(err, ErrResolveBegetHost.Error())
	}

	requestURL := hostURL.ResolveReference(apiURL)

	query := requestURL.Query()
	query.Add(loginField, c.Login)
	query.Add(passwordField, c.Password)
	query.Add(outputFormatField, string(JSON))

	requestURL.RawQuery = query.Encode()

	httpRequest, err := http.NewRequest(method.GetHTTPMethod(), requestURL.String(), http.NoBody)
	if err != nil {
		return nil, errors.Wrap(err, "got error for method: "+string(method.GetName()))
	}

	// Add the user agent to the request.
	httpRequest.Header.Add("User-Agent", UserAgent)

	return NewBegetRequest[Result](httpRequest, client), nil
}

// PrepareRequest is a function to prepare request to Beget.API system.
// The function is an alias in the PrepareRequestWithClient for working with the default http client.
func PrepareRequest[Result any](c Client, method APIMethod[Result]) (*BegetRequest[Result], error) {
	return PrepareRequestWithClient[Result](c, method, nil)
}
