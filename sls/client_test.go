package sls

import "github.com/denverdino/aliyungo/common"


const (
	AccessKeyId = ""
	AccessKeySecret = ""
	Region = common.Hangzhou
)
func DefaultProject() *Project {
	client := NewClient(Region, false, AccessKeyId, AccessKeySecret)
	p, _ := client.Project("yunqi-test")
	return p
}
