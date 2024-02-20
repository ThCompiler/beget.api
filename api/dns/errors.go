package dns

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrTooMuchDNSRecords   = errors.New(fmt.Sprintf("too mach DNS records, max records %d", MaxDNSRecords))
	ErrTooMuchDNSIPRecords = errors.New(fmt.Sprintf("too mach DNS_IP records, max records %d", MaxDNSRecords))
	ErrTooMuchNsRecords    = errors.New(fmt.Sprintf("too mach NS records, max records %d", MaxNSRecords))
	ErrTooMuchCNameRecords = errors.New(fmt.Sprintf("too mach CNAME records, max records %d", MaxCNAMERecords))
	ErrTooMuchARecords     = errors.New(fmt.Sprintf("too mach A records, max records %d", MaxBasicRecords))
	ErrTooMuchAAAARecords  = errors.New(fmt.Sprintf("too mach AAAA records, max records %d", MaxBasicRecords))
	ErrTooMuchMxRecords    = errors.New(fmt.Sprintf("too mach MX records, max records %d", MaxBasicRecords))
	ErrTooMuchTxtRecords   = errors.New(fmt.Sprintf("too mach TXT records, max records %d", MaxBasicRecords))

	ErrNumberDNSRecordsNotEqual = errors.New("DNS records not equal DNS_IP records")
)
