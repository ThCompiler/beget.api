package test

import (
	"bytes"
	"encoding/json"
	"github.com/ThCompiler/go.beget.api/api/dns"
	"github.com/ThCompiler/go.beget.api/api/dns/general"
	"github.com/ThCompiler/go.beget.api/api/dns/record"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
	"github.com/ThCompiler/ts"
	"github.com/stretchr/testify/require"
	"net/url"
	"reflect"
	"testing"
	"text/template"
)

type GetData struct {
	IsUnderControl bool                 // reports whether Beget service the domain.
	IsBegetDNS     bool                 // reports whether the domain is located on the Beget DNS servers.
	IsSubdomain    bool                 // reports whether the domain is subdomain.
	Fqdn           string               // the domain name passed in the request.
	NsRecords      *result.NSRecords    // NS records.
	BasicRecords   *result.BasicRecords // "A, MX, TXT" records.
	CnameRecords   *result.CNAMERecords // CNAME records.
	TypeRecords    result.TypeRecords   // type of DNS-records for the domain.
}

func createByTemplate(t *testing.T, templateName string, templateData string, data any) string {
	var out string
	buf := bytes.NewBufferString(out)

	temp, err := template.New(templateName).Parse(templateData)
	require.NoError(t, err)

	err = temp.Execute(buf, data)
	require.NoError(t, err)

	return buf.String()
}

type getDataRequest struct {
	Fqdn string `json:"fqdn"`
}

func (ap *APISuite) TestGetData() {
	domainName := "domain"
	client := core.Client{
		Login:    "login",
		Password: "password",
	}

	ap.RunTest(
		func(response string) GetData {
			ap.server.DNSGetData(response, func(values url.Values) {
				RequireEqualValues(ap.T(), url.Values{
					"input_format": []string{string(core.JSON)},
				}, values, client, "input_data")

				require.True(ap.T(), values.Has("input_data"))
				request := values.Get("input_data")

				var req getDataRequest
				require.NoError(ap.T(), json.Unmarshal([]byte(request), &req))

				require.Equal(ap.T(), getDataRequest{Fqdn: domainName}, req)
			})

			req, err := core.PrepareRequest[result.GetData](
				client,
				dns.CallGetData(domainName),
			)
			require.NoError(ap.T(), err)

			resp, err := req.Do()
			require.NoError(ap.T(), err)

			answ, err := resp.Get()
			require.NoError(ap.T(), err)

			res, err := answ.Get()
			require.NoError(ap.T(), err)

			return GetData{
				IsUnderControl: res.IsUnderControl(),
				IsBegetDNS:     res.IsBegetDNS(),
				IsSubdomain:    res.IsSubdomain(),
				Fqdn:           res.Fqdn(),
				NsRecords:      res.NSRecords(),
				BasicRecords:   res.BasicRecords(),
				CnameRecords:   res.CNAMERecords(),
				TypeRecords:    res.TypeRecords(),
			}
		},
		ts.TestCase{
			Name:     "BasicRecords",
			Args:     ts.ToTestArgs(createByTemplate(ap.T(), "basic", getDataResponse, getDataResultBasic)),
			Expected: ts.TTVE(getDataResultBasic),
		},
		ts.TestCase{
			Name:     "NsRecords",
			Args:     ts.ToTestArgs(createByTemplate(ap.T(), "ns", getDataResponse, getDataResultNs)),
			Expected: ts.TTVE(getDataResultNs),
		},
		ts.TestCase{
			Name:     "CnameRecords",
			Args:     ts.ToTestArgs(createByTemplate(ap.T(), "cname", getDataResponse, getDataResultCname)),
			Expected: ts.TTVE(getDataResultCname),
		},
	)
}

type changeRecordsRequest struct {
	Fqdn    string              `json:"fqdn"`
	Records dns.SettableRecords `json:"records"`
}

