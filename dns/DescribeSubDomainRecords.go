package dns

import (
	"github.com/denverdino/aliyungo/common"
	"log"
)

type DescribeSubDomainRecordsArgs struct {
	SubDomain string

	//optional
	PageNumber int32
	PageSize   int32
	Type       string
}

type DescribeSubDomainRecordsResponse struct {
	common.Response
	InstanceId    string
	TotalCount    int32
	PageNumber    int32
	PageSize      int32
	DomainRecords struct {
		Record []RecordType
	}
}

// DescribeSubDomainRecords
//
// You can read doc at https://docs.aliyun.com/#/pub/dns/api-reference/record-related&DescribeSubDomainRecords
func (client *Client) DescribeSubDomainRecords(args *DescribeSubDomainRecordsArgs) (response *DescribeSubDomainRecordsResponse, err error) {
	action := "DescribeSubDomainRecords"
	response = &DescribeSubDomainRecordsResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		log.Fatalf("%s error, %v", action, err)
		return response, err
	}
}
