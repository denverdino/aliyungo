package dns

import (
	"log"

	"github.com/denverdino/aliyungo/common"
)

type ModifyHichinaDomainDNSArgs struct {
	DomainName string
}

type ModifyHichinaDomainDNSResponse struct {
	common.Response
	OriginalDnsServers struct {
		DnsServerType []DnsServerType
	}
	NewDnsServers struct {
		DnsServerType []DnsServerType
	}
}

// ModifyHichinaDomainDNS
//
// You can read doc at https://docs.aliyun.com/#/pub/dns/api-reference/.....
func (client *Client) ModifyHichinaDomainDNS(args *ModifyHichinaDomainDNSArgs) (response *ModifyHichinaDomainDNSResponse, err error) {
	action := "ModifyHichinaDomainDNS"
	response = &ModifyHichinaDomainDNSResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		log.Printf("%s error, %v", action, err)
		return response, err
	}
}
