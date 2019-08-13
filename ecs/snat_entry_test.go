package ecs

import "testing"

func TestDescribeSnatTableEntry(t *testing.T) {

	client := NewTestClient()
	args := DescribeSnatTableEntriesArgs{
		RegionId:    "cn-beijing",
		SnatTableId: "stb-abc",
	}
	_, _, err := client.DescribeSnatTableEntries(&args)
	if err != nil {
		t.Fatalf("Failed to DescribeBandwidthPackages: %v", err)
	}
}

func TestCreateSnatEntryWithSourceCIDR(t *testing.T) {
	client := NewTestClient()
	args := CreateSnatEntryArgs{
		RegionId:    "cn-beijing",
		SnatTableId: "stb-xxx",
		SnatIp:      "47.XX.XX.98",
		SourceCIDR:  "192.168.1.1/32",
	}

	_, err := client.CreateSnatEntry(&args)
	if err != nil {
		t.Errorf("failed to CreateSnatEntry with SourceCIDR: %v", err)
	}
}
