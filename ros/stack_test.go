package ros

import (
	"testing"

	"os"
)

func TestClient_CreateStack(t *testing.T) {

}

func TestClient_DeleteStack(t *testing.T) {
	stackName := os.Getenv("StackName")
	stackId := os.Getenv("StackId")

	response, err := debugClientForTestCase.DeleteStack(TestRegionId, stackId, stackName)
	if err != nil {
		t.Fatalf("Failed to DeleteStack %++v", err)
	} else {
		t.Logf("Success %++v", response)
	}
}

func TestClient_AbandonStack(t *testing.T) {
	stackName := os.Getenv("StackName")
	stackId := os.Getenv("StackId")

	response, err := debugClientForTestCase.AbandonStack(TestRegionId, stackId, stackName)
	if err != nil {
		t.Fatalf("Failed to AbandonStack %++v", err)
	} else {
		t.Logf("Success %++v", response)
	}
}

func TestClient_DescribeStacks(t *testing.T) {
	args := &DescribeStacksRequest{
		RegionId: TestRegionId,
	}

	stacks, err := debugClientForTestCase.DescribeStacks(args)
	if err != nil {
		t.Fatalf("Failed to DescribeStacks %++v", err)
	} else {
		t.Logf("Response is %++v", stacks)
	}
}

func TestClient_DescribeStack(t *testing.T) {
	stackName := os.Getenv("StackName")
	stackId := os.Getenv("StackId")

	response, err := debugClientForTestCase.DescribeStack(TestRegionId, stackId, stackName)
	if err != nil {
		t.Fatalf("Failed to DescribeStack %++v", err)
	} else {
		t.Logf("Success %++v", response)
	}
}
