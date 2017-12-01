package nas

import (
	"os"

	"github.com/denverdino/aliyungo/common"
)

var (
	TestAccessKeyId     = os.Getenv("AccessKeyId")
	TestAccessKeySecret = os.Getenv("AccessKeySecret")
	TestSecurityToken   = os.Getenv("SecurityToken")
	TestRegionID        = common.Region(os.Getenv("RegionId"))
	TestAccessGroupName = os.Getenv("AccessGroupName")
	TestFileSystemId    = os.Getenv("FileSystemId")
	TestVpcId           = os.Getenv("VpcId")
	TestVSwitchId       = os.Getenv("VSwitchId")
)

var testDebugClient *Client

func NewTestClientForDebug() *Client {
	if testDebugClient == nil {
		testDebugClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
		testDebugClient.SetDebug(true)
	}
	return testDebugClient
}
