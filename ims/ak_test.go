package ims

import "testing"

func TestImsClient_CreateAccessKey(t *testing.T) {
	args := &CreateAccessKeyRequest{
		UserPrincipalName: "user-1499677896@1312623533146984.onaliyun.com",
	}

	response, err := client.CreateAccessKey(args)
	if err != nil {
		t.Fatalf("Failed to CreateAccessKey %++v", err)
	} else {
		t.Logf("Response = %++v", response)
	}
}

func TestImsClient_DeleteAccessKey(t *testing.T) {
	createArgs := &CreateAccessKeyRequest{
		UserPrincipalName: "user-1499677896@1312623533146984.onaliyun.com",
	}

	response, err := client.CreateAccessKey(createArgs)
	if err != nil {
		t.Fatalf("Failed to CreateAccessKey %++v", err)
	} else {
		args := &DeleteAccessKeyRequest{
			UserAccessKeyId:   response.AccessKey.AccessKeyId,
			UserPrincipalName: createArgs.UserPrincipalName,
		}

		resp, err := client.DeleteAccessKey(args)
		if err != nil {
			t.Fatalf("Failed to DeleteAccessKey %++v", err)
		} else {
			t.Logf("Response = %++v", resp)
		}
	}

}

func TestImsClient_ListAccessKeys(t *testing.T) {
	args := &ListAccessKeysRequest{
		UserPrincipalName: "user-1499677896@1312623533146984.onaliyun.com",
	}

	response, err := client.ListAccessKeys(args)
	if err != nil {
		t.Fatalf("Failed to ListAccessKeys %++v", err)
	} else {
		for index, ak := range response.AccessKeys.AccessKey {
			t.Logf("ak[%d] = %++v", index, ak)
		}
	}
}

func TestImsClient_UpdateAccessKey(t *testing.T) {
	args := &UpdateAccessKeyRequest{
		UserPrincipalName: "user-1499677896@1312623533146984.onaliyun.com",
		UserAccessKeyId:   "LTAIrZw0EpolY1k8",
		Status:            "Active",
	}

	response, err := client.UpdateAccessKey(args)
	if err != nil {
		t.Fatalf("Failed to UpdateAccessKey %++v", err)
	} else {
		t.Logf("Response = %++v", response)
	}
}
