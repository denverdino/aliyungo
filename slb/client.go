package slb

import (
	"os"

	"github.com/denverdino/aliyungo/common"
)

type Client struct {
	common.Client
}

const (
	// SLBDefaultEndpoint is the default API endpoint of SLB services
	SLBDefaultEndpoint = "https://slb.aliyuncs.com"
	SLBAPIVersion      = "2014-05-15"
)

var (
	SLBAPIEndpoint = map[common.Region]string{
		common.APNorthEast1: "https://slb.ap-northeast-1.aliyuncs.com", //日本
		common.EUCentral1:   "https://slb.eu-central-1.aliyuncs.com",   //德国
		common.MEEast1:      "https://slb.me-east-1.aliyuncs.com",      //迪拜
		common.APSouthEast2: "https://slb.ap-southeast-2.aliyuncs.com", //澳洲
		common.ZhangJiaKou:  "https://slb.cn-zhangjiakou.aliyuncs.com", //张家口
	}
)

// NewClient creates a new instance of SLB client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	endpoint := os.Getenv("SLB_ENDPOINT")
	if endpoint == "" {
		endpoint = SLBDefaultEndpoint
	}
	return NewClientWithEndpoint(endpoint, accessKeyId, accessKeySecret)
}

func NewClientWithEndpoint(endpoint string, accessKeyId, accessKeySecret string) *Client {
	client := &Client{}
	client.Init(endpoint, SLBAPIVersion, accessKeyId, accessKeySecret)
	return client
}

func NewClientWithOwnerID(accessKeyId, accessKeySecret string, ownerId string) *Client {
	client := &Client{}
	endpoint := os.Getenv("SLB_ENDPOINT")
	if endpoint == "" {
		endpoint = SLBDefaultEndpoint
	}
	client.InitWithOwnerId(endpoint, SLBAPIVersion, accessKeyId, accessKeySecret, ownerId)
	return client
}

func NewSLBClient(accessKeyId, accessKeySecret, ownerId string, regionId common.Region) *Client {
	client := &Client{}
	endpoint := os.Getenv("SLB_ENDPOINT")
	if endpoint == "" {
		endpoint = SLBDefaultEndpoint
		if v, ok := SLBAPIEndpoint[regionId]; ok {
			endpoint = v
		}
	}

	client.InitWithOwnerId(endpoint, SLBAPIVersion, accessKeyId, accessKeySecret, ownerId)
	return client
}
