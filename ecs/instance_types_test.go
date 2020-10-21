package ecs

import (
	"testing"
)

func TestDescribeInstanceTypes(t *testing.T) {
	client := NewTestClient()
	instanceTypes, err := client.DescribeInstanceTypes()
	if err != nil {
		t.Fatalf("Failed to DescribeInstanceTypes: %v", err)
	}
	for _, instanceType := range instanceTypes {
		t.Logf("InstanceType: %++v", instanceType)
	}

	instanceTypes, err = client.DescribeInstanceTypesNew(&DescribeInstanceTypesArgs{
		InstanceTypes: []string{"ecs.ec5.24xlarge", "ecs.ddh6s.custom.c4m48"},
	})
	if err != nil {
		t.Fatalf("Failed to DescribeInstanceTypesNew: %v", err)
	}
	for _, instanceType := range instanceTypes {
		t.Logf("InstanceType: %++v", instanceType)
	}
}
