package ims

import (
	"os"

	"github.com/denverdino/aliyungo/common"
)

const (
	// IMSDefaultEndpoint is the default API endpoint of RAM services
	IMSDefaultEndpoint = "http://ims.aliyuncs.com"
	IMSAPIVersion      = "2017-04-30"
)

type ImsClient struct {
	common.Client
}

func NewClient(accessKeyId string, accessKeySecret string) *ImsClient {
	endpoint := os.Getenv("RAM_ENDPOINT")
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
