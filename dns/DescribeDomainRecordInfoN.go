package dns

import "github.com/denverdino/aliyungo/common"

// endpoint change to 'http://alidns.aliyuncs.com' then record ttl and priority change to string
type RecordTypeN struct {
	DomainName string
	RecordId   string
	RR         string
	Type       string
	Value      string
	TTL        string
	Priority   string
	Line       string
	Status     string
	Locked     bool
}

type DescribeDomainRecordInfoNArgs struct {
	RecordId string
}

type DescribeDomainRecordInfoNResponse struct {
	common.Response
	RecordTypeN
}

// DescribeDomainRecordInformation
//
// You can read doc at https://docs.aliyun.com/#/pub/dns/api-reference/record-related&DescribeDomainRecordInfo
func (client *Client) DescribeDomainRecordInfoN(args *DescribeDomainRecordInfoNArgs) (response *DescribeDomainRecordInfoNResponse, err error) {
	action := "DescribeDomainRecordInfo"
	response = &DescribeDomainRecordInfoNResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		return nil, err
	}
}
