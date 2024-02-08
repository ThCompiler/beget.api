package records

type CAATag string

const (
	Issue     = CAATag("issue")
	IssueWild = CAATag("issuewild")
	IoDef     = CAATag("iodef")
)

type DNSRecord struct {
	Value string `json:"value"`
}

type DNSIPRecord DNSRecord

type CAARecord struct {
	Flags int64  `json:"flags"`
	Tag   CAATag `json:"tag"`
	Value string `json:"value"`
}

type SRVRecord struct {
	Priority int64  `json:"priority"`
	Weight   int64  `json:"weight"`
	Port     int64  `json:"port"`
	Target   string `json:"target"`
}

type ARecord struct {
	Address string `json:"address"`
}

type AAAARecord ARecord

type MXRecord struct {
	Exchange   string `json:"exchange"`
	Preference int64  `json:"preference"`
}

type TXTRecord struct {
	TxtData string `json:"txtdata"`
}

type NSRecord struct {
	NsdName string `json:"nsdname"`
}

type CNAMERecord struct {
	Cname string `json:"cname"`
}
