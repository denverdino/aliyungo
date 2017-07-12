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

var testRMClient *ResourceManagerClient

func NewTestRMClient() *ResourceManagerClient {
	if testRMClient == nil {
		testRMClient = NewRMClient(TestAccessKeyId, TestAccessKeySecret)
	}
	return testRMClient
}

var testRMClientForDebug *ResourceManagerClient

func NewTestRMClientForDebug() *ResourceManagerClient {
	if testRMClientForDebug == nil {
		testRMClientForDebug = NewRMClient(TestAccessKeyId, TestAccessKeySecret)
	}
	testClient.SetDebug(true)
	return testRMClientForDebug
}
