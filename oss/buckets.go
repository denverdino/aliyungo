package oss

import (
	"net/http"
	"net/url"
	"strings"
)

func (client *Client) BucketOp(method, bucket string, headers http.Header, params url.Values, response interface{}) error {

	if params == nil {
		params = make(url.Values)
	}
	url := "/" + bucket + "?" + params.Encode()

	httpResp, err := client.Invoke(method, url, strings.NewReader(""), headers)

	if err == nil {
		client.decodeResponse(httpResp, response)
	}

	return err
}

func (client *Client) GetService() (bucketList *BucketList, err error) {
	bucketList = &BucketList{}
	httpResp, err := client.Invoke("GET", "/", strings.NewReader(""), nil)
	if err == nil {
		client.decodeResponse(httpResp, &bucketList)
	}
	return
}

func (client *Client) GetBucketAcl(bucket string) (result *AccessControlPolicy, err error) {
	params := make(url.Values)
	params.Add("acl", "")
	result = &AccessControlPolicy{}
	err = client.BucketOp("GET", bucket, nil, params, result)
	if err != nil {
		return nil, err
	}
	return
}

func (client *Client) GetBucket(bucket, prefix, marker, delimiter, maxkeys string) (result *ListBucketResult, err error) {
	return client.listBucket(bucket, prefix, marker, delimiter, maxkeys)
}

func (client *Client) ListBucket(prefix, marker, delimiter, maxkeys string) (result *ListBucketResult, err error) {
	return client.listBucket("", prefix, marker, delimiter, maxkeys)
}

func (client *Client) listBucket(bucket, prefix, marker, delimiter, maxkeys string) (result *ListBucketResult, err error) {
	params := make(url.Values)
	if prefix != "" {
		params.Add("prefix", prefix)
	}
	if marker != "" {
		params.Add("marker", marker)
	}
	if delimiter != "" {
		params.Add("delimiter", delimiter)
	}
	if maxkeys != "" {
		params.Add("max-keys", maxkeys) //TODO: check max-keys <= 1000?
	}
	result = &ListBucketResult{}
	err = client.BucketOp("GET", bucket, nil, params, result)
	if err != nil {
		return nil, err
	}
	return
}

func (client *Client) PutBucket(bucket, acl string) error {
	headers := make(http.Header)
	if acl != "" {
		headers.Add(HEADER_X_OSS_ACL, acl)
	}
	return client.BucketOp("PUT", bucket, headers, nil, nil)
}

func (client *Client) CreateBucket(bucket, acl string) error {
	return client.PutBucket(bucket, acl)
}

func (client *Client) DeleteBucket(bucket string) error {
	return client.BucketOp("DELETE", bucket, nil, nil, nil)
}
