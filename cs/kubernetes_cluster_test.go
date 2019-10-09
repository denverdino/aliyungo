package cs

import "testing"

func Test_ModifyCluster(t *testing.T) {
	client := NewTestClientForDebug()

	args := &ModifyClusterArgs{
		DeletionProtection: false,
	}

	err := client.ModifyCluster(TestClusterId, args)
	if err != nil {
		t.Errorf("Error %++v", err)
	} else {
		t.Logf("OK")
	}
}
