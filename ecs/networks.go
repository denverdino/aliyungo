// API on Network

package ecs

import (
	"github.com/denverdino/aliyungo/util"
)

type AllocatePublicIpAddressArgs struct {
	InstanceId string
}

type AllocatePublicIpAddressResponse struct {
	CommonResponse

	IpAddress string
}

// AllocatePublicIpAddress allocates Public Ip Address
func (client *Client) AllocatePublicIpAddress(instanceId string) (ipAddress string, err error) {
	args := AllocatePublicIpAddressArgs{
		InstanceId: instanceId,
	}
	response := AllocatePublicIpAddressResponse{}
	err = client.Invoke("AllocatePublicIpAddress", &args, &response)
	if err != nil {
		return "", err
	}
	return response.IpAddress, nil
}

type ModifyInstanceNetworkSpec struct {
	InstanceId              string
	InternetMaxBandwidthOut *int
	InternetMaxBandwidthIn  *int
}

type ModifyInstanceNetworkSpecResponse struct {
	CommonResponse
}

// ModifyInstanceNetworkSpec modifies instance network spec
func (client *Client) ModifyInstanceNetworkSpec(args *ModifyInstanceNetworkSpec) error {

	response := ModifyInstanceNetworkSpecResponse{}
	return client.Invoke("ModifyInstanceNetworkSpec", args, &response)
}

type AllocateEipAddressArgs struct {
	RegionId           Region
	Bandwidth          int
	InternetChargeType InternetChargeType
	ClientToken        string
}

type AllocateEipAddressResponse struct {
	CommonResponse
	EipAddress   string
	AllocationId string
}

// AllocateEipAddress allocates Eip Address
func (client *Client) AllocateEipAddress(args *AllocateEipAddressArgs) (EipAddress string, AllocationId string, err error) {
	if args.Bandwidth == 0 {
		args.Bandwidth = 5
	}
	response := AllocateEipAddressResponse{}
	err = client.Invoke("AllocateEipAddress", args, &response)
	if err != nil {
		return "", "", err
	}
	return response.EipAddress, response.AllocationId, nil
}

type AssociateEipAddressArgs struct {
	AllocationId string
	InstanceId   string
}

type AssociateEipAddressResponse struct {
	CommonResponse
}

// AssociateEipAddress associates EIP address to VM instance
func (client *Client) AssociateEipAddress(allocationId string, instanceId string) error {
	args := AssociateEipAddressArgs{
		AllocationId: allocationId,
		InstanceId:   instanceId,
	}
	response := ModifyInstanceNetworkSpecResponse{}
	return client.Invoke("AssociateEipAddress", &args, &response)
}

// Status of disks
type EIPStatus string

const (
	EIPStatusAssociating   = EIPStatus("Associating")
	EIPStatusUnassociating = EIPStatus("Unassociating")
	EIPStatusInUse         = EIPStatus("In_use")
	EIPStatusAvailable     = EIPStatus("Available")
)

type DescribeEipAddressesArgs struct {
	RegionId     Region
	Status       EIPStatus //enum Associating | Unassociating | InUse | Available
	EipAddress   string
	AllocationId string
	Pagination
}

type EipAddressSetType struct {
	RegionId           Region
	IpAddress          string
	AllocationId       string
	Status             EIPStatus
	InstanceId         string
	Bandwidth          int
	InternetChargeType InternetChargeType
	OperationLocks     OperationLocksType
	AllocationTime     util.ISO6801Time
}

type DescribeEipAddressesResponse struct {
	CommonResponse
	PaginationResult
	EipAddresses struct {
		EipAddress []EipAddressSetType
	}
}

// DescribeInstanceStatus describes instance status
func (client *Client) DescribeEipAddresses(args *DescribeEipAddressesArgs) (eipAddresses []EipAddressSetType, pagination *PaginationResult, err error) {
	args.validate()
	response := DescribeEipAddressesResponse{}

	err = client.Invoke("DescribeEipAddresses", args, &response)

	if err == nil {
		return response.EipAddresses.EipAddress, &response.PaginationResult, nil
	}

	return nil, nil, err
}

type ModifyEipAddressAttributeArgs struct {
	AllocationId string
	Bandwidth    int
}

type ModifyEipAddressAttributeResponse struct {
	CommonResponse
}

// ModifyEipAddressAttribute Modifies EIP attribute
func (client *Client) ModifyEipAddressAttribute(allocationId string, bandwidth int) error {
	args := ModifyEipAddressAttributeArgs{
		AllocationId: allocationId,
		Bandwidth:    bandwidth,
	}
	response := ModifyEipAddressAttributeResponse{}
	return client.Invoke("ModifyEipAddressAttribute", &args, &response)
}

type UnallocateEipAddressArgs struct {
	AllocationId string
	InstanceId   string
}

type UnallocateEipAddressResponse struct {
	CommonResponse
}

// UnassociateEipAddress unallocates Eip Address from instance
func (client *Client) UnassociateEipAddress(allocationId string, instanceId string) error {
	args := UnallocateEipAddressArgs{
		AllocationId: allocationId,
		InstanceId:   instanceId,
	}
	response := UnallocateEipAddressResponse{}
	return client.Invoke("UnassociateEipAddress", &args, &response)
}

type ReleaseEipAddressArgs struct {
	AllocationId string
}

type ReleaseEipAddressResponse struct {
	CommonResponse
}

// ReleaseEipAddress releases Eip address
func (client *Client) ReleaseEipAddress(allocationId string) error {
	args := ReleaseEipAddressArgs{
		AllocationId: allocationId,
	}
	response := ReleaseEipAddressResponse{}
	return client.Invoke("ReleaseEipAddress", &args, &response)
}
