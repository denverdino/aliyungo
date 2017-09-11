package ros

import (
	"os"
	"testing"
)

func TestClient_DescribeEvents(t *testing.T) {
	stackName := os.Getenv("StackName")
	stackId := os.Getenv("StackId")

	response, err := debugClientForTestCase.DescribeEvents(stackId, stackName, &DescribeEventsRequest{})
	if err != nil {
		t.Fatalf("Failed to DescribeEvents %++v", err)
	} else {
		t.Logf("Resource = %++v", response)
	}
}
