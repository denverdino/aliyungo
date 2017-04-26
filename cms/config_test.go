package cms

import "github.com/denverdino/aliyungo/common"

const (
	TestAccessKeyId     = "YOUR_ACCESS_KEY_ID"
	TestAccessKeySecret = "YOUR_ACCESS_KEY_SECRET"
	TestRegionID        = common.Hangzhou
)

var testClient *CMSClient

func NewTestClient() *CMSClient {
	if testClient == nil {
		testClient = NewCMSClient(TestAccessKeyId, TestAccessKeySecret)
	}
	return testClient
}

var testDebugClient *CMSClient

func NewTestClientForDebug() *CMSClient {
	if testDebugClient == nil {
		testDebugClient = NewCMSClient(TestAccessKeyId, TestAccessKeySecret)
		testDebugClient.SetDebug(true)
	}
	return testDebugClient
}
