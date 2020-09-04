package standard

import (
	"github.com/denverdino/aliyungo/common"
)

type DeletionProtection string

const (
	DeletionProtectionEnabled  = DeletionProtection("Enabled")
	DeletionProtectionDisabled = DeletionProtection("Disabled")
)

//https://help.aliyun.com/document_detail/28910.html?spm=5176.doc50083.6.580.b5wkQr
type CreateStackRequest struct {
	RegionId           common.Region
	StackName          string
	DisableRollback    bool
	TemplateBody       string
	TemplateURL        string
	Parameters         []Parameter
	StackPolicyURL     string
	TimeoutInMinutes   int
	StackPolicyBody    string
	ClientToken        string
	NotificationURLs   []string
	DeletionProtection DeletionProtection
	RamRoleName        string
}

type CreateStackResponse struct {
	StackId string
	common.Response
}

type ListStackEventsRequest struct {
	common.Pagination
	RegionId          common.Region
	StackId           string
	Status            []string
	ResourceType      []string
	LogicalResourceId []string
}

type ListStackEventsResponse struct {
	common.Response
	common.PaginationResult
	RegionId common.Region
	Events   []Event
}

type Event struct {
	StackId            string
	Status             string
	StackName          string
	StatusReason       string
	EventId            string
	LogicalResourceId  string
	ResourceType       string
	PhysicalResourceId string
	CreateTime         string
}

func (client *Client) ListStackEvents(args *ListStackEventsRequest) (*ListStackEventsResponse, error) {
	response := &ListStackEventsResponse{}
	err := client.Invoke("ListStackEvents", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client *Client) CreateStack(args *CreateStackRequest) (*CreateStackResponse, error) {
	stack := &CreateStackResponse{}
	err := client.Invoke("CreateStack", args, stack)
	if err != nil {
		return nil, err
	}

	return stack, nil
}

//https://help.aliyun.com/document_detail/28911.html?spm=5176.doc28910.6.581.etoi2Z
type DeleteStackRequest struct {
	RegionId           common.Region
	StackId            string
	RetainAllResources bool
	RetainResources    []string
	RamRoleName        string
}

type DeleteStackResponse struct {
	common.Response
}

func (client *Client) DeleteStack(req *DeleteStackRequest) (*DeleteStackResponse, error) {
	response := &DeleteStackResponse{}
	err := client.Invoke("DeleteStack", req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type GetStackRequest struct {
	RegionId    common.Region
	StackId     string
	ClientToken string
}

type GetStackResponse struct {
	CreateTime          string
	Description         string
	DisableRollback     bool
	NotificationURLs    []string
	Outputs             []Output
	ParentStackId       string
	RegionId            common.Region
	Status              string
	StackId             string
	StackName           string
	Parameters          []Parameter
	UpdateTime          string
	StatusReason        string
	TemplateDescription string
	TimeoutInMinutes    int

	RequestId          string
	DeletionProtection DeletionProtection
	DriftDetectionTime string
	RamRoleName        string
	RootStackId        string
	StackDriftStatus   string
	StackType          string
}

type Parameter struct {
	ParameterKey   string
	ParameterValue string
}

type Output struct {
	Description string
	OutputKey   string
	OutputValue interface{}
}

func (client *Client) GetStack(req *GetStackRequest) (*GetStackResponse, error) {
	response := &GetStackResponse{}
	err := client.Invoke("GetStack", req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type ListStacksRequest struct {
	RegionId        common.Region
	StackId         string
	Status          []string
	ParentStackId   string
	StackName       []string
	ShowNestedStack bool
	Tag             []Tag
	common.Pagination
}

type ListStacksResponse struct {
	common.PaginationResult
	common.Response
	Stacks []Stack
}

type Stack struct {
	CreateTime string

	DisableRollback    bool
	DriftDetectionTime string
	ParentStackId      string
	RegionId           common.Region
	StackDriftStatus   string
	StackId            string
	StackName          string
	Status             string
	StatusReason       string

	TimeoutInMinutes int
	UpdateTime       string
}

type Tag struct {
	Key   string
	Value string
}

func (client *Client) ListStacks(req *ListStacksRequest) (*ListStacksResponse, error) {
	response := &ListStacksResponse{}
	err := client.Invoke("ListStacks", req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type UpdateStackRequest struct {
	Parameters                  []Parameter
	RegionId                    string
	StackId                     string
	ClientToken                 string
	StackPolicyDuringUpdateBody string
	TimeoutInMinutes            int
	TemplateBody                string
	StackPolicyURL              string
	StackPolicyDuringUpdateURL  string
	StackPolicyBody             string
	UsePreviousParameters       bool
	DisableRollback             bool
	TemplateURL                 string
	RamRoleName                 string
	ReplacementOption           string
}

type UpdateStackResponse struct {
	StackId string
	common.Response
}

func (client *Client) UpdateStack(req *UpdateStackRequest) (*UpdateStackResponse, error) {
	response := &UpdateStackResponse{}
	err := client.Invoke("UpdateStack", req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type GetStackResourceRequest struct {
	StackId                string
	LogicalResourceId      string
	ClientToken            string
	ShowResourceAttributes bool
	RegionId               common.Region
}

type GetStackResourceResponse struct {
	Status            string
	Description       string
	LogicalResourceId string
	StackId           string

	StackName           string
	StatusReason        string
	PhysicalResourceId  string
	ResourceType        string
	CreateTime          string
	Metadata            map[string]string
	UpdateTime          string
	ResourceAttributes  []ResourceAttribute
	RequestId           string
	DriftDetectionTime  string
	ResourceDriftStatus string
}
type ResourceAttribute struct {
	ResourceAttributeValue string
	ResourceAttributeKey   string
}

func (client *Client) GetStackResource(req *GetStackResourceRequest) (*GetStackResourceResponse, error) {
	response := &GetStackResourceResponse{}
	err := client.Invoke("GetStackResource", req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type ListStackResourcesRequest struct {
	RegionId common.Region
	StackId  string
}

type ListStackResourcesResponse struct {
	common.Response
	Resources []Resource
}

type Resource struct {
	CreateTime         string
	DriftDetectionTime string
	LogicalResourceId  string
	PhysicalResourceId string

	ResourceDriftStatus string
	ResourceType        string
	StackId             string
	StackName           string
	Status              string
	StatusReason        string
	UpdateTime          string
}

func (client *Client) ListStackResources(req *ListStackResourcesRequest) (*ListStackResourcesResponse, error) {
	response := &ListStackResourcesResponse{}
	err := client.Invoke("ListStackResources", req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
