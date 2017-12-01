package nas

import "testing"

func TestClient_CreateMountTarget(t *testing.T) {
	client := NewTestClientForDebug()
	client.SetSecurityToken(TestSecurityToken)

	args := &CreateMountTargetRequest{
		RegionId:        TestRegionID,
		FileSystemId:    TestFileSystemId,
		AccessGroupName: TestAccessGroupName,
		NetworkType:     "Vpc",
		VpcId:           TestVpcId,
		VSwitchId:       TestVSwitchId,
	}

	response, err := client.CreateMountTarget(args)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Result = %++v", response)
	}
}
