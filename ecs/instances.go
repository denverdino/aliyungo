package ecs

import (
	"github.com/denverdino/aliyungo/util"
	"time"
)

type DescribeInstanceStatusArgs struct {
	RegionId string
	ZoneId   string
	Pagination
}

type InstanceStatusItemType struct {
	InstanceId string
	Status     string
}

type DescribeInstanceStatusResponse struct {
	CommonResponse
	PaginationResult
	InstanceStatuses struct {
		InstanceStatus []InstanceStatusItemType
	}
}

func (client *Client) DescribeInstanceStatus(args *DescribeInstanceStatusArgs) (instanceStatuses []InstanceStatusItemType, pagination *PaginationResult, err *ECSError) {
	args.validate()
	response := DescribeInstanceStatusResponse{}

	err = client.Invoke("DescribeInstanceStatus", args, &response)

	if err == nil {
		return response.InstanceStatuses.InstanceStatus, &response.PaginationResult, nil
	}

	return nil, nil, err
}

type StopInstanceArgs struct {
	InstanceId string
	ForceStop  bool
}

type StopInstanceResponse struct {
	CommonResponse
}

func (client *Client) StopInstance(instanceId string, forceStop bool) *ECSError {
	args := StopInstanceArgs{
		InstanceId: instanceId,
		ForceStop:  forceStop,
	}
	response := StopInstanceResponse{}
	err := client.Invoke("StopInstance", &args, &response)
	return err
}

type StartInstanceArgs struct {
	InstanceId string
}

type StartInstanceResponse struct {
	CommonResponse
}

func (client *Client) StartInstance(instanceId string) *ECSError {
	args := StartInstanceArgs{InstanceId: instanceId}
	response := StartInstanceResponse{}
	err := client.Invoke("StartInstance", &args, &response)
	return err
}

type RebootInstanceArgs struct {
	InstanceId string
	ForceStop  bool
}

type RebootInstanceResponse struct {
	CommonResponse
}

func (client *Client) RebootInstance(instanceId string, forceStop bool) *ECSError {
	request := RebootInstanceArgs{
		InstanceId: instanceId,
		ForceStop:  forceStop,
	}
	response := RebootInstanceResponse{}
	err := client.Invoke("RebootInstance", &request, &response)
	return err
}

type DescribeInstanceAttributeArgs struct {
	InstanceId string
}

type OperationLocksType struct {
	LockReason []string //enum for financial, security
}

type SecurityGroupIdSetType struct {
	SecurityGroupId string
}

type IpAddressSetType struct {
	IpAddress []string
}
type VpcAttributesType struct {
	VpcId            string
	VSwitchId        string
	PrivateIpAddress IpAddressSetType
	NatIpAddress     string
}

type EipAddressAssociateType struct {
	AllocationId       string
	IpAddress          string
	Bandwidth          int
	InternetChargeType string
}

type InstanceAttributesType struct {
	InstanceId       string
	InstanceName     string
	Description      string
	ImageId          string
	RegionId         string
	ZoneId           string
	ClusterId        string
	InstanceType     string
	HostName         string
	Status           string
	OperationLocks   OperationLocksType
	SecurityGroupIds struct {
		SecurityGroupId []string
	}
	PublicIpAddress         IpAddressSetType
	InnerIpAddress          IpAddressSetType
	InstanceNetworkType     string //enum Classic | Vpc
	InternetMaxBandwidthIn  int
	InternetMaxBandwidthOut int
	InternetChargeType      string           //enum PayByBandwidth | PayByTraffic
	CreationTime            util.ISO6801Time //time.Time
	VpcAttributes           VpcAttributesType
	EipAddress              EipAddressAssociateType
}

type DescribeInstanceAttributeResponse struct {
	CommonResponse
	InstanceAttributesType
}

func (client *Client) DescribeInstanceAttribute(instanceId string) (instance *InstanceAttributesType, err *ECSError) {
	args := DescribeInstanceAttributeArgs{InstanceId: instanceId}

	response := DescribeInstanceAttributeResponse{}
	err = client.Invoke("DescribeInstanceAttribute", &args, &response)
	if err != nil {
		return nil, err
	}
	return &response.InstanceAttributesType, err
}

const INSTANCE_WAIT_FOR_INVERVAL = 5
const INSTANCE_DEFAULT_TIME_OUT = 60

func (client *Client) WaitForInstance(instanceId string, status string, timeout int) *ECSError {
	if timeout <= 0 {
		timeout = INSTANCE_DEFAULT_TIME_OUT
	}
	for {
		instance, err := client.DescribeInstanceAttribute(instanceId)
		if err != nil {
			return err
		}
		if instance.Status == status {
			break
		}
		timeout = timeout - INSTANCE_WAIT_FOR_INVERVAL
		if timeout <= 0 {
			return getECSErrorFromString("Timeout")
		}
		time.Sleep(INSTANCE_WAIT_FOR_INVERVAL * time.Second)

	}
	return nil
}

type DescribeInstancesArgs struct {
	RegionId            string
	VpcId               string
	VSwitchId           string
	ZoneId              string
	InstanceIds         string
	InstanceNetworkType string
	PrivateIpAddresses  string
	InnerIpAddresses    string
	PublicIpAddresses   string
	SecurityGroupId     string
	Pagination
}

type DescribeInstancesResponse struct {
	CommonResponse
	PaginationResult
	Instances struct {
		Instance []InstanceAttributesType
	}
}

func (client *Client) DescribeInstances(args *DescribeInstancesArgs) (instances []InstanceAttributesType, pagination *PaginationResult, err *ECSError) {
	args.validate()
	response := DescribeInstancesResponse{}

	err = client.Invoke("DescribeInstances", args, &response)

	if err == nil {
		return response.Instances.Instance, &response.PaginationResult, nil
	}

	return nil, nil, err
}

type DeleteInstanceArgs struct {
	InstanceId string
}

type DeleteInstanceResponse struct {
	CommonResponse
}

func (client *Client) DeleteInstance(instanceId string) *ECSError {
	args := DeleteInstanceArgs{InstanceId: instanceId}
	response := DeleteInstanceResponse{}
	err := client.Invoke("DeleteInstance", &args, &response)
	return err
}

type DataDiskType struct {
	Size               int
	Category           string //Enum cloud, ephemeral, ephemeral_ssd
	SnapshotId         string
	DiskName           string
	Description        string
	Device             string
	DeleteWithInstance bool
}

type CreateInstanceArgs struct {
	RegionId                string
	ZoneId                  string
	ImageId                 string
	InstanceType            string
	SecurityGroupId         string
	InstanceName            string
	Description             string
	InternetChargeType      string
	InternetMaxBandwidthIn  int
	InternetMaxBandwidthOut int
	HostName                string
	Password                string
	SystemDisk_Category     string `ArgName:"SystemDisk.Category"`
	SystemDisk_DiskName     string `ArgName:"SystemDisk.DiskName"`
	SystemDisk_Description  string `ArgName:"SystemDisk.Description"`
	DataDisk                []DataDiskType
	VSwitchId               string
	PrivateIpAddress        string
	ClientToken             string
}

type CreateInstanceResponse struct {
	CommonResponse
	InstanceId string
}

func (client *Client) CreateInstance(args *CreateInstanceArgs) (instanceId string, err *ECSError) {
	response := CreateInstanceResponse{}
	err = client.Invoke("CreateInstance", args, &response)
	if err != nil {
		return "", err
	}
	return response.InstanceId, err
}
