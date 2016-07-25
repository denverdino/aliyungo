//go test -v config_info_test.go client.go networks.go networks_test.go instances.go tags.go disks.go
package ecs

import (
	"testing"

	"github.com/hdksky/aliyungo/common"
)

func TestAllocatePublicIpAddress(t *testing.T) {

	client := NewTestClient()
	instance, err := client.DescribeInstanceAttribute(TestInstanceId)
	if err != nil {
		t.Fatalf("Failed to describe instance %s: %v", TestInstanceId, err)
	}
	t.Logf("Instance: %++v  %v", instance, err)
	ipAddr, err := client.AllocatePublicIpAddress(TestInstanceId)
	if err != nil {
		t.Fatalf("Failed to allocate public IP address for instance %s: %v", TestInstanceId, err)
	}
	t.Logf("Public IP address of instance %s: %s", TestInstanceId, ipAddr)

}

func testEipAddress(t *testing.T, client *Client, regionId common.Region, instanceId string) error {

	args := AllocateEipAddressArgs{
		RegionId:           regionId,
		Bandwidth:          5,
		InternetChargeType: common.PayByTraffic,
		ClientToken:        client.GenerateClientToken(),
	}
	ipAddr, allocationId, err := client.AllocateEipAddress(&args)
	if err != nil {
		t.Errorf("Failed to allocate EIP address: %v", err)
		return err
	}
	t.Logf("EIP address: %s, AllocationId: %s", ipAddr, allocationId)

	err = client.WaitForEip(regionId, allocationId, EipStatusAvailable, 0)
	if err != nil {
		t.Errorf("Failed to wait EIP %s: %v", allocationId, err)
	}

	info := AssociateEipAddressArgs{
		AllocationId: allocationId,
		InstanceId:   instanceId,
		InstanceType: "EcsInstance",
	}
	err = client.AssociateEipAddress(&info)
	if err != nil {
		t.Errorf("Failed to associate EIP address: %v", err)
	}
	err = client.WaitForEip(regionId, allocationId, EipStatusInUse, 0)
	if err != nil {
		t.Errorf("Failed to wait EIP %s: %v", allocationId, err)
	}

	data := UnallocateEipAddressArgs{
		AllocationId: allocationId,
		InstanceId:   instanceId,
		InstanceType: "EcsInstance",
	}
	err = client.UnassociateEipAddress(&data)
	if err != nil {
		t.Errorf("Failed to unassociate EIP address: %v", err)
	}
	err = client.WaitForEip(regionId, allocationId, EipStatusAvailable, 0)
	if err != nil {
		t.Errorf("Failed to wait EIP %s: %v", allocationId, err)
	}
	err = client.ReleaseEipAddress(allocationId)
	if err != nil {
		t.Errorf("Failed to release EIP address: %v", err)
	}
	return err
}
