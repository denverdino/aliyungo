package ims

import (
	"fmt"
	"testing"
	"time"
)

func TestResourceManagerClient_CreatePolicy(t *testing.T) {
	policy := "{\"Statement\": [{\"Action\": \"cs:*\",\"Effect\": \"Allow\",\"Resource\": [\"acs:cs:*:*:cluster/c64e1330717b647d891a389464bca4584\"]}],\"Version\": \"1\"}"
	args := &CreatePolicyRequest{
		//AccountId:      "acs-ram-account-c64e1330717b647d891a389464bca4584@1312623533146984.onaliyun.com",
		//PolicyName:     fmt.Sprintf("%s_%d", "AliyunACSResourcesAccess_user", time.Now().Unix()),
		PolicyName:     "AliyunACSDefaultAccess",
		Description:    "aliyun container service use this policy to allocate resources for child accounts",
		PolicyDocument: policy,
	}

	response, err := rmClient.CreatePolicy(args)
	if err != nil {
		t.Fatalf("Failed to CreatePolicy %++v", err)
	} else {
		t.Logf("Response = %++v", response)
	}
}

func TestResourceManagerClient_DeletePolicy(t *testing.T) {
	policy := "{\"Statement\": [{\"Action\": \"cs:*\",\"Effect\": \"Allow\",\"Resource\": [\"acs:cs:*:*:cluster/c64e1330717b647d891a389464bca4584\"]}],\"Version\": \"1\"}"
	args := &CreatePolicyRequest{
		PolicyName:     fmt.Sprintf("%s_%d", "AliyunACSResourcesAccess_user", time.Now().Unix()),
		Description:    "aliyun container service use this policy to allocate resources for child accounts",
		PolicyDocument: policy,
	}

	response, err := rmClient.CreatePolicy(args)
	if err != nil {
		t.Fatalf("Failed to CreatePolicy %++v", err)
	} else {
		t.Logf("CreatePolicy Successfully")
		deletePolicyArgs := &DeletePolicyRequest{
			PolicyName: response.Policy.PolicyName,
		}

		delResponse, err := rmClient.DeletePolicy(deletePolicyArgs)
		if err != nil {
			t.Fatalf("Failed to DeletePolicy %++v", err)
		} else {
			t.Logf("Response = %++v", delResponse)
		}
	}
}

func TestResourceManagerClient_ListPolicies(t *testing.T) {
	args := &ListPoliciesRequest{
		PolicyType: "Custom",
	}
	response, err := rmClient.ListPolicies(args)
	if err != nil {
		t.Fatalf("Failed to ListPolicies %++v", err)
	} else {
		t.Logf("Response =%++v", response)
	}
}

func TestResourceManagerClient_AttachPolicy(t *testing.T) {
	args := &AttachPolicyRequest{
		BasePolicyRequest: BasePolicyRequest{
			ResourceGroupId: "1312623533146984",
			PolicyType:      "Custom",
			PolicyName:      "AliyunACSResourcesAccess_user_1499916220",
			PrincipalType:   "IMSUser",
			PrincipalName:   "user-1499677896@1312623533146984.onaliyun.com",
		},
	}

	response, err := rmClient.AttachPolicy(args)
	if err != nil {
		t.Fatalf("Failed to AttachPolicy %++v", err)
	} else {
		t.Logf("Response = %++v", response)
	}
}

func TestResourceManagerClient_DetachPolicy(t *testing.T) {
	args := &DetachPolicyRequest{
		BasePolicyRequest: BasePolicyRequest{
			ResourceGroupId: "1312623533146984",
			PolicyType:      "Custom",
			PolicyName:      "AliyunACSDefaultAccess",
			PrincipalType:   "IMSUser",
			PrincipalName:   "user-1499677896@1312623533146984.onaliyun.com",
		},
	}

	response, err := rmClient.DetachPolicy(args)
	if err != nil {
		t.Fatalf("Failed to DetachPolicy %++v", err)
	} else {
		t.Logf("Response = %++v", response)
	}
}

