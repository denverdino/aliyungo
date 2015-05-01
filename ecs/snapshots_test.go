package ecs

import (
	"testing"
)

func TestSnapshot(t *testing.T) {

	client := NewClient(TEST_ACCESS_KEY_ID, TEST_ACCESS_KEY_SECRET)

	instance, err := client.DescribeInstanceAttribute(TEST_INSTANCE_ID)
	if err != nil {
		t.Errorf("Failed to DescribeInstanceAttribute for instance %s: %v", TEST_INSTANCE_ID, err)
	}

	args := DescribeSnapshotsArgs{}

	args.InstanceId = TEST_INSTANCE_ID
	args.RegionId = instance.RegionId
	snapshots, _, err := client.DescribeSnapshots(&args)

	if err != nil {
		t.Errorf("Failed to DescribeSnapshots for instance %s: %v", TEST_INSTANCE_ID, err)
	}

	for _, snapshot := range snapshots {
		t.Logf("Snapshot of instance %s: %++v", TEST_INSTANCE_ID, snapshot)
	}
}

func aTestSnapshotCreationAndDeletion(t *testing.T) {

	client := NewClient(TEST_ACCESS_KEY_ID, TEST_ACCESS_KEY_SECRET)

	instance, err := client.DescribeInstanceAttribute(TEST_INSTANCE_ID)
	if err != nil {
		t.Errorf("Failed to DescribeInstanceAttribute for instance %s: %v", TEST_INSTANCE_ID, err)
	}

	//TODO
	diskId := "d-25z6kd44o"

	args := CreateSnapshotArgs{
		DiskId:       diskId,
		SnapshotName: "My_Test_Snapshot",
		Description:  "My Test Snapshot Description",
		ClientToken:  client.GenerateClientToken(),
	}

	snapshotId, err := client.CreateSnapshot(&args)
	if err != nil {
		t.Errorf("Failed to CreateSnapshot for disk %s: %v", diskId, err)
	}
	client.WaitForSnapShotReady(instance.RegionId, snapshotId, 0)

	err = client.DeleteSnapshot(snapshotId)
	if err != nil {
		t.Errorf("Failed to DeleteSnapshot for disk %s: %v", diskId, err)
	}

	t.Logf("Snapshot %s is deleted successfully.", snapshotId)

}
