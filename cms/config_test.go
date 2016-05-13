package cms

const (
	TestAccessKeyId     = "ACSbW2iBbyX0Pk9N"
	TestAccessKeySecret = "TLSwMm5LQU"
)

var testClient CmsClientInterface

func NewTestClient() CmsClientInterface {
	if testClient == nil {
		testClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
	}
	return testClient
}
