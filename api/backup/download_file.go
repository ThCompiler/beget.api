package backup

import (
	"net/http"

	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
)

type downloadFile struct {
	api.BasicMethod
}

// CallDownloadFile is a creation function that returns a [core.APIMethod] corresponding to the method [downloadFile].
// The function expects the backup ID from which files need to be downloaded
// and a list of paths from which files are downloaded.
// If the backup ID is a nil, the databases will be searched in the current copy.
//
// [downloadFile]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#downloadfile
func CallDownloadFile(backupID *result.ID, paths []string) core.APIMethod[result.BoolResult] {
	return &downloadFile{
		BasicMethod: *api.CallMethod(
			DownloadFileMethodPath,
			&downloadFileRequest{BackupID: backupID, Paths: paths},
			nil,
		),
	}
}

// GetHTTPMethod returns name of http method for method [downloadFile].
//
// [downloadFile]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#downloadfile
func (*downloadFile) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [downloadFile].
//
// [downloadFile]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#downloadfile
func (*downloadFile) GetName() core.MethodName {
	return DownloadFileMethodName
}

type downloadFileRequest struct {
	BackupID *result.ID `json:"backup_id,omitempty"`
	Paths    []string   `json:"paths"`
}
