package oss

import (
	"fmt"
)

// Region represents OSS region
type Region string

// Constants of region definition
const (
	Hangzhou      = Region("oss-cn-hangzhou")
	Qingdao       = Region("oss-cn-qingdao")
	Beijing       = Region("oss-cn-beijing")
	Hongkong      = Region("oss-cn-hongkong")
	Shenzhen      = Region("oss-cn-shenzhen")
	USWest1       = Region("oss-us-west-1")
	DefaultRegion = Hangzhou
)

// GetEndpoint returns endpoint of region
func (r Region) GetEndpoint(internal bool, bucket string) string {
	if internal {
		return r.GetInternalEndpoint(bucket)
	}
	return r.GetInternetEndpoint(bucket)
}

// GetInternetEndpoint returns internet endpoint of region
func (r Region) GetInternetEndpoint(bucket string) string {
	if bucket == "" {
		return fmt.Sprintf("http://%s.aliyuncs.com", string(r))
	}
	return fmt.Sprintf("http://%s.%s.aliyuncs.com", bucket, string(r))
}

// GetInternalEndpoint returns internal endpoint of region
func (r Region) GetInternalEndpoint(bucket string) string {
	if bucket == "" {
		return fmt.Sprintf("http://%s-internal.aliyuncs.com", string(r))
	}
	return fmt.Sprintf("http://%s.%s-internal.aliyuncs.com", bucket, string(r))
}
