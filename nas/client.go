package nas

import (
	"github.com/denverdino/aliyungo/common"
	"os"
)

const (
	VERSION            = "2016-02-29"
	END_POINT          = "https://nasservice-inner.aliyuncs.com"
	DEFAULT_POLICY     = "readwrite"
	DEFAULT_SQUASHTYPE = "no_squash"
	DEFAULT_PRIORITY   = "1"
)

type Client struct {
	common.Client
}

// NewClient creates a new instance of NAS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	return NewClientWithSecurityToken(accessKeyId, accessKeySecret,"")
}

func NewClientWithSecurityToken(accessKeyId, accessKeySecret, securityToken string) *Client {
	endpoint := os.Getenv("NAS_ENDPOINT")
	if endpoint == "" {
		endpoint = END_POINT
	}

	return NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, securityToken)
}

func NewClientWithEndpointAndSecurityToken(endpoint string, accessKeyId, accessKeySecret, securityToken string) *Client {
	client := &Client{}
	client.WithEndpoint(endpoint).
		WithVersion(VERSION).
		WithAccessKeyId(accessKeyId).
		WithAccessKeySecret(accessKeySecret).
		WithSecurityToken(securityToken).
		InitClient()
	return client
}