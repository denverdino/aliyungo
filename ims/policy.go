package ims

import "github.com/denverdino/aliyungo/common"

type BasePolicyRequest struct {
	AccountId       string
	ResourceGroupId string //
	PolicyType      string //
	PolicyName      string //
	PrincipalType   string //AliyunAccount/IMSUser/IMSGroup/ServiceRole；
	PrincipalName   string
}

type AttachPolicyRequest struct {
	BasePolicyRequest
}

type AttachPolicyResponse struct {
	common.Response
}

func (client *ResourceManagerClient) AttachPolicy(args *AttachPolicyRequest) (*AttachPolicyResponse, error) {
	response := &AttachPolicyResponse{}

	err := client.Invoke("AttachPolicy", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type DetachPolicyRequest struct {
	BasePolicyRequest
}

type DetachPolicyResponse struct {
	common.Response
}

func (client *ResourceManagerClient) DetachPolicy(args *DetachPolicyRequest) (*DetachPolicyResponse, error) {
	response := &DetachPolicyResponse{}

	err := client.Invoke("DetachPolicy", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type PolicyAttachment struct {
	ResourceGroupId string
	PolicyType      string
	PolicyName      string
	PrincipalType   string
	PrincipalName   string
	AttachDate      string
}

type ListPolicyAttachmentsRequest struct {
	BasePolicyRequest
	PageSize   int
	PageNumber int
}

type ListPolicyAttachmentsResponse struct {
	common.Response
	PageSize          int
	PageNumber        int
	TotalCount        int
	PolicyAttachments struct {
		PolicyAttachment []PolicyAttachment
	}
}

func (client *ResourceManagerClient) ListPolicyAttachments(args *ListPolicyAttachmentsRequest) (*ListPolicyAttachmentsResponse, error) {
	response := &ListPolicyAttachmentsResponse{}

	err := client.Invoke("ListPolicyAttachments", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
