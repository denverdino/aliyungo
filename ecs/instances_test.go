package ecs

import (
	"testing"
)
func ExampleClient_DescribeInstanceStatus(t *testing.T) {
	t.Logf("DescribeInstanceStatus Example\n")

	pagination := &Pagination{1, 1}
	regionId := Region("cn-beijing")
	zoneId := "cn-beijing-b"

	var describeInstanceStatusArgs DescribeInstanceStatusArgs
	describeInstanceStatusArgs.RegionId = regionId
	describeInstanceStatusArgs.ZoneId = zoneId
	describeInstanceStatusArgs.Pagination = *pagination

	client := NewClient(TestAccessKeyId, TestAccessKeySecret)
	instanceStatus, _, err := client.DescribeInstanceStatus(&describeInstanceStatusArgs)

	if err != nil {
		t.Errorf("Failed to describe Instance: %s status:%v", TestInstanceId, err)
	} else {
		for i := 0; i < len(instanceStatus); i++ {
			t.Logf("Instance %s Status: %s ", instanceStatus[i].InstanceId, instanceStatus[i].Status)
		}
	}
}

func ExampleClient_DescribeInstanceAttribute(t *testing.T) {
	t.Logf("DescribeInstanceAttribute Example\n")

	client := NewClient(TestAccessKeyId, TestAccessKeySecret)

	instanceAttributeType, err := client.DescribeInstanceAttribute(TestInstanceId)

	if err != nil {
		t.Errorf("Failed to describe Instance %s attribute: %v", TestInstanceId, err)
	} else {
		t.Logf("Instance Information")
		t.Logf("InstanceId = %s ", instanceAttributeType.InstanceId)
		t.Logf("InstanceName = %s ", instanceAttributeType.InstanceName)
		t.Logf("HostName = %s ", instanceAttributeType.HostName)
		t.Logf("ZoneId = %s ", instanceAttributeType.ZoneId)
		t.Logf("RegionId = %s ", instanceAttributeType.RegionId)
	}
}

func ExampleClient_DescribeInstanceVncUrl(t *testing.T) {
	t.Logf("DescribeInstanceVncUrl Example\n")

	region := Region("cn-beijing")

	var describeInstanceVncUrlArgs DescribeInstanceVncUrlArgs
	describeInstanceVncUrlArgs.RegionId = region
	describeInstanceVncUrlArgs.InstanceId = TestInstanceId

	client := NewClient(TestAccessKeyId, TestAccessKeySecret)

	instanceVncUrl, err := client.DescribeInstanceVncUrl(&describeInstanceVncUrlArgs)

	if err != nil {
		t.Errorf("Failed to describe Instance %s vnc url: %v", TestInstanceId, err)
	} else {
		t.Logf("VNC URL = %s ", instanceVncUrl)
	}
}

func ExampleClient_StopInstance(t *testing.T) {
	t.Logf("Stop Instance Example\n")

	client := NewClient(TestAccessKeyId, TestAccessKeySecret)

	err := client.StopInstance(TestInstanceId, true)

	if err != nil {
		t.Errorf("Failed to stop Instance %s vnc url: %v", TestInstanceId, err)
	}
}

func ExampleClient_DeleteInstance(t *testing.T) {
	t.Logf("Delete Instance Example")
	client := NewClient(TestAccessKeyId, TestAccessKeySecret)

	err := client.DeleteInstance(TestInstanceId)

	if err != nil {
		t.Errorf("Failed to delete Instance %s vnc url: %v", TestInstanceId, err)
	}
}

func TestECSInstance(t *testing.T) {

	client := NewClient(TestAccessKeyId, TestAccessKeySecret)
	instance, err := client.DescribeInstanceAttribute(TestInstanceId)
	if err != nil {
		t.Fatalf("Failed to describe instance %s: %v", TestInstanceId, err)
	}
	t.Logf("Instance: %++v  %v", instance, err)
	err = client.StopInstance(TestInstanceId, true)
	if err != nil {
		t.Errorf("Failed to stop instance %s: %v", TestInstanceId, err)
	}
	err = client.WaitForInstance(TestInstanceId, Stopped, 0)
	if err != nil {
		t.Errorf("Instance %s is failed to stop: %v", TestInstanceId, err)
	}
	t.Logf("Instance %s is stopped successfully.", TestInstanceId)
	err = client.StartInstance(TestInstanceId)
	if err != nil {
		t.Errorf("Failed to start instance %s: %v", TestInstanceId, err)
	}
	err = client.WaitForInstance(TestInstanceId, Running, 0)
	if err != nil {
		t.Errorf("Instance %s is failed to start: %v", TestInstanceId, err)
	}
	t.Logf("Instance %s is running successfully.", TestInstanceId)
	err = client.RebootInstance(TestInstanceId, true)
	if err != nil {
		t.Errorf("Failed to restart instance %s: %v", TestInstanceId, err)
	}
	err = client.WaitForInstance(TestInstanceId, Running, 0)
	if err != nil {
		t.Errorf("Instance %s is failed to restart: %v", TestInstanceId, err)
	}
	t.Logf("Instance %s is running successfully.", TestInstanceId)
}

func TestECSInstanceCreationAndDeletion(t *testing.T) {

	if TestIAmRich == false { // Avoid payment
		return
	}

	client := NewClient(TestAccessKeyId, TestAccessKeySecret)
	instance, err := client.DescribeInstanceAttribute(TestInstanceId)
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

	err = client.WaitForInstance(instanceId, Stopped, 60)

	err = client.StartInstance(instanceId)
	if err != nil {
		t.Errorf("Failed to start instance %s: %v", instanceId, err)
	}
	err = client.WaitForInstance(instanceId, Running, 0)

	err = client.StopInstance(instanceId, true)
	if err != nil {
		t.Errorf("Failed to stop instance %s: %v", instanceId, err)
	}
	err = client.WaitForInstance(instanceId, Stopped, 0)
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
