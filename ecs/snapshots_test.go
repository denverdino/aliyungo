package ecs

import (
	"testing"
	"time"
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
		t.Logf("Snapshot of intance %s: %++v", TEST_INSTANCE_ID, snapshot)
	}
}

func waitSnapShotReady(t *testing.T, client *Client, regionId string, snapshotId string) {
	for {
		args1 := DescribeSnapshotsArgs{}

		args1.InstanceId = TEST_INSTANCE_ID
		args1.RegionId = regionId
		args1.SnapshotIds = []string{snapshotId}
		snapshots, _, err := client.DescribeSnapshots(&args1)

		if err != nil {
			t.Errorf("Failed to DescribeSnapshots for instance %s: %v", TEST_INSTANCE_ID, err)
		}

		for _, snapshot := range snapshots {
			t.Logf("Snapshot of intance %s: %++v", TEST_INSTANCE_ID, snapshot)
			if snapshot.Progress == "100%" {
				return
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func TestSnapshotCreationAndDeletion(t *testing.T) {

	client := NewClient(TEST_ACCESS_KEY_ID, TEST_ACCESS_KEY_SECRET)

	instance, err := client.DescribeInstanceAttribute(TEST_INSTANCE_ID)
	if err != nil {
		t.Errorf("Failed to DescribeInstanceAttribute for instance %s: %v", TEST_INSTANCE_ID, err)
	}

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
	waitSnapShotReady(t, client, instance.RegionId, snapshotId)

	err = client.DeleteSnapshot(snapshotId)
	if err != nil {
		t.Errorf("Failed to DeleteSnapshot for disk %s: %v", diskId, err)
	}

	t.Logf("Snapshot %s is deleted successfully.", snapshotId)

}
