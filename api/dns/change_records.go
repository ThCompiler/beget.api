package dns

import (
	"encoding/json"
	"github.com/ThCompiler/go.beget.api/api"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
	"net/http"
	"reflect"

	"github.com/pkg/errors"
)

type changeRecords struct {
	api.BasicMethod
}

// SettableRecords provides records that can be applied on the DNS server for a specific domain.
type SettableRecords interface {
	// Validate indicates whether the entries are correct from the API's point of view.
	// If the entries are not correct, the method should return an error, otherwise "nil".
	Validate() error
}

// CallChangeRecords is a creation function that returns a [core.APIMethod] corresponding to the method [changeRecords].
// The function expects the domain name for which DNS records need to be changed,
// and the modified set of records that need to be applied.
//
// Three types of records are supported:
//  1. "A, MX, TXT" records ([BasicRecords]).
//  2. NS-records ([NSRecords]).
//  3. CNAME-records ([CNAMERecords]).
//
// Also, each record can contain DNS records ([DNSRecords]).
//
// # Important
//
// The beget system replaces the current DNS server records with the ones passed in the request.
// The old records will be lost.
//
// If you need to change only one record for a domain,
// you should first get full information about the records of this DNS server domain using the [CallGetData] method.
// Then change the necessary information in the received records.
// And pass all this as the body for the change request.
//
// # Noticed
//
// If the domain contains SRV records, it is not possible to change them by [changeRecords] method,
// and if you attempt to change other records for this domain by [changeRecords] method, SRV records will be deleted.
//
// [changeRecords]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#changerecords
func CallChangeRecords(domainName string, records SettableRecords) core.APIMethod[result.BoolResult] {
	request, err := prepareRecords(records, reflect.TypeOf(records).String())
	if err != nil {
		return &getData{
			BasicMethod: *api.CallError(err),
		}
	}

	return &changeRecords{
		BasicMethod: *api.CallMethod(ChangeRecordsMethodPath, &changeRequest{Fqdn: domainName, Records: request}),
	}
}

// GetHTTPMethod returns name of http method for method [changeRecords].
//
// [changeRecords]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#changerecords
func (*changeRecords) GetHTTPMethod() string {
	return http.MethodPost
}

// GetName returns name of method [changeRecords].
//
// [changeRecords]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#changerecords
func (*changeRecords) GetName() core.MethodName {
	return ChangeRecordsMethodName
}

// prepareRecords converts the modified set of domain records to json format, if possible.
func prepareRecords(records SettableRecords, nameRecords string) (json.RawMessage, error) {
	if err := records.Validate(); err != nil {
		return nil, errors.Wrapf(err, "try create request with %s", nameRecords)
	}

	jsonRecords, err := json.Marshal(&records)
	if err != nil {
		return nil, err
	}

	return jsonRecords, nil
}

// changeRequest represents request body for [changeRecords] method.
//
// [changeRecords]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#changerecords
type changeRequest struct {
	Fqdn    string          `json:"fqdn"`
	Records json.RawMessage `json:"records"`
}

//----------------------------------------------------------------------------------------------------------------------
// Records
//----------------------------------------------------------------------------------------------------------------------

// errorWithNumberRecords creates a new error about exceeding the limit on the number of records.
func errorWithNumberRecords(err error, numberRecords int) error {
	return errors.Wrapf(err, "with number records %d", numberRecords)
}

// Limits on the number of specific records.
const (
	MaxDNSRecords   = 4  // limit for DNS-records.
	MaxBasicRecords = 10 // limit for "A, MX, TXT" records.
	MaxCNAMERecords = 1  // limit for CNAME-records.
	MaxNSRecords    = 10 // limit for NS-records.
)

// ChangedRecord represents a single changed record.
type ChangedRecord struct {
	Value    string `json:"value"`
	Priority int64  `json:"priority"`
}

// DNSRecords represents a changed DNS-records.
// If the DNS servers are not proprietary (i.e. they are not located on one of the subdomains of the main domain),
// then the DNS_IP section can be omitted.
type DNSRecords struct {
	DNS   []ChangedRecord `json:"DNS,omitempty"`    // DNS records.
	DNSIP []ChangedRecord `json:"DNS_IP,omitempty"` // DNS_IP records.
}

