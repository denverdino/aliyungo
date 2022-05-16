package common

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
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
			if region == Qingdao || region == HangZhouFinance {
				continue
			}

			client := &Client{}
			client.SetDebug(true)
			client.WithEndpoint(test.endpoint).
				WithVersion(test.version).
				WithAccessKeyId(TestAccessKeyId).
				WithAccessKeySecret(TestAccessKeySecret).
				WithSecurityToken(TestSecurityToken).
				WithServiceCode(test.service).
				WithRegionID(region).
				InitClient4RegionalDomain()

			if endpoint, ok := CentralDomainServices[test.service]; ok {
				domain := fmt.Sprintf("https://%s", endpoint)
				if client.endpoint != domain {
					t.Fail()
				}
				continue
			}

			if ep, ok := UnitRegions[region]; ok {
				domain := fmt.Sprintf("https://%s.%s.aliyuncs.com", test.service, ep)
				if client.endpoint != domain {
					t.Fail()
				}
				continue
			}

			domain := fmt.Sprintf("https://%s%s.%s.aliyuncs.com", test.service, "-vpc", region)
			if client.endpoint != domain {
				t.Fail()
			}
		}

	}
}

func Test_InvokeTracer(t *testing.T) {
	client := NewTestClientForDebug()
	assert.NotNil(t, client)
	args := &DescribeEndpointsArgs{
		Id:          Hangzhou,
		ServiceCode: "ecs",
		Type:        "openAPI",
	}
	// not set global tracer
	resp, err := client.DescribeEndpoints(args)
	t.Log(resp)
	assert.Nil(t, err)

	//set global tracer, no root span
	var cfg = jaegercfg.Configuration{
		ServiceName: "client test",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}
	jLogger := jaegerlog.StdLogger
	tracer, closer, _ := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
	)
	opentracing.InitGlobalTracer(tracer)
	resp, err = client.DescribeEndpoints(args)
	t.Log(resp)
	assert.Nil(t, err)

	// set global tracer, with root span
	parentSpan := tracer.StartSpan("root")
	fmt.Println(parentSpan)
	client.SetSpan(parentSpan)
	resp, err = client.DescribeEndpoints(args)
	t.Log(resp)
	assert.Nil(t, err)

	// set disable trace
	client.SetDisableTrace(true)
	client.SetSpan(parentSpan)
	resp, err = client.DescribeEndpoints(args)
	t.Log(resp)
	assert.Nil(t, err)

	parentSpan.Finish()
	closer.Close()
}
