package ecs

import (
	"github.com/denverdino/aliyungo/common"
	"testing"
	"time"
)

func TestAssignPrivateIPAddresses(t *testing.T) {
	req := AssignPrivateIpAddressesArgs{
		RegionId:           common.Beijing,
		NetworkInterfaceId: "eni-testeni",
		PrivateIpAddress:   []string{"192.168.1.200", "192.168.1.201"},
	}
	client := NewTestClient()
	_, err := client.AssignPrivateIpAddresses(&req)
	if err != nil {
		t.Errorf("Failed to AssignPrivateIpAddresses: %v", err)
	}

	time.Sleep(5 * time.Second)

	req = AssignPrivateIpAddressesArgs{
		RegionId:                       common.Beijing,
		NetworkInterfaceId:             "eni-testeni",
		SecondaryPrivateIpAddressCount: 1,
	}

	_, err = client.AssignPrivateIpAddresses(&req)
	if err != nil {
		t.Errorf("Failed to AssignPrivateIpAddresses: %v", err)
	}
}

func TestUnAssignPrivateIPAddresses(t *testing.T) {
	req := UnassignPrivateIpAddressesArgs{
		RegionId:           common.Beijing,
		NetworkInterfaceId: "eni-testeni",
		PrivateIpAddress:   []string{"192.168.1.200", "192.168.1.201"},
	}
	client := NewTestClient()
	_, err := client.UnassignPrivateIpAddresses(&req)
	if err != nil {
		t.Errorf("Failed to UnAssignPrivateIpAddresses: %v", err)
	}
}
