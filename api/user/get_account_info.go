package user

import (
	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
	"net/http"
)

type getAccountInfo struct {
	api.BasicMethod
}

// CallGetAccountInfo is a creation function that returns a [core.APIMethod] corresponding to the method [getData].
// The function is waiting for the domain name for which it is necessary to get data from the DNS server.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func CallGetAccountInfo() core.APIMethod[result.UserInfo] {
	return &getAccountInfo{
		BasicMethod: *api.CallMethod(GetAccountInfoPath, nil, nil),
	}
}

// GetHTTPMethod returns name of http method for method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*getAccountInfo) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*getAccountInfo) GetName() core.MethodName {
	return GetAccountInfoMethodName
}
