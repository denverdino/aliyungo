package ecs

import (
	"testing"
)

func aTestDisks(t *testing.T) {

	client := NewClient(TEST_ACCESS_KEY_ID, TEST_ACCESS_KEY_SECRET)

	instance, err := client.DescribeInstanceAttribute(TEST_INSTANCE_ID)
	if err != nil {
		t.Fatalf("Failed to DescribeInstanceAttribute for instance %s: %v", TEST_INSTANCE_ID, err)
	}

	args := DescribeDisksArgs{}

	args.InstanceId = TEST_INSTANCE_ID
	args.RegionId = instance.RegionId
	disks, _, err := client.DescribeDisks(&args)

	if err != nil {
		t.Fatalf("Failed to DescribeDisks for instance %s: %v", TEST_INSTANCE_ID, err)
	}

	for _, disk := range disks {
		t.Logf("Disk of instance %s: %++v", TEST_INSTANCE_ID, disk)
	}
}

func TestDiskCreationAndDeletion(t *testing.T) {

	if TEST_I_AM_RICH == false { //Avoid payment
		return
	}

	client := NewClient(TEST_ACCESS_KEY_ID, TEST_ACCESS_KEY_SECRET)

	instance, err := client.DescribeInstanceAttribute(TEST_INSTANCE_ID)
	if err != nil {
		t.Fatalf("Failed to DescribeInstanceAttribute for instance %s: %v", TEST_INSTANCE_ID, err)
	}

	args := CreateDiskArgs{
		RegionId: instance.RegionId,
		ZoneId:   instance.ZoneId,
		DiskName: "test-disk",
		Size:     5,
	}

	diskId, err := client.CreateDisk(&args)
	if err != nil {
		t.Fatalf("Failed to create disk: %v", err)
	}
	t.Logf("Create disk %s successfully", diskId)

	attachArgs := AttachDiskArgs{
		InstanceId: instance.InstanceId,
		DiskId:     diskId,
	}

	err = client.AttachDisk(&attachArgs)
	if err != nil {
		t.Errorf("Failed to create disk: %v", err)
	} else {
		t.Logf("Attach disk %s to instance %s successfully", diskId, instance.InstanceId)

		instance, err = client.DescribeInstanceAttribute(TEST_INSTANCE_ID)
		if err != nil {
			t.Errorf("Failed to DescribeInstanceAttribute for instance %s: %v", TEST_INSTANCE_ID, err)
		} else {
			t.Logf("Instance: %++v  %v", instance, err)
		}
		err = client.WaitForDisk(instance.RegionId, diskId, DISK_STATUS_IN_USE, 0)
		if err != nil {
			t.Fatalf("Failed to wait for disk %s to status %s: %v", diskId, DISK_STATUS_IN_USE, err)
		}
		err = client.DetachDisk(instance.InstanceId, diskId)
		if err != nil {
			t.Errorf("Failed to detach disk: %v", err)
		} else {
			t.Logf("Detach disk %s to instance %s successfully", diskId, instance.InstanceId)
		}

		err = client.WaitForDisk(instance.RegionId, diskId, DISK_STATUS_AVAILABLE, 0)
		if err != nil {
			t.Fatalf("Failed to wait for disk %s to status %s: %v", diskId, DISK_STATUS_AVAILABLE, err)
		}
	}
	err = client.DeleteDisk(diskId)
	if err != nil {
		t.Fatalf("Failed to delete disk %s: %v", diskId, err)
	}
	t.Logf("Delete disk %s successfully", diskId)
}
