package ros

import (
	"github.com/denverdino/aliyungo/common"
	"os"
)

const (
	ROS_RPC_APIVersion = "2019-09-10"
	ROSServiceCode     = "ros"
)

type RpcClient struct {
	common.Client
}

func NewRpcClient(accessKeyId, accessKeySecret string) *RpcClient {
	endpoint := os.Getenv("ROS_ENDPOINT")
	if endpoint == "" {
		endpoint = ROSDefaultEndpoint
	}
	return NewRpcClientWithEndpoint(endpoint, accessKeyId, accessKeySecret)
}

func NewRpcClientWithEndpoint(endpoint string, accessKeyId string, accessKeySecret string) *RpcClient {
	client := &RpcClient{}
	client.Init(endpoint, ROS_RPC_APIVersion, accessKeyId, accessKeySecret)
	return client
}

func NewRosRpcClientWithSecurityToken(accessKeyId string, accessKeySecret string, securityToken string, regionID common.Region) *RpcClient {
	endpoint := os.Getenv("ROS_ENDPOINT")
	if endpoint == "" {
		endpoint = ROSDefaultEndpoint
	}

	return NewRosRpcClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, securityToken, regionID)
}

func NewRosRpcClientWithEndpointAndSecurityToken(endpoint string, accessKeyId string, accessKeySecret string, securityToken string, regionID common.Region) *RpcClient {
	client := &RpcClient{}
	client.WithEndpoint(endpoint).
		WithVersion(ROS_RPC_APIVersion).
		WithAccessKeyId(accessKeyId).
		WithAccessKeySecret(accessKeySecret).
		WithSecurityToken(securityToken).
		WithServiceCode(ROSServiceCode).
		WithRegionID(regionID).
		InitClient()
	return client
}
