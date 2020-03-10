package ros

import (
	"os"

	"github.com/denverdino/aliyungo/common"
)

//Modify with your Access Key Id and Access Key Secret

var (
	TestAccessKeyId     = os.Getenv("AccessKeyId")
	TestAccessKeySecret = os.Getenv("AccessKeySecret")
	TestSecurityToken   = os.Getenv("SecurityToken")
	TestRegionID        = common.Region(os.Getenv("RegionId"))

	debugClientForTestCase    = NewTestClientForDebug()
	debugRpcClientForTestCase = NewTestRpcClientForDebug()
)

var testClient *Client

func NewTestClient() *Client {
	if testClient == nil {
		testClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
	}
	testClient.SetSecurityToken(TestSecurityToken)
	return testClient
}

var testDebugClient *Client
var testDebugRpcClient *RpcClient

func NewTestClientForDebug() *Client {
	if testDebugClient == nil {
		testDebugClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
		testDebugClient.SetDebug(true)
	}
	testDebugClient.SetSecurityToken(TestSecurityToken)
	return testDebugClient
}

func NewTestRpcClientForDebug() *RpcClient {
	if testDebugRpcClient == nil {
		testDebugRpcClient = NewRpcClient(TestAccessKeyId, TestAccessKeySecret)
		testDebugRpcClient.SetDebug(true)
	}
	testDebugRpcClient.SetSecurityToken(TestSecurityToken)
	return testDebugRpcClient
}
