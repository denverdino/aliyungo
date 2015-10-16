package dns

import "github.com/denverdino/aliyungo/common"

type DomainRecordType string

//
//you can read doc at https://docs.aliyun.com/?spm=5176.100054.201.106.OeZ3dN#/pub/dns/api-reference/enum-type&record-format
const (
	ARecord = DomainRecordType("A")
	NSRecord   = DomainRecordType("NS")
	MXRecord = DomainRecordType("MX")
	TXTRecord = DomainRecordType("TXT")
	CNAMERecord = DomainRecordType("CNAME")
	SRVRecord = DomainRecordType("SRV")
	AAAARecord = DomainRecordType("AAAA")
	URLRecord = DomainRecordType("URL")
)

type AddDomainRecordArgs struct {
	DomainName string
	RR string
	Type DomainRecordType
	Value string
	TTL int32
	Priority int
	Line string
}

type AddDomainRecordResponse struct {
	common.Response
	InstanceId string
	RecordId string
}

// AddDomainRecord
//
// You can read doc at https://docs.aliyun.com/?spm=5176.100054.201.106.OeZ3dN#/pub/dns/api-reference/record-related&AddDomainRecord
func (client *Client) AddDomainRecord(args *AddDomainRecordArgs) (InstanceId string, RecordId string, err error) {
	response := AddDomainRecordResponse{}
	err = client.Invoke("AddDomainRecord", args, &response)

	if err == nil {
		return response.InstanceId, response.RecordId, nil
	}

	return "", "", err
}