func (ap *APISuite) TestChangeRecords() {
	domainName := "domain"
	client := core.Client{
		Login:    "login",
		Password: "password",
	}

	ap.RunTest(
		func(records dns.SettableRecords) bool {
			ap.server.DNSChangeRecords(changeRecordsResponse, func(values url.Values) {
				RequireEqualValues(ap.T(), url.Values{
					"input_format": []string{string(core.JSON)},
				}, values, client, "input_data")

				require.True(ap.T(), values.Has("input_data"))
				request := values.Get("input_data")

				req := changeRecordsRequest{
					Fqdn: "",
					Records: reflect.
						New(reflect.ValueOf(records).Elem().Type()).
						Interface().(dns.SettableRecords),
				}
				require.NoError(ap.T(), json.Unmarshal([]byte(request), &req))

				require.Equal(ap.T(), changeRecordsRequest{Fqdn: domainName, Records: records}, req)
			})

			req, err := core.PrepareRequest[result.BoolResult](
				client,
				dns.CallChangeRecords(domainName, records),
			)
			require.NoError(ap.T(), err)

			resp, err := req.Do()
			require.NoError(ap.T(), err)

			answ, err := resp.Get()
			require.NoError(ap.T(), err)

			res, err := answ.Get()
			require.NoError(ap.T(), err)

			return bool(*res)
		},
		ts.TestCase{
			Name:     "BasicRecords",
			Args:     ts.ToTestArgs(&changeRecordsBasicRequest),
			Expected: ts.TTVE(true),
		},
		ts.TestCase{
			Name:     "NsRecords",
			Args:     ts.ToTestArgs(&changeRecordsNSRequest),
			Expected: ts.TTVE(true),
		},
		ts.TestCase{
			Name:     "CnameRecords",
			Args:     ts.ToTestArgs(&changeRecordsCNAMERequest),
			Expected: ts.TTVE(true),
		},
	)
}

var getDataResultBasic = GetData{
	IsUnderControl: false,
	IsBegetDNS:     false,
	IsSubdomain:    true,
	Fqdn:           "some.domain.ru",
	NsRecords:      nil,
	BasicRecords: &result.BasicRecords{
		A: []record.ARecord{
			{TTLRecord: record.TTLRecord{TTL: 600}, ARecord: general.ARecord{Address: "195.19.34.5"}},
			{TTLRecord: record.TTLRecord{TTL: 300}, ARecord: general.ARecord{Address: "195.19.34.6"}},
		},
		AAAA: []record.AAAARecord{
			{TTLRecord: record.TTLRecord{TTL: 300},
				AAAARecord: general.AAAARecord{Address: "2001:0db8:11a3:09d7:1f34:8a2e:07a0:765d"}},
			{TTLRecord: record.TTLRecord{TTL: 300},
				AAAARecord: general.AAAARecord{Address: "2001:0db8:11a3:09d7:1f34:8a2e:07a0:765e"}},
		},
		CAA: []record.CAARecord{
			{TTLRecord: record.TTLRecord{TTL: 300},
				CAARecord: general.CAARecord{Flags: 1, Tag: general.Issue, Value: "uy.ru"}},
			{TTLRecord: record.TTLRecord{TTL: 300},
				CAARecord: general.CAARecord{Flags: 1, Tag: general.Issue, Value: "uy2.ru"}},
		},
		Mx: []record.MXRecord{
			{TTLRecord: record.TTLRecord{TTL: 300}, MXRecord: general.MXRecord{Preference: 10, Exchange: "mail.ru"}},
			{TTLRecord: record.TTLRecord{TTL: 300}, MXRecord: general.MXRecord{Preference: 20, Exchange: "ya.ru"}},
		},
		Txt: []record.TXTRecord{
			{TTLRecord: record.TTLRecord{TTL: 300}, TXTRecord: general.TXTRecord{TxtData: "hello"}},
			{TTLRecord: record.TTLRecord{TTL: 300}, TXTRecord: general.TXTRecord{TxtData: "hello2"}},
		},
		Srv: []record.SRVRecord{
			{TTLRecord: record.TTLRecord{TTL: 300},
				SRVRecord: general.SRVRecord{Priority: 1, Weight: 1, Port: 22, Target: "domain.ru"}},
			{TTLRecord: record.TTLRecord{TTL: 300},
				SRVRecord: general.SRVRecord{Priority: 2, Weight: 1, Port: 34, Target: "domain2.ru"}},
		},
		DNSRecords: result.DNSRecords{
			DNS: []record.DNSRecord{
				{Value: "ns3.domain.com"},
				{Value: "ns3.domain.pro"},
			},
			DNSIP: []record.DNSIPRecord{
				{Value: "1"},
				{Value: "2"},
			},
		},
	},
	CnameRecords: nil,
	TypeRecords:  result.Basic,
}

