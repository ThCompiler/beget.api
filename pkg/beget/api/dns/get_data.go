package dns

import (
	"net/http"

	"github.com/ThCompiler/go.beget.api/pkg/beget/api"
	"github.com/ThCompiler/go.beget.api/pkg/beget/api/dns/result"
	"github.com/ThCompiler/go.beget.api/pkg/beget/core"
)

type getData struct {
	api.BasicMethod
}

func CallGetData(domainName string) core.APIMethod[result.GetData] {
	return &getData{
		BasicMethod: *api.CallMethod(GetDataMethodPath, &getDataRequest{Fqdn: domainName}),
	}
}

func (*getData) GetHTTPMethod() string {
	return http.MethodPost
}

func (*getData) GetName() core.MethodName {
	return GetDataMethodName
}

type getDataRequest struct {
	Fqdn string `json:"fqdn"`
}
