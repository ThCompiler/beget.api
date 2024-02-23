// Package build implements builders of [github.com/ThCompiler/go.beget.api/api/dns.BasicRecords],
// [github.com/ThCompiler/go.beget.api/api/dns.DNSRecords], [github.com/ThCompiler/go.beget.api/api/dns.CNAMERecords]
// and [github.com/ThCompiler/go.beget.api/api/dns.NSRecords].
package build

import (
	"github.com/ThCompiler/go.beget.api/api/dns"
)

// RecordsCreator provides record creator that can create an array of records
// ([github.com/ThCompiler/go.beget.api/api/dns.ChangedRecord]).
type RecordsCreator interface {
	// create creates an array of records.
	create() []dns.ChangedRecord
}

// DNSRecordsCreator creates [github.com/ThCompiler/go.beget.api/api/dns.DNSRecords] as changed DNS-records for
// [github.com/ThCompiler/go.beget.api/api/dns.CallChangeRecords].
//
// Example:
//
//	changeRecords := dns.CallChangeRecords("domain.net",
//		build.NewDNSRecordsCreator().
//			AddDNSRecords(
//				build.NewDNSRecords().
//					AddRecord(10, "ns1.server.com").
//					AddRecord(20, "ns2.server.com"),
//			).AddDNSIPRecords(
//			build.NewDNSIPRecords().
//				AddRecord(10, "187.20.30.10").
//				AddRecord(20, "187.20.30.12"),
//		).Create(),
//	)
type DNSRecordsCreator struct {
	records *dns.DNSRecords
}

// NewDNSRecordsCreator creates new empty [DNSRecordsCreator].
func NewDNSRecordsCreator() *DNSRecordsCreator {
	return &DNSRecordsCreator{
		records: &dns.DNSRecords{},
	}
}

// AddDNSRecords adds an array of DNS-record to the changed DNS-records.
func (d *DNSRecordsCreator) AddDNSRecords(creator RecordsCreator) *DNSRecordsCreator {
	d.records.DNS = creator.create()

	return d
}

// AddDNSIPRecords adds an array of DNS_IP-record to the changed DNS-records.
func (d *DNSRecordsCreator) AddDNSIPRecords(creator RecordsCreator) *DNSRecordsCreator {
	d.records.DNSIP = creator.create()

	return d
}

// Create returns created [github.com/ThCompiler/go.beget.api/api/dns.DNSRecords].
func (d *DNSRecordsCreator) Create() *dns.DNSRecords {
	return d.records
}

// BasicRecordsCreator creates [github.com/ThCompiler/go.beget.api/api/dns.BasicRecords] as changed "A, MX, TXT"-records
// with DNS-records for [github.com/ThCompiler/go.beget.api/api/dns.CallChangeRecords].
//
// Example:
//
//	 changeRecords := dns.CallChangeRecords("domain.net",
//			build.NewBasicRecordsCreator().
//				AddDNSRecords(
//					build.NewDNSRecords().
//						AddRecord(10, "ns1.server.com").
//						AddRecord(20, "ns2.server.com"),
//				).AddDNSIPRecords(
//				build.NewDNSIPRecords().
//					AddRecord(10, "187.20.30.10").
//					AddRecord(20, "187.20.30.12"),
//			).AddARecords(
//				build.NewARecords().
//					AddRecord(10, "187.20.30.10"),
//			).AddAAAARecords(
//				build.NewAAAARecords().
//					AddRecord(10, "2001:0db8:11a3:09d7:1f34:8a2e:07a0:765d"),
//			).AddMXRecords(
//				build.NewMxRecords().
//					AddRecord(10, "mail.domain.com"),
//			).AddTXTRecords(
//				build.NewTxtRecords().
//					AddRecord(10, "This domain name is reserved for use in documentation"),
//			).Create(),
//		)
type BasicRecordsCreator struct {
	dnsRecords *DNSRecordsCreator
	records    *dns.BasicRecords
}

// NewBasicRecordsCreator creates new empty [BasicRecordsCreator].
func NewBasicRecordsCreator() *BasicRecordsCreator {
	return &BasicRecordsCreator{
		dnsRecords: NewDNSRecordsCreator(),
		records:    &dns.BasicRecords{},
	}
}

// AddDNSRecords adds an array of DNS-record to the changed DNS-records.
func (b *BasicRecordsCreator) AddDNSRecords(creator RecordsCreator) *BasicRecordsCreator {
	b.dnsRecords.AddDNSRecords(creator)

	return b
}

// AddDNSIPRecords adds an array of DNS_IP-record to the changed DNS-records.
func (b *BasicRecordsCreator) AddDNSIPRecords(creator RecordsCreator) *BasicRecordsCreator {
	b.dnsRecords.AddDNSIPRecords(creator)

	return b
}

// AddARecords adds an array of A-record to the changed A-records.
func (b *BasicRecordsCreator) AddARecords(creator RecordsCreator) *BasicRecordsCreator {
	b.records.A = creator.create()

	return b
}

// AddAAAARecords adds an array of AAAA-record to the changed AAAA-records.
func (b *BasicRecordsCreator) AddAAAARecords(creator RecordsCreator) *BasicRecordsCreator {
	b.records.AAAA = creator.create()

	return b
}

