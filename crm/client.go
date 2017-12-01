package crm

import "github.com/denverdino/aliyungo/common"

type Client struct {
	common.Client
}

const (
	// CRMDefaultEndpoint is the default API endpoint of CRM services
	CRMDefaultEndpoint = "https://account-crm.aliyuncs.com"
	CRMAPIVersion      = "2016-06-06"
)

// NewClient creates a new instance of CRM client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := &Client{}
	client.Init(CRMDefaultEndpoint, CRMAPIVersion, accessKeyId, accessKeySecret)
	return client
}