var getDataResultNs = GetData{
	IsUnderControl: false,
	IsBegetDNS:     false,
	IsSubdomain:    true,
	Fqdn:           "some.domain.ru",
	NsRecords: &result.NSRecords{
		NSs: []record.NSRecord{
			{TTLRecord: record.TTLRecord{TTL: 600}, NSRecord: general.NSRecord{NsdName: "ns1.kul.ru"}},
			{TTLRecord: record.TTLRecord{TTL: 600}, NSRecord: general.NSRecord{NsdName: "ns2.kul.ru"}},
		},
		DNSRecords: result.DNSRecords{
			DNS: []record.DNSRecord{
				{Value: "ns3.domain.com"},
				{Value: "ns3.domain.pro"},
			},
			DNSIP: []record.DNSIPRecord{
				{Value: "1"},
				{Value: "2"},
			},
		},
	},
	BasicRecords: nil,
	CnameRecords: nil,
	TypeRecords:  result.NS,
}

var getDataResultCname = GetData{
	IsUnderControl: false,
	IsBegetDNS:     false,
	IsSubdomain:    true,
	Fqdn:           "some.domain.ru",
	NsRecords:      nil,
	BasicRecords:   nil,
	CnameRecords: &result.CNAMERecords{
		CNames: []record.CNAMERecord{
			{TTLRecord: record.TTLRecord{TTL: 600}, CNAMERecord: general.CNAMERecord{Cname: "domain.ru"}},
		},
		DNSRecords: result.DNSRecords{
			DNS: []record.DNSRecord{
				{Value: "ns3.domain.com"},
				{Value: "ns3.domain.pro"},
			},
			DNSIP: []record.DNSIPRecord{
				{Value: "1"},
				{Value: "2"},
			},
		},
	},
	TypeRecords: result.CNAME,
}

const getDataResponse = `
{{ define "DNS" }}
				"DNS": [
					{{ range $index, $element := .DNS }}{{ if $index }},{{end}}
					{
						"value": "{{ $element.Value }}"
					}
					{{ end }}
				],
				"DNS_IP": [
					{{ range $index, $element := .DNSIP }}{{ if $index }},{{end}}
					{
						"value": "{{ $element.Value }}"
					}
					{{ end }}
				]
{{ end }}
{
    "status": "success",
    "answer": {
        "status": "success",
        "result": {
            "is_under_control": {{ .IsUnderControl }},
            "is_beget_dns": {{ .IsBegetDNS }},
            "is_subdomain": {{ .IsSubdomain }},
            "fqdn": "{{ .Fqdn }}",
            "records": { {{ with .BasicRecords }}
                "A": [
					{{ range $index, $element := .A }}{{ if $index }},{{end}}
					{
						"ttl": {{ $element.TTL }},
						"address": "{{ $element.Address }}"
					}
					{{ end }}
                ],
                "AAAA": [
                    {{ range $index, $element := .AAAA }}{{ if $index }},{{end}}
					{
						"ttl": {{ $element.TTL }},
						"address": "{{ $element.Address }}"
					}
					{{ end }}
                ],
                "CAA": [
					{{ range $index, $element := .CAA }}{{ if $index }},{{end}}
					{
						"ttl": {{ $element.TTL }},
						"flags": {{ $element.Flags }},
						"tag": "{{ $element.Tag }}",
						"value": "{{ $element.Value }}"
					}
					{{ end }}
                ],
                "MX": [
					{{ range $index, $element := .Mx }}{{ if $index }},{{end}}
					{
						"ttl": {{ $element.TTL }},
						"exchange": "{{ $element.Exchange }}",
						"preference": {{ $element.Preference }}
					}
					{{ end }}
                ],
                "SRV": [
					{{ range $index, $element := .Srv }}{{ if $index }},{{end}}
					{
						"ttl": {{ $element.TTL }},
						"priority": {{ $element.Priority }},
						"weight": {{ $element.Weight }},
						"port": {{ $element.Port }},
						"target": "{{ $element.Target }}"
					}
					{{ end }}
                ],
                "TXT": [
					{{ range $index, $element := .Txt }}{{ if $index }},{{end}}
					{
						"ttl": {{ $element.TTL }},
						"txtdata": "{{ $element.TxtData }}"
					}
					{{ end }}
                ],
				{{ template "DNS" .DNSRecords }}
				{{ end }}{{ with .CnameRecords }}
                "CNAME": [
					{{ range $index, $element := .CNames }}{{ if $index }},{{end}}
					{
						"ttl": {{ $element.TTL }},
						"cname": "{{ $element.Cname }}"
					}
					{{ end }}
                ],
				{{ template "DNS" .DNSRecords }}
				{{ end }} {{ with .NsRecords }}
                "NS": [
					{{ range $index, $element := .NSs }}{{ if $index }},{{end}}
					{
                        "ttl": {{ $element.TTL }},
                        "nsdname": "{{ $element.NsdName }}"
                    }
					{{ end }}
                ],
				{{ template "DNS" .DNSRecords }}
				{{ end }}
            },
            "set_type": {{ .TypeRecords }}
        }
    }
}
`

