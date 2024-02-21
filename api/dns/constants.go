// Package dns implements the functionality of [Dns methods].
// Package implements [getData] and [changeRecords] methods.
// To create the appropriate methods, you must call either [CallGetData] or [CallChangeRecords].
// Also, package contains:
//   - [github.com/ThCompiler/go.beget.api/api/dns/record] package,
//     that implements structures representing records from [getData] response.
//   - [github.com/ThCompiler/go.beget.api/api/dns/general] package,
//     that implements structures with a generalized representation of the DNS server records supported by Beget.
//   - [github.com/ThCompiler/go.beget.api/api/dns/build] package,
//     that implements builders of [BasicRecords], [DNSRecords], [CNAMERecords] and [NSRecords].
//
// [Dns methods]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
// [changeRecords]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#changerecords
package dns

// Constants used to implement the [getData] and [changeRecords] methods.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
// [changeRecords]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#changerecords
const (
	GetDataMethodName       = "GetData"
	GetDataMethodPath       = "dns/getData"
	ChangeRecordsMethodName = "ChangeRecords"
	ChangeRecordsMethodPath = "dns/changeRecords"
)
