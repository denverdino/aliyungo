package dns

import (
	"log"

	"github.com/denverdino/aliyungo/common"
)

type DeleteDomainArgs struct {
	DomainName string
}

type DeleteDomainResponse struct {
	common.Response
	DomainName string
}

// DeleteDomain
//
// You can read doc at https://docs.aliyun.com/#/pub/dns/api-reference/.....
func (client *Client) DeleteDomain(args *DeleteDomainArgs) (response *DeleteDomainResponse, err error) {
	action := "DeleteDomain"
	response = &DeleteDomainResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		log.Printf("%s error, %v", action, err)
		return response, err
	}
}
