package backup

import (
	"net/http"

	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
)

type restoreFile struct {
	api.BasicMethod
}

// CallRestoreFile is a creation function that returns a [core.APIMethod] corresponding to the method [restoreFile].
// The function expects the backup ID from which files need to be restored
// and a list of paths from which files are restored.
//
// [restoreFile]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#restorefile
func CallRestoreFile(backupID result.ID, paths []string) core.APIMethod[result.BoolResult] {
	return &restoreFile{
		BasicMethod: *api.CallMethod(RestoreFileMethodPath, &restoreFileRequest{BackupID: backupID, Paths: paths}, nil),
	}
}

// GetHTTPMethod returns name of http method for method [restoreFile].
//
// [restoreFile]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#restorefile
func (*restoreFile) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [restoreFile].
//
// [restoreFile]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#restorefile
func (*restoreFile) GetName() core.MethodName {
	return RestoreFileMethodName
}

type restoreFileRequest struct {
	BackupID result.ID `json:"backup_id"`
	Paths    []string  `json:"paths"`
}
