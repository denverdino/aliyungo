package slb

import (
	"testing"
)

const LOAD_BALANCER_ID = "lb2ze5338m19366z0sa8"
const PORT = 1022
const NEW_PORT = 1090

func TestListeners(t *testing.T) {

	client := NewTestClientForDebug()
	loadBalancerId := LOAD_BALANCER_ID
	port := PORT

	/*
		//CreateLoadBalancerHTTPListener 通过
		createLoadBalancerHTTPListener := CreateLoadBalancerHTTPListenerArgs{
			LoadBalancerId:    LOAD_BALANCER_ID,
			ListenerPort:      NEW_PORT,
			BackendServerPort: NEW_PORT,
			Bandwidth:         -1,
			StickySession:     OffFlag,
			HealthCheck:       OffFlag,
		}
		err = client.CreateLoadBalancerHTTPListener(&createLoadBalancerHTTPListener)
		if err != nil {
			t.Errorf("Failed to CreateLoadBalancerHTTPListener: %v", err)
		}
		//CreateLoadBalancerHTTPSListener 通过
		createLoadBalancerHTTPSListener := CreateLoadBalancerHTTPSListenerArgs{
			LoadBalancerId:      LOAD_BALANCER_ID,
			ListenerPort:        NEW_PORT,
			BackendServerPort:   NEW_PORT,
			Bandwidth:           -1,
			StickySession:       OffFlag,
			HealthCheck:         OffFlag,
			ServerCertificateId: "1553339444347287_155fbffe45c",
		}

		err = client.CreateLoadBalancerHTTPSListener(&createLoadBalancerHTTPSListener)
		if err != nil {
			t.Errorf("Failed to CreateLoadBalancerHTTPSListener: %v", err)
		}

		//CreateLoadBalancerTCPListener 通过
		createLoadBalancerTCPListener := CreateLoadBalancerTCPListenerArgs{
			LoadBalancerId:    LOAD_BALANCER_ID,
			ListenerPort:      PORT,
			BackendServerPort: PORT,
			Bandwidth:         -1,
		}
		err = client.CreateLoadBalancerTCPListener(&createLoadBalancerTCPListener)
		if err != nil {
			t.Errorf("Failed to CreateLoadBalancerTCPListener: %v", err)
		}


		//CreateLoadBalancerUDPListener 通过
		createLoadBalancerUDPListener := CreateLoadBalancerUDPListenerArgs{
			LoadBalancerId:    LOAD_BALANCER_ID,
			ListenerPort:      PORT,
			BackendServerPort: PORT,
			Bandwidth:         -1,
		}
		err = client.CreateLoadBalancerUDPListener(&createLoadBalancerUDPListener)
		if err != nil {
			t.Errorf("Failed to CreateLoadBalancerUDPListener: %v", err)
		}

		//DeleteLoadBalancerListener 通过
		err = client.DeleteLoadBalancerListener(loadBalancerId, port)

		if err != nil {
			t.Errorf("Failed to DeleteLoadBalancerListener: %v", err)
		}
		t.Logf("Listener: nil")

		//SetListenerAccessControlStatus
		err = client.SetListenerAccessControlStatus(loadBalancerId, port, Close)
		if err != nil {
			t.Errorf("Failed to SetListenerAccessControlStatus: %v", err)
		}

		//AddListenerWhiteListItem
		err = client.AddListenerWhiteListItem(loadBalancerId, port, "1.1.1.1,1.1.1.0/21")
		if err != nil {
			t.Errorf("Failed to AddListenerWhiteListItem: %v", err)
		}

		//RemoveListenerWhiteListItem
		err = client.RemoveListenerWhiteListItem(loadBalancerId, port, "1.1.1.0/21")
		if err != nil {
			t.Errorf("Failed to RemoveListenerWhiteListItem: %v", err)
		}

		//StartLoadBalancerListener
		err = client.StartLoadBalancerListener(loadBalancerId, port)
		if err != nil {
			t.Errorf("Failed to StartLoadBalancerListener: %v", err)
		}

		//StopLoadBalancerListener
		err = client.StopLoadBalancerListener(loadBalancerId, port)
		if err != nil {
			t.Errorf("Failed to StopLoadBalancerListener: %v", err)
		}

		//SetLoadBalancerHTTPListenerAttribute 通过
		setLoadBalancerHTTPListenerAttribute := SetLoadBalancerHTTPListenerAttributeArgs{
			LoadBalancerId:    LOAD_BALANCER_ID,
			ListenerPort:      NEW_PORT,
			BackendServerPort: NEW_PORT,
			Bandwidth:         -1,
			StickySession:     OffFlag,
			HealthCheck:       OffFlag,
		}

		err = client.SetLoadBalancerHTTPListenerAttribute(&setLoadBalancerHTTPListenerAttribute)
		if err != nil {
			t.Errorf("Failed to SetLoadBalancerHTTPListenerAttribute: %v", err)
		}

		//CreateLoadBalancerHTTPSListener 通过
		args := CreateLoadBalancerHTTPSListenerArgs{
			LoadBalancerId:      LOAD_BALANCER_ID,
			ListenerPort:        NEW_PORT,
			BackendServerPort:   NEW_PORT,
			Bandwidth:           -1,
			StickySession:       OffFlag,
			HealthCheck:         OffFlag,
			ServerCertificateId: "1553339444347287_155fbffe45c",
		}
		err = client.CreateLoadBalancerHTTPSListener(&args)
		if err != nil {
			t.Errorf("Failed to CreateLoadBalancerHTTPSListener: %v", err)
		}

		//SetLoadBalancerHTTPSListenerAttribute 通过
		setLoadBalancerHTTPSListenerAttribute := SetLoadBalancerHTTPSListenerAttributeArgs{
			LoadBalancerId:    LOAD_BALANCER_ID,
			ListenerPort:      NEW_PORT,
			BackendServerPort: NEW_PORT,
			Bandwidth:         -1,
			StickySession:     OffFlag,
			HealthCheck:       OffFlag,
		}

		err = client.SetLoadBalancerHTTPSListenerAttribute(&setLoadBalancerHTTPSListenerAttribute)
		if err != nil {
			t.Errorf("Failed to setLoadBalancerHTTPSListenerAttribute: %v", err)
		}

		//SetLoadBalancerTCPListenerAttribute 通过
		setLoadBalancerTCPListenerAttribute := SetLoadBalancerTCPListenerAttribute{
			LoadBalancerId:    LOAD_BALANCER_ID,
			ListenerPort:      NEW_PORT,
			BackendServerPort: NEW_PORT,
			Bandwidth:         -1,
			StickySession:     OffFlag,
			HealthCheck:       OffFlag,
		}

		err = client.SetLoadBalancerTCPListenerAttribute(&setLoadBalancerTCPListenerAttribute)
		if err != nil {
			t.Errorf("Failed to SetLoadBalancerTCPListenerAttribute: %v", err)
		}
	*/

	//DescribeLoadBalancerTCPListenerAttribute
	response, err := client.DescribeLoadBalancerTCPListenerAttribute(loadBalancerId, 1022)
	if err != nil {
		t.Errorf("Failed to DescribeLoadBalancerTCPListenerAttribute: %v", err)
	}
	t.Logf("Listener: %++v", *response)

	//DescribeLoadBalancerUDPListenerAttribute
	args := DescribeLoadBalancerUDPListenerAttributeArgs{
		LoadBalancerId: loadBalancerId,
		ListenerPort:   1021,
		VServerGroupId: "",
	}
	response1, err := client.DescribeLoadBalancerUDPListenerAttribute(&args)
	if err != nil {
		t.Errorf("Failed to DescribeLoadBalancerUDPListenerAttribute: %v", err)
	}
	t.Logf("Listener: %++v", *response1)

	//DescribeLoadBalancerHTTPListenerAttribute
	response2, err := client.DescribeLoadBalancerHTTPListenerAttribute(loadBalancerId, 1025)
	if err != nil {
		t.Errorf("Failed to DescribeLoadBalancerHTTPListenerAttribute: %v", err)
	}
	t.Logf("Listener: %++v", *response2)

	//DescribeLoadBalancerHTTPSListenerAttribute
	response3, err := client.DescribeLoadBalancerHTTPSListenerAttribute(loadBalancerId, 1028)
	if err != nil {
		t.Errorf("Failed to DescribeLoadBalancerHTTPSListenerAttribute: %v", err)
	}
	t.Logf("Listener: %++v", *response3)

	//WaitForListener
	status, err := client.WaitForListener(loadBalancerId, port, TCP)
	if err != nil {
		t.Errorf("Failed to WaitForListener: %v", err)
	}
	t.Logf("Listener status: %s", status)

}
