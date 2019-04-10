package slb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func testListeners(t *testing.T, client *Client, loadBalancerId string) {

	port := 1234

	creationArgs := CreateLoadBalancerTCPListenerArgs{
		LoadBalancerId:    loadBalancerId,
		ListenerPort:      port,
		BackendServerPort: 1234,
		Bandwidth:         -1,
	}

	err := client.CreateLoadBalancerTCPListener(&creationArgs)
	if err != nil {
		t.Errorf("Failed to CreateLoadBalancerTCPListener: %v", err)
	}

	response, err := client.DescribeLoadBalancerTCPListenerAttribute(loadBalancerId, port)
	if err != nil {
		t.Errorf("Failed to DescribeLoadBalancerTCPListenerAttribute: %v", err)
	}
	t.Logf("Listener: %++v", *response)

	err = client.StartLoadBalancerListener(loadBalancerId, port)
	if err != nil {
		t.Errorf("Failed to StartLoadBalancerListener: %v", err)
	}

	status, err := client.WaitForListener(loadBalancerId, port, TCP)
	if err != nil {
		t.Errorf("Failed to WaitForListener: %v", err)
	}
	t.Logf("Listener status: %s", status)

	response, err = client.DescribeLoadBalancerTCPListenerAttribute(loadBalancerId, port)
	if err != nil {
		t.Errorf("Failed to DescribeLoadBalancerTCPListenerAttribute: %v", err)
	}
	t.Logf("Listener: %++v", *response)
}

func TestDescribeListener(t *testing.T) {

	response, err := client.DescribeLoadBalancerTCPListenerAttribute(loadBalancerId, 22)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(PrettyJson(response))
	}
}

func TestDescribeSLB(t *testing.T) {

	response, err := client.DescribeLoadBalancerAttribute(loadBalancerId)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(PrettyJson(response))
	}
}

func PrettyJson(obj interface{}) string {
	pretty := bytes.Buffer{}
	data, err := json.Marshal(obj)
	if err != nil {
		return err.Error()
	}
	err = json.Indent(&pretty, data, "", "    ")

	if err != nil {
		return err.Error()
	}

	return pretty.String()
}
