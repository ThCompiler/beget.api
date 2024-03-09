package user

import (
	"net/http"

	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
)

type getAccountInfo struct {
	api.BasicMethod
}

// CallGetAccountInfo is a creation function
// that returns a [core.APIMethod] corresponding to the method [getAccountInfo].
//
// [getAccountInfo]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom#getaccountinfo
func CallGetAccountInfo() core.APIMethod[result.UserInfo] {
	return &getAccountInfo{
		BasicMethod: *api.CallMethod(GetAccountInfoPath, nil, nil),
	}
}

// GetHTTPMethod returns name of http method for method [getAccountInfo].
//
// [getAccountInfo]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom#getaccountinfo
func (*getAccountInfo) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [getAccountInfo].
//
// [getAccountInfo]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom#getaccountinfo
func (*getAccountInfo) GetName() core.MethodName {
	return GetAccountInfoMethodName
}
