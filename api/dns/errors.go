package dns

import (
	"fmt"

	"github.com/pkg/errors"
)

// Validation errors for new records for DNS record change request ([changeRecords]).
//
// [changeRecords]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#changerecords
var (
	// "too mach DNS records, max records 4"
	ErrTooMuchDNSRecords = errors.New(fmt.Sprintf("too mach DNS records, max records %d", MaxDNSRecords))
	// "too mach DNS_IP records, max records 4"
	ErrTooMuchDNSIPRecords = errors.New(fmt.Sprintf("too mach DNS_IP records, max records %d", MaxDNSRecords))
	// "too mach NS records, max records 10"
	ErrTooMuchNsRecords = errors.New(fmt.Sprintf("too mach NS records, max records %d", MaxNSRecords))
	// "too mach CNAME records, max records 1"
	ErrTooMuchCNameRecords = errors.New(fmt.Sprintf("too mach CNAME records, max records %d", MaxCNAMERecords))
	// "too mach A records, max records 1"
	ErrTooMuchARecords = errors.New(fmt.Sprintf("too mach A records, max records %d", MaxBasicRecords))
	// "too mach AAAA records, max records 10"
	ErrTooMuchAAAARecords = errors.New(fmt.Sprintf("too mach AAAA records, max records %d", MaxBasicRecords))
	// "too mach MX records, max records 10"
	ErrTooMuchMxRecords = errors.New(fmt.Sprintf("too mach MX records, max records %d", MaxBasicRecords))
	// "too mach TXT records, max records 10"
	ErrTooMuchTxtRecords = errors.New(fmt.Sprintf("too mach TXT records, max records %d", MaxBasicRecords))

	// "DNS records not equal DNS_IP records"
	ErrNumberDNSRecordsNotEqual = errors.New("DNS records not equal DNS_IP records")
)
