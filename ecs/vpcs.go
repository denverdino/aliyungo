package ecs

import (
	"github.com/denverdino/aliyungo/util"
)

type CreateVpcArgs struct {
	RegionId    Region
	CidrBlock   string //192.168.0.0/16 or 172.16.0.0/16 (default)
	VpcName     string
	Description string
	ClientToken string
}

type CreateVpcResponse struct {
	CommonResponse
	VpcId        string
	VRouterId    string
	RouteTableId string
}

// CreateVpc creates Virtual Private Cloud
func (client *Client) CreateVpc(args *CreateVpcArgs) (resp *CreateVpcResponse, err error) {
	response := CreateVpcResponse{}
	err = client.Invoke("CreateVpc", args, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

type DeleteVpcArgs struct {
	VpcId string
}

type DeleteVpcResponse struct {
	CommonResponse
}

// DeleteVpc deletes Virtual Private Cloud
func (client *Client) DeleteVpc(vpcId string) error {
	args := DeleteVpcArgs{
		VpcId: vpcId,
	}
	response := DeleteVpcResponse{}
	return client.Invoke("DeleteVpc", &args, &response)
}

type VpcStatus string

const (
	VpcStatusPending   = VpcStatus("Pending")
	VpcStatusAvailable = VpcStatus("Available")
)

var FinalVpcStatus = map[VpcStatus]bool{
	VpcStatusAvailable: true,
}

type DescribeVpcsArgs struct {
	VpcId    string
	RegionId Region
	Pagination
}

type VpcSetType struct {
	VpcId      string
	RegionId   Region
	Status     VpcStatus // enum Pending | Available
	VpcName    string
	VSwitchIds struct {
		VSwitchId []string
	}
	CidrBlock    string
	VRouterId    string
	Description  string
	CreationTime util.ISO6801Time
}

type DescribeVpcsResponse struct {
	CommonResponse
	PaginationResult
	Vpcs struct {
		Vpc []VpcSetType
	}
}

// DescribeInstanceStatus describes instance status
func (client *Client) DescribeVpcs(args *DescribeVpcsArgs) (vpcs []VpcSetType, pagination *PaginationResult, err error) {
	args.validate()
	response := DescribeVpcsResponse{}

	err = client.Invoke("DescribeVpcs", args, &response)

	if err == nil {
		return response.Vpcs.Vpc, &response.PaginationResult, nil
	}

	return nil, nil, err
}

type ModifyVpcAttributeArgs struct {
	VpcId       string
	VpcName     string
	Description string
}

type ModifyVpcAttributeResponse struct {
	CommonResponse
}

// ModifyVpcAttribute modifies attribute of Virtual Private Cloud
func (client *Client) ModifyVpcAttribute(args *ModifyVpcAttributeArgs) error {
	response := ModifyVpcAttributeResponse{}
	return client.Invoke("ModifyVpcAttribute", args, &response)
}

// WaitForVSwitchAvailable waits for VSwitch to given status
func (client *Client) WaitForVpcAvailable(regionId Region, vpcId string, strategy util.AttemptStrategy) (status interface{}, err error) {

	fn := func() (bool, interface{}, error) {

		args := DescribeVpcsArgs{
			RegionId: regionId,
			VpcId:    vpcId,
		}

		vpcs, _, err := client.DescribeVpcs(&args)
		if err != nil {
			return false, util.StatusNotAvailable, err
		}

		if FinalVpcStatus[vpcs[0].Status] {
			return true, vpcs[0].Status, nil
		}
		return false, util.StatusNotAvailable, nil
	}

	status, e1 := util.LoopCall(strategy, fn)

	return status, e1
}
