package ims

import (
	"testing"
)

func TestResourceManagerClient_AttachPolicy(t *testing.T) {
	args := &AttachPolicyRequest{
		BasePolicyRequest: BasePolicyRequest{
			ResourceGroupId: "1312623533146984",
			PolicyType:      "Custom",
			PolicyName:      "AliyunACSDefaultAccess",
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
