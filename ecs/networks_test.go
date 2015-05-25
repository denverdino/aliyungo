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

func _TestEipAddress(t *testing.T) {

	client := NewClient(TestAccessKeyId, TestAccessKeySecret)
	client.SetDebug(true)
	instance, err := client.DescribeInstanceAttribute(TestInstanceId)
	if err != nil {
		t.Fatalf("Failed to describe instance %s: %v", TestInstanceId, err)
	}
	args := AllocateEipAddressArgs{
		RegionId:           instance.RegionId,
		Bandwidth:          5,
		InternetChargeType: PayByTraffic,
		ClientToken:        client.GenerateClientToken(),
	}
	ipAddr, allocationId, err := client.AllocateEipAddress(&args)
	if err != nil {
		t.Fatalf("Failed to allocate EIP address: %v", err)
	}
	t.Logf("EIP address: %s, AllocationId: %s", ipAddr, allocationId)

	err = client.AssociateEipAddress(allocationId, TestInstanceId)
	if err != nil {
		t.Errorf("Failed to associate EIP address: %v", err)
	}
	client.UnassociateEipAddress(allocationId, TestInstanceId)
	if err != nil {
		t.Errorf("Failed to unassociate EIP address: %v", err)
	}
	client.ReleaseEipAddress(allocationId)
	if err != nil {
		t.Errorf("Failed to release EIP address: %v", err)
	}
}
