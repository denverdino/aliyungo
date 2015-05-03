package oss

import (
	"fmt"
	"io"
)

type ErrorResponse struct {
	Code      string
	Message   string
	RequestId string
	HostId    string
}

type OSSError struct {
	ErrorResponse
	StatusCode int
}

func (e *OSSError) Error() string {
	return fmt.Sprintf("OSS Error: Code %s: Message %s", e.Code, e.Message)
}

type BucketList struct {
	Prefix      string
	Marker      string
	MaxKeys     int
	IsTruncated bool
	NextMarker  string
	Owner       Owner    `xml:"Owner"`
	Buckets     []Bucket `xml:"Buckets>Bucket"`
}
type Owner struct {
	ID          string
	DisplayName string
}
type Bucket struct {
	Location     string
	Name         string
	CreationDate string
}

type ListBucketResult struct {
	Name       string
	Prefix     string
	Delimiter  string
	Marker     string
	MaxKeys    int
	NextMarker string
	// IsTruncated is true if the results have been truncated because
	// there are more keys and prefixes than can fit in MaxKeys.
	IsTruncated    bool
	Contents       []Content
	CommonPrefixes []string `xml:">Prefix"`
}

// The Content type holds individual file's information.
type Content struct {
	Key          string
	LastModified string
	ETag         string
	Type         string
	Size         int64
	StorageClass string
	Owner        Owner
}

type ObjectList struct {
	BucketName  string `xml:"Name"`
	Prefix      string
	Marker      string
	MaxKeys     int
	Delimiter   string
	IsTruncated bool
	Object      []Object `xml:"Contents"`
}
type Object struct {
	Key          string
	LastModified string
	ETag         string
	Type         string
	Size         int
	StorageClass string
	Owner        Owner
}

type AccessControlPolicy struct {
	Owner  Owner
	Grants []string `xml:AccessControlList>Grant"`
}
type CopyObjectResult struct {
	LastModified string
	ETag         string
}

type ACL string

const (
	Private           = ACL("private")
	PublicRead        = ACL("public-read")
	PublicReadWrite   = ACL("public-read-write")
	AuthenticatedRead = ACL("authenticated-read")
	BucketOwnerRead   = ACL("bucket-owner-read")
	BucketOwnerFull   = ACL("bucket-owner-full-control")
)

type NOPCloser struct {
	io.Reader
}

func (NOPCloser) Close() error {
	return nil
}
