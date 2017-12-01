package nas

import "github.com/denverdino/aliyungo/common"

type CreateFileSystemRequest struct {
	RegionId     common.Region
	StorageType  string
	ProtocolType string
	Description  string
}

type CreateFileSystemResponse struct {
	common.Response
	FileSystemId string
}

func (client *Client) CreateFileSystem(args *CreateFileSystemRequest) (resp CreateFileSystemResponse, err error) {
	response := CreateFileSystemResponse{}

	err = client.Invoke("CreateFileSystem", args, &response)
	return response, err
}
