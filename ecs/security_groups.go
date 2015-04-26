package ecs

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

// Describe `SecurityGroups`
type DescribeSecurityGroupsArgs struct {
	RegionId string
	VpcId    string
	Pagination
}

type SecurityGroupItemType struct {
	SecurityGroupId string
	Description     string
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
