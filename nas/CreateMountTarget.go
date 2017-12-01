package nas

import "github.com/denverdino/aliyungo/common"

type CreateMountTargetRequest struct {
	RegionId        common.Region
	FileSystemId    string
	AccessGroupName string
	NetworkType     string
	VpcId           string
	VSwitchId       string
}

type CreateMountTargetResponse struct {
	common.Response
	MountTargetDomain string
}

func (client *Client) CreateMountTarget(args *CreateMountTargetRequest) (resp CreateMountTargetResponse, err error) {
	response := CreateMountTargetResponse{}

	err = client.Invoke("CreateMountTarget", args, &response)
	return response, err
}
