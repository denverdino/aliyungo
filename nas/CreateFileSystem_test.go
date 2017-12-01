package nas

import "testing"

func TestClient_CreateFileSystem(t *testing.T) {
	client := NewTestClientForDebug()
	client.SetSecurityToken(TestSecurityToken)

	args := &CreateFileSystemRequest{
		RegionId:     TestRegionID,
		StorageType:  "Performance",
		ProtocolType: "NFS",
		Description:  "my-test-filestyem",
	}

	response, err := client.CreateFileSystem(args)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Result = %++v", response)
	}
}
