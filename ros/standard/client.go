package standard

import (
	"github.com/denverdino/aliyungo/common"

	"os"
)

type Client struct {
	common.Client
}

const (
	// ROSDefaultEndpoint is the default API endpoint of ESS services
	ROSDefaultEndpoint = "https://ros.aliyuncs.com"
	ROSAPIVersion      = "2019-09-10"
	ROSServiceCode     = "ros"
)

// NewClient creates a new instance of RDS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	endpoint := os.Getenv("ROS_ENDPOINT")
	if endpoint == "" {
		endpoint = ROSDefaultEndpoint
	}
	return NewClientWithEndpoint(endpoint, accessKeyId, accessKeySecret)
}

func NewClientWithEndpoint(endpoint string, accessKeyId, accessKeySecret string) *Client {
	client := &Client{}
	client.Init(endpoint, ROSAPIVersion, accessKeyId, accessKeySecret)
	return client
}

func NewROSClient(accessKeyId, accessKeySecret string, regionID common.Region) *Client {
	endpoint := os.Getenv("ROS_ENDPOINT")
	if endpoint == "" {
		endpoint = ROSDefaultEndpoint
	}

	return NewClientWithRegion(endpoint, accessKeyId, accessKeySecret, regionID)
}

func NewClientWithRegion(endpoint string, accessKeyId, accessKeySecret string, regionID common.Region) *Client {
	client := &Client{}
	client.NewInit(endpoint, ROSAPIVersion, accessKeyId, accessKeySecret, ROSServiceCode, regionID)
	return client
}
