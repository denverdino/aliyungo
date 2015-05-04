package oss

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// The Bucket type encapsulates operations with an OSS bucket.
type Bucket struct {
	*Client
	Name string
}

// Bucket returns a Bucket with the given name.
func (client *Client) Bucket(name string) *Bucket {
	return &Bucket{
		Client: client,
		Name:   name,
	}
}

func (b *Bucket) bucketOp(method string, headers http.Header, params url.Values, response interface{}) error {

	if params == nil {
		params = make(url.Values)
	}
	url := "/" + b.Name + "?" + params.Encode()

	httpResp, err := b.Client.Invoke(method, url, strings.NewReader(""), headers)

	if err == nil {
		b.Client.decodeResponse(httpResp, response)
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
func (b *Bucket) GetBucketAcl() (result *AccessControlPolicy, err error) {
	params := make(url.Values)
	params.Add("acl", "")
	result = &AccessControlPolicy{}
	err = b.bucketOp("GET", nil, params, result)
	if err != nil {
		return nil, err
	}
	return
}

// List returns list of buckets
func (b *Bucket) List(prefix, marker, delimiter string, maxkeys int) (result *ObjectList, err error) {

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
	if maxkeys != 0 {
		params.Add("max-keys", strconv.Itoa(maxkeys)) //TODO: check max-keys <= 1000?
	}
	result = &ObjectList{}
	err = b.bucketOp("GET", nil, params, result)
	if err != nil {
		return nil, err
	}
	return
}

// PutBucket creates bucket with optional ACL setting
func (b *Bucket) PutBucket(acl ACL) error {
	headers := make(http.Header)
	if acl != "" {
		headers.Add(HeaderOSSACL, string(acl))
	}
	return b.bucketOp("PUT", headers, nil, nil)
}

// CreateBucket creates bucket with optional ACL setting
func (b *Bucket) CreateBucket(acl ACL) error {
	return b.PutBucket(acl)
}

// DeleteBucket deletes bucket
func (b *Bucket) DeleteBucket() error {
	return b.bucketOp("DELETE", nil, nil, nil)
}
