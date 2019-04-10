package acm

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func getClient() *Client {
	client, err := NewClient(func(c *Client) {
		c.AccessKey = os.Getenv("AccessKeyId")
		c.SecretKey = os.Getenv("AccessKeySecret")
		c.EndPoint = "acm.aliyun.com"
		c.NameSpace = "test"
	})

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func RunWithTest(t *testing.T, test func(client *Client, t *testing.T)) {
	client := getClient()
	defer client.Delete("test", "test")

	_, err := client.Publish("test", "test", "test测试")

	if err != nil {
		t.Fatalf("pulish error:%s", err)
	}

	test(client, t)
}

func TestNewClient(t *testing.T) {
	client := getClient()
	servers := client.GetServers()

	if len(servers) == 0 {
		t.Error("get server error")
	}
}

func TestClient_GetConfig(t *testing.T) {
	RunWithTest(t, func(client *Client, t *testing.T) {
		ret, err := client.GetConfig("test", "test")
		if err != nil {
			t.Error(err)
		}
		if ret != "test测试" {
			t.Error("wrong respond content")
		}
		fmt.Println(ret)
	})
}

func TestClient_Subscribe(t *testing.T) {
	RunWithTest(t, func(client *Client, t *testing.T) {
		_, err := client.Subscribe("test", "test", "")
		if err != nil {
			t.Error(err)
		}
	})
}
