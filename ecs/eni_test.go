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

func TestModifyNetworkInterfaceAttribute(t *testing.T) {
	args := &ModifyNetworkInterfaceAttributeArgs{
		RegionId:           common.Shanghai,
		NetworkInterfaceId: "eni-testeni",
		SecurityGroupId:    []string{"sg-xxx", "sg-yyy"},
	}

	client := NewTestClient()
	_, err := client.ModifyNetworkInterfaceAttribute(args)
	if err != nil {
		t.Errorf("failed to ModifyNetworkInterfaceAttribute: %v", err)
	}
}

func TestCreateNetworkInterface(t *testing.T) {
	args := &CreateNetworkInterfaceArgs{
		RegionId:                       common.Shanghai,
		VSwitchId:                      "vsw-xxx",
		SecurityGroupIds:               []string{"sg-xxx", "sg-yyy"},
		SecondaryPrivateIpAddressCount: 9,
	}
	client := NewTestClient()
	resp, err := client.CreateNetworkInterface(args)
	if err != nil {
		t.Errorf("failed to CreateNetworkInterface: %v", err)
	}
	t.Logf("new eni info: %+v", resp)
}
