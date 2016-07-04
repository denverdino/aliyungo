package slb

//Modify with your Access Key Id and Access Key Secret

const (
	// BID 小号
	//	TestAccessKeyId     = "MY_ACCESS_KEY_ID"
	//	TestAccessKeySecret = "MY_ACCESS_KEY_SECRET"
	//BID 大账号
	TestAccessKeyId     = "MY_ACCESS_KEY_ID"
	TestAccessKeySecret = "MY_ACCESS_KEY_SECRET"
	TestInstanceId      = "MY_INSTANCE_ID"
	TestOwnerId         = ""
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

var testBIDDebugClient *Client

func NewTestBIDClientForDebug() *Client {
	if testBIDDebugClient == nil {
		testBIDDebugClient = NewClientWithOwnerID(TestAccessKeyId, TestAccessKeySecret, TestOwnerId)
		testBIDDebugClient.SetDebug(true)
	}
	return testBIDDebugClient
}
