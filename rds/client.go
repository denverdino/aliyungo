package rds

import (
	"github.com/denverdino/aliyungo/common"

	"os"
)

type Client struct {
	common.Client
}

const (
	// ECSDefaultEndpoint is the default API endpoint of RDS services
	RDSDefaultEndpoint = "https://rds.aliyuncs.com"
	RDSAPIVersion      = "2014-08-15"
)

var (
	RDSAPIEndpoint = map[common.Region]string{
		common.APNorthEast1: "https://rds.ap-northeast-1.aliyuncs.com", //日本
		common.EUCentral1:   "https://rds.eu-central-1.aliyuncs.com",   //德国
		common.MEEast1:      "https://rds.me-east-1.aliyuncs.com",      //迪拜
		common.APSouthEast2: "https://rds.ap-southeast-2.aliyuncs.com", //澳洲
		common.ZhangJiaKou:  "https://rds.cn-zhangjiakou.aliyuncs.com", //张家口
		common.APSouthEast3: "https://rds.ap-southeast-3.aliyuncs.com", //马来西亚
	}
)

// NewClient creates a new instance of RDS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	endpoint := os.Getenv("RDS_ENDPOINT")
	if endpoint == "" {
		endpoint = RDSDefaultEndpoint
	}
	return NewClientWithEndpoint(endpoint, accessKeyId, accessKeySecret)
}

func NewClientWithEndpoint(endpoint string, accessKeyId, accessKeySecret string) *Client {
	client := &Client{}
	client.Init(endpoint, RDSAPIVersion, accessKeyId, accessKeySecret)
	return client
}

func NewRDSClient(accessKeyId, accessKeySecret string, regionId common.Region) *Client {
	endpoint := os.Getenv("RDS_ENDPOINT")
	if endpoint == "" {
		endpoint = RDSDefaultEndpoint
		if v, ok := RDSAPIEndpoint[regionId]; ok {
			endpoint = v
		}
	}

	return NewClientWithEndpoint(endpoint, accessKeyId, accessKeySecret)
}

func CreateRDSClient(accessKeyId, accessKeySecret string, securityToken string, regionId common.Region) *Client {
	client := &Client{}
	endpoint := os.Getenv("RDS_ENDPOINT")
	if endpoint == "" {
		endpoint = RDSDefaultEndpoint
		if v, ok := RDSAPIEndpoint[regionId]; ok {
			endpoint = v
		}
	}

	client.InitForAssumeRole(endpoint, RDSAPIVersion, accessKeyId, accessKeySecret, securityToken, "")
	return client
}
