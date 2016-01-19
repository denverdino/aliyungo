package ram

import (
	"github.com/denverdino/aliyungo/common"
)

const (
	// RAMDefaultEndpoint is the default API endpoint of RAM services
	RAMDefaultEndpoint = "https://ram.aliyuncs.com"
	RAMAPIVersion      = "2015-05-01"
)

type RamClient struct {
	common.Client
}

func NewClient(accessKeyId string, accessKeySecret string) RamClientInterface {
	client := &RamClient{}
	client.Init(RAMDefaultEndpoint, RAMAPIVersion, accessKeyId, accessKeySecret)
	return client
}
