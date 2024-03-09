package backup

import (
	"net/http"

	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
)

type restoreMysql struct {
	api.BasicMethod
}

// CallRestoreMysql is a creation function that returns a [core.APIMethod] corresponding to the method [restoreMysql].
// The function expects the backup ID from which data needs to be restored,
// and a list of databases whose data is being restored.
//
// [restoreMysql]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#restoremysql
func CallRestoreMysql(backupID result.ID, bases []result.DatabaseName) core.APIMethod[result.BoolResult] {
	return &restoreMysql{
		BasicMethod: *api.CallMethod(
			RestoreMysqlMethodPath,
			&restoreMysqlRequest{BackupID: backupID, Bases: bases},
			nil,
		),
	}
}

// GetHTTPMethod returns name of http method for method [restoreMysql].
//
// [restoreMysql]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#restoremysql
func (*restoreMysql) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [restoreMysql].
//
// [restoreMysql]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#restoremysql
func (*restoreMysql) GetName() core.MethodName {
	return RestoreMysqlMethodName
}

type restoreMysqlRequest struct {
	BackupID result.ID             `json:"backup_id"`
	Bases    []result.DatabaseName `json:"bases"`
}
