package backup

import (
	"net/http"

	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
)

type downloadMysql struct {
	api.BasicMethod
}

// CallDownloadMysql is a creation function that returns a [core.APIMethod] corresponding to the method [downloadMysql].
// The function expects the backup ID from which data needs to be downloaded,
// and a list of databases whose data is being downloaded.
// If the backup ID is a nil, the databases will be searched in the current copy.
//
// [downloadMysql]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#downloadmysql
func CallDownloadMysql(backupID *result.ID, bases []result.DatabaseName) core.APIMethod[result.BoolResult] {
	return &downloadMysql{
		BasicMethod: *api.CallMethod(
			DownloadMysqlMethodPath,
			&downloadMysqlRequest{BackupID: backupID, Bases: bases},
			nil,
		),
	}
}

// GetHTTPMethod returns name of http method for method [downloadMysql].
//
// [downloadMysql]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#downloadmysql
func (*downloadMysql) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [downloadMysql].
//
// [downloadMysql]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#downloadmysql
func (*downloadMysql) GetName() core.MethodName {
	return DownloadMysqlMethodName
}

type downloadMysqlRequest struct {
	BackupID *result.ID            `json:"backup_id,omitempty"`
	Bases    []result.DatabaseName `json:"bases"`
}
