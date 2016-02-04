package sls

import "github.com/denverdino/aliyungo/common"


const(
	AccessKeySecret =""
	AccessKeyId = ""
	Region = common.Hangzhou
)
func DefaultProject() *Project {
	client := NewClient(Region, false, AccessKeyId, AccessKeySecret)
	return client.Project("yunqi-test")
}
