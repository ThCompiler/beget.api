package set

import "github.com/pkg/errors"

func errorWithNumberRecords(err error, numberRecords int) error {
	return errors.Wrapf(err, "with number records %d", numberRecords)
}

const (
	MaxDNSRecords   = 4
	MaxBasicRecords = 10
	MaxCNAMERecords = 1
	MaxNSRecords    = 10
)

type ChangedRecord struct {
	Value    string `json:"value"`
	Priority int64  `json:"priority"`
}

type DNSRecords struct {
	DNS   []ChangedRecord `json:"DNS,omitempty"`
	DNSIP []ChangedRecord `json:"DNS_IP,omitempty"`
}

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

type BasicRecords struct {
	DNSRecords
	A    []ChangedRecord `json:"A,omitempty"`
	AAAA []ChangedRecord `json:"AAAA,omitempty"`
	Mx   []ChangedRecord `json:"MX,omitempty"`
	Txt  []ChangedRecord `json:"TXT,omitempty"`
}

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

type NSRecords struct {
	DNSRecords
	Ns []ChangedRecord `json:"NS,omitempty"`
}

func (n *NSRecords) Validate() error {
	if len(n.Ns) > MaxNSRecords {
		return errorWithNumberRecords(ErrTooMuchNsRecords, len(n.Ns))
	}

	return n.DNSRecords.Validate()
}

type CNAMERecords struct {
	DNSRecords
	CName []ChangedRecord `json:"CNAME,omitempty"`
}

func (c *CNAMERecords) Validate() error {
	if len(c.CName) > MaxCNAMERecords {
		return errorWithNumberRecords(ErrTooMuchCNameRecords, len(c.CName))
	}

	return c.DNSRecords.Validate()
}
