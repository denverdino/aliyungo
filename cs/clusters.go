package cs

import (
	"net/http"
	"net/url"

	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/ecs"
	"github.com/denverdino/aliyungo/util"
)

type ClusterState string

const (
	Initial      = ClusterState("Initial")
	Running      = ClusterState("Running")
	Updating     = ClusterState("Updating")
	Scaling      = ClusterState("Scaling")
	Failed       = ClusterState("Failed")
	Deleting     = ClusterState("Deleting")
	DeleteFailed = ClusterState("DeleteFailed")
	Deleted      = ClusterState("Deleted")
)

type NodeStatus struct {
	Health   int64 `json:"health"`
	Unhealth int64 `json:"unhealth"`
}

type NetworkModeType string

const (
	ClassicNetwork = NetworkModeType("classic")
	VPCNetwork     = NetworkModeType("vpc")
)

// https://help.aliyun.com/document_detail/26053.html
type ClusterType struct {
	AgentVersion           string           `json:"agent_version"`
	ClusterID              string           `json:"cluster_id"`
	Name                   string           `json:"name"`
	Created                util.ISO6801Time `json:"created"`
	ExternalLoadbalancerID string           `json:"external_loadbalancer_id"`
	MasterURL              string           `json:"master_url"`
	NetworkMode            NetworkModeType  `json:"network_mode"`
	RegionID               common.Region    `json:"region_id"`
	SecurityGroupID        string           `json:"security_group_id"`
	Size                   int64            `json:"size"`
	State                  ClusterState     `json:"state"`
	Updated                util.ISO6801Time `json:"updated"`
	VPCID                  string           `json:"vpc_id"`
	VSwitchID              string           `json:"vswitch_id"`
	NodeStatus             NodeStatus       `json:"node_status"`
	DockerVersion          string           `json:"docker_version"`
}

func (client *Client) DescribeClusters(nameFilter string) (clusters []ClusterType, err error) {
	query := make(url.Values)

	if nameFilter != "" {
		query.Add("name", nameFilter)
	}

	err = client.Invoke("", http.MethodGet, "/clusters", query, nil, &clusters)
	return
}

func (client *Client) DescribeCluster(id string) (cluster ClusterType, err error) {
	err = client.Invoke("", http.MethodGet, "/clusters/"+id, nil, nil, &cluster)
	return
}

type ClusterCreationArgs struct {
	Name             string           `json:"name"`
	Size             int64            `json:"size"`
	NetworkMode      NetworkModeType  `json:"network_mode"`
	SubnetCIDR       string           `json:"subnet_cidr,omitempty"`
	InstanceType     string           `json:"instance_type"`
	VPCID            string           `json:"vpc_id,omitempty"`
	VSwitchID        string           `json:"vswitch_id,omitempty"`
	Password         string           `json:"password"`
	DataDiskSize     int64            `json:"data_disk_size"`
	DataDiskCategory ecs.DiskCategory `json:"data_disk_category"`
	ECSImageID       string           `json:"ecs_image_id,omitempty"`
	IOOptimized      ecs.IoOptimized  `json:"io_optimized"`
}

type ClusterCreationResponse struct {
	Response
	ClusterID string `json:"cluster_id"`
}

func (client *Client) CreateCluster(region common.Region, args *ClusterCreationArgs) (cluster ClusterCreationResponse, err error) {
	err = client.Invoke(region, http.MethodPost, "/clusters", nil, args, &cluster)
	return
}

func (client *Client) DeleteCluster(clusterID string) error {
	return client.Invoke("", http.MethodDelete, "/clusters/"+clusterID, nil, nil, nil)
}

type ClusterCerts struct {
	CA   string `json:"ca,omitempty"`
	Key  string `json:"key,omitempty"`
	Cert string `json:"cert,omitempty"`
}

func (client *Client) GetClusterCerts(id string) (certs ClusterCerts, err error) {
	err = client.Invoke("", http.MethodGet, "/clusters/"+id+"/certs", nil, nil, &certs)
	return
}