// AddMXRecords adds an array of MX-record to the changed MX-records.
func (b *BasicRecordsCreator) AddMXRecords(creator RecordsCreator) *BasicRecordsCreator {
	b.records.Mx = creator.create()

	return b
}

// AddTXTRecords adds an array of TXT-record to the changed TXT-records.
func (b *BasicRecordsCreator) AddTXTRecords(creator RecordsCreator) *BasicRecordsCreator {
	b.records.A = creator.create()

	return b
}

// Create returns created [github.com/ThCompiler/go.beget.api/api/dns.BasicRecords].
func (b *BasicRecordsCreator) Create() *dns.BasicRecords {
	b.records.DNSRecords = *b.dnsRecords.Create()

	return b.records
}

// NsRecordsCreator creates [github.com/ThCompiler/go.beget.api/api/dns.NSRecords] as changed NS-records
// with DNS-records for [github.com/ThCompiler/go.beget.api/api/dns.CallChangeRecords].
//
// Example:
//
//	 changeRecords := dns.CallChangeRecords("domain.net",
//			build.NewNsRecordsCreator().
//				AddDNSRecords(
//					build.NewDNSRecords().
//						AddRecord(10, "ns1.server.com").
//						AddRecord(20, "ns2.server.com"),
//				).AddDNSIPRecords(
//				build.NewDNSIPRecords().
//					AddRecord(10, "187.20.30.10").
//					AddRecord(20, "187.20.30.12"),
//			).AddNSRecords(
//				build.NewNsRecords().
//					AddRecord(10, "ns3.server.com"),
//			).Create(),
//		)
type NsRecordsCreator struct {
	dnsRecords *DNSRecordsCreator
	records    *dns.NSRecords
}

// NewNsRecordsCreator creates new empty [NsRecordsCreator].
func NewNsRecordsCreator() *NsRecordsCreator {
	return &NsRecordsCreator{
		dnsRecords: NewDNSRecordsCreator(),
		records:    &dns.NSRecords{},
	}
}

// AddDNSRecords adds an array of DNS-record to the changed DNS-records.
func (n *NsRecordsCreator) AddDNSRecords(creator RecordsCreator) *NsRecordsCreator {
	n.dnsRecords.AddDNSRecords(creator)

	return n
}

// AddDNSIPRecords adds an array of DNS_IP-record to the changed DNS-records.
func (n *NsRecordsCreator) AddDNSIPRecords(creator RecordsCreator) *NsRecordsCreator {
	n.dnsRecords.AddDNSIPRecords(creator)

	return n
}

func (n *NsRecordsCreator) AddNSRecords(creator RecordsCreator) *NsRecordsCreator {
	n.records.Ns = creator.create()

	return n
}

// Create returns created [github.com/ThCompiler/go.beget.api/api/dns.NSRecords].
func (n *NsRecordsCreator) Create() *dns.NSRecords {
	n.records.DNSRecords = *n.dnsRecords.Create()

	return n.records
}

// CNameRecordsCreator creates [github.com/ThCompiler/go.beget.api/api/dns.NSRecords] as changed CNAME-records
// with DNS-records for [github.com/ThCompiler/go.beget.api/api/dns.CallChangeRecords].
//
// Example:
//
//	 changeRecords := dns.CallChangeRecords("domain.net",
//			build.NewCNameRecordsCreator().
//				AddDNSRecords(
//					build.NewDNSRecords().
//						AddRecord(10, "ns1.server.com").
//						AddRecord(20, "ns2.server.com"),
//				).AddDNSIPRecords(
//				build.NewDNSIPRecords().
//					AddRecord(10, "187.20.30.10").
//					AddRecord(20, "187.20.30.12"),
//			).AddCNameRecords(
//				build.NewCNameRecords().
//					AddRecord(10, "alias.domain.com"),
//			).Create(),
//		)
type CNameRecordsCreator struct {
	dnsRecords *DNSRecordsCreator
	records    *dns.CNAMERecords
}

// NewCNameRecordsCreator creates new empty [CNameRecordsCreator].
func NewCNameRecordsCreator() *CNameRecordsCreator {
	return &CNameRecordsCreator{
		dnsRecords: NewDNSRecordsCreator(),
		records:    &dns.CNAMERecords{},
	}
}

// AddDNSRecords adds an array of DNS-record to the changed DNS-records.
func (c *CNameRecordsCreator) AddDNSRecords(creator RecordsCreator) *CNameRecordsCreator {
	c.dnsRecords.AddDNSRecords(creator)

	return c
}

// AddDNSIPRecords adds an array of DNS_IP-record to the changed DNS-records.
func (c *CNameRecordsCreator) AddDNSIPRecords(creator RecordsCreator) *CNameRecordsCreator {
	c.dnsRecords.AddDNSIPRecords(creator)

	return c
}

func (c *CNameRecordsCreator) AddCNameRecords(creator RecordsCreator) *CNameRecordsCreator {
	c.records.CName = creator.create()

	return c
}

// Create returns created [github.com/ThCompiler/go.beget.api/api/dns.CNAMERecords].
func (c *CNameRecordsCreator) Create() *dns.CNAMERecords {
	c.records.DNSRecords = *c.dnsRecords.Create()

	return c.records
}
