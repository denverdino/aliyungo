package nas

import (
	"github.com/denverdino/aliyungo/common"
)

const (
	VERSION            = "2017-06-26"
	END_POINT          = "https://nas.aliyuncs.com"
	DEFAULT_POLICY     = "readwrite"
	DEFAULT_SQUASHTYPE = "no_squash"
	DEFAULT_PRIORITY   = "1"
)

type Client struct {
	common.Client
}

// NewClient creates a new instance of NAS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := &Client{}
	client.Init(END_POINT, VERSION, accessKeyId, accessKeySecret)
	return client
}
