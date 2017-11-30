package rds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_ModifySecurityIps(t *testing.T) {

	if TestAccessKeyId == "" {
		t.SkipNow()
	}
	client := NewClient(TestAccessKeyId, TestAccessKeySecret)
	client.SetDebug(true)

	// TODO:
	args := &ModifySecurityIpsArgs{
		DBInstanceId: "xxxx",
		SecurityIps:  "x.x.x.x,x.x.x.x",
	}
	resp, err := client.ModifySecurityIps(args)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	t.Logf("the result is %++v ", resp)
}

func TestClient_DescribeRegions(t *testing.T) {
	client := NewTestClientForDebug()
	client.SetSecurityToken(TestSecurityToken)

	regions, err := client.DescribeRegions()
	if err != nil {
		t.Fatal("Error %++v", err)
	} else {
		t.Logf("Result = %++v", regions)
	}
}
