package ecs

import (
	"testing"
)

func TestGenerateClientToken(t *testing.T) {
	client := NewClient(TEST_ACCESS_KEY_ID, TEST_ACCESS_KEY_SECRET)
	for i := 0; i < 10; i++ {
		t.Log("GenerateClientToken: ", client.GenerateClientToken())
	}

}

func TestECSDescribe(t *testing.T) {
	client := NewClient(TEST_ACCESS_KEY_ID, TEST_ACCESS_KEY_SECRET)

	regions, err := client.DescribeRegions()

	t.Log("regions: ", regions, err)

	for _, region := range regions {
		zones, err := client.DescribeZones(region.RegionId)
		t.Log("zones: ", zones, err)
		for _, zone := range zones {
			args := DescribeInstanceStatusArgs{
				RegionId: region.RegionId,
				ZoneId:   zone.ZoneId,
			}
			instanceStatuses, pagination, err := client.DescribeInstanceStatus(&args)
			t.Log("instanceStatuses: ", instanceStatuses, pagination, err)
			for _, instanceStatus := range instanceStatuses {
				instance, err := client.DescribeInstanceAttribute(instanceStatus.InstanceId)
				t.Logf("Intances: %++v", instance)
				t.Logf("Error: %++v", err)
			}
			args1 := DescribeInstancesArgs{
				RegionId: region.RegionId,
				ZoneId:   zone.ZoneId,
			}
			instances, _, err := client.DescribeInstances(&args1)
			if err != nil {
				t.Errorf("Failed to describe instance %s %s", region.RegionId, zone.ZoneId)
			} else {
				for _, instance := range instances {
					t.Logf("Intances: %++v", instance)
				}
			}

		}
		args := DescribeImagesArgs{RegionId: region.RegionId}

		images, pagination, err := client.DescribeImages(&args)
		t.Logf("Total image count for region %s: %d\n", region.RegionId, pagination.TotalCount)

		for _, image := range images {
			t.Logf("Image: %++v", image)
		}
		t.Logf("Error: %++v", err)
	}
}

func TestECSInstance(t *testing.T) {

	client := NewClient(TEST_ACCESS_KEY_ID, TEST_ACCESS_KEY_SECRET)
	instance, err := client.DescribeInstanceAttribute(TEST_INSTANCE_ID)
	t.Logf("Intance: %++v  %v\n", instance, err)
	err = client.StopInstance(TEST_INSTANCE_ID, false)
	if err != nil {
		t.Errorf("Failed to stop instance %s: %v", TEST_INSTANCE_ID, err)
	}
	err = client.WaitForInstance(TEST_INSTANCE_ID, "Stopped")
	if err != nil {
		t.Errorf("Intance %s is failed to stop: %v", TEST_INSTANCE_ID, err)
	}
	t.Logf("Intance %s is stopped successfully.", TEST_INSTANCE_ID)
	err = client.StartInstance(TEST_INSTANCE_ID)
	if err != nil {
		t.Errorf("Failed to start instance %s: %v", TEST_INSTANCE_ID, err)
	}
	err = client.WaitForInstance(TEST_INSTANCE_ID, "Running")
	if err != nil {
		t.Errorf("Intance %s is failed to start: %v", TEST_INSTANCE_ID, err)
	}
	t.Logf("Intance %s is running successfully.", TEST_INSTANCE_ID)
	err = client.RebootInstance(TEST_INSTANCE_ID, true)
	if err != nil {
		t.Errorf("Failed to restart instance %s: %v", TEST_INSTANCE_ID, err)
	}
	err = client.WaitForInstance(TEST_INSTANCE_ID, "Running")
	if err != nil {
		t.Errorf("Intance %s is failed to restart: %v", TEST_INSTANCE_ID, err)
	}
	t.Logf("Intance %s is running successfully.", TEST_INSTANCE_ID)
}
