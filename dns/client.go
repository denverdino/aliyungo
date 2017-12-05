package dns

import (
	"github.com/denverdino/aliyungo/common"
	"os"
)

type Client struct {
	common.Client
}

const (
	// DNSDefaultEndpoint is the default API endpoint of DNS services
	DNSDefaultEndpoint = "http://dns.aliyuncs.com"
	DNSAPIVersion      = "2015-01-09"

	DNSDefaultEndpointNew = "http://alidns.aliyuncs.com"
)

// NewClient creates a new instance of DNS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	endpoint := os.Getenv("DNS_ENDPOINT")
	if endpoint == "" {
		endpoint = DNSDefaultEndpoint
	}
	return NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, "")
}

// NewClientNew creates a new instance of DNS client, with http://alidns.aliyuncs.com as default endpoint
func NewClientNew(accessKeyId, accessKeySecret string) *Client {
	return NewClientWithSecurityToken(accessKeyId, accessKeySecret,"")
}

// NewCustomClient creates a new instance of ECS client with customized API endpoint
func NewCustomClient(accessKeyId, accessKeySecret string, endpoint string) *Client {
	return NewClientWithEndpoint(endpoint, accessKeyId, accessKeySecret)
}

func NewClientWithEndpoint(endpoint, accessKeyId, accessKeySecret string) *Client {
	return NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, "")
}

func NewClientWithSecurityToken(accessKeyId, accessKeySecret, securityToken string) *Client {
	endpoint := os.Getenv("DNS_ENDPOINT")
	if endpoint == "" {
		endpoint = DNSDefaultEndpointNew
	}

	return NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, securityToken)
}

func NewClientWithEndpointAndSecurityToken(endpoint string, accessKeyId, accessKeySecret, securityToken string) *Client {
	client := &Client{}
	client.WithEndpoint(endpoint).
		WithVersion(DNSAPIVersion).
		WithAccessKeyId(accessKeyId).
		WithAccessKeySecret(accessKeySecret).
		WithSecurityToken(securityToken).
		InitClient()
	return client
}