package nas

import "testing"

func TestClient_CreateAccessRule(t *testing.T) {
	client := NewTestClientForDebug()
	client.SetSecurityToken(TestSecurityToken)

	args := &CreateAccessRuleRequest{
		AccessGroupName: TestAccessGroupName,
		RegionId:        TestRegionID,
		SourceCidrIp:    "172.16.1.0/24",
	}

	response, err := client.CreateAccessRule(args)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Result = %++v", response)
	}
}
