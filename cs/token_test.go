package cs

import (
	"testing"
	"time"
)

func Test_CreateClusterToken(t *testing.T) {
	client := NewTestClientForDebug()

	req := &ClusterTokenReqeust{
		Expired:       time.Now().Unix() + 86400,
		IsPermanently: false,
	}

	token, err := client.CreateClusterToken(TestClusterId, req)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Token = %++v", token)
	}
}

func Test_RevokeClusterToken(t *testing.T) {
	client := NewTestClientForDebug()
	req := &ClusterTokenReqeust{
		Expired:       time.Now().Unix() + 86400,
		IsPermanently: false,
	}

	token, err := client.CreateClusterToken(TestClusterId, req)
	if err != nil {
		t.Fatalf("Error = %++v", err)
	} else {
		err = client.RevokeToken(token.Token)
		if err != nil {
			t.Fatalf("Error = %++v", err)
		} else {
			tokens, err := client.DescribeClusterTokens(TestClusterId)
			if err != nil {
				t.Fatalf("Error %++v", err)
			} else {
				for _, token := range tokens {
					t.Logf("Token = %++v", token)
				}
			}
		}
	}
}

func Test_DescribeClusterTokens(t *testing.T) {
	client := NewTestClientForDebug()

	tokens, err := client.DescribeClusterTokens(TestClusterId)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		for _, token := range tokens {
			t.Logf("Token = %++v", token)
		}
	}
}

func Test_DescribeClusterToken(t *testing.T) {
	client := NewTestClientForDebug()

	token, err := client.DescribeClusterToken(TestClusterId, TestToken)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Token = %++v", token)
	}
}
