package ecs

import "github.com/hdksky/aliyungo/common"

type DescribeInstanceTypesArgs struct {
	InstanceTypeFamily string
}

//
// You can read doc at http://docs.aliyun.com/#/pub/ecs/open-api/datatype&instancetypeitemtype
type InstanceTypeItemType struct {
	InstanceTypeId     string
	CpuCoreCount       int
	MemorySize         float64
	InstanceTypeFamily string
}

type DescribeInstanceTypesResponse struct {
	common.Response
	InstanceTypes struct {
		InstanceType []InstanceTypeItemType
	}
}

// DescribeInstanceTypes describes all instance types
//
// You can read doc at http://docs.aliyun.com/#/pub/ecs/open-api/other&describeinstancetypes
func (client *Client) DescribeInstanceTypes(args *DescribeInstanceTypesArgs) (instanceTypes []InstanceTypeItemType, err error) {
	response := DescribeInstanceTypesResponse{}

	err = client.Invoke("DescribeInstanceTypes", args, &response)

	if err != nil {
		return []InstanceTypeItemType{}, err
	}
	return response.InstanceTypes.InstanceType, nil

}

//DescribeInstanceTypeFamilies
type GenerationType string

const (
	GenerationTypeEcs_1 = GenerationType("ecs-1")
	GenerationTypeEcs_2 = GenerationType("ecs-2")
)

type DescribeInstanceTypeFamiliesArgs struct {
	RegionId   common.Region
	Generation GenerationType
}

type InstanceTypeFamilyItemType struct {
	InstanceTypeFamilyId string
	Generation           GenerationType
}

type DescribeInstanceTypeFamiliesResponse struct {
	common.Response
	InstanceTypeFamilies struct {
		InstanceTypeFamily []InstanceTypeFamilyItemType
	}
}

func (client *Client) DescribeInstanceTypeFamilies(args *DescribeInstanceTypeFamiliesArgs) (InstanceTypeFamilies []InstanceTypeFamilyItemType, err error) {
	response := DescribeInstanceTypeFamiliesResponse{}

	err = client.Invoke("DescribeInstanceTypeFamilies", args, &response)

	if err != nil {
		return []InstanceTypeFamilyItemType{}, err
	}
	return response.InstanceTypeFamilies.InstanceTypeFamily, nil

}
