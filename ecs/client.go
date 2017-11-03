package ecs

import (
	"os"

	"github.com/denverdino/aliyungo/common"
)

// Interval for checking status in WaitForXXX method
const DefaultWaitForInterval = 5

// Default timeout value for WaitForXXX method
const DefaultTimeout = 60

type Client struct {
	common.Client
}

const (
	// ECSDefaultEndpoint is the default API endpoint of ECS services
	ECSDefaultEndpoint = "https://ecs-cn-hangzhou.aliyuncs.com"
	ECSAPIVersion      = "2014-05-26"
)

var (
	ECSAPIEndpoint = map[common.Region]string{
		common.APNorthEast1: "https://ecs.ap-northeast-1.aliyuncs.com", //日本
		common.EUCentral1:   "https://ecs.eu-central-1.aliyuncs.com",   //德国
		common.MEEast1:      "https://ecs.me-east-1.aliyuncs.com",      //迪拜
		common.APSouthEast2: "https://ecs.ap-southeast-2.aliyuncs.com", //澳洲
		common.ZhangJiaKou:  "https://ecs.cn-zhangjiakou.aliyuncs.com", //华北三
		common.APSouthEast3: "https://ecs.ap-southeast-3.aliyuncs.com", //马来西亚
	}
)

// NewClient creates a new instance of ECS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	endpoint := os.Getenv("ECS_ENDPOINT")
	if endpoint == "" {
		endpoint = ECSDefaultEndpoint
	}
	return NewClientWithEndpoint(endpoint, accessKeyId, accessKeySecret)
}

func NewClientWithEndpoint(endpoint string, accessKeyId, accessKeySecret string) *Client {
	client := &Client{}
	client.Init(endpoint, ECSAPIVersion, accessKeyId, accessKeySecret)
	return client
}

func NewClientWithOwnerID(accessKeyId, accessKeySecret string, ownerId string) *Client {
	client := &Client{}
	endpoint := os.Getenv("ECS_ENDPOINT")
	if endpoint == "" {
		endpoint = ECSDefaultEndpoint
	}
	client.InitWithOwnerId(endpoint, ECSAPIVersion, accessKeyId, accessKeySecret, ownerId)
	return client
}

func NewECSClient(accessKeyId, accessKeySecret, ownerId string, regionId common.Region) *Client {
	client := &Client{}
	endpoint := os.Getenv("ECS_ENDPOINT")
	if endpoint == "" {
		endpoint = ECSDefaultEndpoint
		if v, ok := ECSAPIEndpoint[regionId]; ok {
			endpoint = v
		}
	}

	client.InitWithOwnerId(endpoint, ECSAPIVersion, accessKeyId, accessKeySecret, ownerId)
	return client
}
