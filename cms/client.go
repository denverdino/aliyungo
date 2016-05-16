// 云监控数据上报接口  Cloud Monitoring Service
package cms

import (
	"github.com/denverdino/aliyungo/common"
)

const (
	// CMSDefaultEndpoint is the default API endpoint of Cloud Monitoring Services
	CMSDefaultEndpoint = "http://metrics.aliyuncs.com"
	CMSAPIVersion      = "2015-10-20"
)

type CmsClient struct {
	common.Client
}

func NewClient(accessKeyId string, accessKeySecret string) CmsClientInterface {
	client := &CmsClient{}
	client.Init(CMSDefaultEndpoint, CMSAPIVersion, accessKeyId, accessKeySecret)
	client.SetDebug(true)
	return client
}
