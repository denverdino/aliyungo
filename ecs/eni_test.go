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

func TestDescribeENI(t *testing.T) {
	req := DescribeNetworkInterfacesArgs{
		RegionId:           common.Beijing,
		NetworkInterfaceId: []string{"eni-testeni"},
	}
	client := NewTestClient()
	resp, err := client.DescribeNetworkInterfaces(&req)
	if err != nil {
		t.Errorf("Failed to DescribeNetworkInterfaces: %v", err)
	}
	if len(resp.NetworkInterfaceSets.NetworkInterfaceSet[0].PrivateIpSets.PrivateIpSet) != 4 {
		t.Errorf("assert network private ip count be 4, %+v", resp.NetworkInterfaceSets.NetworkInterfaceSet[0].PrivateIpSets.PrivateIpSet)
	}
	t.Logf("%+v", resp.NetworkInterfaceSets.NetworkInterfaceSet[0])
}

func TestFindENIByPrivateIP(t *testing.T) {
	req := DescribeNetworkInterfacesArgs{
		RegionId:         common.Shanghai,
		VpcID:            "vpc-xxx",
		PrivateIpAddress: []string{"192.168.108.191"},
	}
	client := NewTestClient()
	resp, err := client.DescribeNetworkInterfaces(&req)
	if err != nil {
		t.Errorf("Failed to DescribeNetworkInterfaces: %v", err)
	}
	t.Logf("%+v", resp.NetworkInterfaceSets.NetworkInterfaceSet)
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
