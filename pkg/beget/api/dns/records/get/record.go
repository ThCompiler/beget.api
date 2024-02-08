package get

import "github.com/ThCompiler/go.beget.api/pkg/beget/api/dns/records"

type TTLRecord struct {
	TTL int64 `json:"ttl"`
}

type DNSRecord records.DNSRecord

type DNSIPRecord records.DNSIPRecord

type CAARecord struct {
	TTLRecord
	records.CAARecord
}

type SRVRecord struct {
	TTLRecord
	records.SRVRecord
}

type ARecord struct {
	TTLRecord
	records.ARecord
}

type AAAARecord struct {
	TTLRecord
	records.AAAARecord
}

type MXRecord struct {
	TTLRecord
	records.MXRecord
}

type TXTRecord struct {
	TTLRecord
	records.TXTRecord
}

type NSRecord struct {
	TTLRecord
	records.NSRecord
}

type CNAMERecord struct {
	TTLRecord
	records.CNAMERecord
}
