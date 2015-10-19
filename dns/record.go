package dns

//
//you can read doc at https://docs.aliyun.com/?spm=5176.100054.201.106.OeZ3dN#/pub/dns/api-reference/enum-type&record-format
const (
	ARecord           = "A"
	NSRecord          = "NS"
	MXRecord          = "MX"
	TXTRecord         = "TXT"
	CNAMERecord       = "CNAME"
	SRVRecord         = "SRV"
	AAAARecord        = "AAAA"
	RedirectURLRecord = "REDIRECT_URL"
	ForwordURLRecord  = "FORWORD_URL"
)

type RecordType struct {
	DomainName string
	RecordId   string
	RR         string
	Type       string
	Value      string
	TTL        int32
	Priority   int32
	Line       string
	Status     string
	Locked     bool
}
