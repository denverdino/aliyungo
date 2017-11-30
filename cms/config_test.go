package cms

import (
	"os"

	"github.com/denverdino/aliyungo/common"
)

var (
	TestAccessKeyId     = os.Getenv("AccessKeyId")
	TestAccessKeySecret = os.Getenv("AccessKeySecret")
	TestSecurityToken   = os.Getenv("SecurityToken")
	TestRegionID        = common.Region(os.Getenv("RegionId"))
)

var testClient CmsClientInterface

func NewTestClient() CmsClientInterface {
	if testClient == nil {
		testClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
	}

	return testClient
}

var testClientForDebug CmsClientInterface

func NewTestClientForDebug() CmsClientInterface {
	if testClientForDebug == nil {
		testClientForDebug = CreateCMSClient(TestAccessKeyId, TestAccessKeySecret, TestSecurityToken)
	}

	return testClientForDebug
}
