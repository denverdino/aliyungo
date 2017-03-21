package ecs

import "testing"


func TestDescribeSnatTableEntry(t *testing.T) {

	client := NewTestClient()
	args := DescribeSnatTableEntriesArgs{
		RegionId: "cn-beijing",
		SnatTableId: "stb-2zeku7yxumzob3dsxytoy",
	}
	snatEntrySetType, _, err := client.DescribeSnatTableEntries(&args)
	if err != nil {
		t.Fatalf("Failed to DescribeBandwidthPackages: %v", err)
	}
}
