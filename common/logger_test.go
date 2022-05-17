package common

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Client_Logger(t *testing.T) {
	client := NewTestClientForDebug()
	assert.NotNil(t, client)
	args := &DescribeEndpointsArgs{
		Id:          Hangzhou,
		ServiceCode: "ecs",
		Type:        "openAPI",
	}
	// without logger
	resp, err := client.DescribeEndpoints(args)
	t.Log(resp)
	assert.Nil(t, err)

	// with logger
	wr := new(bytes.Buffer)
	assert.Nil(t, err)
	template := `{time} {channel}: {method} {host} {uri} HTTP/{version} {code} {cost} {hostname} {req_headers} {error} {res_body}`
	client.SetLogger("", "openapi", wr, template)
	resp, err = client.DescribeEndpoints(args)
	t.Log(wr.String())
	t.Log(resp)
	assert.Nil(t, err)
}
