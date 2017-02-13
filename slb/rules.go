package slb

import "github.com/denverdino/aliyungo/common"

type CreateRulesArgs struct {
	RegionId       common.Region
	LoadBalancerId string
	ListenerPort   string
	RuleList       string
}

type DeleteRulesArgs struct {
	RegionId common.Region
	RuleIds  string
}

type SetRuleArgs struct {
	RegionId       common.Region
	RuleId         string
	VServerGroupId string
}

type DescribeRuleAttributeArgs struct {
	RegionId common.Region
	RuleId   string
}

type DescribeRulesArgs struct {
	RegionId       common.Region
	LoadBalancerId string
	ListenerPort   string
}

type CreateRulesResponse struct {
	common.Response
	Rules []struct {
		RuleId   string
		RuleName string
	}
}

type DeleteRulesResponse struct {
	common.Response
}

type SetRuleResponse DeleteRulesResponse

type DescribeRuleAttributeResponse struct {
	common.Response
	RuleName       string
	LoadBalancerId string
	ListenerPort   int
	Domain         string
	Url            string
	VServerGroupId string
}

type DescribeRulesResponse struct {
	common.Response
	RuleList string
}

type RuleType struct {
	RuleName       string
	Domain         string
	Url            string
	VServerGroupId string
}

// CreateRules create rules
//
// You can read doc at https://help.aliyun.com/document_detail/35226.html?spm=5176.doc27635.6.672.hc3DsR
func (client *Client) CreateRules(args *CreateRulesArgs) (response *CreateRulesResponse, err error) {
	response = &CreateRulesResponse{}
	err = client.Invoke("CreateRules", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

//DeleteRules delete rules
//
// You can read doc at https://help.aliyun.com/document_detail/35227.html?spm=5176.doc35227.6.673.Z45yIM
func (client *Client) DeleteRules(args *DeleteRulesArgs) (response *DeleteRulesResponse, err error) {
	response = &DeleteRulesResponse{}
	err = client.Invoke("DeleteRules", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

//SetRule set rule
//
// You can read doc at https://help.aliyun.com/document_detail/35228.html?spm=5176.doc35227.6.674.mrK2jY
func (client *Client) SetRule(args *SetRuleArgs) (response *SetRuleResponse, err error) {
	response = &SetRuleResponse{}
	err = client.Invoke("SetRule", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

//DescribeRuleAttribute describe RuleAttribute
//
// You can read doc at https://help.aliyun.com/document_detail/35229.html?spm=5176.doc35228.6.675.8suGoF
func (client *Client) DescribeRuleAttribute(args *DescribeRuleAttributeArgs) (response *DescribeRuleAttributeResponse, err error) {
	response = &DescribeRuleAttributeResponse{}
	err = client.Invoke("DescribeRuleAttribute", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

//DescribeRules describe rules
//
// You can read doc at https://help.aliyun.com/document_detail/35230.html?spm=5176.doc35229.6.676.9Kx6tg
func (client *Client) DescribeRules(args *DescribeRulesArgs) (response *DescribeRulesResponse, err error) {
	response = &DescribeRulesResponse{}
	err = client.Invoke("DescribeRules", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}
