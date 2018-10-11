package cs

import (
	"testing"

	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/ecs"
)

func TestClient_DescribeClusters(t *testing.T) {

	client := NewTestDebugAussumeRoleClient()

	clusters, err := client.DescribeClusters("")
	if err != nil {
		t.Fatalf("Failed to DescribeClusters: %v", err)
	}

	for _, cluster := range clusters {
		t.Logf("Cluster: %++v", cluster)
		c, err := client.DescribeCluster(cluster.ClusterID)
		if err != nil {
			t.Errorf("Failed to DescribeCluster: %v", err)
		}
		t.Logf("Cluster Describe: %++v", c)
		certs, err := client.GetClusterCerts(cluster.ClusterID)
		if err != nil {
			t.Errorf("Failed to GetClusterCerts: %v", err)
		}
		t.Logf("Cluster certs: %++v", certs)

	}
}

func TestClient_DescribeKubernetesClusters(t *testing.T) {

	client := NewTestDebugAussumeRoleClient()

	clusters, err := client.DescribeClusters("")
	if err != nil {
		t.Fatalf("Failed to DescribeClusters: %v", err)
	}

	for _, cluster := range clusters {
		if cluster.ClusterType != "Kubernetes" {
			continue
		}
		t.Logf("Cluster: %++v", cluster)
		c, err := client.DescribeKubernetesCluster(cluster.ClusterID)
		if err != nil {
			t.Errorf("Failed to DescribeCluster: %v", err)
		}
		t.Logf("Cluster Describe: %++v", c)
		t.Logf("Cluster KeyPair %v", c.Parameters.KeyPair)
		t.Logf("Cluster RawWorkerDataDisk %v", c.Parameters.RawWorkerDataDisk)
		t.Logf("Cluster WorkerDataDisk %v", c.Parameters.WorkerDataDisk)
		if c.Parameters.WorkerDataDisk {
			t.Logf("Cluster WorkerDataDiskSize %v", c.Parameters.WorkerDataDiskSize)
			t.Logf("Cluster WorkerDataDiskCategory %v", c.Parameters.WorkerDataDiskCategory)
		}
		t.Logf("Cluster NodeCIDRMask %v", c.Parameters.NodeCIDRMask)
		t.Logf("Cluster LoggingType %v", c.Parameters.LoggingType)
		t.Logf("Cluster SLSProjectName %v", c.Parameters.SLSProjectName)
		t.Logf("Cluster PublicSLB %v", c.Parameters.PublicSLB)

		if c.MetaData.MultiAZ || c.MetaData.SubClass == "3az" {
			t.Logf("%v is a MultiAZ kubernetes cluster", c.ClusterID)
			t.Logf("Cluster VSWA ID %v", c.Parameters.VSwitchIdA)
			t.Logf("Cluster VSWB ID %v", c.Parameters.VSwitchIdB)
			t.Logf("Cluster VSWC ID %v", c.Parameters.VSwitchIdC)
			t.Logf("Cluster MasterInstanceTypeA %v", c.Parameters.MasterInstanceTypeA)
			t.Logf("Cluster MasterInstanceTypeB %v", c.Parameters.MasterInstanceTypeB)
			t.Logf("Cluster MasterInstanceTypeC %v", c.Parameters.MasterInstanceTypeC)
			t.Logf("Cluster NumOfNodeA %v", c.Parameters.NumOfNodesA)
			t.Logf("Cluster NumOfNodeB %v", c.Parameters.NumOfNodesB)
			t.Logf("Cluster NumOfNodeC %v", c.Parameters.NumOfNodesC)
		} else {
			t.Logf("%v is a single kubernetes cluster", c.ClusterID)
			t.Logf("Cluster VSW ID %v", c.Parameters.VSwitchID)
			t.Logf("Cluster MasterInstanceType %v", c.Parameters.MasterInstanceType)
			t.Logf("Cluster NumOfNode %v", c.Parameters.NumOfNodes)
		}
	}
}

func TestListClusters(t *testing.T) {

	client := NewTestClientForDebug()

	clusters, err := client.DescribeClusters("")
	if err != nil {
		t.Fatalf("Failed to DescribeClusters: %v", err)
	}

	for _, cluster := range clusters {
		t.Logf("Cluster: %++v", cluster)
		c, err := client.DescribeCluster(cluster.ClusterID)
		if err != nil {
			t.Errorf("Failed to DescribeCluster: %v", err)
		}
		t.Logf("Cluster Describe: %++v", c)
		certs, err := client.GetClusterCerts(cluster.ClusterID)
		if err != nil {
			t.Errorf("Failed to GetClusterCerts: %v", err)
		}
		t.Logf("Cluster certs: %++v", certs)

	}
}

