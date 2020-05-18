package cen

import (
	"os"

	"github.com/denverdino/aliyungo/common"
)

// Interval for checking status in WaitForXXX method
const DefaultWaitForInterval = 5

// Default timeout value for WaitForXXX method
const DefaultTimeout = 60

type Client struct {
	common.Client
}

const (
	// CENDefaultEndpoint is the default API endpoint of CEN services
	CENDefaultEndpoint = "https://cbn.aliyuncs.com"
	CENAPIVersion      = "2017-09-12"
	CENServiceCode     = "cen"
)

// ---------------------------------------
// NewCENClient creates a new instance of CEN client
// ---------------------------------------
func NewCENClient(accessKeyId, accessKeySecret string, regionID common.Region) *Client {
	return NewCENClientWithSecurityToken(accessKeyId, accessKeySecret, "", regionID)
}

func NewCENClientWithSecurityToken(accessKeyId string, accessKeySecret string, securityToken string, regionID common.Region) *Client {
	endpoint := os.Getenv("CEN_ENDPOINT")
	if endpoint == "" {
		endpoint = CENDefaultEndpoint
	}

	return NewCENClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, securityToken, regionID)
}

//only for Hangzhou Regional Domain
func NewCENClientWithSecurityToken4RegionalDomain(accessKeyId string, accessKeySecret string, securityToken string, regionID common.Region) *Client {
	endpoint := os.Getenv("CEN_ENDPOINT")
	if endpoint == "" {
		endpoint = CENDefaultEndpoint
	}

	return NewCENClientWithEndpointAndSecurityToken4RegionalDomain(endpoint, accessKeyId, accessKeySecret, securityToken, regionID)
}

func NewCENClientWithEndpointAndSecurityToken(endpoint string, accessKeyId string, accessKeySecret string, securityToken string, regionID common.Region) *Client {
	client := &Client{}
	client.WithEndpoint(endpoint).
		WithVersion(CENAPIVersion).
		WithAccessKeyId(accessKeyId).
		WithAccessKeySecret(accessKeySecret).
		WithSecurityToken(securityToken).
		WithServiceCode(CENServiceCode).
		WithRegionID(regionID).
		InitClient()
	return client
}

func NewCENClientWithEndpointAndSecurityToken4RegionalDomain(endpoint string, accessKeyId string, accessKeySecret string, securityToken string, regionID common.Region) *Client {
	client := &Client{}
	client.WithEndpoint(endpoint).
		WithVersion(CENAPIVersion).
		WithAccessKeyId(accessKeyId).
		WithAccessKeySecret(accessKeySecret).
		WithSecurityToken(securityToken).
		WithServiceCode(CENServiceCode).
		WithRegionID(regionID).
		InitClient4RegionalDomain()
	return client
}
