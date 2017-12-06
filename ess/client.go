package ess

import (
	"github.com/denverdino/aliyungo/common"

	"os"
)

type Client struct {
	common.Client
}

const (
	// ESSDefaultEndpoint is the default API endpoint of ESS services
	ESSDefaultEndpoint = "https://ess.aliyuncs.com"
	ESSAPIVersion      = "2014-08-28"
	ESSServiceCode     = "ess"
)

// NewClient creates a new instance of RDS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	return NewClientWithSecurityToken(accessKeyId, accessKeySecret, "", "")
}

func NewClientWithEndpoint(endpoint string, accessKeyId, accessKeySecret string) *Client {
	return NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, "", "")
}

func NewESSClient(accessKeyId, accessKeySecret string, regionID common.Region) *Client {
	return NewClientWithSecurityToken(accessKeyId, accessKeySecret, "", regionID)
}

func NewClientWithRegion(endpoint string, accessKeyId, accessKeySecret string, regionID common.Region) *Client {
	return NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, "", regionID)
}

func NewClientWithSecurityToken(accessKeyId, accessKeySecret, securityToken string, regionID common.Region) *Client {
	endpoint := os.Getenv("ESS_ENDPOINT")
	if endpoint == "" {
		endpoint = ESSDefaultEndpoint
	}

	return NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, securityToken, regionID)
}

func NewClientWithEndpointAndSecurityToken(endpoint string, accessKeyId, accessKeySecret, securityToken string, regionID common.Region) *Client {
	client := &Client{}
	client.WithEndpoint(endpoint).
		WithVersion(ESSAPIVersion).
		WithAccessKeyId(accessKeyId).
		WithAccessKeySecret(accessKeySecret).
		WithSecurityToken(securityToken).
		WithServiceCode(ESSServiceCode).
		WithRegionID(regionID).
		InitClient()
	return client
}