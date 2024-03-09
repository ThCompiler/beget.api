package backup

import (
	"net/http"

	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
)

type getFileList struct {
	api.BasicMethod
}

// CallGetFileList is a creation function that returns a [core.APIMethod] corresponding to the method [getFileList].
// The function expects the backup ID, for which it is necessary to get a list of files
// and the path to the directory whose contents are being requested.
// If the backup ID is a nil, the directory will be searched in the current copy.
//
// [getFileList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getfilelist
func CallGetFileList(backupID *result.ID, path string) core.APIMethod[result.FileList] {
	return &getFileList{
		BasicMethod: *api.CallMethod(GetFileListMethodPath, &getFileListRequest{BackupID: backupID, Path: path}, nil),
	}
}

// GetHTTPMethod returns name of http method for method [getFileList].
//
// [getFileList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getfilelist
func (*getFileList) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [getFileList].
//
// [getFileList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getfilelist
func (*getFileList) GetName() core.MethodName {
	return GetFileListMethodName
}

type getFileListRequest struct {
	BackupID *result.ID `json:"backup_id,omitempty"`
	Path     string     `json:"path"`
}
