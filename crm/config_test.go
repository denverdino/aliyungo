package crm

import (
	"os"

	"github.com/denverdino/aliyungo/common"
	"github.com/menglingwei/aliyungo/crm"
)

var (
	TestAccessKeyId     = os.Getenv("AccessKeyId")
	TestAccessKeySecret = os.Getenv("AccessKeySecret")
	TestSecurityToken   = os.Getenv("SecurityToken")
	TestRegionID        = common.Region(os.Getenv("RegionId"))
)

var testDebugClient *crm.Client

func NewTestClientForDebug() *crm.Client {
	if testDebugClient == nil {
		testDebugClient = crm.NewClient(TestAccessKeyId, TestAccessKeySecret)
		testDebugClient.SetDebug(true)
	}
	return testDebugClient
}
