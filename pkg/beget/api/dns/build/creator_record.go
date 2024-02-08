package build

import (
	"github.com/ThCompiler/go.beget.api/pkg/beget/api/dns/records/set"
)

type RecordCreator struct {
	maxRecords    int
	overflowError error
	records       []set.ChangedRecord
}

func (r *RecordCreator) AddRecord(priority int64, value string) *RecordCreator {
	if len(r.records) >= r.maxRecords {
		panic(r.overflowError)
	}

	r.records = append(r.records, set.ChangedRecord{Priority: priority, Value: value})

	return r
}

func (r *RecordCreator) create() []set.ChangedRecord {
	return r.records
}

func NewARecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    set.MaxBasicRecords,
		overflowError: set.ErrTooMuchARecords,
		records:       make([]set.ChangedRecord, 0),
	}
}

func NewAAAARecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    set.MaxBasicRecords,
		overflowError: set.ErrTooMuchAAAARecords,
		records:       make([]set.ChangedRecord, 0),
	}
}

func NewMxRecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    set.MaxBasicRecords,
		overflowError: set.ErrTooMuchMxRecords,
		records:       make([]set.ChangedRecord, 0),
	}
}

func NewTxtRecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    set.MaxBasicRecords,
		overflowError: set.ErrTooMuchTxtRecords,
		records:       make([]set.ChangedRecord, 0),
	}
}

func NewNsRecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    set.MaxNSRecords,
		overflowError: set.ErrTooMuchNsRecords,
		records:       make([]set.ChangedRecord, 0),
	}
}

func NewCNameRecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    set.MaxCNAMERecords,
		overflowError: set.ErrTooMuchCNameRecords,
		records:       make([]set.ChangedRecord, 0),
	}
}

func NewDNSIPRecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    set.MaxDNSRecords,
		overflowError: set.ErrTooMuchDNSIPRecords,
		records:       make([]set.ChangedRecord, 0),
	}
}

func NewDNSRecordCreator() *RecordCreator {
	return &RecordCreator{
		maxRecords:    set.MaxDNSRecords,
		overflowError: set.ErrTooMuchDNSRecords,
		records:       make([]set.ChangedRecord, 0),
	}
}
