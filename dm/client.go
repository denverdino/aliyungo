package dm

import (
	"github.com/denverdino/aliyungo/common"
	"os"
)

// --------------------------------------------
// DM is the short name of Direct Mail
// See: https://help.aliyun.com/product/29412.html
// --------------------------------------------

const (
	EmailEndPoint   = "https://dm.aliyuncs.com/"
	SingleSendMail  = "SingleSendMail"
	BatchSendMail   = "BatchSendMail"
	EmailAPIVersion = "2015-11-23"
)

type Client struct {
	common.Client
}

func NewClient(accessKeyId, accessKeySecret string) *Client {
	return NewClientWithSecurityToken(accessKeyId, accessKeySecret,"")
}

func NewClientWithSecurityToken(accessKeyId, accessKeySecret, securityToken string) *Client {
	endpoint := os.Getenv("DM_ENDPOINT")
	if endpoint == "" {
		endpoint = EmailEndPoint
	}

	return NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, securityToken)
}

func NewClientWithEndpointAndSecurityToken(endpoint string, accessKeyId, accessKeySecret, securityToken string) *Client {
	client := &Client{}
	client.WithEndpoint(endpoint).
		WithVersion(EmailAPIVersion).
		WithAccessKeyId(accessKeyId).
		WithAccessKeySecret(accessKeySecret).
		WithSecurityToken(securityToken).
		InitClient()
	return client
}