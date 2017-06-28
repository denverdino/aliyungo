package dns

import (
	"log"

	"github.com/denverdino/aliyungo/common"
)

type DescribeDomainInfoArgs struct {
	DomainName string
}

type DescribeDomainInfoResponse struct {
	response common.Response
	DomainType
}

// DescribeDomainInfo
//
// You can read doc at https://docs.aliyun.com/#/pub/dns/api-reference/.....
func (client *Client) DescribeDomainInfo(args *DescribeDomainInfoArgs) (response *DescribeDomainInfoResponse, err error) {
	action := "DescribeDomainInfo"
	response = &DescribeDomainInfoResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		log.Printf("%s error, %v", action, err)
		return response, err
	}
}
