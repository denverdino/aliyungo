package ecs

import (
	"testing"
)

func TestAllocatePublicIpAddress(t *testing.T) {

	client := NewClient(TestAccessKeyId, TestAccessKeySecret)
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

func testEipAddress(t *testing.T, client *Client, regionId Region, instanceId string) error {

	args := AllocateEipAddressArgs{
		RegionId:           regionId,
		Bandwidth:          5,
		InternetChargeType: PayByTraffic,
		ClientToken:        client.GenerateClientToken(),
	}
	ipAddr, allocationId, err := client.AllocateEipAddress(&args)
	if err != nil {
		t.Errorf("Failed to allocate EIP address: %v", err)
		return err
	}
	t.Logf("EIP address: %s, AllocationId: %s", ipAddr, allocationId)

	status, err := client.WaitForEip(regionId, allocationId, DefaultStrategy)
	if err != nil || status != EipStatusAvailable {
		t.Errorf("Failed to wait EIP %s: %v", allocationId, err)
	}

	err = client.AssociateEipAddress(allocationId, instanceId)
	if err != nil {
		t.Errorf("Failed to associate EIP address: %v", err)
	}
	status, err = client.WaitForEip(regionId, allocationId, DefaultStrategy)
	if err != nil || status != EipStatusInUse {
		t.Errorf("Failed to wait EIP %s: %v", allocationId, err)
	}
	err = client.UnassociateEipAddress(allocationId, instanceId)
	if err != nil {
		t.Errorf("Failed to unassociate EIP address: %v", err)
	}
	status, err = client.WaitForEip(regionId, allocationId, DefaultStrategy)
	if err != nil || status != EipStatusAvailable {
		t.Errorf("Failed to wait EIP %s: %v", allocationId, err)
	}
	err = client.ReleaseEipAddress(allocationId)
	if err != nil {
		t.Errorf("Failed to release EIP address: %v", err)
	}
	return err
}
