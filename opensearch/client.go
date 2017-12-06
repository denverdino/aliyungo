package opensearch

import (
	"github.com/denverdino/aliyungo/common"
	"os"
	"fmt"
)

const (
	Internet = ""
	Intranet = "intranet."
	VPC = "vpc."
	APIVersion = "v2"
)

type Client struct {
	common.Client
}

//OpenSearch的API比较奇怪，action不在公共参数里面
type OpenSearchArgs struct {
	Action string `ArgName:"action"`
}

func NewClient(networkType string, regionID common.Region, accessKeyId, accessKeySecret string) *Client {
	return NewClientWithSecurityToken(accessKeyId, accessKeySecret, "", regionID, networkType)
}

func NewClientWithSecurityToken(accessKeyId, accessKeySecret, securityToken string, regionID common.Region, networkType string) *Client {
	endpoint := os.Getenv("ESS_ENDPOINT")
	if endpoint == "" {
		endpoint = buildEndpoint(networkType, regionID)
	}

	return NewClientWithEndpointAndSecurityToken(endpoint, accessKeyId, accessKeySecret, securityToken, regionID)
}

func NewClientWithEndpointAndSecurityToken(endpoint string, accessKeyId, accessKeySecret, securityToken string, regionID common.Region) *Client {
	client := &Client{}
	client.WithEndpoint(endpoint).
		WithVersion(APIVersion).
		WithAccessKeyId(accessKeyId).
		WithAccessKeySecret(accessKeySecret).
		WithSecurityToken(securityToken).
		WithRegionID(regionID).
		InitClient()
	return client
}

func buildEndpoint(networkType string, regionID common.Region) string {
	return fmt.Sprintf("http://%sopensearch-%s.aliyuncs.com", networkType, regionID)
}