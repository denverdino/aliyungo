package ecs

import "github.com/denverdino/aliyungo/common"

//Modify with your Access Key Id and Access Key Secret

const (
	// BID 小号
	//	TestAccessKeyId     = "MY_ACCESS_KEY_ID"
	//	TestAccessKeySecret = "MY_ACCESS_KEY_SECRET"
	//BID 大账号
	TestAccessKeyId     = "MY_ACCESS_KEY_ID"
	TestAccessKeySecret = "MY_ACCESS_KEY_SECRET"
	TestInstanceId      = "MY_ECS_INSTANCE_ID"
	TestSecurityGroupId = "MY_SECURITY_GROUP_ID"
	TestInstanceType    = "ecs.n4.small"
	TestRegionId        = common.APNorthEast1
	TestImageId         = "ubuntu1404_64_40G_cloudinit_20160727.raw"
	TestOwnerId         = ""
	TestAccountId       = "MY_TEST_ACCOUNT_ID" //Get from https://account.console.aliyun.com

	TestIAmRich = false
	TestQuick   = false
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

var testBIDDebugClient *Client

func NewTestBIDClientForDebug() *Client {
	if testBIDDebugClient == nil {
		testBIDDebugClient = NewClientWithOwnerID(TestAccessKeyId, TestAccessKeySecret, TestOwnerId)
		testBIDDebugClient.SetDebug(true)
	}

	return testBIDDebugClient
}

var testDebugECSClient *Client

func NewTestECSClientForDebug() *Client {
	if testDebugECSClient == nil {
		testDebugECSClient = NewECSClient(TestAccessKeyId, TestAccessKeySecret, TestOwnerId, TestRegionId)
		testDebugECSClient.SetDebug(true)
	}

	return testDebugECSClient
}
