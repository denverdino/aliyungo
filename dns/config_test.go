package dns

import "os"

//Modify with your Access Key Id and Access Key Secret
var (
	TestAccessKeyId     = os.Getenv("MY_ACCESS_KEY_ID")
	TestAccessKeySecret = os.Getenv("MY_ACCESS_KEY_SECRET")
	TestDomainName      = os.Getenv("MY_TEST_TOP_DOMAIN")
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
