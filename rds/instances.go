package rds

import (
	"fmt"

	"github.com/denverdino/aliyungo/common"
)

// ref: https://help.aliyun.com/document_detail/26242.html
type ModifySecurityIpsArgs struct {
	DBInstanceId               string
	SecurityIps                string
	DBInstanceIPArrayName      string
	DBInstanceIPArrayAttribute string
}

type DescribeDBInstanceIPArrayListArgs struct {
	DBInstanceId          string
	DBInstanceIPArrayName string
}

type DBInstanceIPs struct {
	DBInstanceIPArrayName      string
	DBInstanceIPArrayAttribute string
	SecurityIPList             string
}

type DBInstanceIPsItems struct {
	DBInstanceIPArray []DBInstanceIPs
}

type DescribeDBInstanceIPArrayListResponse struct {
	common.Response

	Items *DBInstanceIPsItems
}

func (client *Client) ModifySecurityIps(args *ModifySecurityIpsArgs) (resp *common.Response, err error) {
	response := &common.Response{}
	if args.SecurityIps == "" {
		return response, nil
	}
	//Query security ips and add new ips
	request := &DescribeDBInstanceIPArrayListArgs{
		DBInstanceId:          args.DBInstanceId,
		DBInstanceIPArrayName: args.DBInstanceIPArrayName,
	}
	descResponse, err := client.DescribeDBInstanceIPArrayList(request)
	if err != nil {
		return response, err
	}
	fmt.Printf(" the result is %++v", descResponse)
	if err == nil && descResponse.Items != nil {
		for _, item := range descResponse.Items.DBInstanceIPArray {
			if item.DBInstanceIPArrayName == args.DBInstanceIPArrayName && item.SecurityIPList != "" {
				args.SecurityIps = args.SecurityIps + "," + item.SecurityIPList
			}
		}
	}
	fmt.Printf(" the args is %++v", args)
	err = client.Invoke("ModifySecurityIps", args, &response)
	return response, err
}

func (client *Client) DescribeDBInstanceIPArrayList(args *DescribeDBInstanceIPArrayListArgs) (*DescribeDBInstanceIPArrayListResponse, error) {
	resp := &DescribeDBInstanceIPArrayListResponse{}
	err := client.Invoke("DescribeDBInstanceIPArrayList", args, resp)
	return resp, err
}

type DescribeRegionsArgs struct {
}

type DescribeRegionsResponse struct {
	Regions struct {
		RDSRegion []RDSRegion
	}
}

type RDSRegion struct {
	RegionId string
	ZoneId   string
}

// DescribeRegions describe rds regions
//
// You can read doc at https://help.aliyun.com/document_detail/26243.html?spm=5176.doc26244.6.715.OSNUa8
func (client *Client) DescribeRegions() (resp *DescribeRegionsResponse, err error) {
	args := DescribeRegionsArgs{}
	response := DescribeRegionsResponse{}
	err = client.Invoke("DescribeRegions", &args, &response)

	if err != nil {
		return nil, err
	}
	return &response, nil
}
