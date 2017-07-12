package ims

import "testing"

var (
	client   = NewTestClientForDebug()
	rmClient = NewTestRMClientForDebug()
)

func TestImsClient_CheckDirectoryEnabled(t *testing.T) {
	args := &CheckDirectoryEnabledRequest{}

	response, err := client.CheckDirectoryEnabled(args)
	if err != nil {
		t.Fatalf("Failed to CheckDirectoryEnabled")
	} else {
		t.Logf("Response is %++v", response)
	}
}

func TestImsClient_CheckDirectoryEnabled_V2(t *testing.T) {
	args := &CheckDirectoryEnabledRequest{}

	response, err := client.CheckDirectoryEnabled(args)
	if err != nil {
		t.Fatalf("Failed to CheckDirectoryEnabled")
	} else {
		t.Logf("Response is %++v", response)
	}
}

func TestImsClient_GetDirectory(t *testing.T) {
	args := &GetDirectoryRequest{}
	response, err := client.GetDirectory(args)
	if err != nil {
		t.Fatalf("Failed to GetDirectory %++v", err)
	} else {
		t.Logf("Response is %++v", response)
	}
}
