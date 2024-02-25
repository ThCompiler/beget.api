package user

import (
	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
	"net/http"
	"net/url"
	"strconv"
)

const (
	statusField   = "status"
	fTPLoginField = "ftplogin"
)

type Status int64

const (
	ENABLE  Status = 1
	DISABLE Status = 0
)

type toggleSSH struct {
	api.BasicMethod
}

// CallToggleSSH is a creation function that returns a [core.APIMethod] corresponding to the method [getData].
// The function is waiting for the domain name for which it is necessary to get data from the DNS server.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func CallToggleSSH(status Status) core.APIMethod[result.BoolResult] {
	return &toggleSSH{
		BasicMethod: *api.CallMethod(
			ToggleSSHMethodPath,
			nil,
			url.Values{
				statusField: []string{strconv.FormatInt(int64(status), 10)},
			},
		),
	}
}

// CallToggleSSHFTP is a creation function that returns a [core.APIMethod] corresponding to the method [getData].
// The function is waiting for the domain name for which it is necessary to get data from the DNS server.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func CallToggleSSHFTP(status Status, FTPLogin string) core.APIMethod[result.GetData] {
	return &toggleSSH{
		BasicMethod: *api.CallMethod(
			ToggleSSHMethodPath,
			nil,
			url.Values{
				statusField:   []string{strconv.FormatInt(int64(status), 10)},
				fTPLoginField: []string{FTPLogin},
			},
		),
	}
}

// GetHTTPMethod returns name of http method for method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*toggleSSH) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*toggleSSH) GetName() core.MethodName {
	return ToggleSSHMethodName
}
