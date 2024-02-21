// Package general implements structures with a generalized representation of the DNS server records
// supported by Beget.
package general

// CAATag is a tag for controlling the issuance of domain name certificates as well S/MIME certificates for CAA record.
type CAATag string

// Values of CAATag
const (
	Issue     = CAATag("issue")     // tag sets a policy for domain single-name and TLS/SSL wildcard certificate issuance
	IssueWild = CAATag("issuewild") // tag sets a policy for domain wildcard certificate issuance
	IoDef     = CAATag("iodef")     // tag sets a policy for S/MIME certificate issuance
)

// DNSRecord is a representation of a DNS record.
type DNSRecord struct {
	Value string `json:"value"` // DNS server name
}

// DNSIPRecord is a representation of a DNS_IP record.
// [DNSRecord.Value] stores the IP address of the DNS server.
type DNSIPRecord DNSRecord

// CAARecord is a representation of a CAA record.
// More information can be found in the [CAA Records].
//
// [CAA Records]: https://support.dnsimple.com/articles/caa-record/
type CAARecord struct {
	Flags int64  `json:"flags"` // flags of CAA-record.
	Tag   CAATag `json:"tag"`   // tag of CAA-record.
	Value string `json:"value"` // value of CAA-record.
}

// SRVRecord is a representation of a SVR record.
// More information can be found in the [SVR Records].
//
// # Note
//
// Beget supports only tcp protocol and sip service for SRV-record.
//
// [SVR Records]: https://support.dnsimple.com/articles/srv-record/
type SRVRecord struct {
	Priority int64  `json:"priority"` // priority of SRV-record.
	Weight   int64  `json:"weight"`   // weight of SRV-record.
	Port     int64  `json:"port"`     // port of SRV-record.
	Target   string `json:"target"`   // target of SRV-record.
}

// ARecord is a representation of a A record.
// More information can be found in the [A Records].
//
// [A Records]: https://support.dnsimple.com/articles/a-record/
type ARecord struct {
	Address string `json:"address"` // IPv4 address.
}

// AAAARecord is a representation of a AAAA record.
// More information can be found in the [AAAA Records].
//
// Field [ARecord.Address] stores the IPv6 address.
//
// [AAAA Records]: https://support.dnsimple.com/articles/aaaa-record/
type AAAARecord ARecord

// MXRecord is a representation of a MX record.
// More information can be found in the [MX Records].
//
// [MX Records]: https://support.dnsimple.com/articles/mx-record/
type MXRecord struct {
	Exchange   string `json:"exchange"`   // exchange of MX-record (domain name of the mail server).
	Preference int64  `json:"preference"` // preference of MX-record.
}

// TXTRecord is a representation of a TXT record.
// More information can be found in the [TXT Records].
//
// [TXT Records]: https://support.dnsimple.com/articles/txt-record/
type TXTRecord struct {
	TxtData string `json:"txtdata"` // text value of TXT-record.
}

// NSRecord is a representation of a NS record.
// More information can be found in the [NS Records].
//
// [NS Records]: https://support.dnsimple.com/articles/ns-record/
type NSRecord struct {
	NsdName string `json:"nsdname"` // name of NS server for NS-record.
}

// CNAMERecord is a representation of a CNAME record.
// More information can be found in the [CNAME Records].
//
// [CNAME Records]: https://support.dnsimple.com/articles/cname-record/
type CNAMERecord struct {
	Cname string `json:"cname"` // domain name  for which the alias is set in CNAME-record.
}
