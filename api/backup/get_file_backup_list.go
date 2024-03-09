package backup

import (
	"net/http"

	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
)

type getFileBackupList struct {
	api.BasicMethod
}

// CallGetFileBackupList is a creation function
// that returns a [core.APIMethod] corresponding to the method [getFileBackupList].
//
// [getFileBackupList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getfilebackuplist
func CallGetFileBackupList() core.APIMethod[result.FileBackupList] {
	return &getFileBackupList{
		BasicMethod: *api.CallMethod(GetFileBackupListMethodPath, nil, nil),
	}
}

// GetHTTPMethod returns name of http method for method [getFileBackupList].
//
// [getFileBackupList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getfilebackuplist
func (*getFileBackupList) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [getFileBackupList].
//
// [getFileBackupList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getfilebackuplist
func (*getFileBackupList) GetName() core.MethodName {
	return GetFileBackupListMethodName
}
