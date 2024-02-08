package build

import (
	"github.com/ThCompiler/go.beget.api/pkg/beget/api/dns/records/set"
)

type RecordsCreator interface {
	create() []set.ChangedRecord
}

type DNSRecordsCreator struct {
	records *set.DNSRecords
}

func NewDNSRecordsCreator() *DNSRecordsCreator {
	return &DNSRecordsCreator{
		records: &set.DNSRecords{},
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

func (d *DNSRecordsCreator) Create() *set.DNSRecords {
	return d.records
}

type BasicRecordsCreator struct {
	DNSRecordsCreator
	records *set.BasicRecords
}

func NewBasicRecordsCreator() *BasicRecordsCreator {
	return &BasicRecordsCreator{
		DNSRecordsCreator: *NewDNSRecordsCreator(),
		records:           &set.BasicRecords{},
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

func (b *BasicRecordsCreator) Create() *set.BasicRecords {
	b.records.DNSRecords = *b.DNSRecordsCreator.Create()

	return b.records
}

type NsRecordsCreator struct {
	DNSRecordsCreator
	records *set.NSRecords
}

func NewNsRecordsCreator() *NsRecordsCreator {
	return &NsRecordsCreator{
		DNSRecordsCreator: *NewDNSRecordsCreator(),
		records:           &set.NSRecords{},
	}
}

func (n *NsRecordsCreator) AddNSRecords(creator RecordsCreator) *NsRecordsCreator {
	n.records.Ns = creator.create()

	return n
}

func (n *NsRecordsCreator) Create() *set.NSRecords {
	n.records.DNSRecords = *n.DNSRecordsCreator.Create()

	return n.records
}

type CNameRecordsCreator struct {
	DNSRecordsCreator
	records *set.CNAMERecords
}

func NewCNameRecordsCreator() *CNameRecordsCreator {
	return &CNameRecordsCreator{
		DNSRecordsCreator: *NewDNSRecordsCreator(),
		records:           &set.CNAMERecords{},
	}
}

func (c *CNameRecordsCreator) AddCNameRecords(creator RecordsCreator) *CNameRecordsCreator {
	c.records.CName = creator.create()

	return c
}

func (c *CNameRecordsCreator) Create() *set.CNAMERecords {
	c.records.DNSRecords = *c.DNSRecordsCreator.Create()

	return c.records
}
