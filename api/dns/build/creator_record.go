package build

import (
	"github.com/ThCompiler/go.beget.api/api/dns"
)

// RecordCreator is the creator of an array of any records ([github.com/ThCompiler/go.beget.api/api/dns.ChangedRecord])
// for RecordsCreator.
type RecordCreator struct {
	maxRecords    int                 // limit the number of records
	overflowError error               // error overflowing the number of records
	records       []dns.ChangedRecord // added records
}

// AddRecord adds a record to array.
//
// Panics if the number of records of a certain type exceeds the allowed value.
//
// Panics errors:
//   - [github.com/ThCompiler/go.beget.api/api/dns.ErrTooMuchARecords].
//   - [github.com/ThCompiler/go.beget.api/api/dns.ErrTooMuchAAAARecords].
//   - [github.com/ThCompiler/go.beget.api/api/dns.ErrTooMuchMxRecords].
//   - [github.com/ThCompiler/go.beget.api/api/dns.ErrTooMuchTxtRecords].
//   - [github.com/ThCompiler/go.beget.api/api/dns.ErrTooMuchNsRecords].
//   - [github.com/ThCompiler/go.beget.api/api/dns.ErrTooMuchCNameRecords].
//   - [github.com/ThCompiler/go.beget.api/api/dns.ErrTooMuchDNSIPRecords].
//   - [github.com/ThCompiler/go.beget.api/api/dns.ErrTooMuchDNSRecords].
func (r *RecordCreator) AddRecord(priority int64, value string) *RecordCreator {
	if len(r.records) >= r.maxRecords {
		panic(r.overflowError)
	}

	r.records = append(r.records, dns.ChangedRecord{Priority: priority, Value: value})

	return r
}

// create creates an array of records.
// Uses only RecordsCreator.
func (r *RecordCreator) create() []dns.ChangedRecord {
	return r.records
}

// NewARecords creates [RecordCreator] for A-records.
func NewARecords() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxBasicRecords,
		overflowError: dns.ErrTooMuchARecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}

// NewAAAARecords creates [RecordCreator] for AAAA-records.
func NewAAAARecords() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxBasicRecords,
		overflowError: dns.ErrTooMuchAAAARecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}

// NewMxRecords creates [RecordCreator] for MX-records.
func NewMxRecords() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxBasicRecords,
		overflowError: dns.ErrTooMuchMxRecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}

// NewTxtRecords creates [RecordCreator] for TXT-records.
func NewTxtRecords() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxBasicRecords,
		overflowError: dns.ErrTooMuchTxtRecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}

// NewNsRecords creates [RecordCreator] for NS-records.
func NewNsRecords() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxNSRecords,
		overflowError: dns.ErrTooMuchNsRecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}

// NewCNameRecords creates [RecordCreator] for CNAME-records.
func NewCNameRecords() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxCNAMERecords,
		overflowError: dns.ErrTooMuchCNameRecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}

// NewDNSIPRecords creates [RecordCreator] for DNS_IP-records.
func NewDNSIPRecords() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxDNSRecords,
		overflowError: dns.ErrTooMuchDNSIPRecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}

// NewDNSRecords creates [RecordCreator] for DNS-records.
func NewDNSRecords() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxDNSRecords,
		overflowError: dns.ErrTooMuchDNSRecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}
