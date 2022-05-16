package common

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
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
