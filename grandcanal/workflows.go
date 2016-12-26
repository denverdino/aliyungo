package grandcanal

import (
	"github.com/denverdino/aliyungo/common"
)

// StartWorkflowArgs represents arguments to start a new workflow instance
type StartWorkflowArgs struct {
	DomainName          string
	WorkflowTypeName    string
	WorkflowTypeVersion string
	ParentId            string
	BizId               string
	Tag1                string
	Tag2                string
	Tag3                string
	Input               string
	Timeout             int32
	Priority            int32
	DeciderTimeout      int32
	TaskQueue           string
}

type StartWorkflowResponse struct {
	common.Response
	WorkflowId string
}

// StartWorkflow start a new workflow instance
func (client *Client) StartWorkflow(args *StartWorkflowArgs) (workflowId string, err error) {
	response := &StartWorkflowResponse{}
	err = client.Invoke("StartWorkflow", args, &response)
	if err != nil {
		return "", err
	}
	return response.WorkflowId, nil
}

// SignalWorkflowArgs represents arguments to signal a workflow
type SignalWorkflowArgs struct {
	WorkflowId string
	SignalName string
	Input      string
}

type SignalWorkflowResponse struct {
	RequestId string
}

// SignalWorkflow send a signal to a workflow instance
func (client *Client) SignalWorkflow(args *SignalWorkflowArgs) (requestId string, err error) {
	response := &SignalWorkflowResponse{}
	err = client.Invoke("SignalWorkflow", args, &response)
	if err != nil {
		return "", err
	}
	return response.RequestId, nil
}

// RollbackWorkflowArg represents arguments to rollback a workflow
type RollbackWorkflowArgs struct {
	WorkflowId string
}

type RollbackWorkflowResponse struct {
	RequestId string
}

// RollbackWorkflow rollback an existing workflow instance
func (client *Client) RollbackWorkflow(args *RollbackWorkflowArgs) (requestId string, err error) {
	response := &RollbackWorkflowResponse{}
	err = client.Invoke("RollbackWorkflow", args, &response)
	if err != nil {
		return "", err
	}
	return response.RequestId, nil
}

// TerminateWorkflowArg represents arguments to terminate a workflow
type TerminateWorkflowArgs struct {
	WorkflowId string
}

type TerminateWorkflowResponse struct {
	RequestId string
}

// TerminateWorkflow terminate an existing workflow instance
func (client *Client) TerminateWorkflow(args *TerminateWorkflowArgs) (requestId string, err error) {
	response := &TerminateWorkflowResponse{}
	err = client.Invoke("TerminateWorkflow", args, &response)
	if err != nil {
		return "", err
	}
	return response.RequestId, nil
}

// SuspendWorkflowArg represents arguments to suspend a workflow
type SuspendWorkflowArgs struct {
	WorkflowId string
}

type SuspendWorkflowResponse struct {
	RequestId string
}

// SuspendWorkflow suspend a running workflow instance
func (client *Client) SuspendWorkflow(args *SuspendWorkflowArgs) (requestId string, err error) {
	response := &SuspendWorkflowResponse{}
	err = client.Invoke("SuspendWorkflow", args, &response)
	if err != nil {
		return "", err
	}
	return response.RequestId, nil
}

// ResumeWorkflowArg represents arguments to resume a workflow
type ResumeWorkflowArgs struct {
	WorkflowId string
}

type ResumeWorkflowResponse struct {
	RequestId string
}

// ResumeWorkflow resume a workflow instance
func (client *Client) ResumeWorkflow(args *ResumeWorkflowArgs) (requestId string, err error) {
	response := &ResumeWorkflowResponse{}
	err = client.Invoke("ResumeWorkflow", args, &response)
	if err != nil {
		return "", err
	}
	return response.RequestId, nil
}

// QueryWorkflowInstanceArgs represents arguments to query workflow instances
type QueryWorkflowInstancesArgs struct {
	WorkflowId          string
	BizId               string
	Tag                 string
	WorkflowTypeName    string
	WorkflowTypeVersion string
	StartTime           int64
	EndTime             int64
	common.Pagination
}

type QueryWorkflowInstancesResponse struct {
	common.Response
	common.PaginationResult
	Workflows struct {
		Workflow []WorkflowInstance
	}
}

