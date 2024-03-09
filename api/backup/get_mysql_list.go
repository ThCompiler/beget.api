package backup

import (
	"net/http"

	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
)

type getMysqlList struct {
	api.BasicMethod
}

// CallGetMysqlList is a creation function that returns a [core.APIMethod] corresponding to the method [getMysqlList].
// The function expects the backup ID, for which it is necessary to get a list of databases.
// If the backup ID is a nil, the databases will be searched in the current copy.
//
// [getMysqlList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getmysqllist
func CallGetMysqlList(backupID *result.ID) core.APIMethod[result.MYSQLList] {
	return &getMysqlList{
		BasicMethod: *api.CallMethod(GetMysqlListMethodPath, &getMysqlListRequest{BackupID: backupID}, nil),
	}
}

// GetHTTPMethod returns name of http method for method [getMysqlList].
//
// [getMysqlList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getmysqllist
func (*getMysqlList) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [getMysqlList].
//
// [getMysqlList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getmysqllist
func (*getMysqlList) GetName() core.MethodName {
	return GetMysqlListMethodName
}

type getMysqlListRequest struct {
	BackupID *result.ID `json:"backup_id,omitempty"`
}
