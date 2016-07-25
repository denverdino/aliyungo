package slb

import (
	"github.com/hdksky/aliyungo/common"
)

//CreateProtectedDomain

type WafPolicyLevelType string

const (
	HighWafPolicyLevelType   = WafPolicyLevelType("high")
	MiddleWafPolicyLevelType = WafPolicyLevelType("middle")
	LowWafPolicyLevelType    = WafPolicyLevelType("low")
)

type CreateProtectedDomainArgs struct {
	RegionId       common.Region
	DomainName     string
	WafPolicyLevel WafPolicyLevelType
}

type CreateProtectedDomainResponse struct {
	common.Response
}

func (client *Client) CreateProtectedDomain(args *CreateProtectedDomainArgs) (err error) {
	response := &CreateProtectedDomainResponse{}
	err = client.Invoke("CreateProtectedDomain", args, response)
	if err != nil {
		return err
	}
	return err
}

//SetProtectedDomainStatus

type WafStatusType string

const (
	ActiveWafStatusType   = WafStatusType("active")
	InactiveWafStatusType = WafStatusType("inactive")
)

type CcStatusType string

const (
	ActiveCcStatusType   = CcStatusType("active")
	InactiveCcStatusType = CcStatusType("inactive")
)

type SetProtectedDomainStatusArgs struct {
	RegionId   common.Region
	DomainName string
	WafStatus  WafStatusType
	CcStatus   CcStatusType
}

type SetProtectedDomainStatusResponse struct {
	common.Response
}

func (client *Client) SetProtectedDomainStatus(args *SetProtectedDomainStatusArgs) (err error) {
	response := &SetProtectedDomainStatusResponse{}
	err = client.Invoke("SetProtectedDomainStatus", args, response)
	if err != nil {
		return err
	}
	return err
}

//ModifyProtectedDomainSpec
type ModifyProtectedDomainSpecArgs CreateProtectedDomainArgs
type ModifyProtectedDomainSpecResponse CreateProtectedDomainResponse

func (client *Client) ModifyProtectedDomainSpec(args *ModifyProtectedDomainSpecArgs) (err error) {
	response := &ModifyProtectedDomainSpecResponse{}
	err = client.Invoke("ModifyProtectedDomainSpec", args, response)
	if err != nil {
		return err
	}
	return err
}

//DescribeProtectedDomain
type DescribeProtectedDomainArgs struct {
	RegionId   common.Region
	DomainName string
}

type DescribeProtectedDomainResponse struct {
	common.Response
	DomainType
}