func _TestCreateClusters(t *testing.T) {

	client := NewTestClientForDebug()

	args := ClusterCreationArgs{
		Name:             "test",
		Size:             1,
		NetworkMode:      ClassicNetwork,
		DataDiskCategory: ecs.DiskCategoryCloud,
		InstanceType:     "ecs.s2.small",
		Password:         "just$test",
	}
	cluster, err := client.CreateCluster(common.Beijing, &args)
	if err != nil {
		t.Fatalf("Failed to CreateCluster: %v", err)
	}

	t.Logf("Cluster: %++v", cluster)
}

func _TestDeleteClusters(t *testing.T) {

	client := NewTestClientForDebug()
	clusterId := "c14601b7676204f73b838329685704902"
	err := client.DeleteCluster(clusterId)
	if err != nil {
		t.Fatalf("Failed to CreateCluster: %v", err)
	}
	t.Logf("Cluster %s is deleting", clusterId)
}

func _TestCreateKubernetesCluster(t *testing.T) {

	client := NewTestClientForDebug()

	args := KubernetesCreationArgs{
		Name:            "single-az-k8s",
		ClusterType:     "Kubernetes",
		DisableRollback: true,
		//VPCID:                    "vpc-id",
		//VSwitchId:                "vsw-id",
		ZoneId:                   "cn-hangzhou-g",
		SNatEntry:                true,
		NumOfNodes:               1,
		MasterInstanceType:       "ecs.sn1ne.large",
		MasterSystemDiskCategory: "cloud_efficiency",
		MasterSystemDiskSize:     40,
		WorkerInstanceType:       "ecs.sn1ne.large",
		WorkerSystemDiskCategory: "cloud_efficiency",
		WorkerSystemDiskSize:     40,
		SSHFlags:                 true,
		ContainerCIDR:            "172.16.0.0/16",
		ServiceCIDR:              "172.19.0.0/20",
		LoginPassword:            "test-password123",
		WorkerDataDisk:           true,
		WorkerDataDiskCategory:   "cloud_efficiency",
		WorkerDataDiskSize:       100,
		PublicSLB:                true,
		NodeCIDRMask:             "25",
		LoggingType:              "SLS",
		SLSProjectName:           "k8s-test-my-terraform-singleaz",
	}
	cluster, err := client.CreateKubernetesCluster(common.Hangzhou, &args)
	if err != nil {
		t.Fatalf("Failed to CreateKubernetesCluster: %v", err)
	}

	t.Logf("Cluster: %++v", cluster)
}

func _TestCreateKubernetesMultiAZCluster(t *testing.T) {

	client := NewTestClientForDebug()

	args := KubernetesMultiAZCreationArgs{
		Name:                     "multiaz-test1",
		ClusterType:              "Kubernetes",
		DisableRollback:          true,
		MultiAZ:                  true,
		VPCID:                    "vpc-id",
		VSwitchIdA:               "vsw-id1",
		VSwitchIdB:               "vsw-id2",
		VSwitchIdC:               "vsw-id3",
		NumOfNodesA:              1,
		NumOfNodesB:              1,
		NumOfNodesC:              1,
		MasterInstanceTypeA:      "ecs.sn1ne.large",
		MasterInstanceTypeB:      "ecs.sn1ne.large",
		MasterInstanceTypeC:      "ecs.sn1ne.large",
		MasterSystemDiskCategory: "cloud_efficiency",
		MasterSystemDiskSize:     40,
		WorkerInstanceTypeA:      "ecs.sn1ne.large",
		WorkerInstanceTypeB:      "ecs.sn1ne.large",
		WorkerInstanceTypeC:      "ecs.sn1ne.large",
		WorkerSystemDiskCategory: "cloud_efficiency",
		WorkerSystemDiskSize:     40,
		SSHFlags:                 true,
		ContainerCIDR:            "172.17.0.0/16",
		ServiceCIDR:              "172.20.0.0/20",
		LoginPassword:            "test-password123",
		WorkerDataDisk:           true,
		WorkerDataDiskCategory:   "cloud_efficiency",
		WorkerDataDiskSize:       100,
		NodeCIDRMask:             "25",
		LoggingType:              "SLS",
		SLSProjectName:           "k8s-test-my-terraform",
	}
	cluster, err := client.CreateKubernetesMultiAZCluster(common.Hangzhou, &args)
	if err != nil {
		t.Fatalf("Failed to CreateKubernetesMultiAZCluster: %v", err)
	}

	t.Logf("Cluster: %++v", cluster)
}

func _TestResizeKubernetesCluster(t *testing.T) {
	client := NewTestClientForDebug()

	args := KubernetesClusterResizeArgs{
		DisableRollback:    true,
		TimeoutMins:        60,
		LoginPassword:      "test-password123",
		WorkerInstanceType: "ecs.sn1ne.large",
		NumOfNodes:         2,
	}

	err := client.ResizeKubernetesCluster("c3419330390a94906bcacc55bfa9dd21f", &args)
	if err != nil {
		t.Fatalf("Failed to TestResizeKubernetesCluster: %v", err)
	}
}
