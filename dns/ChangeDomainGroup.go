package dns

import (
	"log"

	"github.com/denverdino/aliyungo/common"
)

type ChangeDomainGroupArgs struct {
	DomainName string
	GroupId string
}

type ChangeDomainGroupResponse struct {
	common.Response
	GroupId   string
	GroupName string
}

// ChangeDomainGroup
//
// You can read doc at https://docs.aliyun.com/#/pub/dns/api-reference/.......
func (client *Client) ChangeDomainGroup(args *ChangeDomainGroupArgs) (response *ChangeDomainGroupResponse, err error) {
	action := "AddDomain"
	response = &ChangeDomainGroupResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		log.Printf("%s error, %v", action, err)
		return response, err
	}
}