var changeRecordsBasicRequest = dns.BasicRecords{
	A: []dns.ChangedRecord{
		{Value: "192.165.1.3", Priority: 10},
		{Value: "192.165.1.5", Priority: 2},
	},
	AAAA: []dns.ChangedRecord{
		{Value: "2001:0db8:11a3:09d7:1f34:8a2e:07a0:765d", Priority: 10},
		{Value: "2001:0db8:11a3:09d7:1f34:8a2e:07a0:765d", Priority: 2},
	},
	Mx: []dns.ChangedRecord{
		{Value: "mail.ru", Priority: 10},
		{Value: "ya.ru", Priority: 2},
	},
	Txt: []dns.ChangedRecord{
		{Value: "hello", Priority: 10},
		{Value: "hello 2", Priority: 2},
	},
	DNSRecords: dns.DNSRecords{
		DNS: []dns.ChangedRecord{
			{Value: "ns.mail.ru", Priority: 10},
			{Value: "ns.ya.ru", Priority: 2},
		},
		DNSIP: []dns.ChangedRecord{
			{Value: "192.165.1.3", Priority: 10},
			{Value: "192.165.1.5", Priority: 2},
		},
	},
}

var changeRecordsNSRequest = dns.NSRecords{
	Ns: []dns.ChangedRecord{
		{Value: "a.mail.ru", Priority: 10},
		{Value: "a.ya.ru", Priority: 2},
	},
	DNSRecords: dns.DNSRecords{
		DNS: []dns.ChangedRecord{
			{Value: "ns.mail.ru", Priority: 10},
			{Value: "ns.ya.ru", Priority: 2},
		},
		DNSIP: []dns.ChangedRecord{
			{Value: "192.165.1.3", Priority: 10},
			{Value: "192.165.1.5", Priority: 2},
		},
	},
}

var changeRecordsCNAMERequest = dns.CNAMERecords{
	CName: []dns.ChangedRecord{
		{Value: "mail.ru", Priority: 10},
	},
	DNSRecords: dns.DNSRecords{
		DNS: []dns.ChangedRecord{
			{Value: "ns.mail.ru", Priority: 10},
			{Value: "ns.ya.ru", Priority: 2},
		},
		DNSIP: []dns.ChangedRecord{
			{Value: "192.165.1.3", Priority: 10},
			{Value: "192.165.1.5", Priority: 2},
		},
	},
}

const changeRecordsResponse = `
{
    "status": "success",
    "answer": {
        "status": "success",
        "result": true
    }
}
`
