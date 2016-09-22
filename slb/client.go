package slb

import (
	"os"

	"github.com/denverdino/aliyungo/common"
)

type Client struct {
	common.Client
}

const (
	// SLBDefaultEndpoint is the default API endpoint of SLB services
	SLBDefaultEndpoint = "https://slb.aliyuncs.com"
	SLBAPIVersion      = "2014-05-15"
)

// NewClient creates a new instance of SLB client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	endpoint := os.Getenv("SLB_ENDPOINT")
	if endpoint == "" {
		endpoint = SLBDefaultEndpoint
	}
	return NewClientWithEndpoint(endpoint, accessKeyId, accessKeySecret)
}

func NewClientWithEndpoint(endpoint string, accessKeyId, accessKeySecret string) *Client {
	client := &Client{}
	client.Init(endpoint, SLBAPIVersion, accessKeyId, accessKeySecret)
	return client
}

func NewClientWithOwnerID(accessKeyId, accessKeySecret string, ownerId string) *Client {
	client := &Client{}
	endpoint := os.Getenv("SLB_ENDPOINT")
	if endpoint == "" {
		endpoint = SLBDefaultEndpoint
	}
	client.InitWithOwnerId(endpoint, SLBAPIVersion, accessKeyId, accessKeySecret, ownerId)
	return client
}
