package sls

import (
	"github.com/hdksky/aliyungo/common"
	"testing"
	"github.com/golang/protobuf/proto"
	"time"
)

const (
	AccessKeyId      = ""
	AccessKeySecret  = ""
	Region           = common.Hangzhou
	TestProjectName  = "test-project123"
	TestLogstoreName = "test-logstore"
)

func DefaultProject(t *testing.T) *Project {
	client := NewClient(Region, false, AccessKeyId, AccessKeySecret)
	err := client.CreateProject(TestProjectName, "description")
	if err != nil {
		if e, ok := err.(*Error); ok && e.Code != "ProjectAlreadyExist" {
			t.Fatalf("create project fail: %s", err.Error())
		}
	}
	p, err := client.Project(TestProjectName)
	if err != nil {
		t.Fatalf("get project fail: %s", err.Error())
	}
	//Create default logstore

	logstore := &Logstore{
		TTL:   2,
		Shard: 3,
		Name:  TestLogstoreName,
	}
	err = p.CreateLogstore(logstore)
	if err != nil {
		if e, ok := err.(*Error); ok && e.Code != "LogStoreAlreadyExist" {
			t.Fatalf("create logstore fail: %s", err.Error())
		}
	}

	return p
}

func TestClient_PutLogs(t *testing.T) {
	region           := common.Beijing
	project  := "testych"
	logStore := "test1"

	client := NewClient(region, false, AccessKeyId, AccessKeySecret)

	contents := []*Log_Content{}
	key := "log1"
	value := "value1"
	contents = append( contents, &Log_Content{
		Key: &key,
		Value: &value,
	})
	key2 := "log2"
	value2 := "value2"
	contents = append( contents, &Log_Content{
		Key: &key2,
		Value: &value2,
	})
	
	logs := []*Log{}
	logs = append(logs, &Log{
						Time: proto.Uint32(uint32(time.Now().Unix())),
						Contents: contents,
					})
	
	
	request := &PutLogsRequest{
		Project : project,
		LogStore: logStore,
		LogItems : LogGroup{
			Logs: logs,
		},
	}
	
	err:=client.PutLogs( request )
	if err!= nil {
		t.Errorf( "get the error %v", err )
	}
}