package ecs

import (
	"github.com/denverdino/aliyungo/util"
)

type DescribeSecurityGroupAttributeArgs struct {
	SecurityGroupId string
	RegionId        string
	NicType         string //Enum internet (default) |intranet
}

type PermissionType struct {
	IpProtocol              string
	PortRange               string
	SourceCidrIp            string
	SourceGroupId           string
	SourceGroupOwnerAccount string
	Policy                  string
	NicType                 string
}

type DescribeSecurityGroupAttributeResponse struct {
	CommonResponse

	SecurityGroupId   string
	SecurityGroupName string
	RegionId          string
	Description       string
	Permissions       struct {
		Permission []PermissionType
	}
	VpcId string
}

func (client *Client) DescribeSecurityGroupAttribute(args *DescribeSecurityGroupAttributeArgs) (response *DescribeSecurityGroupAttributeResponse, err *ECSError) {
	response = &DescribeSecurityGroupAttributeResponse{}
	err = client.Invoke("DescribeSecurityGroupAttribute", args, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Describe Security Groups
type DescribeSecurityGroupsArgs struct {
	RegionId string
	VpcId    string
	Pagination
}

type SecurityGroupItemType struct {
	SecurityGroupId   string
	SecurityGroupName string
	Description       string
	VpcId             string
	CreationTime      util.ISO6801Time
}

type DescribeSecurityGroupsResponse struct {
	CommonResponse

	PaginationResult
	RegionId       string
	SecurityGroups struct {
		SecurityGroup []SecurityGroupItemType
	}
}

func (client *Client) DescribeSecurityGroups(args *DescribeSecurityGroupsArgs) (securityGroupItems []SecurityGroupItemType, pagination *PaginationResult, err *ECSError) {
	args.validate()
	response := DescribeSecurityGroupsResponse{}

	err = client.Invoke("DescribeSecurityGroups", args, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.SecurityGroups.SecurityGroup, &response.PaginationResult, nil
}

type CreateSecurityGroupArgs struct {
	RegionId          string
	SecurityGroupName string
	Description       string
	VpcId             string
	ClientToken       string
}

type CreateSecurityGroupResponse struct {
	CommonResponse

	SecurityGroupId string
}

func (client *Client) CreateSecurityGroup(args *CreateSecurityGroupArgs) (securityGroupId string, err *ECSError) {
	response := CreateSecurityGroupResponse{}
	err = client.Invoke("CreateSecurityGroup", args, &response)
	if err != nil {
		return "", err
	}
	return response.SecurityGroupId, err
}

type DeleteSecurityGroupArgs struct {
	RegionId        string
	SecurityGroupId string
}

type DeleteSecurityGroupResponse struct {
	CommonResponse
}

func (client *Client) DeleteSecurityGroup(regionId string, securityGroupId string) *ECSError {
	args := DeleteSecurityGroupArgs{
		RegionId:        regionId,
		SecurityGroupId: securityGroupId,
	}
	response := DeleteSecurityGroupResponse{}
	err := client.Invoke("DeleteSecurityGroup", &args, &response)
	return err
}

type ModifySecurityGroupAttributeArgs struct {
	RegionId          string
	SecurityGroupId   string
	SecurityGroupName string
	Description       string
}

type ModifySecurityGroupAttributeResponse struct {
	CommonResponse
}

func (client *Client) ModifySecurityGroupAttribute(args *ModifySecurityGroupAttributeArgs) *ECSError {
	response := ModifySecurityGroupAttributeResponse{}
	err := client.Invoke("ModifySecurityGroupAttribute", args, &response)
	return err
}
