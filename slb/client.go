package slb

import "github.com/denverdino/aliyungo/common"

type Client struct {
	common.Client
}

const (
	// SLBDefaultEndpoint is the default API endpoint of SLB services
	SLBDefaultEndpoint = "https://slb.aliyuncs.com"
	SLBAPIVersion      = "2014-05-15"
)

// NewClient creates a new instance of ECS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := &Client{}
	client.Init(SLBDefaultEndpoint, SLBAPIVersion, accessKeyId, accessKeySecret)
	return client
}
