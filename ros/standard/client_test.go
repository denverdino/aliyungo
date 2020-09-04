package standard

import (
	"fmt"
	"github.com/denverdino/aliyungo/common"
	"testing"
)

const (
	TestAccessKeyId     = ""
	TestAccessKeySecret = ""
	Region              = common.Shanghai
	StackID             = "c7f1fed9-0104-4596-aae6-aa5215f5a793"
)

func NewTestClient() *Client {
	return NewROSClient(TestAccessKeyId, TestAccessKeySecret, Region)
}

func TestGetStack(t *testing.T) {

	client := NewTestClient()
	req := GetStackRequest{
		RegionId: Region,
		StackId:  StackID,
	}
	res, err := client.GetStack(&req)
	if err != nil {
		t.Fail()
	}
	fmt.Printf("Response: %+v\n", res)
}

func TestListStack(t *testing.T) {
	client := NewTestClient()
	req := &ListStacksRequest{}
	res, err := client.ListStacks(req)
	if err != nil {
		t.Fail()
	}
	fmt.Printf("ListResponse: %+v\n", res)
}

func TestListStackEvent(t *testing.T) {
	client := NewTestClient()
	req := &ListStackEventsRequest{
		StackId: StackID,
	}
	res, err := client.ListStackEvents(req)
	if err != nil {
		t.Fail()
	}
	fmt.Printf("ListEventResponse: %+v\n", res)
}
