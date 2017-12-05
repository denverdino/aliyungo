package crm

import (
	"github.com/denverdino/aliyungo/common"
	"os"
)

type Client struct {
	common.Client
}

const (
	// CRMDefaultEndpoint is the default API endpoint of CRM services
	CRMDefaultEndpoint = "https://crm-cn-hangzhou.aliyuncs.com"
	CRMAPIVersion      = "2015-04-08"
	CRMServiceCode     = "crm"
)

// NewClient creates a new instance of CRM client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	return NewClientWithSecurityToken(accessKeyId, accessKeySecret,"")
}

func NewClientWithSecurityToken(accessKeyId, accessKeySecret, securityToken string) *Client {
	endpoint := os.Getenv("CRM_ENDPOINT")
	if endpoint == "" {
		endpoint = CRMDefaultEndpoint
	}

	return NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, securityToken)
}

func NewClientWithEndpointAndSecurityToken(endpoint string, accessKeyId, accessKeySecret, securityToken string) *Client {
	client := &Client{}
	client.WithEndpoint(endpoint).
		WithVersion(CRMAPIVersion).
		WithAccessKeyId(accessKeyId).
		WithAccessKeySecret(accessKeySecret).
		WithSecurityToken(securityToken).
		WithServiceCode(CRMServiceCode).
		InitClient()
	return client
}