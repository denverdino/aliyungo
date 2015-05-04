package oss

import (
	"fmt"
)

// An ErrorResponse represents OSS error response
type ErrorResponse struct {
	Code      string
	Message   string
	RequestId string
	HostId    string
}

// An Error represents custom error for OSS failure response
type Error struct {
	ErrorResponse
	StatusCode int
}

func (e *Error) Error() string {
	return fmt.Sprintf("OSS Error: Code %s: Message %s", e.Code, e.Message)
}

// BucketList represents response for GetServices and GetBucket
type BucketList struct {
	Prefix      string
	Marker      string
	MaxKeys     int
	IsTruncated bool
	NextMarker  string
	Owner       Owner        `xml:"Owner"`
	Buckets     []BucketInfo `xml:"Buckets>Bucket"`
}

type Owner struct {
	ID          string
	DisplayName string
}

type BucketInfo struct {
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
	Object      []ObjectInfo `xml:"Contents"`
}
type ObjectInfo struct {
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
	Grants []string `xml:"AccessControlList>Grant"`
}

// CopyObjectResult is the output from a Copy request
type CopyObjectResult struct {
	ETag         string
	LastModified string
}

// An ACL presents the ACL setting
type ACL string

// Constants of ACLs
const (
	Private           = ACL("private")
	PublicRead        = ACL("public-read")
	PublicReadWrite   = ACL("public-read-write")
	AuthenticatedRead = ACL("authenticated-read")
	BucketOwnerRead   = ACL("bucket-owner-read")
	BucketOwnerFull   = ACL("bucket-owner-full-control")
)
