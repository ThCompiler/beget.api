package user

import (
	"net/http"
	"net/url"

	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
)

const (
	statusField   = "status"
	fTPLoginField = "ftplogin"
)

type Status rune

const (
	ENABLE  Status = '1'
	DISABLE Status = '0'
)

type toggleSSH struct {
	api.BasicMethod
}

// CallToggleSSH is a creation function that returns a [core.APIMethod] corresponding to the method [getData].
// The function is waiting for the domain name for which it is necessary to get data from the DNS server.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func CallToggleSSH(status Status) core.APIMethod[result.SSHToggle] {
	return &toggleSSH{
		BasicMethod: *api.CallMethod(
			ToggleSSHMethodPath,
			nil,
			url.Values{
				statusField: []string{string(status)},
			},
		),
	}
}

// CallToggleSSHFTP is a creation function that returns a [core.APIMethod] corresponding to the method [getData].
// The function is waiting for the domain name for which it is necessary to get data from the DNS server.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func CallToggleSSHFTP(status Status, fTPLogin string) core.APIMethod[result.SSHToggle] {
	return &toggleSSH{
		BasicMethod: *api.CallMethod(
			ToggleSSHMethodPath,
			nil,
			url.Values{
				statusField:   []string{string(status)},
				fTPLoginField: []string{fTPLogin},
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
