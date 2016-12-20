package mq

import "testing"


// 您在控制台创建的Topic
var Topic = ""
// 公测集群URL
var ENDPOINT = ""
// 阿里云官网身份验证访问码
var Ak = ""
// 阿里云身份验证密钥
var Sk = ""
// MQ控制台创建的Producer ID
var ProducerID =""
// MQ控制台创建的Consumer ID
var ConsumerID = ""


func TestNewClient(t *testing.T) {
	client := NewClient(Ak, Sk, ENDPOINT, Topic, ProducerID, ConsumerID, "http", "http")
	msgId, err := client.send(GetCurrentUnixMicro(), "hello world")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("The message id successfully send is %v", msgId)
	}
}



