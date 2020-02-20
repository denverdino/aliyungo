package common

import (
	"fmt"
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

func Test_InitClient4RegionalDomain(t *testing.T) {

	var tests = []struct {
		service  string
		version  string
		endpoint string
	}{
		{"ecs", "2014-05-26", "https://ecs-cn-hangzhou.aliyuncs.com"},
		{"pvtz", "2018-01-01", "https://pvtz.aliyuncs.com"},
		{"slb", "2014-05-15", "https://slb.aliyuncs.com"},
		{"vpc", "2016-04-28", "https://vpc.aliyuncs.com"},
	}

	for _, test := range tests {
		for _, region := range ValidRegions {
			if region == Qingdao || region == ShenZhenFinance || region == ShanghaiFinance {
				continue
			}

			client := &Client{}
			client.SetDebug(true)
			client.WithEndpoint(test.endpoint).
				WithVersion(test.version).
				WithAccessKeyId(TestAccessKeyId).
				WithAccessKeySecret(TestAccessKeySecret).
				WithServiceCode(test.service).
				WithRegionID(region).
				InitClient4RegionalDomain()

			domain := fmt.Sprintf("https://%s.%s.aliyuncs.com", test.service, region)

			if client.endpoint != domain {
				if test.service == "vpc" && (region == Beijing || region == Hongkong || region == Shanghai) {
					continue
				}
				t.Fail()
			}
		}

	}
}
