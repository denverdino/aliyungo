package ecs

import ()

type Region string

const (
	HANGZHOU  = Region("cn-hangzhou")
	QINGDAO   = Region("cn-qingdao")
	BEIJING   = Region("cn-beijing")
	HONGKONG  = Region("cn-hongkong")
	SHENZHEN  = Region("cn-shenzhen")
	US_WEST_1 = Region("us-west-1")
)

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

func (client *Client) DescribeRegions() (regions []RegionType, err error) {
	response := DescribeRegionsRespones{}

	err = client.Invoke("DescribeRegions", &DescribeRegionsArgs{}, &response)

	if err != nil {
		return []RegionType{}, err
	}
	return response.Regions.Region, nil
}
