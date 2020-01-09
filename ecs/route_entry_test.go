package ecs

import (
	"testing"
)

func TestClient_DescribeRouteEntry(t *testing.T) {
	client := NewVpcTestClientForDebug()

	nextHopId := "i-xxxx"
	destinationCidrBlock := "172.xxx/x"
	args := &DescribeRouteEntryListArgs{
		RegionId:             "cn-hangzhou",
		RouteTableId:         "vtb-xxxxx",
		DestinationCidrBlock: destinationCidrBlock,
		NextHopId:            nextHopId,
		RouteEntryType:       "Custom",
	}

	response, err := client.DescribeRouteEntryList(args)

	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Result %++v", response)
	}
}