package oss

import (
	"net/http"
	"net/url"
	"strings"
)

func (client *Client) bucketOp(method, bucket string, headers http.Header, params url.Values, response interface{}) error {

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

// GetService returns list of all buckets
func (client *Client) GetService() (bucketList *BucketList, err error) {
	bucketList = &BucketList{}
	httpResp, err := client.Invoke("GET", "/", strings.NewReader(""), nil)
	if err == nil {
		client.decodeResponse(httpResp, &bucketList)
	}
	return
}

// GetBucketAcl returns ACL of bucket
func (client *Client) GetBucketAcl(bucket string) (result *AccessControlPolicy, err error) {
	params := make(url.Values)
	params.Add("acl", "")
	result = &AccessControlPolicy{}
	err = client.bucketOp("GET", bucket, nil, params, result)
	if err != nil {
		return nil, err
	}
	return
}

// GetBucket returns list of buckets
func (client *Client) GetBucket(bucket, prefix, marker, delimiter, maxkeys string) (result *ObjectList, err error) {

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
	result = &ObjectList{}
	err = client.bucketOp("GET", bucket, nil, params, result)
	if err != nil {
		return nil, err
	}
	return
}

// PutBucket creates bucket with optional ACL setting
func (client *Client) PutBucket(bucket, acl string) error {
	headers := make(http.Header)
	if acl != "" {
		headers.Add(HeaderOSSACL, acl)
	}
	return client.bucketOp("PUT", bucket, headers, nil, nil)
}

// CreateBucket creates bucket with optional ACL setting
func (client *Client) CreateBucket(bucket, acl string) error {
	return client.PutBucket(bucket, acl)
}

// DeleteBucket deletes bucket
func (client *Client) DeleteBucket(bucket string) error {
	return client.bucketOp("DELETE", bucket, nil, nil, nil)
}
