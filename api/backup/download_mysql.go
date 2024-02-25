package backup

import (
	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
	"net/http"
)

type downloadMysql struct {
	api.BasicMethod
}

// CallDownloadMysql is a creation function that returns a [core.APIMethod] corresponding to the method [getData].
// The function is waiting for the domain name for which it is necessary to get data from the DNS server.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func CallDownloadMysql(backupId result.ID, bases []result.DatabaseName) core.APIMethod[result.BoolResult] {
	return &downloadMysql{
		BasicMethod: *api.CallMethod(
			DownloadMysqlMethodPath,
			&downloadMysqlRequest{BackupId: backupId, Bases: bases},
			nil,
		),
	}
}

// GetHTTPMethod returns name of http method for method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*downloadMysql) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*downloadMysql) GetName() core.MethodName {
	return DownloadMysqlMethodName
}

type downloadMysqlRequest struct {
	BackupId result.ID             `json:"backup_id"`
	Bases    []result.DatabaseName `json:"bases"`
}
