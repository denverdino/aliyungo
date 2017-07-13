package ims

import "github.com/denverdino/aliyungo/common"

type Policy struct {
	PolicyName     string
	PolicyType     string
	Description    string
	DefaultVersion string
	CreateDate     string
}

type CreatePolicyRequest struct {
	AccountId      string
	PolicyName     string
	Description    string
	PolicyDocument string
}

type CreatePolicyResponse struct {
	common.Response
	Policy Policy
}

func (client *ResourceManagerClient) CreatePolicy(args *CreatePolicyRequest) (*CreatePolicyResponse, error) {
	response := &CreatePolicyResponse{}
	err := client.Invoke("CreatePolicy", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type DeletePolicyRequest struct {
	AccountId  string
	PolicyName string
}

type DeletePolicyResponse struct {
	common.Response
}

func (client *ResourceManagerClient) DeletePolicy(args *DeletePolicyRequest) (*DeletePolicyResponse, error) {
	response := &DeletePolicyResponse{}
	err := client.Invoke("DeletePolicy", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type PolicyItem struct {
	PolicyName      string
	PolicyType      string
	Description     string
	DefaultVersion  string
	CreateDate      string
	UpdateDate      string
	AttachmentCount int
}

type ListPoliciesRequest struct {
	AccountId        string
	PolicyType       string
	ResourceGroupId  string
	IncludeInherited bool
	PageSize         int
	PageNumber       int
}

type ListPoliciesResponse struct {
	common.Response
	PageNumber int
	PageSize   int
	TotalCount int
	Policies   struct {
		Policy []PolicyItem
	}
}

func (client *ResourceManagerClient) ListPolicies(args *ListPoliciesRequest) (*ListPoliciesResponse, error) {
	response := &ListPoliciesResponse{}
	err := client.Invoke("ListPolicies", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type GetPolicyRequest struct {
	AccountId  string
	PolicyType string
	PolicyName string
}

type GetPolicyResponse struct {
	common.Response
	PolicyItem
}

func (client *ResourceManagerClient) GetPolicy(args *GetPolicyRequest) (*GetPolicyResponse, error) {
	response := &GetPolicyResponse{}
	err := client.Invoke("GetPolicy", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type PolicyVersion struct {
	VersionId        string
	IsDefaultVersion bool
	CreateDate       string
}

type CreatePolicyVersionRequest struct {
	AccountId      string
	PolicyName     string
	PolicyDocument string
	SetAsDefault   string
}

type CreatePolicyVersionResponse struct {
	common.Response
	PolicyVersion PolicyVersion
}

func (client *ResourceManagerClient) CreatePolicyVersion(args *CreatePolicyVersionRequest) (*CreatePolicyVersionResponse, error) {
	response := &CreatePolicyVersionResponse{}
	err := client.Invoke("CreatePolicyVersion", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type DeletePolicyVersionRequest struct {
	AccountId  string
	PolicyName string
	VersionId  string
}

type DeletePolicyVersionResponse struct {
	common.Response
}

func (client *ResourceManagerClient) DeletePolicyVersion(args *DeletePolicyVersionRequest) (*DeletePolicyVersionResponse, error) {
	response := &DeletePolicyVersionResponse{}
	err := client.Invoke("DeletePolicyVersion", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type ListPolicyVersionsRequest struct {
	AccountId  string
	PolicyType string
	PolicyName string
}

type ListPolicyVersionsResponse struct {
	PolicyVersions struct {
		PolicyVersion []PolicyVersion
	}
}

func (client *ResourceManagerClient) ListPolicyVersions(args *ListPolicyVersionsRequest) (*ListPolicyVersionsResponse, error) {
	response := &ListPolicyVersionsResponse{}
	err := client.Invoke("ListPolicyVersions", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type PolicyVersionItem struct {
	PolicyVersion
	PolicyDocument string
}

type GetPolicyVersionRequest struct {
	AccountId  string
	PolicyType string
	PolicyName string
	VersionId  string
}

type GetPolicyVersionResponse struct {
	common.Response
	PolicyVersion PolicyVersionItem
}

func (client *ResourceManagerClient) GetPolicyVersion(args *GetPolicyVersionRequest) (*GetPolicyVersionResponse, error) {
	response := &GetPolicyVersionResponse{}
	err := client.Invoke("GetPolicyVersion", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type SetDefaultPolicyVersionRequest struct {
	AccountId  string
	PolicyName string
	VersionId  string
}

type SetDefaultPolicyVersionResponse struct {
	common.Response
}

func (client *ResourceManagerClient) SetDefaultPolicyVersion(args *SetDefaultPolicyVersionRequest) (*SetDefaultPolicyVersionResponse, error) {

	response := &SetDefaultPolicyVersionResponse{}
	err := client.Invoke("SetDefaultPolicyVersion", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

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
