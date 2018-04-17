package ecs

import (
	"testing"

	"github.com/denverdino/aliyungo/common"
)

func Test_DescribeAvailableResource(t *testing.T) {
	client := NetTestLocationClientForDebug()

	args := &DescribeAvailableResourceArgs{
		RegionId:            TestRegionID,
		DestinationResource: common.DestinationResource_InstanceType,
		ZoneId:              TestZoneId,
		InstanceChargeType:  common.PostPaid,
		IoOptimized:         "optimized",
		//InstanceType:        "ecs.n4.small",
		NetworkCategory: "Vpc",
	}

	response, err := client.DescribeAvailableResource(args)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Result = %++v", response)
	}
}
