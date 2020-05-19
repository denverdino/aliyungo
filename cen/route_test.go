package cen

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	ak  = ""
	sec = ""
)

func TestDescribePublishedRoute(t *testing.T) {
	client := NewCENClient(ak, sec, "cn-shanghai")
	res, err := client.DescribePublishedRouteEntries(
		&DescribePublishedRouteEntriesArgs{
			CenId:                     "cen-qhu4rn3cknrg5o4qhl",
			ChildInstanceType:         "VPC",
			ChildInstanceRegionId:     "cn-shanghai",
			ChildInstanceRouteTableId: "vtb-uf699blmsutb4wkbzqcmt",
			ChildInstanceId:           "vpc-uf6ch2jfder4r0z51vtox",
		},
	)
	if err != nil {
		t.Errorf("describe: %s", err.Error())
		t.FailNow()
	}
	fmt.Printf("Result: %+v", res)
	b, err := json.MarshalIndent(res, "", "  ")
	fmt.Printf("%s", b)
}

func TestPublishedRoute(t *testing.T) {
	client := NewCENClient(ak, sec, "cn-shanghai")
	err := client.PublishRouteEntries(
		&PublishRouteEntriesArgs{
			CenId:                     "cen-qhu4rn3cknrg5o4qhl",
			ChildInstanceType:         "VPC",
			ChildInstanceRegionId:     "cn-shanghai",
			ChildInstanceRouteTableId: "vtb-uf6nco4vj87ly556c589f",
			ChildInstanceId:           "vpc-uf6ch2jfder4r0z51vtox",
			DestinationCidrBlock:      "192.168.0.0/26",
		},
	)
	if err != nil {
		t.Errorf("publish: %s", err.Error())
		t.FailNow()
	}
}
