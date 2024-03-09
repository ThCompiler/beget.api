package backup

import (
	"net/http"

	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
)

type getMysqlBackupList struct {
	api.BasicMethod
}

// CallGetMysqlBackupList is a creation function
// that returns a [core.APIMethod] corresponding to the method [getMysqlBackupList].
//
// [getMysqlBackupList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getmysqlbackuplist
func CallGetMysqlBackupList() core.APIMethod[result.MYSQLBackupList] {
	return &getMysqlBackupList{
		BasicMethod: *api.CallMethod(GetMysqlBackupListMethodPath, nil, nil),
	}
}

// GetHTTPMethod returns name of http method for method [getMysqlBackupList].
//
// [getMysqlBackupList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getmysqlbackuplist
func (*getMysqlBackupList) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [getMysqlBackupList].
//
// [getMysqlBackupList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getmysqlbackuplist
func (*getMysqlBackupList) GetName() core.MethodName {
	return GetMysqlBackupListMethodName
}
