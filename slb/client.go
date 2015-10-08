package slb

import "github.com/denverdino/aliyungo/common"

type Client struct {
	common.Client
}

const (
	// ECSDefaultEndpoint is the default API endpoint of ECS services
	ECSDefaultEndpoint = "https://slb.aliyuncs.com"
	ECSAPIVersion      = "2014-05-15"
)

// NewClient creates a new instance of ECS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := &Client{}
	client.Init(ECSDefaultEndpoint, ECSAPIVersion, accessKeyId, accessKeySecret)
	return client
}
