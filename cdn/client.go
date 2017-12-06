package cdn

import (
	"os"

	"github.com/denverdino/aliyungo/common"
)

const (
	// CDNDefaultEndpoint is the default API endpoint of CDN services
	CDNDefaultEndpoint = "https://cdn.aliyuncs.com"
	CDNAPIVersion      = "2014-11-11"
	CDNServiceCode     = "cdn"
)

type CdnClient struct {
	common.Client
}

func NewClient(accessKeyId string, accessKeySecret string) *CdnClient {
	return NewClientWithSecurityToken(accessKeyId, accessKeySecret, "")
}

func NewClientWithSecurityToken(accessKeyId, accessKeySecret, securityToken string) *CdnClient {
	endpoint := os.Getenv("CDN_ENDPOINT")
	if endpoint == "" {
		endpoint = CDNDefaultEndpoint
	}

	return NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, securityToken)
}

func NewClientWithEndpoint(endpoint string, accessKeyId string, accessKeySecret string) *CdnClient {
	client := &CdnClient{}
	client.Init(endpoint, CDNAPIVersion, accessKeyId, accessKeySecret)
	return client
}

func NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, securityToken string) *CdnClient {
	client := &CdnClient{}
	client.WithEndpoint(endpoint).
		WithVersion(CDNAPIVersion).
		WithAccessKeyId(accessKeyId).
		WithAccessKeySecret(accessKeySecret).
		WithSecurityToken(securityToken).
		WithServiceCode(CDNServiceCode).
		InitClient()
	return client
}