func (client *Client) DescribeProtectedDomain(args *DescribeProtectedDomainArgs) (response *DescribeProtectedDomainResponse, err error) {
	response = &DescribeProtectedDomainResponse{}
	err = client.Invoke("DescribeProtectedDomain", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

//DescribeProtectedDomains

type CcBlackListType struct {
	IpItem []string
}
type CcWhiteList CcBlackListType
type CcBlackList CcBlackListType
type WafWhiteList CcBlackListType
type WafBlackList CcBlackListType

type DomainType struct {
	DomainName     string
	WafStatus      WafStatusType
	CcStatus       CcStatusType
	WafPolicyLevel WafPolicyLevelType
	CcWhiteList
	CcBlackList
	WafWhiteList
	WafBlackList
}

type DescribeProtectedDomainsArgs struct {
	RegionId common.Region
}

type DescribeProtectedDomainsResponse struct {
	common.Response
	Domains struct {
		Domain []DomainType
	}
}

func (client *Client) DescribeProtectedDomains(args *DescribeProtectedDomainsArgs) (response *DescribeProtectedDomainsResponse, err error) {
	response = &DescribeProtectedDomainsResponse{}
	err = client.Invoke("DescribeProtectedDomains", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

//DeleteProtectedDomain
type DeleteProtectedDomainArgs struct {
	RegionId   common.Region
	DomainName string
}

type DeleteProtectedDomainResponse struct {
	common.Response
}

func (client *Client) DeleteProtectedDomain(RegionId common.Region, DomainName string) (err error) {
	args := &DeleteProtectedDomainArgs{
		RegionId:   RegionId,
		DomainName: DomainName,
	}
	response := &DeleteProtectedDomainResponse{}
	err = client.Invoke("DeleteProtectedDomain", args, response)
	if err != nil {
		return err
	}
	return err
}

//AddDomainCcBlackList
type AddDomainCcBlackListArgs struct {
	RegionId   common.Region
	DomainName string
	List       string
}

type AddDomainCcBlackListResponse struct {
	common.Response
}

func (client *Client) AddDomainCcBlackList(args *AddDomainCcBlackListArgs) (err error) {
	response := &AddDomainCcBlackListResponse{}
	err = client.Invoke("AddDomainCcBlackList", args, response)
	if err != nil {
		return err
	}
	return err
}

//RemoveDomainCcBlackList
type RemoveDomainCcBlackListArgs AddDomainCcBlackListArgs
type RemoveDomainCcBlackListResponse AddDomainCcBlackListResponse

func (client *Client) RemoveDomainCcBlackList(args *RemoveDomainCcBlackListArgs) (err error) {
	response := &RemoveDomainCcBlackListResponse{}
	err = client.Invoke("RemoveDomainCcBlackList", args, response)
	if err != nil {
		return err
	}
	return err
}

//DescribeDomainCcBlackList
type DescribeDomainCcBlackListArgs DescribeProtectedDomainArgs
type DescribeDomainCcBlackListResponse struct {
	common.Response
	CcBlackList
}

func (client *Client) DescribeDomainCcBlackList(args *DescribeDomainCcBlackListArgs) (response *DescribeDomainCcBlackListResponse, err error) {
	response = &DescribeDomainCcBlackListResponse{}
	err = client.Invoke("DescribeDomainCcBlackList", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

//AddDomainCcWhiteList
type AddDomainCcWhiteListArgs AddDomainCcBlackListArgs
type AddDomainCcWhiteListResponse AddDomainCcBlackListResponse

func (client *Client) AddDomainCcWhiteList(args *AddDomainCcWhiteListArgs) (err error) {
	response := &AddDomainCcWhiteListResponse{}
	err = client.Invoke("AddDomainCcWhiteList", args, response)
	if err != nil {
		return err
	}
	return err
}

//RemoveDomainCcWhiteList
type RemoveDomainCcWhiteListArgs AddDomainCcBlackListArgs
type RemoveDomainCcWhiteListResponse AddDomainCcBlackListResponse

func (client *Client) RemoveDomainCcWhiteList(args *RemoveDomainCcWhiteListArgs) (err error) {
	response := &RemoveDomainCcWhiteListResponse{}
	err = client.Invoke("RemoveDomainCcWhiteList", args, response)
	if err != nil {
		return err
	}
	return err
}

//DescribeDomainCcWhiteList
type DescribeDomainCcWhiteListArgs DescribeProtectedDomainArgs
type DescribeDomainCcWhiteListResponse struct {
	common.Response
	CcWhiteList
}

func (client *Client) DescribeDomainCcWhiteList(args *DescribeDomainCcWhiteListArgs) (response *DescribeDomainCcWhiteListResponse, err error) {
	response = &DescribeDomainCcWhiteListResponse{}
	err = client.Invoke("DescribeDomainCcWhiteList", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

//AddDomainWafBlackList
type AddDomainWafBlackListArgs AddDomainCcBlackListArgs
type AddDomainWafBlackListResponse AddDomainCcBlackListResponse

func (client *Client) AddDomainWafBlackList(args *AddDomainWafBlackListArgs) (err error) {
	response := &AddDomainWafBlackListResponse{}
	err = client.Invoke("AddDomainWafBlackList", args, response)
	if err != nil {
		return err
	}
	return err
}

//RemoveDomainWafBlackList
type RemoveDomainWafBlackListArgs AddDomainCcBlackListArgs
type RemoveDomainWafBlackListResponse AddDomainCcBlackListResponse

func (client *Client) RemoveDomainWafBlackList(args *RemoveDomainWafBlackListArgs) (err error) {
	response := &RemoveDomainWafBlackListResponse{}
	err = client.Invoke("RemoveDomainWafBlackList", args, response)
	if err != nil {
		return err
	}
	return err
}

//DescribeDomainWafBlackList
type DescribeDomainWafBlackListArgs DescribeProtectedDomainArgs
type DescribeDomainWafBlackListResponse struct {
	common.Response
	WafBlackList
}

func (client *Client) DescribeDomainWafBlackList(args *DescribeDomainWafBlackListArgs) (response *DescribeDomainWafBlackListResponse, err error) {
	response = &DescribeDomainWafBlackListResponse{}
	err = client.Invoke("DescribeDomainWafBlackList", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

//AddDomainWafWhiteList
type AddDomainWafWhiteListArgs AddDomainCcBlackListArgs
type AddDomainWafWhiteListResponse AddDomainCcBlackListResponse

func (client *Client) AddDomainWafWhiteList(args *AddDomainWafWhiteListArgs) (err error) {
	response := &AddDomainWafWhiteListResponse{}
	err = client.Invoke("AddDomainWafWhiteList", args, response)
	if err != nil {
		return err
	}
	return err
}

//RemoveDomainWafWhiteList
type RemoveDomainWafWhiteListArgs AddDomainCcBlackListArgs
type RemoveDomainWafWhiteListResponse AddDomainCcBlackListResponse

func (client *Client) RemoveDomainWafWhiteList(args *RemoveDomainWafWhiteListArgs) (err error) {
	response := &RemoveDomainWafWhiteListResponse{}
	err = client.Invoke("RemoveDomainWafWhiteList", args, response)
	if err != nil {
		return err
	}
	return err
}

//DescribeDomainWafWhiteList
type DescribeDomainWafWhiteListArgs DescribeProtectedDomainArgs
type DescribeDomainWafWhiteListResponse struct {
	common.Response
	WafWhiteList
}

func (client *Client) DescribeDomainWafWhiteList(args *DescribeDomainWafWhiteListArgs) (response *DescribeDomainWafWhiteListResponse, err error) {
	response = &DescribeDomainWafWhiteListResponse{}
	err = client.Invoke("DescribeDomainWafWhiteList", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

//SetListenerSecurityStatus

type SecurityStatusType string

const (
	OnSecurityStatusType  = SecurityStatusType("on")
	OffSecurityStatusType = SecurityStatusType("off")
)

type SetListenerSecurityStatusArgs struct {
	LoadBalancerId string
	ListenerPort   int
	SecurityStatus SecurityStatusType
}

type SetListenerSecurityStatusResponse struct {
	common.Response
}

func (client *Client) SetListenerSecurityStatus(args *SetListenerSecurityStatusArgs) (err error) {
	response := &SetListenerSecurityStatusResponse{}
	err = client.Invoke("SetListenerSecurityStatus", args, response)
	if err != nil {
		return err
	}
	return err
}
