package dns

import (
	"github.com/denverdino/aliyungo/common"
	"log"
)

type DescribeDomainRecordsArgs struct {
	DomainName string

	//optional
	PageNumber   int32
	PageSize     int32
	RRKeyWord    string
	TypeKeyWord  string
	ValueKeyWord string
}

type DescribeDomainRecordsResponse struct {
	common.Response
	InstanceId    string
	TotalCount    int32
	PageNumber    int32
	PageSize      int32
	DomainRecords struct {
		Record []RecordType
	}
}

// DescribeDomainRecords
//
// You can read doc at https://docs.aliyun.com/?spm=5176.100054.201.106.OeZ3dN#/pub/dns/api-reference/record-related&DescribeDomainRecords
func (client *Client) DescribeDomainRecords(args *DescribeDomainRecordsArgs) (response *DescribeDomainRecordsResponse, err error) {
	action := "DescribeDomainRecords"
	response = &DescribeDomainRecordsResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		log.Fatalf("%s error, %v", action, err)
		return response, err
	}
}
