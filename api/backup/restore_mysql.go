package backup

import (
	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
	"net/http"
)

type restoreMysql struct {
	api.BasicMethod
}

// CallRestoreMysql is a creation function that returns a [core.APIMethod] corresponding to the method [getData].
// The function is waiting for the domain name for which it is necessary to get data from the DNS server.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func CallRestoreMysql(backupId result.ID, bases []result.DatabaseName) core.APIMethod[result.BoolResult] {
	return &restoreMysql{
		BasicMethod: *api.CallMethod(RestoreMysqlMethodPath, &restoreMysqlRequest{BackupId: backupId, Bases: bases}, nil),
	}
}

// GetHTTPMethod returns name of http method for method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*restoreMysql) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*restoreMysql) GetName() core.MethodName {
	return RestoreMysqlMethodName
}

type restoreMysqlRequest struct {
	BackupId result.ID             `json:"backup_id"`
	Bases    []result.DatabaseName `json:"bases"`
}
