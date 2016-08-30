package vpc

import "github.com/hdksky/aliyungo/common"

type DescribeRegionsArgs struct {
}

//
// You can read doc at http://docs.aliyun.com/#/pub/ecs/open-api/datatype&regiontype
type RegionType struct {
	RegionId  common.Region
	LocalName string
}

type DescribeRegionsResponse struct {
	common.Response
	Regions struct {
		Region []RegionType
	}
}

// DescribeRegions describes regions
//
// You can read doc at http://docs.aliyun.com/#/pub/slb/api-reference/api-related-loadbalancer&DescribeRegions
func (client *Client) DescribeRegions() (regions []RegionType, err error) {
	response := DescribeRegionsResponse{}

	err = client.Invoke("DescribeRegions", &DescribeRegionsArgs{}, &response)

	if err != nil {
		return []RegionType{}, err
	}
	return response.Regions.Region, nil
}

// cookie add func DescribeZones on 2016-07-15
type DescribeZonesArgs struct {
	RegionId common.Region
}

type ZoneType struct {
	ZoneId                    string
	LocalName                 string
	AvailableResourceCreation struct {
		ResourceTypes []string
	}
	AvailableDiskCategories struct {
		DiskCategories []string
	}
}

type DescribeZonesResponse struct {
	common.Response
	Zones struct {
		Zone []ZoneType
	}
}

// DescribeZones describes zones
//
func (client *Client) DescribeZones(regionId common.Region) (zones []ZoneType, err error) {
	args := DescribeZonesArgs{
		RegionId: regionId,
	}
	response := DescribeZonesResponse{}

	err = client.Invoke("DescribeZones", &args, &response)

	if err == nil {
		return response.Zones.Zone, nil
	}

	return []ZoneType{}, err
}
