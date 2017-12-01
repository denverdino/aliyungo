package nas

import "github.com/denverdino/aliyungo/common"

type DescribeAccessRulesRequest struct {
	RegionId        common.Region
	AccessGroupName string
	AccessRuleId    string
	PageSize        int
	PageNumber      int
}

type DescribeAccessRulesResponse struct {
	common.Response
	TotalCount  int
	PageSize    int
	PageNumber  int
	AccessRules struct {
		AccessRule []AccessRule
	}
}

type AccessRule struct {
	RWAccess     string
	UserAccess   string
	Priority     int
	SourceCidrIp string
	AccessRuleId string
}

func (client *Client) DescribeAccessRules(args *DescribeAccessRulesRequest) (resp DescribeAccessRulesResponse, err error) {
	response := DescribeAccessRulesResponse{}
	err = client.Invoke("DescribeAccessRules", args, &response)
	return response, err
}
