package ecs

import (
	"testing"
)

func TestAllocatePublicIpAddress(t *testing.T) {

	client := NewClient(TEST_ACCESS_KEY_ID, TEST_ACCESS_KEY_SECRET)
	instance, err := client.DescribeInstanceAttribute(TEST_INSTANCE_ID)
	if err != nil {
		t.Fatalf("Failed to describe instance %s: %v", TEST_INSTANCE_ID, err)
	}
	t.Logf("Instance: %++v  %v", instance, err)
	ipAddr, err := client.AllocatePublicIpAddress(TEST_INSTANCE_ID)
	if err != nil {
		t.Fatalf("Failed to allocate public IP address for instance %s: %v", TEST_INSTANCE_ID, err)
	}
	t.Logf("Public IP address of instance %s: %s", TEST_INSTANCE_ID, ipAddr)

}
