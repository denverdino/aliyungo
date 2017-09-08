package smsv2

import "github.com/denverdino/aliyungo/common"

const (
	SmsEndPoint   = "https://dysmsapi.aliyuncs.com/"
	SendSMS       = "SendSms"
	SmsAPIVersion = "2017-05-25"
)

type Client struct {
	common.Client
}

func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := new(Client)
	client.Init(SmsEndPoint, SmsAPIVersion, accessKeyId, accessKeySecret)
	return client
}
