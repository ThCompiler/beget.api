package backup

import (
	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
	"net/http"
)

type getMysqlBackupList struct {
	api.BasicMethod
}

// CallGetMysqlBackupList is a creation function that returns a [core.APIMethod] corresponding to the method [getData].
// The function is waiting for the domain name for which it is necessary to get data from the DNS server.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func CallGetMysqlBackupList() core.APIMethod[result.MYSQLBackupList] {
	return &getMysqlBackupList{
		BasicMethod: *api.CallMethod(GetMysqlBackupListMethodPath, nil, nil),
	}
}

// GetHTTPMethod returns name of http method for method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*getMysqlBackupList) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [getData].
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (*getMysqlBackupList) GetName() core.MethodName {
	return GetMysqlBackupListMethodName
}
