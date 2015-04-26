package ecs

import ()

type DescribeRegionsArgs struct {
}

type RegionType struct {
	RegionId  string
	LocalName string
}

type DescribeRegionsRespones struct {
	CommonResponse
	Regions struct {
		Region []RegionType
	}
}

func (client *Client) DescribeRegions() (regions []RegionType, err *ECSError) {
	response := DescribeRegionsRespones{}

	err = client.Invoke("DescribeRegions", &DescribeRegionsArgs{}, &response)

	if err != nil {
		return []RegionType{}, err
	}
	return response.Regions.Region, nil
}
