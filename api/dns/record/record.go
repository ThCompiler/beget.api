// Package record implements structures representing records from [getData] response.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
package record

import "github.com/ThCompiler/go.beget.api/api/dns/general"

// TTLRecord is a representation of a TTL field in any records.
type TTLRecord struct {
	TTL int64 `json:"ttl"`
}

// DNSRecord is alias of [general.DNSRecord] for [getData] response.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type DNSRecord general.DNSRecord

// DNSIPRecord is alias of [general.DNSIPRecord] for [getData] response.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type DNSIPRecord general.DNSIPRecord

// CAARecord is alias of [general.CAARecord] for [getData] response with TTL field.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type CAARecord struct {
	TTLRecord
	general.CAARecord
}

// SRVRecord is alias of [general.SRVRecord] for [getData] response with TTL field.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type SRVRecord struct {
	TTLRecord
	general.SRVRecord
}

// ARecord is alias of [general.ARecord] for [getData] response with TTL field.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type ARecord struct {
	TTLRecord
	general.ARecord
}

// AAAARecord is alias of [general.AAAARecord] for [getData] response with TTL field.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type AAAARecord struct {
	TTLRecord
	general.AAAARecord
}

// MXRecord is alias of [general.MXRecord] for [getData] response with TTL field.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type MXRecord struct {
	TTLRecord
	general.MXRecord
}

// TXTRecord is alias of [general.TXTRecord] for [getData] response with TTL field.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type TXTRecord struct {
	TTLRecord
	general.TXTRecord
}

// NSRecord is alias of [general.NSRecord] for [getData] response with TTL field.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type NSRecord struct {
	TTLRecord
	general.NSRecord
}

// CNAMERecord is alias of [general.CNAMERecord] for [getData] response with TTL field.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
type CNAMERecord struct {
	TTLRecord
	general.CNAMERecord
}
