package backup

import (
	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
	"net/http"
)

type getMysqlList struct {
	api.BasicMethod
}

// CallGetMysqlList is a creation function that returns a [core.APIMethod] corresponding to the method [getData].
// The function is waiting for the domain name for which it is necessary to get data from the DNS server.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func CallGetMysqlList(backupId result.ID) core.APIMethod[result.MYSQLList] {
	return &getMysqlList{
		BasicMethod: *api.CallMethod(GetMysqlListMethodPath, &getMysqlListRequest{BackupId: backupId}, nil),
	}
}

// GetHTTPMethod returns name of http method for method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*getMysqlList) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*getMysqlList) GetName() core.MethodName {
	return GetMysqlListMethodName
}

type getMysqlListRequest struct {
	BackupId result.ID `json:"backup_id"`
}
