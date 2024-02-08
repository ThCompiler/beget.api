package core

import (
	"net/http"
	"net/url"

	infoapi "github.com/ThCompiler/go.beget.api/pkg/beget/core/info/api"
	"github.com/ThCompiler/go.beget.api/pkg/beget/core/request"
	"github.com/pkg/errors"
)

const (
	loginField        = "login"
	passwordField     = "passwd"
	outputFormatField = "output_format"
)

type Client struct {
	Login    string
	Password string
}

func PrepareRequest[Result any](c Client, method APIMethod[Result]) (*request.BegetRequest[Result], error) {
	if method.Error() != nil {
		return nil, errors.Wrap(method.Error(), ErrFromAPIMethod.Error())
	}

	apiURL := method.GetURL()

	hostURL, err := url.Parse(Host)
	if err != nil {
		return nil, errors.Wrap(err, ErrResolveBegetHost.Error())
	}

	requestURL := hostURL.ResolveReference(apiURL)

	query := requestURL.Query()
	query.Add(loginField, c.Login)
	query.Add(passwordField, c.Password)
	query.Add(outputFormatField, string(infoapi.JSON))

	requestURL.RawQuery = query.Encode()

	httpRequest, err := http.NewRequest(method.GetHTTPMethod(), requestURL.String(), http.NoBody)
	if err != nil {
		return nil, errors.Wrap(err, "got error for method: "+string(method.GetName()))
	}

	// Add the user agent to the request.
	httpRequest.Header.Add("User-Agent", UserAgent)

	return request.NewBegetRequest[Result](httpRequest), nil
}
