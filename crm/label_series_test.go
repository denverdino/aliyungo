package crm

import "testing"

func TestClient_QueryCustomerLabel(t *testing.T) {
	client := NewTestClientForDebug()
	client.SetSecurityToken(TestSecurityToken)

	response, err := client.QueryCustomerLabel(FINANCE_SERIES)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Result = %++v", response)
	}
}
