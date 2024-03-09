package backup

import (
	"net/http"

	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
)

type getLog struct {
	api.BasicMethod
}

// CallGetLog is a creation function that returns a [core.APIMethod] corresponding to the method [getLog].
//
// [getLog]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getlog
func CallGetLog() core.APIMethod[result.Log] {
	return &getLog{
		BasicMethod: *api.CallMethod(GetLogMethodPath, nil, nil),
	}
}

// GetHTTPMethod returns name of http method for method [getLog].
//
// [getLog]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getlog
func (*getLog) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [getData].
//
// [getLog]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getlog
func (*getLog) GetName() core.MethodName {
	return GetLogMethodName
}
