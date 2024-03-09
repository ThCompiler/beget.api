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

// CallGetFileList is a creation function that returns a [core.APIMethod] corresponding to the method [getData].
// The function is waiting for the domain name for which it is necessary to get data from the DNS server.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func CallGetFileList(backupID *result.ID, path string) core.APIMethod[result.FileList] {
	return &getFileList{
		BasicMethod: *api.CallMethod(GetFileListMethodPath, &getFileListRequest{BackupID: backupID, Path: path}, nil),
	}
}

// GetHTTPMethod returns name of http method for method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*getFileList) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*getFileList) GetName() core.MethodName {
	return GetFileListMethodName
}

type getFileListRequest struct {
	BackupID *result.ID `json:"backup_id,omitempty"`
	Path     string     `json:"path"`
}
