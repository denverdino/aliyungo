package oss

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

// GetInternetEndpoint return internet endpoint of region
func (r Region) GetInternetEndpoint() string {
	return "http://" + string(r) + ".aliyuncs.com"
}

// GetInternalEndpoint return internal endpoint of region
func (r Region) GetInternalEndpoint() string {
	return "http://" + string(r) + "-internal.aliyuncs.com"
}
