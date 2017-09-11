package ros

import (
	"os"

	"github.com/denverdino/aliyungo/common"
)

//Modify with your Access Key Id and Access Key Secret

var (
	TestAccessKeyId     = os.Getenv("AccessKeyId")
	TestAccessKeySecret = os.Getenv("AccessKeySecret")
	TestRegionId        = common.Region(os.Getenv("RegionId"))

	clientForTestCase      = NewTestClient()
	debugClientForTestCase = NewTestClientForDebug()
)

var testClient *Client

func NewTestClient() *Client {
	if testClient == nil {
		testClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
	}
	return testClient
}

var testDebugClient *Client

func NewTestClientForDebug() *Client {
	if testDebugClient == nil {
		testDebugClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
		testDebugClient.SetDebug(true)
	}
	return testDebugClient
}
