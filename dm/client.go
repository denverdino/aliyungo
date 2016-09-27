package dm

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"
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
	return &Client{accessKeyId, accessKeySecret}
}

func (this Client) newParamMap() *url.Values {
	ret := &url.Values{}
	ret.Add("AccessKeyId", this.accessKeyId)
	ret.Add("Format", AcceptJson)
	ret.Add("Version", APIVersion)
	ret.Add("SignatureMethod", SignatureMethod)
	ret.Add("SignatureVersion", SignatureVersion)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret.Add("SignatureNonce", fmt.Sprintf("%d", r.Int63()))
	ret.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	return ret
}
