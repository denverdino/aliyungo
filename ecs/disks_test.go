package ecs

import (
	"testing"
)

func TestDisks(t *testing.T) {

	client := NewClient(TEST_ACCESS_KEY_ID, TEST_ACCESS_KEY_SECRET)

	instance, err := client.DescribeInstanceAttribute(TEST_INSTANCE_ID)
	if err != nil {
		t.Errorf("Failed to DescribeInstanceAttribute for instance %s: %v", TEST_INSTANCE_ID, err)
	}

	args := DescribeDisksArgs{}

	args.InstanceId = TEST_INSTANCE_ID
	args.RegionId = instance.RegionId
	disks, _, err := client.DescribeDisks(&args)

	if err != nil {
		t.Errorf("Failed to DescribeDisks for instance %s: %v", TEST_INSTANCE_ID, err)
	}

	for _, disk := range disks {
		t.Logf("Disk of intance %s: %++v", TEST_INSTANCE_ID, disk)
	}
}
