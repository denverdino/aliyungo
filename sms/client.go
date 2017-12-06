package sms

import (
	"github.com/denverdino/aliyungo/common"
	"os"
)

//dysms是阿里云通讯的短信服务，sms是旧版本邮件推送产品短信功能，两者互不相通，必须使用对应的API。
//目前阿里已经把短信服务整合至MNS中，2017年6月22日以后开通的用户请使用MNS来发送短信。
//Refer to: https://help.aliyun.com/product/44282.html

const (
	DYSmsEndPoint   = "https://dysmsapi.aliyuncs.com/"
	SendSms = "SendSms"
	DYSmsAPIVersion = "2017-05-25"

	SmsEndPoint = "https://sms.aliyuncs.com/"
	SingleSendSms = "SingleSendSms"
	SmsAPIVersion = "2016-09-27"
)

type Client struct {
	common.Client
}

func NewClient(accessKeyId, accessKeySecret string) *Client {
	return NewSMSClientWithSecurityToken(accessKeyId, accessKeySecret,"")
}

type DYSmsClient struct {
	common.Client
	Region common.Region
}

func NewDYSmsClient(accessKeyId, accessKeySecret string) *DYSmsClient {
	return NewDYSmsClientWithSecurityToken(accessKeyId, accessKeySecret, "", common.Hangzhou)
}

func NewSMSClientWithSecurityToken(accessKeyId, accessKeySecret, securityToken string) *Client {
	endpoint := os.Getenv("SMS_ENDPOINT")
	if endpoint == "" {
		endpoint = SmsEndPoint
	}

	return NewSMSClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, securityToken)
}

func NewSMSClientWithEndpointAndSecurityToken(endpoint string, accessKeyId, accessKeySecret, securityToken string) *Client {
	client := &Client{}
	client.WithEndpoint(endpoint).
		WithVersion(SmsAPIVersion).
		WithAccessKeyId(accessKeyId).
		WithAccessKeySecret(accessKeySecret).
		WithSecurityToken(securityToken).
		InitClient()
	return client
}

func NewDYSmsClientWithSecurityToken(accessKeyId, accessKeySecret, securityToken string, regionID common.Region) *DYSmsClient {
	endpoint := os.Getenv("DYSMS_ENDPOINT")
	if endpoint == "" {
		endpoint = DYSmsEndPoint
	}

	return NewDYSmsClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, securityToken, regionID)
}

func NewDYSmsClientWithEndpointAndSecurityToken(endpoint string, accessKeyId, accessKeySecret, securityToken string, regionID common.Region) *DYSmsClient {
	client := &DYSmsClient{}
	client.WithEndpoint(endpoint).
		WithVersion(DYSmsAPIVersion).
		WithAccessKeyId(accessKeyId).
		WithAccessKeySecret(accessKeySecret).
		WithSecurityToken(securityToken).
		WithRegionID(regionID).
		InitClient()
	return client
}