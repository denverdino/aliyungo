package dns

import "github.com/denverdino/aliyungo/common"


type DomainRecordType struct {
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

type DescribeDomainRecordInformationArgs struct {
	RecordId string
}

type DescribeDomainRecordInformationResponse struct {
	common.Response
	DomainRecordType
}

// DescribeDomainRecordInformation
//
// You can read doc at https://docs.aliyun.com/#/pub/dns/api-reference/record-related&DescribeDomainRecordInfo
func (client *Client) DescribeDomainRecordInformation(args *DescribeDomainRecordInformationArgs) (response *DescribeDomainRecordInformationResponse, err error) {
	action := "DescribeDomainRecordInfo"
	response = &DescribeDomainRecordInformationResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		return nil, err
	}
}
