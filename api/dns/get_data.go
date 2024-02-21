package dns

import (
	"github.com/ThCompiler/go.beget.api/api/result"
	"net/http"

	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/core"
)

type getData struct {
	api.BasicMethod
}

// CallGetData is a creation function that returns a [core.APIMethod] corresponding to the method [getData].
// The function is waiting for the domain name for which it is necessary to get data from the DNS server.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func CallGetData(domainName string) core.APIMethod[result.GetData] {
	return &getData{
		BasicMethod: *api.CallMethod(GetDataMethodPath, &getDataRequest{Fqdn: domainName}),
	}
}

// GetHTTPMethod returns name of http method for method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*getData) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*getData) GetName() core.MethodName {
	return GetDataMethodName
}

// getDataRequest represents request body for [getData] method.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type getDataRequest struct {
	Fqdn string `json:"fqdn"`
}
