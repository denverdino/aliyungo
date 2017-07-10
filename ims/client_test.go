package ims

import "os"

var (
	TestAccessKeyId     = os.Getenv("AccessKeyId")
	TestAccessKeySecret = os.Getenv("AccessKeySecret")
)

var testClient *ImsClient

func NewTestClient() *ImsClient {
	if testClient == nil {
		testClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
	}
	return testClient
}

var testClientForDebug *ImsClient

func NewTestClientForDebug() *ImsClient {
	if testClient == nil {
		testClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
	}
	testClient.SetDebug(true)
	return testClient
}
