package ecs

import "github.com/denverdino/aliyungo/common"

type DescribeAvailableResourceArgs struct {
	RegionId            common.Region
	DestinationResource common.DestinationResource
	ZoneId              string
	InstanceChargeType  common.InstanceChargeType
	SpotStrategy        common.SpotStrategy
	IoOptimized         string
	InstanceType        string
	SystemDiskCategory  string
	DataDiskCategory    string
	NetworkCategory     string
}

type DescribeAvailableResourceResponse struct {
	common.Response
	AvailableZones struct {
		AvailableZone []AvailableZone
	}
}

type AvailableZone struct {
	RegionId           common.Region
	ZoneId             string
	Status             string
	AvailableResources struct {
		AvailableResource []AvailableResourcesTypeItem
	}
}

type AvailableResourcesTypeItem struct {
	Type               string
	SupportedResources struct {
		SupportedResource []SupportedResourcesType
	}
}

type SupportedResourcesType struct {
	Value  string
	Status string
	Min    int64
	Max    int64
	Unit   int64
}

// docs https://help.aliyun.com/document_detail/66186.html?spm=a2c4g.11186623.6.963.y4ccIA
func (client *Client) DescribeAvailableResource(args *DescribeAvailableResourceArgs) (*DescribeAvailableResourceResponse, error) {
	response := &DescribeAvailableResourceResponse{}
	err := client.Invoke("DescribeAvailableResource", args, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
