package location

import (
	"os"

	"github.com/denverdino/aliyungo/common"
)

type Client struct {
	common.Client
}

const (
	// LocationDefaultEndpoint is the default API endpoint of Location services
	LocationDefaultEndpoint = "https://location.aliyuncs.com"
	LocationAPIVersion      = "2015-06-12"
)

// NewClient creates a new instance of Location client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	return NewClientWithSecurityToken(accessKeyId, accessKeySecret, "")
}

func NewClientWithEndpoint(endpoint string, accessKeyId, accessKeySecret string) *Client {
	return NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, "")
}

func NewClientWithSecurityToken(accessKeyId, accessKeySecret, securityToken string) *Client {
	endpoint := os.Getenv("LOCATION_ENDPOINT")
	if endpoint == "" {
		endpoint = LocationDefaultEndpoint
	}

	return NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, securityToken)
}

func NewClientWithEndpointAndSecurityToken(endpoint string, accessKeyId, accessKeySecret, securityToken string) *Client {
	client := &Client{}
	client.WithEndpoint(endpoint).
		WithVersion(LocationAPIVersion).
		WithAccessKeyId(accessKeyId).
		WithAccessKeySecret(accessKeySecret).
		WithSecurityToken(securityToken).
		InitClient()
	return client
}