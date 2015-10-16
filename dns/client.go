package dns

import "github.com/denverdino/aliyungo/common"

type Client struct {
	common.Client
}

const (
	// DNSDefaultEndpoint is the default API endpoint of DNS services
	DNSDefaultEndpoint = "https://dns.aliyuncs.com"
	DNSAPIVersion      = "2015-01-09"
)

// NewClient creates a new instance of DNS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := &Client{}
	client.Init(DNSDefaultEndpoint, DNSAPIVersion, accessKeyId, accessKeySecret)
	return client
}
