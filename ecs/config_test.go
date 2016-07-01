package ecs

//Modify with your Access Key Id and Access Key Secret

const (
	// BID 小号
	//	TestAccessKeyId     = "sczRHxPls3VC1O2S"
	//	TestAccessKeySecret = "xg6t4cvsLiOtN7hutd5xhReppxei7x"
	//BID 大账号
	TestAccessKeyId     = "ACSbW2iBbyX0Pk9N"
	TestAccessKeySecret = "TLSwMm5LQU"
	TestInstanceId      = "i-25xna88dv"
	TestSecurityGroupId = "sg-252okhmti"
	TestImageId         = "suse12sp1_64_40G_cloudinit_20160520.vhd"
	TestOwnerId         = ""                   //1542159133705458"
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
