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
		if cluster.ClusterType != "Kubernetes" && cluster.ClusterType != "ManagedKubernetes" {
			continue
		}
		t.Logf("Cluster: %++v", cluster)
		c, err := client.DescribeKubernetesCluster(cluster.ClusterID)
		if err != nil {
			t.Errorf("Failed to DescribeCluster: %v", err)
		}
		t.Logf("Cluster Describe: %++v", c)

		if c.MetaData.MultiAZ || c.MetaData.SubClass == "3az" {
			t.Logf("%v is a MultiAZ kubernetes cluster", c.ClusterID)
		} else {
			if cluster.ClusterType == "ManagedKubernetes" {
				t.Logf("%v is a Managed kubernetes cluster", c.ClusterID)
			} else {
				t.Logf("%v is a SingleAZ kubernetes cluster", c.ClusterID)
			}
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

func _TestGetClusterEndpoints(t *testing.T) {
	client := NewTestClientForDebug()
	clusterId := "c213c31b97430433c87afe4852b6a08ef"
	clusterEndpoints, err := client.GetClusterEndpoints(clusterId)
	if err != nil {
		t.Fatalf("Failed to GetClusterEndpoints: %v", err)
	}
	t.Logf("Succeed getting clusterEndpoints %v", clusterEndpoints)
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

func _TestCreateManagedKubernetesCluster(t *testing.T) {

	client := NewTestClientForDebug()

	args := KubernetesCreationArgs{
		Name:            "single-managed-az-k8s",
		ClusterType:     "ManagedKubernetes",
		DisableRollback: true,
		//VPCID:                    "vpc-id",
		//VSwitchId:                "vsw-id",
		ZoneId:                   "cn-hangzhou-g",
		SNatEntry:                true,
		NumOfNodes:               2,
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
		Name:            "multiaz-test",
		ClusterType:     "Kubernetes",
		DisableRollback: true,
		MultiAZ:         true,
		VPCID:           "vpc-id",
		VSwitchIdA:      "vsw-id1",
		VSwitchIdB:      "vsw-id2",
		VSwitchIdC:      "vsw-id3",
		NumOfNodesA:     1,
		NumOfNodesB:     1,
		NumOfNodesC:     1,

		MasterInstanceTypeA:      "ecs.ic5.xlarge",
		MasterInstanceTypeB:      "ecs.ic5.xlarge",
		MasterInstanceTypeC:      "ecs.ic5.xlarge",
		MasterSystemDiskCategory: "cloud_efficiency",
		MasterSystemDiskSize:     40,

		MasterInstanceChargeType: "PrePaid",
		MasterPeriodUnit:         "Week",
		MasterPeriod:             1,
		MasterAutoRenew:          true,
		MasterAutoRenewPeriod:    1,

		WorkerInstanceTypeA:      "ecs.ic5.xlarge",
		WorkerInstanceTypeB:      "ecs.ic5.xlarge",
		WorkerInstanceTypeC:      "ecs.ic5.xlarge",
		WorkerSystemDiskCategory: "cloud_efficiency",
		WorkerSystemDiskSize:     40,

		WorkerInstanceChargeType: "PrePaid",
		WorkerPeriodUnit:         "Week",
		WorkerPeriod:             1,
		WorkerAutoRenew:          true,
		WorkerAutoRenewPeriod:    1,

		SSHFlags:               true,
		ContainerCIDR:          "172.20.0.0/16",
		ServiceCIDR:            "172.21.0.0/20",
		LoginPassword:          "test-password123",
		ImageId:                "centos_7_03_64_20G_alibase_20170818.vhd",
		KubernetesVersion:      "1.11.2",
		WorkerDataDisk:         true,
		WorkerDataDiskCategory: "cloud_efficiency",
		WorkerDataDiskSize:     100,
		NodeCIDRMask:           "25",
		LoggingType:            "SLS",
		SLSProjectName:         "k8s-test-my-terraform",
	}
	cluster, err := client.CreateKubernetesMultiAZCluster(common.Hangzhou, &args)
	if err != nil {
		t.Fatalf("Failed to CreateKubernetesMultiAZCluster: %v", err)
	}

	t.Logf("Cluster: %++v", cluster)
}

func TestScaleKubernetesCluster(t *testing.T) {
	// Comment below to test
	//t.SkipNow()

	client := NewTestClientForDebug()

	args := KubernetesClusterScaleArgs{
		LoginPassword:            "test-password123",
		WorkerInstanceTypes:      []string{"ecs.sn1ne.large"},
		WorkerSystemDiskCategory: "cloud_ssd",
		WorkerSystemDiskSize:     int64(40),
		Count:                    2,
		WorkerDataDisk:           true,
		WorkerDataDiskCategory:   "cloud_ssd",
		WorkerDataDiskSize:       int64(200),
	}

	err := client.ScaleKubernetesCluster(TestClusterId, &args)
	if err != nil {
		t.Fatalf("Failed to TestScaleKubernetesCluster: %v", err)
	}
}
