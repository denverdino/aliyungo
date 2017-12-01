package nas

import "github.com/denverdino/aliyungo/common"

type CreateAccessRuleRequest struct {
	RegionId        common.Region
	AccessGroupName string
	SourceCidrIp    string
	RWAccessType    string
	UserAccessType  string
	Priority        int
}

type CreateAccessRuleResponse struct {
	common.Response
	AccessRuleId string
}

func (client *Client) CreateAccessRule(args *CreateAccessRuleRequest) (resp CreateAccessRuleResponse, err error) {
	response := CreateAccessRuleResponse{}

	err = client.Invoke("CreateAccessRule", args, &response)
	return response, err
}
