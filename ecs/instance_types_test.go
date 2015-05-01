package ecs

import (
	"testing"
)

func TestDescribeInstanceTypes(t *testing.T) {

	client := NewClient(TEST_ACCESS_KEY_ID, TEST_ACCESS_KEY_SECRET)
	instanceTypes, err := client.DescribeInstanceTypes()
	if err != nil {
		t.Errorf("Failed to DescribeInstanceTypes: %v", err)
	}
	for _, instanceType := range instanceTypes {
		t.Logf("InstanceType: %++v", instanceType)
	}
}
