package nas

import "testing"

func TestClient_DescribeAccessRules(t *testing.T) {
	client := NewTestClientForDebug()
	client.SetSecurityToken(TestSecurityToken)

	args := &DescribeAccessRulesRequest{
		AccessGroupName: TestAccessGroupName,
		RegionId:        TestRegionID,
	}

	response, err := client.DescribeAccessRules(args)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Result = %++v", response)
	}
}
