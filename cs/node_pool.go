package cs

import (
	"fmt"
	"github.com/denverdino/aliyungo/common"
	"net/http"
	"time"
)

type NodePoolInfo struct {
	NodePoolId      string        `json:"nodepool_id"`
	RegionId        common.Region `json:"region_id"`
	Name            string        `json:"name"`
	Created         time.Time     `json:"created"`
	Updated         time.Time     `json:"updated"`
	IsDefault       bool          `json:"is_default"`
	NodePoolType    string        `json:"type"`
	ResourceGroupId string        `json:"resource_group_id"`
}

type ScalingGroup struct {
	VpcId              string             `json:"vpc_id"`
	VswitchIds         []string           `json:"vswitch_ids"`
	InstanceTypes      []string           `json:"instance_types"`
	LoginPassword      string             `json:"login_password"`
	KeyPair            string             `json:"key_pair"`
	SecurityGroupId    string             `json:"security_group_id"`
	SystemDiskCategory string             `json:"system_disk_category"`
	SystemDiskSize     int64              `json:"system_disk_size"`
	DataDisks          []NodePoolDataDisk `json:"data_disks"` //支持多个数据盘
	Tags               []Tag              `json:"tags"`
	ImageId            string             `json:"image_id"`
	Platform           string             `json:"platform"`
}

type AutoScaling struct {
	EnableAutoScaling bool   `json:"enable_auto_scaling"`
	MaxInstance       int64  `json:"max_instance"`
	MinInstance       int64  `json:"min_instance"`
	Type              string `json:"type"`
}

type KubernetesConfig struct {
	NodeNameMode string  `json:"node_name_mode"`
	Taints       []Taint `json:"taints"`
	Labels       []Label `json:"labels"`
	CpuPolicy    string  `json:"cpu_policy"`
	UserData     string  `json:"user_data"`
}

type CreateNodePoolRequest struct {
	RegionId         common.Region `json:"region_id"`
	Count            int64         `json:"count"`
	NodePoolInfo     `json:"nodepool_info"`
	ScalingGroup     `json:"scaling_group"`
	KubernetesConfig `json:"kubernetes_config"`
	AutoScaling      `json:"auto_scaling"`
}

type BasicNodePool struct {
	NodePoolInfo   `json:"nodepool_info"`
	NodePoolStatus `json:"status"`
}

type NodePoolStatus struct {
	TotalNodes   int `json:"total_nodes"`
	OfflineNodes int `json:"offline_nodes"`
	ServingNodes int `json:"serving_nodes"`
	//DesiredNodes int  `json:"desired_nodes"`
	RemovingNodes int    `json:"removing_nodes"`
	FailedNodes   int    `json:"failed_nodes"`
	InitialNodes  int    `json:"initial_nodes"`
	HealthyNodes  int    `json:"healthy_nodes"`
	State         string `json:"state"`
}

type NodePoolDetail struct {
	BasicNodePool
	KubernetesConfig `json:"kubernetes_config"`
	ScalingGroup     `json:"scaling_group"`
	AutoScaling      `json:"auto_scaling"`
}

type CreateNodePoolResponse struct {
	Response
	NodePoolID string `json:"nodepool_id"`
	Message    string `json:"Message"`
}

type UpdateNodePoolRequest struct {
	RegionId         common.Region `json:"region_id"`
	Count            int64         `json:"count"`
	NodePoolInfo     `json:"nodepool_info"`
	ScalingGroup     `json:"scaling_group"`
	KubernetesConfig `json:"kubernetes_config"`
	AutoScaling      `json:"auto_scaling"`
}

func (client *Client) CreateNodePool(request *CreateNodePoolRequest, clusterId string) (*CreateNodePoolResponse, error) {
	response := &CreateNodePoolResponse{}
	err := client.Invoke(request.RegionId, http.MethodPost, fmt.Sprintf("/clusters/%s/nodepools", clusterId), nil, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client *Client) DescribeNodePoolDetail(clusterId, nodePoolId string) (*NodePoolDetail, error) {
	nodePool := &NodePoolDetail{}
	err := client.Invoke("", http.MethodGet, fmt.Sprintf("/clusters/%s/nodepools/%s", clusterId, nodePoolId), nil, nil, nodePool)
	if err != nil {
		return nil, err
	}
	return nodePool, nil
}

func (client *Client) UpdateNodePool(clusterId string, nodePoolId string, request *UpdateNodePoolRequest) (*Response, error) {
	response := &Response{}
	err := client.Invoke(request.RegionId, http.MethodPut, fmt.Sprintf("/clusters/%s/nodepools/%s", clusterId, nodePoolId), nil, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client *Client) DeleteNodePool(clusterId, nodePoolId string) error {
	return client.Invoke("", http.MethodDelete, fmt.Sprintf("/clusters/%s/nodepools/%s", clusterId, nodePoolId), nil, nil, nil)
}
