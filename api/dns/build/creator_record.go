package build

import (
	"github.com/ThCompiler/go.beget.api/api/dns"
)

type RecordCreator struct {
	maxRecords    int
	overflowError error
	records       []dns.ChangedRecord
}

func (r *RecordCreator) AddRecord(priority int64, value string) *RecordCreator {
	if len(r.records) >= r.maxRecords {
		panic(r.overflowError)
	}

	r.records = append(r.records, dns.ChangedRecord{Priority: priority, Value: value})

	return r
}

func (r *RecordCreator) create() []dns.ChangedRecord {
	return r.records
}

func NewARecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxBasicRecords,
		overflowError: dns.ErrTooMuchARecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}

func NewAAAARecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxBasicRecords,
		overflowError: dns.ErrTooMuchAAAARecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}

func NewMxRecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxBasicRecords,
		overflowError: dns.ErrTooMuchMxRecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}

func NewTxtRecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxBasicRecords,
		overflowError: dns.ErrTooMuchTxtRecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}

func NewNsRecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxNSRecords,
		overflowError: dns.ErrTooMuchNsRecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}

func NewCNameRecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxCNAMERecords,
		overflowError: dns.ErrTooMuchCNameRecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}

func NewDNSIPRecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxDNSRecords,
		overflowError: dns.ErrTooMuchDNSIPRecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}

func NewDNSRecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    dns.MaxDNSRecords,
		overflowError: dns.ErrTooMuchDNSRecords,
		records:       make([]dns.ChangedRecord, 0),
	}
}
