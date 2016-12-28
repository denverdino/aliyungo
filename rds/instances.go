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
	DBInstanceId string
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
		DBInstanceId: args.DBInstanceId,
	}
	descResponse, err := client.DescribeDBInstanceIPArrayList(request)
	if err != nil {
		return response, err
	}
	fmt.Printf(" the result is %++v", descResponse)
	if err == nil && descResponse.Items != nil {
		for _, item := range descResponse.Items.DBInstanceIPArray {
			if item.SecurityIPList != "" {
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
