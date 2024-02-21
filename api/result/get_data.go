package result

import (
	"encoding/json"
	"github.com/ThCompiler/go.beget.api/api/dns/record"
	"github.com/pkg/errors"
)

// The names json fields of the [getData] response.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
const (
	isUnderControlField = "is_under_control"
	isBegetDNSField     = "is_beget_dns"
	isSubdomainField    = "is_subdomain"
	fqdnField           = "fqdn"
	setTypeField        = "set_type"
	recordsField        = "records"
)

// TypeRecords is a type, that specifies the type of the records in the [getData] response.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type TypeRecords int64

const (
	Basic = TypeRecords(1) // type of "A, MX, TXT" records.
	NS    = TypeRecords(2) // type of NS-records.
	CNAME = TypeRecords(3) // type of CNAME-records.
)

// DNSRecords is a representation of the DNS records information in the [getData] response.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type DNSRecords struct {
	DNS   []record.DNSRecord   `json:"DNS"`              // DNS records.
	DNSIP []record.DNSIPRecord `json:"DNS_IP,omitempty"` // DNS_IP records.
}

// NSRecords is a representation of the NS records information in the [getData] response.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type NSRecords struct {
	DNSRecords
	NSs []record.NSRecord `json:"NS,omitempty"` // NS records.
}

// CNAMERecords is a representation of the CNAME records information in the [getData] response.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type CNAMERecords struct {
	DNSRecords
	CNames []record.CNAMERecord `json:"CNAME,omitempty"` // CNAME records.
}

// BasicRecords is a representation of the "A, MX, TXT" records information in the [getData] response.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type BasicRecords struct {
	DNSRecords
	A    []record.ARecord    `json:"A,omitempty"`    // A records.
	AAAA []record.AAAARecord `json:"AAAA,omitempty"` // AAAA records.
	CAA  []record.CAARecord  `json:"CAA,omitempty"`  // CAA records.
	Mx   []record.MXRecord   `json:"MX,omitempty"`   // MX records.
	Txt  []record.TXTRecord  `json:"TXT,omitempty"`  // TXT records.
	Srv  []record.SRVRecord  `json:"SRV,omitempty"`  // SRV records.
}

// GetData is the result of a successful call to the [getData] method.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type GetData struct {
	isUnderControl bool          // reports whether Beget service the domain.
	isBegetDNS     bool          // reports whether the domain is located on the Beget DNS servers.
	isSubdomain    bool          // reports whether the domain is subdomain.
	fqdn           string        // the domain name passed in the request.
	nsRecords      *NSRecords    // NS records.
	basicRecords   *BasicRecords // "A, MX, TXT" records.
	cnameRecords   *CNAMERecords // CNAME records.
	typeRecords    TypeRecords   // type of DNS-records for the domain.
}

// IsUnderControl reports whether Beget service the domain.
func (gdr *GetData) IsUnderControl() bool {
	return gdr.isUnderControl
}

// IsBegetDNS reports whether the domain is located on the Beget DNS servers.
func (gdr *GetData) IsBegetDNS() bool {
	return gdr.isBegetDNS
}

// IsSubdomain reports whether the domain is subdomain.
func (gdr *GetData) IsSubdomain() bool {
	return gdr.isSubdomain
}

// Fqdn returns the domain name passed in the request to [getData] method.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (gdr *GetData) Fqdn() string {
	return gdr.fqdn
}

// TypeRecords returns type of DNS-records for the domain.
func (gdr *GetData) TypeRecords() TypeRecords {
	return gdr.typeRecords
}

// BasicRecords returns "A, MX, TXT" records if they are provided in the response.
// Otherwise, it returns nil
func (gdr *GetData) BasicRecords() *BasicRecords {
	return gdr.basicRecords
}

// CNAMERecords returns CNAME records if they are provided in the response.
// Otherwise, it returns nil
func (gdr *GetData) CNAMERecords() *CNAMERecords {
	return gdr.cnameRecords
}

// NSRecords returns NS records if they are provided in the response.
// Otherwise, it returns nil
func (gdr *GetData) NSRecords() *NSRecords {
	return gdr.nsRecords
}

func unmarshalField(field string, fieldMaps map[string]json.RawMessage, out any) error {
	jsonField, ok := fieldMaps[field]
	if !ok {
		return errors.Wrapf(ErrNotFoundField, "with field %s", field)
	}

	if err := json.Unmarshal(jsonField, &out); err != nil {
		return errors.Wrap(err, "can't not get field '"+field+"' for get data result")
	}

	return nil
}

func (gdr *GetData) unmarshalSimpleFields(resultMap map[string]json.RawMessage) error {
	if err := unmarshalField(isUnderControlField, resultMap, &gdr.isUnderControl); err != nil {
		return err
	}

	if err := unmarshalField(isBegetDNSField, resultMap, &gdr.isBegetDNS); err != nil {
		return err
	}

	if err := unmarshalField(isSubdomainField, resultMap, &gdr.isSubdomain); err != nil {
		return err
	}

	if err := unmarshalField(fqdnField, resultMap, &gdr.fqdn); err != nil {
		return err
	}

	return unmarshalField(setTypeField, resultMap, &gdr.typeRecords)
}

// UnmarshalJSON is functions for [encoding/json] to unmarshal [getData] response from json format.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (gdr *GetData) UnmarshalJSON(data []byte) error {
	*gdr = GetData{}

	var resultMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &resultMap); err != nil {
		return errors.Wrap(err, "can't not get fields for get data result")
	}

	if err := gdr.unmarshalSimpleFields(resultMap); err != nil {
		return err
	}

	var readRecord any

	switch gdr.typeRecords {
	case Basic:
		gdr.basicRecords = &BasicRecords{}
		readRecord = gdr.basicRecords
	case NS:
		gdr.nsRecords = &NSRecords{}
		readRecord = gdr.nsRecords
	case CNAME:
		gdr.cnameRecords = &CNAMERecords{}
		readRecord = gdr.cnameRecords
	}

	if err := unmarshalField(recordsField, resultMap, readRecord); err != nil {
		if errors.Is(err, errors.Unwrap(err)) {
			return err
		}

		return errors.Wrapf(err, "getted fields %v not for type records %d", gdr.typeRecords, resultMap["records"])
	}

	return nil
}
