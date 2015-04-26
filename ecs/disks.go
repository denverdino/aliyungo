package ecs

import (
	"github.com/denverdino/aliyun-go/util"
)

// Describe Disks
const (
	DISK_TYPE_ALL    = "all" //Default
	DISK_TYPE_SYSTEM = "system"
	DISK_TYPE_DATA   = "data"

	DISK_CATEGORY_ALL       = "all" //Default
	DISK_CATEGORY_CLOUD     = "cloud"
	DISK_CATEGORY_EPHEMERAL = "ephemeral"

	DISK_STATUS_IN_USE    = "In_use"
	DISK_STATUS_AVAILABLE = "Available"
	DISK_STATUS_ATTACHING = "Attaching"
	DISK_STATUS_DETACHING = "Detaching"
	DISK_STATUS_CREATING  = "Creating"
	DISK_STATUS_REINITING = "ReIniting"
	DISK_STATUS_ALL       = "All" //Default
)

type DescribeDisksArgs struct {
	RegionId           string
	ZoneId             string
	DiskIds            []string
	InstanceId         string
	DiskType           string //all(default) | system | data
	Category           string //all(default) | cloud | ephemeral
	Status             string //In_use | Available | Attaching | Detaching | Creating | ReIniting | All(default)
	SnapshotId         string
	Name               string
	Portable           *bool
	DeleteWithInstance *bool
	DeleteAutoSnapshot *bool
	Pagination
}

type DiskItemType struct {
	DiskId             string
	RegionId           string
	ZoneId             string
	DiskName           string
	Description        string
	Type               string
	Category           string
	Size               int
	ImageId            string
	SourceSnapshotId   string
	ProductCode        string
	Portable           bool
	Status             string
	OperationLocks     OperationLocksType
	InstanceId         string
	Device             string
	DeleteWithInstance bool
	DeleteAutoSnapshot bool
	EnableAutoSnapshot bool
	CreationTime       util.ISO6801Time
	AttachedTime       util.ISO6801Time
	DetachedTime       util.ISO6801Time
}

type DescribeDisksResponse struct {
	CommonResponse

	RegionId string
	PaginationResult
	Disks struct {
		Disk []DiskItemType
	}
}

func (client *Client) DescribeDisks(args *DescribeDisksArgs) (disks []DiskItemType, pagination *PaginationResult, err *ECSError) {
	response := DescribeDisksResponse{}

	err = client.Invoke("DescribeDisks", args, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Disks.Disk, &response.PaginationResult, err
}
