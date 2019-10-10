package cs

import (
	"fmt"
	"testing"
	"time"
)

func Test_CreateServelessKubernetesCluster(t *testing.T) {
	client := NewTestClientForDebug()

	tags := make([]Tag, 0)
	tags = append(tags, Tag{
		Key:   "key-a",
		Value: "key-a",
	})

	tags = append(tags, Tag{
		Key:   "key-b",
		Value: "key-b",
	})

	tags = append(tags, Tag{
		Key:   "key-c",
		Value: "key-c",
	})

	args := &ServerlessCreationArgs{
		ClusterType:          ClusterTypeServelessKubernetes,
		Name:                 fmt.Sprintf("serveless-cluster-%d", time.Now().Unix()),
		RegionId:             string(TestRegionID),
		VpcId:                TestVpcId,
		VSwitchId:            TestVSwitchId,
		PrivateZone:          true,
		EndpointPublicAccess: true,
		NatGateway:           true,
		DeletionProtection:   true,
		Tags:                 tags,
	}

	response, err := client.CreateServelessKubernetesCluster(args)
	if err != nil {
		t.Errorf("Error %++v", err)
	} else {
		t.Logf("Response = %++v", response)
	}
}

func Test_DescribeCluster(t *testing.T) {
	client := NewTestClientForDebug()

	cluster, err := client.DescribeServelessKubernetesCluster(TestClusterId)
	if err != nil {
		t.Errorf("Error = %++v", err)
	} else {
		t.Logf("Cluster = %#v", cluster)
		t.Logf("%v ", cluster.Created)
	}
}

func Test_DescribeClusterUserConfig(t *testing.T) {
	client := NewTestClientForDebug()

	config, err := client.DescribeClusterUserConfig(TestClusterId, TestPrivateIpAddress)
	if err != nil {
		t.Errorf("Error = %++v", err)
	} else {
		t.Logf("Config = %#v", config)
	}
}

func Test_DeleteServelessCluster(t *testing.T) {
	client := NewTestClientForDebug()

	err := client.DeleteCluster(TestClusterId)
	if err != nil {
		t.Errorf("Error = %++v", err)
	} else {
		t.Logf("OK")
	}
}

func Test_DescribeServelessKubernetesClusters(t *testing.T) {
	client := NewTestClientForDebug()
	clusters, err := client.DescribeServelessKubernetesClusters()
	if err != nil {
		t.Errorf("Error = %++v", err)
	} else {
		for _, cluster := range clusters {
			t.Logf("Cluster = %#v", cluster)
		}
	}
}
