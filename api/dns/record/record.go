package record

import "github.com/ThCompiler/go.beget.api/api/dns/general"

//----------------------------------------------------------------------------------------------------------------------
// Records
//----------------------------------------------------------------------------------------------------------------------

type TTLRecord struct {
	TTL int64 `json:"ttl"`
}

type DNSRecord general.DNSRecord

type DNSIPRecord general.DNSIPRecord

type CAARecord struct {
	TTLRecord
	general.CAARecord
}

type SRVRecord struct {
	TTLRecord
	general.SRVRecord
}

type ARecord struct {
	TTLRecord
	general.ARecord
}

type AAAARecord struct {
	TTLRecord
	general.AAAARecord
}

type MXRecord struct {
	TTLRecord
	general.MXRecord
}

type TXTRecord struct {
	TTLRecord
	general.TXTRecord
}

type NSRecord struct {
	TTLRecord
	general.NSRecord
}

type CNAMERecord struct {
	TTLRecord
	general.CNAMERecord
}
