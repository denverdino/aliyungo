package ecs

import (
	"testing"
)

func TestECSInstance(t *testing.T) {

	client := NewClient(TEST_ACCESS_KEY_ID, TEST_ACCESS_KEY_SECRET)
	instance, err := client.DescribeInstanceAttribute(TEST_INSTANCE_ID)
	t.Logf("Instance: %++v  %v", instance, err)
	err = client.StopInstance(TEST_INSTANCE_ID, false)
	if err != nil {
		t.Errorf("Failed to stop instance %s: %v", TEST_INSTANCE_ID, err)
	}
	err = client.WaitForInstance(TEST_INSTANCE_ID, "Stopped", 0)
	if err != nil {
		t.Errorf("Instance %s is failed to stop: %v", TEST_INSTANCE_ID, err)
	}
	t.Logf("Instance %s is stopped successfully.", TEST_INSTANCE_ID)
	err = client.StartInstance(TEST_INSTANCE_ID)
	if err != nil {
		t.Errorf("Failed to start instance %s: %v", TEST_INSTANCE_ID, err)
	}
	err = client.WaitForInstance(TEST_INSTANCE_ID, "Running", 0)
	if err != nil {
		t.Errorf("Instance %s is failed to start: %v", TEST_INSTANCE_ID, err)
	}
	t.Logf("Instance %s is running successfully.", TEST_INSTANCE_ID)
	err = client.RebootInstance(TEST_INSTANCE_ID, true)
	if err != nil {
		t.Errorf("Failed to restart instance %s: %v", TEST_INSTANCE_ID, err)
	}
	err = client.WaitForInstance(TEST_INSTANCE_ID, "Running", 0)
	if err != nil {
		t.Errorf("Instance %s is failed to restart: %v", TEST_INSTANCE_ID, err)
	}
	t.Logf("Instance %s is running successfully.", TEST_INSTANCE_ID)
}

func TestECSInstanceCreationAndDeletion(t *testing.T) {

	if TEST_I_AM_RICH == false { // Avoid payment
		return
	}

	client := NewClient(TEST_ACCESS_KEY_ID, TEST_ACCESS_KEY_SECRET)
	instance, err := client.DescribeInstanceAttribute(TEST_INSTANCE_ID)
	t.Logf("Instance: %++v  %v", instance, err)

	args := CreateInstanceArgs{
		RegionId:        instance.RegionId,
		ImageId:         instance.ImageId,
		InstanceType:    "ecs.t1.small",
		SecurityGroupId: instance.SecurityGroupIds.SecurityGroupId[0],
	}

	instanceId, err := client.CreateInstance(&args)
	if err != nil {
		t.Errorf("Failed to create instance from Image %s: %v", args.ImageId, err)
	}
	t.Logf("Instance %s is created successfully.", instanceId)

	instance, err = client.DescribeInstanceAttribute(instanceId)
	t.Logf("Instance: %++v  %v", instance, err)
	err = client.StartInstance(instanceId)
	if err != nil {
		t.Errorf("Failed to start instance %s: %v", instanceId, err)
	}
	err = client.WaitForInstance(instanceId, "Running", 0)

	err = client.StopInstance(instanceId, true)
	if err != nil {
		t.Errorf("Failed to stop instance %s: %v", instanceId, err)
	}
	err = client.WaitForInstance(instanceId, "Stopped", 0)
	if err != nil {
		t.Errorf("Instance %s is failed to stop: %v", instanceId, err)
	}
	t.Logf("Instance %s is stopped successfully.", instanceId)

	err = client.DeleteInstance(instanceId)

	if err != nil {
		t.Errorf("Failed to delete instance %s: %v", instanceId, err)
	}
	t.Logf("Instance %s is deleted successfully.", instanceId)
}
