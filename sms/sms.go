package sms

import (
	"net/http"

	"github.com/denverdino/aliyungo/common"
)

//阿里云通信
type SendSmsArgs struct {
	PhoneNumbers    string
	SignName        string
	TemplateCode    string
	TemplateParam   string
	SmsUpExtendCode string `ArgName:"smsUpExtendCode"`
	OutId           string
}

type SendSmsResponse struct {
	common.Response
	Code string
	Message string
	BizId string
}

func (this *DYSmsClient) SendSms(args *SendSmsArgs) (*SendSmsResponse, error) {
	resp := SendSmsResponse{}
	return &resp, this.InvokeByAnyMethod(http.MethodGet, SendSms, "", args, &resp)
}

//邮件推送产品短信功能
type SingleSendSmsArgs struct {
	SignName string
	TemplateCode  string
	RecNum        string
	ParamString   string
}

func (this *Client) SingleSendSms(args *SingleSendSmsArgs) error {
	return this.InvokeByAnyMethod(http.MethodPost, SingleSendSms, "", args, &common.Response{})
}
