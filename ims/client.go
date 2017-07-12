package ims

import (
	"os"

	"github.com/denverdino/aliyungo/common"
)

const (
	// IMSDefaultEndpoint is the default API endpoint of RAM services
	IMSDefaultEndpoint = "http://ims.aliyuncs.com"
	IMSAPIVersion      = "2017-04-30"

	ResourceManagerDefaultEndpoint = "https://resourcemanager.aliyuncs.com"
	ResourceManagerAPIVersion      = "2016-11-11"
)

type ImsClient struct {
	common.Client
}

type ResourceManagerClient struct {
	common.Client
}

func NewClient(accessKeyId string, accessKeySecret string) *ImsClient {
	endpoint := os.Getenv("IMS_ENDPOINT")
	if endpoint == "" {
		endpoint = IMSDefaultEndpoint
	}
	return NewClientWithEndpoint(endpoint, accessKeyId, accessKeySecret)
}

func NewClientWithEndpoint(endpoint string, accessKeyId string, accessKeySecret string) *ImsClient {
	client := &ImsClient{}
	client.Init(endpoint, IMSAPIVersion, accessKeyId, accessKeySecret)
	return client
}

func NewRMClient(accessKeyId string, accessKeySecret string) *ResourceManagerClient {
	endpoint := os.Getenv("RESOURCEMANAGER_ENDPOINT")
	if endpoint == "" {
		endpoint = ResourceManagerDefaultEndpoint
	}
	return NewRMClientWithEndpoint(endpoint, accessKeyId, accessKeySecret)
}

func NewRMClientWithEndpoint(endpoint string, accessKeyId string, accessKeySecret string) *ResourceManagerClient {
	client := &ResourceManagerClient{}
	client.Init(endpoint, ResourceManagerAPIVersion, accessKeyId, accessKeySecret)
	return client
}
