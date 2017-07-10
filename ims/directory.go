package ims

import "github.com/denverdino/aliyungo/common"

type CheckDirectoryEnabledRequest struct {
}

type CheckDirectoryEnabledResponse struct {
	common.Response

	Enabled bool
}

func (client *ImsClient) CheckDirectoryEnabled(args *CheckDirectoryEnabledRequest) (*CheckDirectoryEnabledResponse, error) {
	response := &CheckDirectoryEnabledResponse{}
	err := client.Invoke("CheckDirectoryEnabled", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type GetDirectoryRequest struct {
	DirectoryId string
}

type GetDirectoryResponse struct {
	common.Response
	Directory struct {
		DirectoryId       string
		DisplayName       string
		OwnerId           string
		DefaultDomainName string
	}
}

func (client *ImsClient) GetDirectory(args *GetDirectoryRequest) (*GetDirectoryResponse, error) {
	response := &GetDirectoryResponse{}
	err := client.Invoke("GetDirectory", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
