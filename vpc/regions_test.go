package vpc

import "testing"

func TestDescribeRegions(t *testing.T) {

	client := NewTestClient()

	regions, err := client.DescribeRegions()

	if err == nil {
		t.Logf("regions: %v", regions)
	} else {
		t.Errorf("Failed to DescribeRegions: %v", err)
	}

	zones, err := client.DescribeZones("cn-beijing")

	if err == nil {
		t.Logf("zones: %v", zones)
	} else {
		t.Errorf("Failed to DescribeZones: %v", err)
	}

}
