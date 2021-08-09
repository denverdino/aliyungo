package slb

import (
	"testing"

	"github.com/denverdino/aliyungo/common"
)

func TestLoadBalancer(t *testing.T) {

	client := NewTestClientForDebug()

	creationArgs := CreateLoadBalancerArgs{
		RegionId:         common.Beijing,
		LoadBalancerName: "test-slb",
		LoadBalancerSpec: S2Medium, // eni not support slb.s0.share slb(default slb.s0.share)
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

func TestLoadBalancerIPv6(t *testing.T) {

	client := NewTestClientForDebug()

	creationArgs := CreateLoadBalancerArgs{
		RegionId:         common.Hangzhou,
		LoadBalancerName: "test-slb-ipv6",
		AddressType:      InternetAddressType,
		MasterZoneId:     "cn-hangzhou-e",
		SlaveZoneId:      "cn-hangzhou-f",
		ClientToken:      client.GenerateClientToken(),
		AddressIPVersion: IPv6,
	}

	response, err := client.CreateLoadBalancer(&creationArgs)
	if err != nil {
		t.Fatalf("Failed to CreateLoadBalancer: %v", err)
	}

	t.Logf("CreateLoadBalancer result: %v", *response)
	lbId := response.LoadBalancerId

	describeLoadBalancersArgs := DescribeLoadBalancersArgs{
		RegionId: common.Hangzhou,
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

func TestClient_DescribeLoadBalancers(t *testing.T) {
	client := NewTestNewSLBClientForDebug()
	//client.SetSecurityToken(TestSecurityToken)

	args := &DescribeLoadBalancersArgs{
		RegionId: TestRegionID,
		//SecurityToken: TestSecurityToken,
	}

	slbs, err := client.DescribeLoadBalancers(args)
	if err != nil {
		t.Fatalf("Failed %++v", err)
	} else {
		t.Logf("Result = %++v", slbs)
	}
}

func TestClient_SetLoadBalancerDeleteProtection(t *testing.T) {
	client := NewTestNewSLBClientForDebug()

	creationArgs := CreateLoadBalancerArgs{
		RegionId:         common.Beijing,
		LoadBalancerName: "test-slb",
		LoadBalancerSpec: S2Medium,
		AddressType:      InternetAddressType,
		ClientToken:      client.GenerateClientToken(),
	}

	response, err := client.CreateLoadBalancer(&creationArgs)
	if err != nil {
		t.Fatalf("Failed to CreateLoadBalancer: %v", err)
	}

	t.Logf("CreateLoadBalancer result: %v", *response)
	lbId := response.LoadBalancerId

	args := &SetLoadBalancerDeleteProtectionArgs{
		LoadBalancerId:   lbId,
		DeleteProtection: OnFlag,
		RegionId:         common.Beijing,
	}

	err = client.SetLoadBalancerDeleteProtection(args)
	if err != nil {
		t.Fatalf("Failed %++v", err)
	}
	t.Logf("SetLoadBalancerDeleteProtection result: %v", *response)

	err = client.DeleteLoadBalancer(lbId)
	if err != nil {
		t.Logf("DeleteLoadBalancer result: %++v", err)
	} else {
		t.Fatalf("Failed to set LoadBalancer delete protection.")
	}
}

func TestClient_SetLoadBalancerModificationProtection(t *testing.T) {
	client := NewTestNewSLBClientForDebug()

	creationArgs := CreateLoadBalancerArgs{
		RegionId:                     common.Beijing,
		LoadBalancerName:             "test-slb-modification-protection",
		LoadBalancerSpec:             S1Small,
		AddressType:                  InternetAddressType,
		ModificationProtectionStatus: ConsoleProtection,
		ModificationProtectionReason: "kubernetes.do.not.delete",
		ClientToken:                  client.GenerateClientToken(),
	}
	response, err := client.CreateLoadBalancer(&creationArgs)
	if err != nil {
		t.Fatalf("Failed to CreateLoadBalancer: %v", err)
	}

	t.Logf("CreateLoadBalancer result: %v", *response)
	lbId := response.LoadBalancerId

	lb, err := client.DescribeLoadBalancerAttribute(lbId)
	if err != nil {
		t.Fatalf("Failed to DescribeLoadBalancerAttribute: %v", err)
	}

	if lb.ModificationProtectionStatus != ConsoleProtection {
		t.Fatalf("Failed to SetLoadBalancerModificationProtection, slb %s, expected %s got %s",
			lbId, ConsoleProtection, lb.ModificationProtectionStatus)
	}

	args := SetLoadBalancerModificationProtectionArgs{
		RegionId:                     common.Beijing,
		LoadBalancerId:               lbId,
		ModificationProtectionStatus: NonProtection,
	}

	err = client.SetLoadBalancerModificationProtection(&args)
	if err != nil {
		t.Fatalf("Failed to SetLoadBalancerModificationProtection: %v", err)
	}

	lb, err = client.DescribeLoadBalancerAttribute(lbId)
	if err != nil {
		t.Fatalf("Failed to DescribeLoadBalancerAttribute: %v", err)
	}

	if lb.ModificationProtectionStatus != NonProtection {
		t.Fatalf("Failed to SetLoadBalancerModificationProtection, slb %s, expected %s got %s",
			lbId, ConsoleProtection, lb.ModificationProtectionStatus)
	}

	// Delete Slb
	err = client.DeleteLoadBalancer(lbId)
	if err != nil {
		t.Fatalf("Failed to DeleteLoadBalancer: %v", err)
	}

}

func TestClient_ServiceManagedControl(t *testing.T) {
	client := NewTestNewSLBClientForDebug()

	creationArgs := CreateLoadBalancerArgs{
		RegionId:                     common.Beijing,
		LoadBalancerName:             "test-slb-modification-protection",
		LoadBalancerSpec:             S1Small,
		AddressType:                  InternetAddressType,
		ModificationProtectionStatus: ConsoleProtection,
		ModificationProtectionReason: "kubernetes.do.not.delete",
		ClientToken:                  client.GenerateClientToken(),
	}
	response, err := client.CreateLoadBalancer(&creationArgs)
	if err != nil {
		t.Fatalf("Failed to CreateLoadBalancer: %v", err)
	}

	t.Logf("CreateLoadBalancer result: %v", *response)
	lbId := response.LoadBalancerId

	lb, err := client.DescribeLoadBalancerAttribute(lbId)
	if err != nil {
		t.Fatalf("Failed to DescribeLoadBalancerAttribute: %v", err)
	}

	resource := make([]ManagedResourceModel,0)
	resource = append(resource, ManagedResourceModel{ResourceId: lb.LoadBalancerId})
	args := ServiceManagedControlArgs{
		RegionId:			common.Beijing,
		ServiceManagedMode: Managed,
		ResourceType: 		ManagedLoadBalancer,
		Resources: 			resource,
	}

	err = client.ServiceManagedControl(&args)
	if err != nil {
		t.Fatalf("Failed to modify resource managed status: %v", err)
	}

	lb, err = client.DescribeLoadBalancerAttribute(lbId)
	if err != nil {
		t.Fatalf("Failed to DescribeLoadBalancerAttribute: %v", err)
	}

	if lb.ServiceManagedMode != Managed {
		t.Fatalf("Failed to modify resource managed status, slb %s, expected %s got %s",
			lbId, Managed, lb.ServiceManagedMode)
	}

	// Delete Slb
	err = client.DeleteLoadBalancer(lbId)
	if err != nil {
		t.Fatalf("Failed to DeleteLoadBalancer: %v", err)
	}

}
