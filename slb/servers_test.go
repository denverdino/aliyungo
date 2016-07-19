package slb

import (
	"testing"
)

const LOAD_BALANCER_ID = "lb2ze5338m19366z0sa8"

func TestServers(t *testing.T) {

	client := NewTestClientForDebug()
	loadBalancerId := LOAD_BALANCER_ID
	TestInstanceId := ""

	backendServers := []BackendServerType{
		BackendServerType{
			ServerId: TestInstanceId,
			Weight:   100,
		},
	}

	//AddBackendServers
	servers, err := client.AddBackendServers(loadBalancerId, backendServers)
	if err != nil {
		t.Errorf("Failed to AddBackendServers: %v", err)
	}
	t.Logf("Backend servers: %++v", servers)

	/*
		//SetBackendServers
		servers, err = client.SetBackendServers(loadBalancerId, backendServers)
		if err != nil {
			t.Errorf("Failed to SetBackendServers: %v", err)
		}
		t.Logf("Backend servers: %++v", servers)
	*/

	//DescribeHealthStatus
	args := DescribeHealthStatusArgs{
		LoadBalancerId: LOAD_BALANCER_ID,
		ListenerPort:   1234,
	}
	describeHealthStatusResponse, describeHealthStatusErr := client.DescribeHealthStatus(&args)
	if describeHealthStatusErr != nil {
		t.Errorf("Failed to DescribeHealthStatus: %v", describeHealthStatusErr)
	}
	t.Logf("Backend servers: %++v", describeHealthStatusResponse)

	//RemoveBackendServers
	servers, err = client.RemoveBackendServers(loadBalancerId, []string{TestInstanceId})
	if err != nil {
		t.Errorf("Failed to RemoveBackendServers: %v", err)
	}
	t.Logf("Backend servers: %++v", servers)

}
