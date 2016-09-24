package dm

import (
	"github.com/denverdino/aliyungo/common"
)

const (
	Url              = "https://dm.aliyuncs.com/"
	SingleSMS        = "SingleSendSms"
	SingleMail       = "SingleSendMail"
	BathMail         = "BatchSendMail"
	APIVersion       = "2015-11-23"
	AcceptXML        = "XML"
	AcceptJson       = "JSON"
	SignatureMethod  = "HMAC-SHA1"
	TimestampFormat  = "YYYY-MM-DDThh:mm:ssZ"
	SignatureVersion = "1.0"
)

type Client struct {
	accessKeyId, accessKeySecret string
}

// NewClient creates a new instance of dm client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	return NewClientWithEndpoint(accessKeyId, accessKeySecret)
}
