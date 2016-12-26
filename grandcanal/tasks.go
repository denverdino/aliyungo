package grandcanal

import (
	"github.com/denverdino/aliyungo/common"
)

// PollForWorkflowTaskArgs represents arguments to poll workflow decision task
type PollForWorkflowTaskArgs struct {
	Domain   string
	Queue    string
	Identity string
}

type PollForWorkflowTaskResponse struct {
	common.Response
	TaskType      string
	WorkflowId    string
	NextPageToken string
	TaskToken     string
	Activities    struct {
		Activity []Activity
	}
	WorkflowType  SimpleWorkflowType
	TaskAttribute PollTaskAttribute
}

type Activity struct {
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

type SimpleWorkflowType struct {
	Name    string
	Version string
}

type SimpleActivityType struct {
	Name    string
	Version string
}

type PollTaskAttribute struct {
	RunId               string
	TaskType            string
	Name                string
	Result              string
	Error               string
	ActivityState       string
	Input               string
	Method              string
	ActivityTypeName    string
	ActivityTypeVersion string
	TaskQueue           string
	Priority            int32
	Delay               int32
	Timeout             int32
	HeartbeatTimeout    int32
	WorkflowState       string
	SignalName          string
	WorkflowTimeout     int64
	BizId               string
	ParentWorkflowId    string
	Tags                Tags
}

type Tags struct {
	Tag []string
}

// PollForWorkflowTask poll workflow task
func (client *Client) PollForWorkflowTask(args *PollForWorkflowTaskArgs) (resp *PollForWorkflowTaskResponse, err error) {
	response := &PollForWorkflowTaskResponse{}
	err = client.Invoke("PollForWorkflowTask", args, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// PollForActivityTaskArgs represents arguments to poll activity task
type PollForActivityTaskArgs struct {
	Domain   string
	Queue    string
	Identity string
}

type PollForActivityTaskResponse struct {
	common.Response
	WorkflowId   string
	TaskToken    string
	Input        string
	RunId        string
	Name         string
	Method       string
	Priority     int32
	ActivityType SimpleActivityType
}

// PollForActivityTask poll activity task
func (client *Client) PollForActivityTask(args *PollForActivityTaskArgs) (resp *PollForActivityTaskResponse, err error) {
	response := &PollForActivityTaskResponse{}
	err = client.Invoke("PollForActivityTask", args, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// RespondWorkflowTaskArgs represents arguments to respond a workflow task
type RespondWorkflowTaskArgs struct {
	TaskToken      string
	Identity       string
	TaskAttributes RespondTaskAttributes
}

type RespondTaskAttributes struct {
	TaskAttribute []RespondTaskAttribute
}

type RespondTaskAttribute struct {
	RunId               string
	TaskType            string
	Name                string
	Result              string
	Error               string
	ActivityState       string
	Input               string
	Method              string
	ActivityTypeName    string
	ActivityTypeVersion string
	TaskQueue           string
	Priority            int32
	Delay               int32
	Timeout             int32
	HeartbeatTimeout    int32
	WorkflowState       string
	SignalName          string
	WorkflowTimeout     int64
	Bizid               string
	ParentWorkflowId    string
}

type RespondWorkflowTaskResponse struct {
	common.Response
}

// RespondWorkflowTask respond a workflow task
func (client *Client) RespondWorkflowTask(args *RespondWorkflowTaskArgs) (requestId string, err error) {
	response := &RespondWorkflowTaskResponse{}
	err = client.Invoke("RespondWorkflowTask", args, &response)
	if err != nil {
		return "", err
	}
	return response.Response.RequestId, nil
}

// RespondActivityTaskArgs represents arguments to respond an activity task
type RespondActivityTaskArgs struct {
	TaskToken     string
	Result        string
	Error         string
	ActivityState string
	Identity      string
}

type RespondActivityTaskResponse struct {
	common.Response
}

// RespondActivityTask respond an activity task
func (client *Client) RespondActivityTask(args *RespondActivityTaskArgs) (requestId string, err error) {
	response := &RespondActivityTaskResponse{}
	err = client.Invoke("RespondActivityTask", args, &response)
	if err != nil {
		return "", err
	}
	return response.Response.RequestId, nil
}

// ReceiveHeartbeatArgs represents arguments to send a heartbeat message
type ReceiveHeartbeatArgs struct {
	TaskToken string
}

type ReceiveHeartbeatResponse struct {
	common.Response
	TaskCancelled string
}

// ReceiveHeartbeat send a heartbeat message to indicate the task is still alive
func (client *Client) ReceiveHeartbeat(args *ReceiveHeartbeatArgs) (taskCancelled string, err error) {
	response := &ReceiveHeartbeatResponse{}
	err = client.Invoke("ReceiveHeartbeat", args, &response)
	if err != nil {
		return "", err
	}
	return response.TaskCancelled, nil
}
