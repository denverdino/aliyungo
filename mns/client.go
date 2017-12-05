package mns

import (
	"github.com/denverdino/aliyungo/common"
)

// ----------------------------------
// MNS is the short name of Message Service
// See: https://help.aliyun.com/product/27412.html
// ----------------------------------

const (
	MNSAPIVersion = "2015-06-06"
	MNSServiceCode     = "mns"
)

type Client struct {
	common.Client
}

func NewClient(accessKeyId, accessKeySecret, endpoint string) (client *Client) {
	return NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret,"")
}

func NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, securityToken string) *Client {
	client := &Client{}
	client.WithEndpoint(endpoint).
		WithVersion(MNSAPIVersion).
		WithAccessKeyId(accessKeyId).
		WithAccessKeySecret(accessKeySecret).
		WithSecurityToken(securityToken).
		WithServiceCode(MNSServiceCode).
		InitClient()
	return client
}