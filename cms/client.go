package cms

import (
	"os"

	"github.com/denverdino/aliyungo/common"
)

type Client struct {
	common.Client
	internal bool
}

type CMSClient struct {
	common.Client
}

const (
	//TODO 旧的API，暂时保留
	DefaultEndpoint = "http://alert.aliyuncs.com"
	APIVersion      = "2015-08-15"
	METHOD_GET      = "GET"
	METHOD_POST     = "POST"
	METHOD_PUT      = "PUT"
	METHOD_DELETE   = "DELETE"

	CMSDefaultEndpoint = "http://metrics.cn-hangzhou.aliyuncs.com"
	CMSAPIVersion      = "2017-03-01"
	CMSServiceCode     = "cms"
)

// NewClient creates a new instance of ECS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	return NewClientWithSecurityToken(accessKeyId, accessKeySecret, "")
}

func NewClientWithSecurityToken(accessKeyId, accessKeySecret, securityToken string) *Client {
	endpoint := os.Getenv("CMS_ENDPOINT")
	if endpoint == "" {
		endpoint = CMSDefaultEndpoint
	}

	return NewClientWithEndpointAndSecurityToken(endpoint, false, accessKeyId, accessKeySecret, securityToken)
}

func NewClientWithEndpointAndSecurityToken(endpoint string, internal bool, accessKeyId, accessKeySecret, securityToken string) *Client {
	client := &Client{
		internal: internal,
	}
	client.WithEndpoint(endpoint).
		WithVersion(CMSAPIVersion).
		WithAccessKeyId(accessKeyId).
		WithAccessKeySecret(accessKeySecret).
		WithSecurityToken(securityToken).
		WithServiceCode(CMSServiceCode).
		InitClient()
	return client
}

func (client *Client) GetApiUri() string {
	return client.Endpoint()
}

func (client *Client) GetAccessKey() string {
	return client.AccessKeyId
}

func (client *Client) GetAccessSecret() string {
	return client.AccessKeySecret
}

// NewClient creates a new instance of CMS client
func NewCMSClient(accessKeyId, accessKeySecret string) *CMSClient {
	endpoint := os.Getenv("CMS_ENDPOINT")
	if endpoint == "" {
		endpoint = CMSDefaultEndpoint
	}
	return NewClientWithEndpoint(endpoint, accessKeyId, accessKeySecret)
}

func NewClientWithEndpoint(endpoint string, accessKeyId, accessKeySecret string) *CMSClient {
	client := &CMSClient{}
	client.Init(endpoint, CMSAPIVersion, accessKeyId, accessKeySecret)
	return client
}

func NewCMSRegionClient(accessKeyId, accessKeySecret string, regionID common.Region) *CMSClient {
	endpoint := os.Getenv("CMS_ENDPOINT")
	if endpoint == "" {
		endpoint = CMSDefaultEndpoint
	}

	return NewClientWithRegion(endpoint, accessKeyId, accessKeySecret, regionID)
}

func NewClientWithRegion(endpoint string, accessKeyId, accessKeySecret string, regionID common.Region) *CMSClient {
	client := &CMSClient{}
	client.NewInit(endpoint, CMSAPIVersion, accessKeyId, accessKeySecret, CMSServiceCode, regionID)
	return client
}
