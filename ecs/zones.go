package ecs

import "github.com/hdksky/aliyungo/common"

type ResourceType string

const (
	ResourceTypeInstance            = ResourceType("Instance")
	ResourceTypeDisk                = ResourceType("Disk")
	ResourceTypeVSwitch             = ResourceType("VSwitch")
	ResourceTypeIOOptimizedInstance = ResourceType("IoOptimized")
)

type DescribeZonesArgs struct {
	RegionId common.Region
}

//
// You can read doc at http://docs.aliyun.com/#/pub/ecs/open-api/datatype&availableresourcecreationtype
type AvailableResourceCreationType struct {
	ResourceTypes []ResourceType //enum for Instance, Disk, VSwitch
}

//
// You can read doc at http://docs.aliyun.com/#/pub/ecs/open-api/datatype&availablediskcategoriestype
type AvailableDiskCategoriesType struct {
	DiskCategories []DiskCategory //enum for cloud, ephemeral, ephemeral_ssd
}

type InstanceType InstanceTypeItemType

type SystemDiskCategory string

const (
	SystemDiskCategoryEphemeralSSD    = SystemDiskCategory("ephemeral_ssd")
	SystemDiskCategoryCloudEfficiency = SystemDiskCategory("cloud_efficiency")
	SystemDiskCategoryCloudSSD        = SystemDiskCategory("cloud_ssd")
)

type AvailableResourcesType struct {
	IoOptimized  bool
	NetworkTypes struct {
		SupportedNetworkCategory []string
	}
	InstanceGenerations struct {
		SupportedInstanceGeneration []string
	}
	InstanceTypeFamilies struct {
		SupportedInstanceTypeFamily []string
	}
	SystemDiskCategories struct {
		SupportedSystemDiskCategory []string
	}
	DataDiskCategories struct {
		SupportedDataDiskCategory []string
	}
	InstanceTypes struct {
		SupportedInstanceType []string
	}
}

//
// You can read doc at http://docs.aliyun.com/#/pub/ecs/open-api/datatype&zonetype
type ZoneType struct {
	ZoneId             string
	LocalName          string
	AvailableResources struct {
		ResourcesInfo []AvailableResourcesType
	}
	AvailableInstanceTypes struct {
		InstanceTypes []string
	}
	AvailableResourceCreation struct {
		ResourceTypes []ResourceType
	}
	AvailableDiskCategories struct {
		DiskCategories []DiskCategory
	}
}

type DescribeZonesResponse struct {
	common.Response
	Zones struct {
		Zone []ZoneType
	}
}

// DescribeZones describes zones
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
