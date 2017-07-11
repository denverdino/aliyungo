package dns

import "github.com/denverdino/aliyungo/common"

type DescribeDomainRecordsNArgs struct {
	DomainName string

	//optional
	common.Pagination
	RRKeyWord    string
	TypeKeyWord  string
	ValueKeyWord string
}

type DescribeDomainRecordsNResponse struct {
	common.Response
	common.PaginationResult
	InstanceId    string
	DomainRecords struct {
		Record []RecordTypeN
	}
}

// DescribeDomainRecordsNew
//
// You can read doc at https://docs.aliyun.com/#/pub/dns/api-reference/record-related&DescribeDomainRecords
func (client *Client) DescribeDomainRecordsN(args *DescribeDomainRecordsNArgs) (response *DescribeDomainRecordsNResponse, err error) {
	action := "DescribeDomainRecords"
	response = &DescribeDomainRecordsNResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		return nil, err
	}
}
