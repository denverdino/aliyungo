package dns

import (
	"log"

	"github.com/denverdino/aliyungo/common"
)

type AddDomainGroupArgs struct {
	GroupName string
}

type AddDomainGroupResponse struct {
	common.Response
	GroupId   string
	GroupName string
}

// AddDomainGroup
//
// You can read doc at https://docs.aliyun.com/#/pub/dns/api-reference/.......
func (client *Client) AddDomainGroup(args *AddDomainGroupArgs) (response *AddDomainGroupResponse, err error) {
	action := "AddDomainGroup"
	response = &AddDomainGroupResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		log.Printf("%s error, %v", action, err)
		return response, err
	}
}
