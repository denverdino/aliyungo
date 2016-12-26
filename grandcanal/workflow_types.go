package grandcanal

import (
	"github.com/denverdino/aliyungo/common"
)

// CreateWorkflowTypeArgs represents arguments to create a workflow type
type CreateWorkflowTypeArgs struct {
	Domain                 string
	WorkflowTypeName       string
	WorkflowTypeVersion    string
	Description            string
	DefaultWorkflowTimeout int64
	DefaultPriority        int64
	DefaultDeciderTimeout  int64
}

type CreateWorkflowTypeResponse struct {
	common.Response
}

// CreateWorkflowType create a new workflow type
func (client *Client) CreateWorkflowType(args *CreateWorkflowTypeArgs) (requestId string, err error) {
	response := &CreateWorkflowTypeResponse{}
	err = client.Invoke("CreateWorkflowType", args, &response)
	if err != nil {
		return "", err
	}
	return response.RequestId, nil
}

// DescribeWorkflowTypeArgs represents arguments to describe workflow type
type DescribeWorkflowTypeArgs struct {
	Domain              string
	WorkflowTypeName    string
	WorkflowTypeVersion string
}

type DescribeWorkflowTypeResponse struct {
	common.Response
	WorkflowType WorkflowType
}

type WorkflowType struct {
	WorkflowTypeName       string
	WorkflowTypeVersion    string
	CreationDate           string
	DisableDate            string
	Description            string
	Status                 string
	DefaultWorkflowTimeout int64
	DefaultPriority        int64
	DefaultDeciderTimeout  int64
}

// DescribeWorkflowType describe the detailed information of a workflow type
func (client *Client) DescribeWorkflowType(args *DescribeWorkflowTypeArgs) (workflowType *WorkflowType, err error) {
	response := &DescribeWorkflowTypeResponse{}
	err = client.Invoke("DescribeWorkflowType", args, &response)
	if err != nil {
		return nil, err
	}
	return &response.WorkflowType, nil
}

// DeleteWorkflowTypeArgs represents arguments to delete a workflow type
type DeleteWorkflowTypeArgs struct {
	Domain              string
	WorkflowTypeName    string
	WorkflowTypeVersion string
}

type DeleteWorkflowTypeResponse struct {
	common.Response
}

// DeleteWorkflowType delete a workflow type object
func (client *Client) DeleteWorkflowType(args *DeleteWorkflowTypeArgs) (requestId string, err error) {
	response := &DeleteWorkflowTypeResponse{}
	err = client.Invoke("DeleteWorkflowType", args, &response)
	if err != nil {
		return "", err
	}
	return response.RequestId, nil
}
