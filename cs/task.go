package cs

import (
	"fmt"
	"net/http"
	"time"
)

type TaskType string

type TaskError struct {
	Code      string      `json:"code,omitempty"`
	Message   interface{} `json:"message,omitempty"`
	RequestId string      `json:"requestId,omitempty"`
	Status    int         `json:"status,omitempty"`
	Method    string      `json:"-"`
}

type TaskResults []TaskResult
type TaskResult struct {
	Status  string      `json:"status"`
	Phase   string      `json:"phase"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TaskTargetType string
type TaskTarget struct {
	Type TaskTargetType `json:"type"`
	Id   string         `json:"id"`
	Name string         `json:"name,omitempty"`
}

type DescribeTaskInfoResponse struct {
	TaskID     string       `json:"task_id"`
	ClusterID  string       `json:"cluster_id"`
	TaskType   TaskType     `json:"task_type"`
	State      TaskState    `json:"state"`
	Created    time.Time    `json:"created_at"`
	Updated    time.Time    `json:"updated_at"`
	TaskError  *TaskError   `json:"error,omitempty"`
	TaskResult *TaskResults `json:"task_result,omitempty"`
	Target     *TaskTarget  `json:"target,omitempty"`
	Params     interface{}  `json:"parameters,omitempty"`
}

func (client *Client) DescribeTaskInfo(taskId string) (*DescribeTaskInfoResponse, error) {
	taskInfo := &DescribeTaskInfoResponse{}
	err := client.Invoke("", http.MethodGet, fmt.Sprintf("/tasks/%s", taskId), nil, nil, taskInfo)
	if err != nil {
		return nil, err
	}
	return taskInfo, nil
}
