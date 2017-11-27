package nas

import "testing"

func TestClient_DescribeFileSystems(t *testing.T) {
	client := NewTestClientForDebug()
	client.SetSecurityToken(TestSecurityToken)

	args := &DescribeFileSystemsRequest{
		RegionId: TestRegionID,
	}

	response, err := client.DescribeFileSystems(args)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Result = %++v", response)
	}
}
