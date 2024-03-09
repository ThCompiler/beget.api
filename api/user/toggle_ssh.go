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

// Status represents the status of SSH permissions.
type Status rune

const (
	ENABLE  Status = '1' // SSH connection is enabled for an account
	DISABLE Status = '0' // SSH connection is disabled for an account
)

type toggleSSH struct {
	api.BasicMethod
}

// CallToggleSSH is a creation function that returns a [core.APIMethod] corresponding to the method [toggleSsh].
// The function is waiting for SSH status to be received for the main Beget account.
//
// [toggleSsh]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom#togglessh
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

// CallToggleSSHFTP is a creation function that returns a [core.APIMethod] corresponding to the method [toggleSsh].
// The function waits for the FTP account login and SSH status for this account.
//
// [toggleSsh]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom#togglessh
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

// GetHTTPMethod returns name of http method for method [toggleSsh].
//
// [toggleSsh]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom#togglessh
func (*toggleSSH) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [toggleSsh].
//
// [toggleSsh]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom#togglessh
func (*toggleSSH) GetName() core.MethodName {
	return ToggleSSHMethodName
}
