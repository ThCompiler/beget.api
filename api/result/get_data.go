package result

import (
	"encoding/json"
	"github.com/ThCompiler/go.beget.api/api/dns/record"
	"github.com/pkg/errors"
)

const (
	isUnderControlField = "is_under_control"
	isBegetDNSField     = "is_beget_dns"
	isSubdomainField    = "is_subdomain"
	fqdnField           = "fqdn"
	setTypeField        = "set_type"
	recordsField        = "records"
)

type TypeRecords int64

const (
	Basic = TypeRecords(1)
	NS    = TypeRecords(2)
	CNAME = TypeRecords(3)
)

type DNSRecords struct {
	DNS   []record.DNSRecord   `json:"DNS"`
	DNSIP []record.DNSIPRecord `json:"DNS_IP"`
}

type NSRecords struct {
	DNSRecords
	NSs []record.NSRecord `json:"NS,omitempty"`
}

type CNAMERecords struct {
	DNSRecords
	CNames []record.CNAMERecord `json:"CNAME,omitempty"`
}

type BasicRecords struct {
	DNSRecords
	A    []record.ARecord    `json:"A,omitempty"`
	AAAA []record.AAAARecord `json:"AAAA,omitempty"`
	CAA  []record.CAARecord  `json:"CAA,omitempty"`
	Mx   []record.MXRecord   `json:"MX,omitempty"`
	Txt  []record.TXTRecord  `json:"TXT,omitempty"`
	Srv  []record.SRVRecord  `json:"SRV,omitempty"`
}

type GetData struct {
	isUnderControl bool
	isBegetDNS     bool
	isSubdomain    bool
	fqdn           string
	nsRecords      *NSRecords
	basicRecords   *BasicRecords
	cnameRecords   *CNAMERecords
	typeRecords    TypeRecords
}

func (gdr *GetData) IsUnderControl() bool {
	return gdr.isUnderControl
}

func (gdr *GetData) IsBegetDNS() bool {
	return gdr.isBegetDNS
}

func (gdr *GetData) IsSubdomain() bool {
	return gdr.isSubdomain
}

func (gdr *GetData) Fqdn() string {
	return gdr.fqdn
}

func (gdr *GetData) TypeRecords() TypeRecords {
	return gdr.typeRecords
}

func (gdr *GetData) BasicRecords() *BasicRecords {
	return gdr.basicRecords
}

func (gdr *GetData) CNAMERecords() *CNAMERecords {
	return gdr.cnameRecords
}

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
