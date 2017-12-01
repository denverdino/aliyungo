package nas

import "testing"

func TestClient_DescribeMountTargets(t *testing.T) {

	client := NewTestClientForDebug()
	client.SetSecurityToken(TestSecurityToken)

	args := &DescribeMountTargetsRequest{
		FileSystemId: TestFileSystemId,
		RegionId:     TestRegionID,
	}

	response, err := client.DescribeMountTargets(args)

	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Result = %++v", response)
	}
}
