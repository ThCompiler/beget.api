package backup

import (
	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
	"net/http"
)

type restoreFile struct {
	api.BasicMethod
}

// CallRestoreFile is a creation function that returns a [core.APIMethod] corresponding to the method [getData].
// The function is waiting for the domain name for which it is necessary to get data from the DNS server.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func CallRestoreFile(backupId result.ID, paths []string) core.APIMethod[result.BoolResult] {
	return &restoreFile{
		BasicMethod: *api.CallMethod(RestoreFileMethodPath, &restoreFileRequest{BackupId: backupId, Paths: paths}, nil),
	}
}

// GetHTTPMethod returns name of http method for method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*restoreFile) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*restoreFile) GetName() core.MethodName {
	return RestoreFileMethodName
}

type restoreFileRequest struct {
	BackupId result.ID `json:"backup_id"`
	Paths    []string  `json:"paths"`
}
