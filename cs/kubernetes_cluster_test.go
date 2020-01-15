package cs

import (
	"testing"
)

func Test_ModifyCluster(t *testing.T) {
	client := NewTestClientForDebug()

	args := &ModifyClusterArgs{
		DeletionProtection: false,
	}

	err := client.ModifyCluster(TestClusterId, args)
	if err != nil {
		t.Errorf("Error %++v", err)
	} else {
		t.Logf("OK")
	}
}

func Test_UpgradeCluster(t *testing.T) {
	client := NewTestClientForDebug()

	args := &UpgradeClusterArgs{
		Version: "1.14.8-aliyun.1",
	}

	err := client.UpgradeCluster(TestClusterId, args)
	if err != nil {
		t.Errorf("Error %++v", err)
	} else {
		t.Logf("OK")
	}
}

func Test_CancelUpgradeCluster(t *testing.T) {
	client := NewTestClientForDebug()

	err := client.CancelUpgradeCluster(TestClusterId)
	if err != nil {
		t.Errorf("Error %++v", err)
	} else {
		t.Logf("OK")
	}
}

func Test_QueryUpgradeClusterResult(t *testing.T) {
	client := NewTestClientForDebug()

	result, err := client.QueryUpgradeClusterResult(TestClusterId)
	if err != nil {
		t.Errorf("Error %++v", err)
	} else {
		t.Logf("OK, result: %++v", result)
	}
}

func Test_CreateDelicatedKubernetesCluster(t *testing.T) {
	t.SkipNow()
	client := NewTestClientForDebug()

	request := &DelicatedKubernetesClusterCreationRequest{}
	response, err := client.CreateDelicatedKubernetesCluster(request)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Response %++v", response)
	}
}

func Test_CreateManagedKubernetesCluster(t *testing.T) {
	t.SkipNow()
	client := NewTestClientForDebug()

	request := &ManagedKubernetesClusterCreationRequest{}

	response, err := client.CreateManagedKubernetesCluster(request)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Response %++v", response)
	}
}

func Test_DescribeKubernetesClusterDetail(t *testing.T) {
	client := NewTestClientForDebug()

	cluster, err := client.DescribeKubernetesClusterDetail(TestClusterId)

	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Response = %++v", cluster)
		t.Logf("MetaData = %++v", cluster.GetMetaData())
	}
}

func Test_ScaleOutKubernetesCluster(t *testing.T) {
	client := NewTestClientForDebug()

	request := &ScaleOutKubernetesClusterRequest{
		LoginPassword:            "Hello1234",
		WorkerVSwitchIds:         []string{"vsw-xxxx"},
		WorkerInstanceTypes:      []string{"ecs.n4.xlarge"},
		WorkerInstanceChargeType: "PostPaid",
		WorkerPeriod:             1,
		WorkerPeriodUnit:         "1",
		WorkerAutoRenew:          true,
		WorkerAutoRenewPeriod:    1,
		WorkerDataDisk:           true,
		WorkerDataDisks: []DataDisk{
			{
				Category: "cloud_ssd",
				Size:     "200",
			},
			{
				Category: "cloud_ssd",
				Size:     "300",
			},
		},
		Tags: []Tag{
			{Key: "k-aaa",
				Value: "v-aaa",
			},
		},
		Count: 2,
	}

	response, err := client.ScaleOutKubernetesCluster(TestClusterId, request)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Response %++v", response)
	}
}

func Test_DeleteKubernetesClusterNodes(t *testing.T) {
	client := NewTestClientForDebug()

	request := &DeleteKubernetesClusterNodesRequest{
		ReleaseNode: false,
		Nodes:       []string{"cn-beijing.192.168.0.128"},
	}

	response, err := client.DeleteKubernetesClusterNodes(TestClusterId, request)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Response %++v", response)
	}
}
