package grandcanal

import (
	"os"
	"strings"

	"github.com/google/uuid"

	"github.com/denverdino/aliyungo/common"
)

type Client struct {
	common.Client
}

const (
	// ECSDefaultEndpoint is the default API endpoint of ECS services
	ECSDefaultEndpoint = "https://grandcanal-inner.aliyuncs.com"
	ECSAPIVersion      = "2016-07-15"
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

// Generate a random string
func RandString() string {
	s := strings.Replace(uuid.New().String(), "-", "", -1)
	return strings.Join([]string{"A", s}, "")
}
