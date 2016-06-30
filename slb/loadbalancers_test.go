package slb

import (
	"testing"

	"github.com/denverdino/aliyungo/common"
)

func TestLoadBlancer(t *testing.T) {

	client := NewTestClientForDebug()

	creationArgs := CreateLoadBalancerArgs{
		RegionId:         common.Beijing,
		LoadBalancerName: "test-slb",
		AddressType:      InternetAddressType,
		ClientToken:      client.GenerateClientToken(),
	}

	response, err := client.CreateLoadBalancer(&creationArgs)
	if err != nil {
		t.Fatalf("Failed to CreateLoadBalancer: %v", err)
	}

	t.Logf("CreateLoadBalancer result: %v", *response)
	lbId := response.LoadBalancerId

	testBackendServers(t, client, lbId)
	testListeners(t, client, lbId)

	describeLoadBalancersArgs := DescribeLoadBalancersArgs{
		RegionId: common.Beijing,
	}

	loadBalancers, err := client.DescribeLoadBalancers(&describeLoadBalancersArgs)

	if err != nil {
		t.Fatalf("Failed to DescribeLoadBalancers: %v", err)
	}
	t.Logf("DescribeLoadBalancers result: %++v", loadBalancers)

	err = client.SetLoadBalancerStatus(lbId, InactiveStatus)
	if err != nil {
		t.Fatalf("Failed to SetLoadBalancerStatus: %v", err)
	}
	err = client.SetLoadBalancerName(lbId, "test-slb2")
	if err != nil {
		t.Fatalf("Failed to SetLoadBalancerName: %v", err)
	}
	loadBalancer, err := client.DescribeLoadBalancerAttribute(lbId)

	if err != nil {
		t.Fatalf("Failed to DescribeLoadBalancerAttribute: %v", err)
	}
	t.Logf("DescribeLoadBalancerAttribute result: %++v", loadBalancer)

	err = client.DeleteLoadBalancer(lbId)
	if err != nil {
		t.Errorf("Failed to DeleteLoadBalancer: %v", err)
	}

	t.Logf("DeleteLoadBalancer successfully: %s", lbId)

}

//create slb
func TestCreateLoadBalancer(t *testing.T) {
	client := NewTestBIDClientForDebug()

	args := CreateLoadBalancerArgs{
		RegionId:         common.Beijing,
		LoadBalancerName: "test-slb",
		AddressType:      InternetAddressType,
		ClientToken:      client.GenerateClientToken(),
	}

	response, err := client.CreateLoadBalancer(&args)
	if err != nil {
		t.Fatalf("Failed to CreateLoadBalancer: %v", err)
	} else {
		t.Logf("CreateLoadBalancer result: %v", *response)
	}
}

func TestDescribeLoadBalancers(t *testing.T) {
	client := NewTestBIDClientForDebug()

	args := DescribeLoadBalancersArgs{
		RegionId: common.Beijing,
	}

	lbs, err := client.DescribeLoadBalancers(&args)
	if err != nil {
		t.Fatalf("Failed to DescribeLoadBalancers: %v", err)
	} else {
		t.Logf("LoadBalancers %v is described successfully.", lbs)
	}
}
