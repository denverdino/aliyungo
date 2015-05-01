package ecs

import (
	"github.com/denverdino/aliyungo/util"
	"time"
)

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

// Describe Disks
type DescribeDisksArgs struct {
	RegionId           string
	ZoneId             string
	DiskIds            []string
	InstanceId         string
	DiskType           string //enum for all(default) | system | data
	Category           string //enum for all(default) | cloud | ephemeral
	Status             string //enum for In_use | Available | Attaching | Detaching | Creating | ReIniting | All(default)
	SnapshotId         string
	Name               string
	Portable           *bool //optional
	DeleteWithInstance *bool //optional
	DeleteAutoSnapshot *bool //optional
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

// Describe Disks
func (client *Client) DescribeDisks(args *DescribeDisksArgs) (disks []DiskItemType, pagination *PaginationResult, err *ECSError) {
	response := DescribeDisksResponse{}

	err = client.Invoke("DescribeDisks", args, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Disks.Disk, &response.PaginationResult, err
}

// Create Disk
type CreateDiskArgs struct {
	RegionId    string
	ZoneId      string
	DiskName    string
	Description string
	Size        int
	SnapshotId  string
	ClientToken string
}

type CreateDisksResponse struct {
	CommonResponse
	DiskId string
}

// Create Disk
func (client *Client) CreateDisk(args *CreateDiskArgs) (diskId string, err *ECSError) {
	response := CreateDisksResponse{}
	err = client.Invoke("CreateDisk", args, &response)
	if err != nil {
		return "", err
	}
	return response.DiskId, err
}

type DeleteDiskArgs struct {
	DiskId string
}

type DeleteDiskResponse struct {
	CommonResponse
}

func (client *Client) DeleteDisk(diskId string) *ECSError {
	args := DeleteDiskArgs{
		DiskId: diskId,
	}
	response := DeleteDiskResponse{}
	err := client.Invoke("DeleteDisk", &args, &response)
	return err
}

type ReInitDiskArgs struct {
	DiskId string
}

type ReInitDiskResponse struct {
	CommonResponse
}

func (client *Client) ReInitDisk(diskId string) *ECSError {
	args := ReInitDiskArgs{
		DiskId: diskId,
	}
	response := ReInitDiskResponse{}
	err := client.Invoke("ReInitDisk", &args, &response)
	return err
}

type AttachDiskArgs struct {
	InstanceId         string
	DiskId             string
	Device             string
	DeleteWithInstance bool
}

type AttachDiskResponse struct {
	CommonResponse
}

func (client *Client) AttachDisk(args *AttachDiskArgs) *ECSError {
	response := AttachDiskResponse{}
	err := client.Invoke("AttachDisk", args, &response)
	return err
}

type DetachDiskArgs struct {
	InstanceId string
	DiskId     string
}

type DetachDiskResponse struct {
	CommonResponse
}

func (client *Client) DetachDisk(instanceId string, diskId string) *ECSError {
	args := DetachDiskArgs{
		InstanceId: instanceId,
		DiskId:     diskId,
	}
	response := DetachDiskResponse{}
	err := client.Invoke("DetachDisk", &args, &response)
	return err
}

type ResetDiskArgs struct {
	DiskId     string
	SnapshotId string
}

type ResetDiskResponse struct {
	CommonResponse
}

func (client *Client) ResetDisk(diskId string, snapshotId string) *ECSError {
	args := ResetDiskArgs{
		SnapshotId: snapshotId,
		DiskId:     diskId,
	}
	response := ResetDiskResponse{}
	err := client.Invoke("ResetDisk", &args, &response)
	return err
}

type ModifyDiskAttributeArgs struct {
	DiskId             string
	DiskName           string
	Description        string
	DeleteWithInstance *bool
	DeleteAutoSnapshot *bool
	EnableAutoSnapshot *bool
}

type ModifyDiskAttributeResponse struct {
	CommonResponse
}

func (client *Client) ModifyDiskAttribute(args *ModifyDiskAttributeArgs) *ECSError {
	response := ModifyDiskAttributeResponse{}
	err := client.Invoke("ModifyDiskAttribute", &args, &response)
	return err
}

const DISK_WAIT_FOR_INVERVAL = 5
const DISK_DEFAULT_TIME_OUT = 60

func (client *Client) WaitForDisk(regionId string, diskId string, status string, timeout int) *ECSError {
	if timeout <= 0 {
		timeout = DISK_DEFAULT_TIME_OUT
	}
	args := DescribeDisksArgs{
		RegionId: regionId,
		DiskIds:  []string{diskId},
	}

	for {
		disks, _, err := client.DescribeDisks(&args)
		if err != nil {
			return err
		}
		if disks == nil || len(disks) == 0 {
			return getECSErrorFromString("Not found")
		}
		if disks[0].Status == status {
			break
		}
		timeout = timeout - DISK_WAIT_FOR_INVERVAL
		if timeout <= 0 {
			return getECSErrorFromString("Timeout")
		}
		time.Sleep(DISK_WAIT_FOR_INVERVAL * time.Second)
	}
	return nil
}