func (client *Client) QueryWorkflowInstances(args *QueryWorkflowInstancesArgs) (resp *QueryWorkflowInstancesResponse, err error) {
	response := &QueryWorkflowInstancesResponse{}
	err = client.Invoke("QueryWorkflowInstances", args, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// QueryActivityInstanceArgs represents arguments to query activity instances
type QueryActivityInstancesArgs struct {
	WorkflowId string
	ActivityId string
	common.Pagination
}

type QueryActivityInstancesResponse struct {
	common.Response
	common.PaginationResult
	ActivityInstances struct {
		ActivityInstance []ActivityInstance
	}
}

func (client *Client) QueryActivityInstances(args *QueryActivityInstancesArgs) (resp *QueryActivityInstancesResponse, err error) {
	response := &QueryActivityInstancesResponse{}
	err = client.Invoke("QueryActivityInstances", args, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// DescribeWorkflowInstanceArgs represnets arguments to describe a workflow instance
type DescribeWorkflowInstanceArgs struct {
	WorkflowId string
}

type DescribeWorkflowInstanceResponse struct {
	common.Response
	WorkflowInstance WorkflowInstance
}

type WorkflowInstance struct {
	WorkflowId          string
	ParentId            string
	DomainName          string
	WorkflowTypeName    string
	WorkflowTypeVersion string
	Status              string
	BizId               string
	Input               string
	Priority            int32
	Timeout             int32
	DeciderTimeout      int32
	TaskQueue           string
	StartTime           string
	EndTime             string
	Duration            int64
	Tags                Tags
}

func (client *Client) DescribeWorkflowInstance(args *DescribeWorkflowInstanceArgs) (instance *WorkflowInstance, err error) {
	response := &DescribeWorkflowInstanceResponse{}
	err = client.Invoke("DescribeWorkflowInstance", args, &response)
	if err != nil {
		return nil, err
	}
	return &response.WorkflowInstance, nil
}

// DescribeActivityInstanceRequest represents arguments to describe an activity instance
type DescribeActivityInstanceArgs struct {
	ActivityId string
}

type DescribeActivityInstanceResponse struct {
	common.Response
	ActivityInstance ActivityInstance
}

type ActivityInstance struct {
	DomainName          string
	ActivityId          string
	WorkflowId          string
	ActivityTypeName    string
	ActivityTypeVersion string
	ActivityName        string
	Method              string
	Status              string
	Priority            int32
	Timeout             int32
	HeartbeatTimeout    int32
	TaskQueue           string
	Input               string
	Output              string
	Error               string
	StartTime           string
	EndTime             string
	Duration            int64
}

func (client *Client) DescribeActivityInstance(args *DescribeActivityInstanceArgs) (instance *ActivityInstance, err error) {
	response := &DescribeActivityInstanceResponse{}
	err = client.Invoke("DescribeActivityInstance", args, &response)
	if err != nil {
		return nil, err
	}
	return &response.ActivityInstance, nil
}

// CountWorkflowInstanceArgs represnets arguments to count workflow instance
type CountWorkflowInstanceArgs struct {
	WorkflowId          string
	BizId               string
	Tag                 string
	WorkflowTypeName    string
	WorkflowTypeVersion string
	StartTime           int64
	EndTime             int64
}

type CountWorkflowInstanceResponse struct {
	common.Response
	Count int32
}

func (client *Client) CountWorkflowInstance(args *CountWorkflowInstanceArgs) (instance int32, err error) {
	response := &CountWorkflowInstanceResponse{}
	err = client.Invoke("CountWorkflowInstance", args, &response)
	if err != nil {
		return 0, err
	}
	return response.Count, nil
}

// CountActivityInstanceArgs represnets arguments to count activity instance
type CountActivityInstanceArgs struct {
	WorkflowId string
	Name       string
}

type CountActivityInstanceResponse struct {
	common.Response
	Count int32
}

func (client *Client) CountActivityInstance(args *CountActivityInstanceArgs) (instance int32, err error) {
	response := &CountActivityInstanceResponse{}
	err = client.Invoke("CountActivityInstance", args, &response)
	if err != nil {
		return 0, err
	}
	return response.Count, nil
}

// ListAvailableActivityHistoryArgs represents arguments to list available activity history
type ListAvailableActivityHistoryArgs struct {
	WorkflowId string
	ActivityId string
	common.Pagination
}

type ListAvailableActivityHistoryResponse struct {
	common.Response
	common.PaginationResult
	Activities struct {
		Activity []ActivityHistory
	}
}

type ActivityHistory struct {
	RunId        string
	State        string
	Input        string
	Output       string
	Error        string
	StartTime    string
	EndTime      string
	Method       string
	Name         string
	ActivityType SimpleActivityType
}

func (client *Client) ListAvailableActivityHistory(args *ListAvailableActivityHistoryArgs) (resp *ListAvailableActivityHistoryResponse, err error) {
	response := &ListAvailableActivityHistoryResponse{}
	err = client.Invoke("ListAvailableActivityHistory", args, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
