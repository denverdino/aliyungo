package dns

import (
	"log"

	"github.com/denverdino/aliyungo/common"
)

type DescribeDomainsArgs struct {
	PageNumber int32
	PageSize   int32
	KeyWord    string
	GroupId    string
}

type DescribeDomainsResponse struct {
	response   common.Response
	TotalCount int32
	PageNumber int32
	PageSize   int32
	Domains struct {
		DomainType []DomainType
	}
}

// DescribeDomains
//
// You can read doc at https://docs.aliyun.com/#/pub/dns/api-reference/.....
func (client *Client) DescribeDomains(args *DescribeDomainsArgs) (response *DescribeDomainsResponse, err error) {
	action := "DescribeDomains"
	response = &DescribeDomainsResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		log.Printf("%s error, %v", action, err)
		return response, err
	}
}