func TestResourceManagerClient_GetPolicy(t *testing.T) {
	args := &GetPolicyRequest{
		PolicyType: "Custom",
		PolicyName: "AliyunACSDefaultAccess",
	}

	response, err := rmClient.GetPolicy(args)
	if err != nil {
		t.Fatalf("Failed to GetPolicy %++v", err)
	} else {
		t.Logf("Response = %++v", response)
	}
}

func TestResourceManagerClient_CreatePolicyVersion(t *testing.T) {
	policy := "{\"Statement\": [{\"Action\": \"cs:*\",\"Effect\": \"Allow\",\"Resource\": [\"acs:cs:*:*:cluster/c64e1330717b647d891a389464bca4584\"]}],\"Version\": \"1\"}"
	args := &CreatePolicyVersionRequest{
		PolicyName:     "AliyunACSResourcesAccess_user-1499677896_1",
		PolicyDocument: policy,
		SetAsDefault:   "true",
	}

	response, err := rmClient.CreatePolicyVersion(args)
	if err != nil {
		t.Fatalf("Failed to CreatePolicyVersion %++v", err)
	} else {
		t.Logf("Response = %++v", response)
	}

}

func TestResourceManagerClient_DeletePolicyVersion(t *testing.T) {
	args := &DeletePolicyVersionRequest{
		PolicyName: "AliyunACSResourcesAccess_user-1499677896_1",
		VersionId:  "v3",
	}

	response, err := rmClient.DeletePolicyVersion(args)
	if err != nil {
		t.Fatalf("Failed to DeletePolicyVersion %++v", err)
	} else {
		t.Logf("Response = %++v", response)
	}

}

func TestResourceManagerClient_ListPolicyVersions(t *testing.T) {
	args := &ListPolicyVersionsRequest{
		PolicyName: "AliyunACSResourcesAccess_user-1499677896_1",
		PolicyType: "Custom",
	}

	response, err := rmClient.ListPolicyVersions(args)
	if err != nil {
		t.Fatalf("Failed to ListPolicyVersions %++v", err)
	} else {
		t.Logf("Response = %++v", response)
	}
}

func TestResourceManagerClient_SetDefaultPolicyVersion(t *testing.T) {
	args := &SetDefaultPolicyVersionRequest{
		PolicyName: "AliyunACSResourcesAccess_user-1499677896_1",
		VersionId:  "v5",
	}

	response, err := rmClient.SetDefaultPolicyVersion(args)
	if err != nil {
		t.Fatalf("Failed to SetDefaultPolicyVersion %++v", err)
	} else {
		t.Logf("Response = %++v", response)
	}
}

func TestResourceManagerClient_GetPolicyVersion(t *testing.T) {
	args := &GetPolicyVersionRequest{
		PolicyName: "AliyunACSResourcesAccess_user-1499677896_1",
		PolicyType: "Custom",
		VersionId:  "v4",
	}

	response, err := rmClient.GetPolicyVersion(args)
	if err != nil {
		t.Fatalf("Failed to GetPolicyVersion %++v", err)
	} else {
		t.Logf("Response = %++v", response)
	}
}

func TestResourceManagerClient_ListPolicyAttachments(t *testing.T) {
	args := &ListPolicyAttachmentsRequest{
		BasePolicyRequest: BasePolicyRequest{
			ResourceGroupId: "1312623533146984",
			PolicyType:      "Custom",
			PolicyName:      "AliyunACSDefaultAccess",
			PrincipalType:   "IMSUser",
			PrincipalName:   "user-1499677896@1312623533146984.onaliyun.com",
		},
	}

	response, err := rmClient.ListPolicyAttachments(args)
	if err != nil {
		t.Fatalf("Failed to ListPolicyAttachments %++v", err)
	} else {
		for index, policy := range response.PolicyAttachments.PolicyAttachment {
			t.Logf("PolicyAttachment[%d] = %++v", index, policy)
		}
	}
}
