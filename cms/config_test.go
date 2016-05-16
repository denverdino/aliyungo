package cms

const (
	TestAccessKeyId     = "<YOUR ACCESS_KEY_ID>"
	TestAccessKeySecret = "<YOUR ACCESS_KEY_SECRET>"
)

var testClient CmsClientInterface

func NewTestClient() CmsClientInterface {
	if testClient == nil {
		testClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
	}
	return testClient
}
