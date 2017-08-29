package slb

import "github.com/denverdino/aliyungo/common"

//Modify with your Access Key Id and Access Key Secret

const (
	TestAccessKeyId     = "MY_ACCESS_KEY_ID"
	TestAccessKeySecret = "MY_ACCESS_KEY_SECRET"
	TestLoadBlancerID   = "MY_LOADBALANCEID"
	TestVServerGroupID  = "MY_VSERVER_GROUPID"
	TestListenerPort    = 9000
	TestInstanceId      = "MY_INSTANCE_ID"
	TestRegionID        = common.APNorthEast1
	TestIAmRich         = false
	TestQuick           = false
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

var testDebugNewSLBClient *Client

func NewTestNewSLBClientForDebug() *Client {
	if testDebugNewSLBClient == nil {
		testDebugNewSLBClient = NewSLBClient(TestAccessKeyId, TestAccessKeySecret, TestRegionID)
		testDebugNewSLBClient.SetDebug(true)
	}
	return testDebugNewSLBClient
}
