package common

import (
	"net/http"
	"os"
	"testing"
)

var (
	TestAccessKeyId     = os.Getenv("AccessKeyId")
	TestAccessKeySecret = os.Getenv("AccessKeySecret")
	TestSecurityToken   = os.Getenv("SecurityToken")
	TestRegionId        = os.Getenv("RegionId")
	TestServiceCode     = os.Getenv("ServiceCode")
)
var testDebugClient *LocationClient

func NewTestClientForDebug() *LocationClient {
	if testDebugClient == nil {
		testDebugClient = NewLocationClient(TestAccessKeyId, TestAccessKeySecret, TestSecurityToken)
		testDebugClient.SetDebug(true)
	}
	return testDebugClient
}

func TestClient_SetTransport(t *testing.T) {
	client := NewTestClientForDebug()
	transport := &myTransport{}
	client.SetTransport(transport)
	if client.httpClient.Transport.(*myTransport) != transport {
		t.Fail()
	}
}

type myTransport struct{}

func (m *myTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return http.DefaultTransport.RoundTrip(req)
}
