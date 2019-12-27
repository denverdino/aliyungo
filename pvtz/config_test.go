package pvtz

import (
	"github.com/denverdino/aliyungo/common"
	"os"
)

//Modify with your Access Key Id and Access Key Secret

var (
	TestAccessKeyId     = os.Getenv("AccessKeyId")
	TestAccessKeySecret = os.Getenv("AccessKeySecret")
	TestVPCId           = os.Getenv("VPCId")
	TestSecurityToken   = os.Getenv("SecurityToken")
	TestRegionID        = common.Region(os.Getenv("RegionId"))
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

var testDebugLocationClient *Client

func NewTestLocationClientForDebug() *Client {
	if testDebugLocationClient == nil {
		testDebugLocationClient = NewPVTZClientWithSecurityToken4RegionalDomain(TestAccessKeyId, TestAccessKeySecret, TestSecurityToken, TestRegionID)
		testDebugLocationClient.SetDebug(true)
	}
	return testDebugLocationClient
}
