package oss

type Region string

const (
	HANGZHOU       = Region("oss-cn-hangzhou")
	QINGDAO        = Region("oss-cn-qingdao")
	BEIJING        = Region("oss-cn-beijing")
	HONGKONG       = Region("oss-cn-hongkong")
	SHENZHEN       = Region("oss-cn-shenzhen")
	US_WEST_1      = Region("oss-us-west-1")
	DEFAULT_REGION = HANGZHOU
)

func (r Region) GetInternetEndpoint() string {
	return "http://" + string(r) + ".aliyuncs.com"
}

func (r Region) GetInternalEndpoint() string {
	return "http://" + string(r) + "-internal.aliyuncs.com"
}
