package mns

import (
	"net/http"
)

const (
	MNSAPIVersion = "2015-06-06"
)

type Client struct {
	AccessKeyId     string
	AccessKeySecret string
	Endpoint        string
	Version         string
	httpClient      *http.Client
}

func NewClient(accessKeyId, accessKeySecret, endpoint string) (client *Client) {
	client = &Client{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Endpoint:        endpoint,
		Version:         MNSAPIVersion,
		httpClient:      &http.Client{},
	}
	return client
}
