// Package user implements the functionality of [User methods].
// Package implements [getAccountInfo] and [toggleSsh] methods.
// To create the appropriate methods, you need to call either [CallGetAccountInfo] or [CallToggleSSH].
//
// [User methods]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom
// [getAccountInfo]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom#getaccountinfo
// [toggleSsh]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom#togglessh
package user

// Constants used to implement the [getAccountInfo] and [toggleSsh] methods.
//
// [getAccountInfo]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom#getaccountinfo
// [toggleSsh]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom#togglessh
const (
	GetAccountInfoMethodName = "GetAccountInfo"
	GetAccountInfoPath       = "user/getAccountInfo"

	ToggleSSHMethodName = "ToggleSsh"
	ToggleSSHMethodPath = "user/toggleSsh"
)
