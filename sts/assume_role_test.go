package sts

import (
	"os"
	"testing"
	"time"

	"fmt"
)

func TestSTSClient_AssumeRole(t *testing.T) {
	client := NewTestClient()

	roleArn := os.Getenv("RoleArn")

	req := AssumeRoleRequest{
		RoleArn:         roleArn,
		RoleSessionName: fmt.Sprintf("commander-role-%d", time.Now().Unix()),
		DurationSeconds: 3600,
	}

	response, err := client.AssumeRole(req)
	if err != nil {
		t.Fatalf("%++v", err)
	} else {
		t.Logf("Response=%++v", response)
	}
}
