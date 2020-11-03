package cs

import (
	"os"
	"strings"
	"testing"
)

var (
	VpcId         = os.Getenv("VpcId")
	VswitchIds    = os.Getenv("VswitchIds")
	LoginPassword = os.Getenv("LoginPassword")
	NodePoolId    = os.Getenv("NodePoolId")
)

func Test_CreateNodePool(t *testing.T) {
	client := NewTestClientForDebug()

	args := &CreateNodePoolRequest{
		Count:    1,
		RegionId: TestRegionID,
		NodePoolInfo: NodePoolInfo{
			Name:         "test-npl",
			NodePoolType: "ess",
		},
		ScalingGroup: ScalingGroup{
			VpcId:              VpcId,
			VswitchIds:         strings.Split(VswitchIds, ","),
			InstanceTypes:      []string{"ecs.n6.large"},
			LoginPassword:      LoginPassword,
			SystemDiskCategory: "cloud_efficiency",
			SystemDiskSize:     120,
			DataDisks:          []NodePoolDataDisk{{Size: 100, Category: "cloud_ssd"}},
		},
		KubernetesConfig: KubernetesConfig{
			NodeNameMode: "customized,aliyun.com,5,test",
		},
	}

	resp, err := client.CreateNodePool(args, TestClusterId)
	if err != nil {
		t.Errorf("Error %++v", err)
	} else {
		t.Logf("response: %++v", resp)
	}
}

func Test_DescribeNodePoolDetail(t *testing.T) {
	client := NewTestClientForDebug()

	resp, err := client.DescribeNodePoolDetail(TestClusterId, NodePoolId)
	if err != nil {
		t.Errorf("Error %++v", err)
	} else {
		t.Logf("response: %++v", resp)
	}
}

func Test_UpdateNodePool(t *testing.T) {
	client := NewTestClientForDebug()

	args := &UpdateNodePoolRequest{
		Count:    2,
		RegionId: TestRegionID,
		ScalingGroup: ScalingGroup{
			InstanceTypes: []string{"ecs.n4.large"},
		},
	}

	resp, err := client.UpdateNodePool(TestClusterId, NodePoolId, args)
	if err != nil {
		t.Errorf("Error %++v", err)
	} else {
		t.Logf("response: %++v", resp)
	}
}

func Test_DeleteNodePool(t *testing.T) {
	client := NewTestClientForDebug()

	err := client.DeleteNodePool(TestClusterId, NodePoolId)
	if err != nil {
		t.Errorf("Error %++v", err)
	} else {
		t.Logf("success")
	}
}

func Test_DescribeClusterNodePools(t *testing.T) {
	client := NewTestClientForDebug()

	resp, err := client.DescribeClusterNodePools(TestClusterId)
	if err != nil {
		t.Errorf("Error %++v", err)
	} else {
		t.Logf("response: %++v", resp)
	}
}
