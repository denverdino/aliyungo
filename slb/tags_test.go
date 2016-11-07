package slb

import (
	"encoding/json"
	"testing"

	"github.com/denverdino/aliyungo/common"
)

func getTestTags() string {
	a := make([]TagItem, 0)
	a = append(a, TagItem{TagKey: "provider", TagValue: "acs"})
	a = append(a, TagItem{TagKey: "clusterid", TagValue: "cxxxxxxxx"})
	a = append(a, TagItem{TagKey: "types", TagValue: "hybird"})
	b, _ := json.Marshal(a)
	return string(b)
}

func TestAddTags(t *testing.T) {
	client := NewTestClientForDebug()

	args := AddTagsArgs{
		RegionId:       common.Beijing,
		LoadBalancerID: TestLoadBalancerID,
		Tags:           getTestTags(),
	}
	err := client.AddTags(&args)

	if err != nil {
		t.Errorf("Failed to AddTags for instance %s: %v", TestLoadBalancerID, err)
	}

}

func TestDescribeTags(t *testing.T) {
	client := NewTestClientForDebug()

	args := DescribeTagsArgs{
		RegionId: common.Beijing,
		//LoadBalancerID: TestLoadBalancerID,
		Tags: getTestTags(),
	}
	result, _, err := client.DescribeTags(&args)

	if err != nil {
		t.Errorf("Failed to DescribeTags: %v", err)
	} else {
		t.Logf("result: %v", result)
	}
}

func TestRemoveTags(t *testing.T) {

	client := NewTestClientForDebug()

	args := RemoveTagsArgs{
		RegionId:       common.Beijing,
		LoadBalancerID: TestLoadBalancerID,
		Tags:           getTestTags(),
	}
	err := client.RemoveTags(&args)

	if err != nil {
		t.Errorf("Failed to RemoveTags for instance %s: %v", TestInstanceId, err)
	}

}