// Validate checks records for limiting their number
// and for matching the number of DNS records to the number of DNS_IP records (if the latter are not omitted).
//
// Can return errors:
//   - [ErrTooMuchDNSRecords].
//   - [ErrTooMuchDNSIPRecords].
//   - [ErrNumberDNSRecordsNotEqual].
func (d *DNSRecords) Validate() error {
	switch {
	case len(d.DNS) > MaxDNSRecords:
		return errorWithNumberRecords(ErrTooMuchDNSRecords, len(d.DNS))
	case len(d.DNSIP) > MaxDNSRecords:
		return errorWithNumberRecords(ErrTooMuchDNSIPRecords, len(d.DNSIP))
	case len(d.DNSIP) != len(d.DNS) && len(d.DNSIP) != 0:
		return ErrNumberDNSRecordsNotEqual
	}

	return nil
}

// BasicRecords represents a changed "A, MX, TXT" records.
// There is no information about AAAA-records in the documentation of the Beget.API,
// but if this field is set, the requests are executed successfully
type BasicRecords struct {
	DNSRecords
	A    []ChangedRecord `json:"A,omitempty"`    // A records.
	AAAA []ChangedRecord `json:"AAAA,omitempty"` // AAAA records.
	Mx   []ChangedRecord `json:"MX,omitempty"`   // MX records.
	Txt  []ChangedRecord `json:"TXT,omitempty"`  // TXT records.
}

// Validate checks records for limiting their number and checks embedded [DNSRecords].
//
// Can return errors:
//   - [ErrTooMuchARecords].
//   - [ErrTooMuchAAAARecords].
//   - [ErrTooMuchMxRecords].
//   - [ErrTooMuchTxtRecords].
//   - [ErrTooMuchDNSRecords].
//   - [ErrTooMuchDNSIPRecords].
//   - [ErrNumberDNSRecordsNotEqual].
func (b *BasicRecords) Validate() error {
	switch {
	case len(b.A) > MaxBasicRecords:
		return errorWithNumberRecords(ErrTooMuchARecords, len(b.A))
	case len(b.AAAA) > MaxBasicRecords:
		return errorWithNumberRecords(ErrTooMuchAAAARecords, len(b.AAAA))
	case len(b.Mx) > MaxBasicRecords:
		return errorWithNumberRecords(ErrTooMuchMxRecords, len(b.Mx))
	case len(b.Txt) > MaxBasicRecords:
		return errorWithNumberRecords(ErrTooMuchTxtRecords, len(b.Txt))
	}

	return b.DNSRecords.Validate()
}

// NSRecords represents a changed NS-records.
type NSRecords struct {
	DNSRecords
	Ns []ChangedRecord `json:"NS,omitempty"` // NS records.
}

// Validate checks records for limiting their number and checks embedded [DNSRecords].
//
// Can return errors:
//   - [ErrTooMuchNsRecords].
//   - [ErrTooMuchDNSRecords].
//   - [ErrTooMuchDNSIPRecords].
//   - [ErrNumberDNSRecordsNotEqual].
func (n *NSRecords) Validate() error {
	if len(n.Ns) > MaxNSRecords {
		return errorWithNumberRecords(ErrTooMuchNsRecords, len(n.Ns))
	}

	return n.DNSRecords.Validate()
}

// CNAMERecords represents a changed CNAME-records.
type CNAMERecords struct {
	DNSRecords
	CName []ChangedRecord `json:"CNAME,omitempty"` // CNAME records.
}

// Validate checks records for limiting their number and checks embedded [DNSRecords].
//
// Can return errors:
//   - [ErrTooMuchCNameRecords].
//   - [ErrTooMuchDNSRecords].
//   - [ErrTooMuchDNSIPRecords].
//   - [ErrNumberDNSRecordsNotEqual].
func (c *CNAMERecords) Validate() error {
	if len(c.CName) > MaxCNAMERecords {
		return errorWithNumberRecords(ErrTooMuchCNameRecords, len(c.CName))
	}

	return c.DNSRecords.Validate()
}
