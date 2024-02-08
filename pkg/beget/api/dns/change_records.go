package dns

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/ThCompiler/go.beget.api/pkg/beget/api"
	"github.com/ThCompiler/go.beget.api/pkg/beget/api/dns/records/set"
	"github.com/ThCompiler/go.beget.api/pkg/beget/api/result"
	"github.com/ThCompiler/go.beget.api/pkg/beget/core"
	"github.com/pkg/errors"
)

type changeRecords struct {
	api.BasicMethod
}

type SettingRecords func() (json.RawMessage, error)

func CallChangeRecords(domainName string, records SettingRecords) core.APIMethod[result.BoolResult] {
	request, err := records()
	if err != nil {
		return &getData{
			BasicMethod: *api.CallError(err),
		}
	}

	return &changeRecords{
		BasicMethod: *api.CallMethod(ChangeRecordsMethodPath, &ChangeRequest{Fqdn: domainName, Records: request}),
	}
}

func (*changeRecords) GetHTTPMethod() string {
	return http.MethodPost
}

func (*changeRecords) GetName() core.MethodName {
	return ChangeRecordsMethodName
}

func SetRecords(records settableRecords) SettingRecords {
	return setRecords(records, reflect.TypeOf(records).String())
}

func SetBasicRecords(records *set.BasicRecords) SettingRecords {
	return setRecords(records, "basic records")
}

func SetNsRecords(records *set.NSRecords) SettingRecords {
	return setRecords(records, "ns records")
}

func SetCNameRecords(records *set.CNAMERecords) SettingRecords {
	return setRecords(records, "cname records")
}

type settableRecords interface {
	Validate() error
}

func setRecords(records settableRecords, nameRecords string) SettingRecords {
	if err := records.Validate(); err != nil {
		panic(errors.Wrapf(err, "try create request with %s", nameRecords))
	}

	return func() (json.RawMessage, error) {
		jsonRecords, err := json.Marshal(&records)
		if err != nil {
			return nil, err
		}

		return jsonRecords, nil
	}
}

type ChangeRequest struct {
	Fqdn    string          `json:"fqdn"`
	Records json.RawMessage `json:"records"`
}
