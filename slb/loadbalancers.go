package slb

import (
	"github.com/hdksky/aliyungo/common"
	"github.com/hdksky/aliyungo/util"
)

const (
	InternetAddressType = "internet"
	IntranetAddressType = "intranet"
)

type CreateLoadBalancerArgs struct {
	RegionId           common.Region
	LoadBalancerName   string
	AddressType        string
	VSwitchId          string
	InternetChargeType common.InternetChargeType
	Bandwidth          int
	ClientToken        string
	MasterZoneId       string
	SlaveZoneId        string
} // cookie add last two attr on 2016-07-15

type CreateLoadBalancerResponse struct {
	common.Response
	LoadBalancerId   string
	Address          string
	VpcId            string
	VSwitchId        string
	MasterZoneId     string
	SlaveZoneId      string
	NetworkType      string //not found in aliyun api doc
	LoadBalancerName string
} // cookie add attr MasterZoneId & SlaveZoneId on 2016-07-15

// CreateLoadBalancer create loadbalancer
//
// You can read doc at http://docs.aliyun.com/#/pub/slb/api-reference/api-related-loadbalancer&CreateLoadBalancer
func (client *Client) CreateLoadBalancer(args *CreateLoadBalancerArgs) (response *CreateLoadBalancerResponse, err error) {
	response = &CreateLoadBalancerResponse{}
	err = client.Invoke("CreateLoadBalancer", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

type DeleteLoadBalancerArgs struct {
	LoadBalancerId string
}

type DeleteLoadBalancerResponse struct {
	common.Response
}

// DeleteLoadBalancer delete loadbalancer
//
// You can read doc at http://docs.aliyun.com/#/pub/slb/api-reference/api-related-loadbalancer&DeleteLoadBalancer
func (client *Client) DeleteLoadBalancer(loadBalancerId string) (err error) {
	args := &DeleteLoadBalancerArgs{
		LoadBalancerId: loadBalancerId,
	}
	response := &DeleteLoadBalancerResponse{}
	err = client.Invoke("DeleteLoadBalancer", args, response)
	if err != nil {
		return err
	}
	return err
}

type ModifyLoadBalancerInternetSpecArgs struct {
	LoadBalancerId     string
	InternetChargeType common.InternetChargeType
	Bandwidth          int
}

type ModifyLoadBalancerInternetSpecResponse struct {
	common.Response
}

// ModifyLoadBalancerInternetSpec Modify loadbalancer internet spec
//
// You can read doc at http://docs.aliyun.com/#/pub/slb/api-reference/api-related-loadbalancer&ModifyLoadBalancerInternetSpec

func (client *Client) ModifyLoadBalancerInternetSpec(args *ModifyLoadBalancerInternetSpecArgs) (err error) {
	response := &ModifyLoadBalancerInternetSpecResponse{}
	err = client.Invoke("ModifyLoadBalancerInternetSpec", args, response)
	if err != nil {
		return err
	}
	return err
}

type Status string

const InactiveStatus = Status("inactive")
const ActiveStatus = Status("active")
const LockedStatus = Status("locked")

type SetLoadBalancerStatusArgs struct {
	LoadBalancerId     string
	LoadBalancerStatus Status
}

type SetLoadBalancerStatusResponse struct {
	common.Response
}

// SetLoadBalancerStatus Set loadbalancer status
//
// You can read doc at http://docs.aliyun.com/#/pub/slb/api-reference/api-related-loadbalancer&SetLoadBalancerStatus

func (client *Client) SetLoadBalancerStatus(loadBalancerId string, status Status) (err error) {
	args := &SetLoadBalancerStatusArgs{
		LoadBalancerId:     loadBalancerId,
		LoadBalancerStatus: status,
	}
	response := &SetLoadBalancerStatusResponse{}
	err = client.Invoke("SetLoadBalancerStatus", args, response)
	if err != nil {
		return err
	}
	return err
}

type SetLoadBalancerNameArgs struct {
	LoadBalancerId   string
	LoadBalancerName string
}

type SetLoadBalancerNameResponse struct {
	common.Response
}

// SetLoadBalancerName Set loadbalancer name
//
// You can read doc at http://docs.aliyun.com/#/pub/slb/api-reference/api-related-loadbalancer&SetLoadBalancerName

func (client *Client) SetLoadBalancerName(loadBalancerId string, name string) (err error) {
	args := &SetLoadBalancerNameArgs{
		LoadBalancerId:   loadBalancerId,
		LoadBalancerName: name,
	}
	response := &SetLoadBalancerNameResponse{}
	err = client.Invoke("SetLoadBalancerName", args, response)
	if err != nil {
		return err
	}
	return err
}

type DescribeLoadBalancersArgs struct {
	RegionId           common.Region
	LoadBalancerId     string
	AddressType        string
	NetworkType        string
	VpcId              string
	VSwitchId          string
	Address            string
	InternetChargeType common.InternetChargeType
	ServerId           string
	MasterZoneId       string
	SlaveZoneId        string
}

type ListenerPortAndProtocolType struct {
	ListenerPort     int
	ListenerProtocol string
}

type BackendServerType struct {
	ServerId string
	Weight   int
}

type LoadBalancerType struct {
	LoadBalancerId     string
	LoadBalancerName   string
	LoadBalancerStatus string
	Address            string
	RegionId           common.Region
	RegionIdAlias      string
	AddressType        string
	VSwitchId          string
	VpcId              string
	NetworkType        string
	Bandwidth          int
	InternetChargeType common.InternetChargeType //not found in aliyun api doc
	CreateTime         string                    //Why not ISO 6801
	MasterZoneId       string
	SlaveZoneId        string
	CreateTimeStamp    util.ISO6801Time //not found in aliyun api doc
	ListenerPorts      struct {
		ListenerPort []int
	}
	ListenerPortsAndProtocol struct {
		ListenerPortAndProtocol []ListenerPortAndProtocolType
	}
	BackendServers struct {
		BackendServer []BackendServerType
	}
}

type DescribeLoadBalancersResponse struct {
	common.Response
	LoadBalancers struct {
		LoadBalancer []LoadBalancerType
	}
}

// DescribeLoadBalancers Describe loadbalancers
//
// You can read doc at http://docs.aliyun.com/#/pub/slb/api-reference/api-related-loadbalancer&DescribeLoadBalancers

func (client *Client) DescribeLoadBalancers(args *DescribeLoadBalancersArgs) (loadBalancers []LoadBalancerType, err error) {
	response := &DescribeLoadBalancersResponse{}
	err = client.Invoke("DescribeLoadBalancers", args, response)
	if err != nil {
		return nil, err
	}
	return response.LoadBalancers.LoadBalancer, err
}

type DescribeLoadBalancerAttributeArgs struct {
	LoadBalancerId string
	MasterZoneId   string
	SlaveZoneId    string
}

type DescribeLoadBalancerAttributeResponse struct {
	common.Response
	LoadBalancerType
}

// DescribeLoadBalancerAttribute Describe loadbalancer attribute
//
// You can read doc at http://docs.aliyun.com/#/pub/slb/api-reference/api-related-loadbalancer&DescribeLoadBalancerAttribute

func (client *Client) DescribeLoadBalancerAttribute(loadBalancerId string) (loadBalancer *LoadBalancerType, err error) {
	args := &DescribeLoadBalancersArgs{
		LoadBalancerId: loadBalancerId,
	}
	response := &DescribeLoadBalancerAttributeResponse{}
	err = client.Invoke("DescribeLoadBalancerAttribute", args, response)
	if err != nil {
		return nil, err
	}
	return &response.LoadBalancerType, err
}
