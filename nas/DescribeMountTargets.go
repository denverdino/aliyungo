package nas

import "github.com/denverdino/aliyungo/common"

type DescribeMountTargetsRequest struct {
	FileSystemId      string
	RegionId          common.Region
	MountTargetDomain string

	PageSize   int
	PageNumber int
}

type DescribeMountTargetsResponse struct {
	common.Response
	MountTargets struct {
		MountTarget []MountTarget
	}
}

type MountTarget struct {
	AccessGroupName   string
	MountTargetDomain string
	VpcId             string
	VSwitchId         string
	NetworkType       string
	Status            string
}

func (client *Client) DescribeMountTargets(args *DescribeMountTargetsRequest) (resp DescribeMountTargetsResponse, err error) {
	response := DescribeMountTargetsResponse{}
	err = client.Invoke("DescribeMountTargets", args, &response)
	return response, err
}
