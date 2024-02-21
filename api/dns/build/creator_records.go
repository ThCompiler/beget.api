// Package build implements builders of [github.com/ThCompiler/go.beget.api/api/dns.BasicRecords],
// [github.com/ThCompiler/go.beget.api/api/dns.DNSRecords], [github.com/ThCompiler/go.beget.api/api/dns.CNAMERecords]
// and [github.com/ThCompiler/go.beget.api/api/dns.NSRecords].
package build

import (
	"github.com/ThCompiler/go.beget.api/api/dns"
)

type RecordsCreator interface {
	create() []dns.ChangedRecord
}

type DNSRecordsCreator struct {
	records *dns.DNSRecords
}

func NewDNSRecordsCreator() *DNSRecordsCreator {
	return &DNSRecordsCreator{
		records: &dns.DNSRecords{},
	}
}

func (d *DNSRecordsCreator) AddDNSRecords(creator RecordsCreator) *DNSRecordsCreator {
	d.records.DNS = creator.create()

	return d
}

func (d *DNSRecordsCreator) AddDNSIPRecords(creator RecordsCreator) *DNSRecordsCreator {
	d.records.DNSIP = creator.create()

	return d
}

func (d *DNSRecordsCreator) Create() *dns.DNSRecords {
	return d.records
}

type BasicRecordsCreator struct {
	DNSRecordsCreator
	records *dns.BasicRecords
}

func NewBasicRecordsCreator() *BasicRecordsCreator {
	return &BasicRecordsCreator{
		DNSRecordsCreator: *NewDNSRecordsCreator(),
		records:           &dns.BasicRecords{},
	}
}

func (b *BasicRecordsCreator) AddARecords(creator RecordsCreator) *BasicRecordsCreator {
	b.records.A = creator.create()

	return b
}

func (b *BasicRecordsCreator) AddAAAARecords(creator RecordsCreator) *BasicRecordsCreator {
	b.records.AAAA = creator.create()

	return b
}

func (b *BasicRecordsCreator) AddMXRecords(creator RecordsCreator) *BasicRecordsCreator {
	b.records.Mx = creator.create()

	return b
}

func (b *BasicRecordsCreator) AddTXTRecords(creator RecordsCreator) *BasicRecordsCreator {
	b.records.A = creator.create()

	return b
}

func (b *BasicRecordsCreator) Create() *dns.BasicRecords {
	b.records.DNSRecords = *b.DNSRecordsCreator.Create()

	return b.records
}

type NsRecordsCreator struct {
	DNSRecordsCreator
	records *dns.NSRecords
}

func NewNsRecordsCreator() *NsRecordsCreator {
	return &NsRecordsCreator{
		DNSRecordsCreator: *NewDNSRecordsCreator(),
		records:           &dns.NSRecords{},
	}
}

func (n *NsRecordsCreator) AddNSRecords(creator RecordsCreator) *NsRecordsCreator {
	n.records.Ns = creator.create()

	return n
}

func (n *NsRecordsCreator) Create() *dns.NSRecords {
	n.records.DNSRecords = *n.DNSRecordsCreator.Create()

	return n.records
}

type CNameRecordsCreator struct {
	DNSRecordsCreator
	records *dns.CNAMERecords
}

func NewCNameRecordsCreator() *CNameRecordsCreator {
	return &CNameRecordsCreator{
		DNSRecordsCreator: *NewDNSRecordsCreator(),
		records:           &dns.CNAMERecords{},
	}
}

func (c *CNameRecordsCreator) AddCNameRecords(creator RecordsCreator) *CNameRecordsCreator {
	c.records.CName = creator.create()

	return c
}

func (c *CNameRecordsCreator) Create() *dns.CNAMERecords {
	c.records.DNSRecords = *c.DNSRecordsCreator.Create()

	return c.records
}
