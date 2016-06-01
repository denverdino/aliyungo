package rds

import (
	"github.com/denverdino/aliyungo/common"

)




type Client struct {
	common.Client
}

const (
	// ECSDefaultEndpoint is the default API endpoint of RDS services
	RDSDefaultEndpoint = "https://rds.aliyuncs.com"
	RDSAPIVersion      = "2014-08-15"
)

// NewClient creates a new instance of RDS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := &Client{}
	client.Init(RDSDefaultEndpoint, RDSAPIVersion, accessKeyId, accessKeySecret)
	return client
}




