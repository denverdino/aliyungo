package dns

import (
	"log"

	"github.com/denverdino/aliyungo/common"
)

type DescribeDomainGroupsArgs struct {
	PageNumber int32
	PageSize   int32
	KeyWord    string
}

type DescribeDomainGroupsResponse struct {
	response   common.Response
	TotalCount int32
	PageNumber int32
	PageSize   int32
	DomainGroups struct {
		DomainGroupType []DomainGroupType
	}
}

// DescribeDomainGroups
//
// You can read doc at https://docs.aliyun.com/#/pub/dns/api-reference/.....
func (client *Client) DescribeDomainGroups(args *DescribeDomainGroupsArgs) (response *DescribeDomainGroupsResponse, err error) {
	action := "DescribeDomainGroups"
	response = &DescribeDomainGroupsResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		log.Printf("%s error, %v", action, err)
		return response, err
	}
}
