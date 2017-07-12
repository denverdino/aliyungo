package ims

import "github.com/denverdino/aliyungo/common"

type CreateAccessKeyRequest struct {
	UserPrincipalName string
}

type CreateAccessKeyResponse struct {
	common.Response
	AccessKey AccessKey
}

type AccessKey struct {
	AccessKeyId     string
	AccessKeySecret string
	CreateDate      string
	Status          string
}

func (client *ImsClient) CreateAccessKey(args *CreateAccessKeyRequest) (*CreateAccessKeyResponse, error) {
	response := &CreateAccessKeyResponse{}

	err := client.Invoke("CreateAccessKey", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type DeleteAccessKeyRequest struct {
	UserAccessKeyId   string
	UserPrincipalName string
}

type DeleteAccessKeyResponse struct {
	common.Response
}

func (client *ImsClient) DeleteAccessKey(args *DeleteAccessKeyRequest) (*DeleteAccessKeyResponse, error) {

	response := &DeleteAccessKeyResponse{}

	err := client.Invoke("DeleteAccessKey", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type ListAccessKeysRequest struct {
	UserPrincipalName string
}

type ListAccessKeysResponse struct {
	common.Response

	AccessKeys struct {
		AccessKey []AccessKey
	}
}

func (client *ImsClient) ListAccessKeys(args *ListAccessKeysRequest) (*ListAccessKeysResponse, error) {
	response := &ListAccessKeysResponse{}

	err := client.Invoke("ListAccessKeys", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type UpdateAccessKeyRequest struct {
	UserPrincipalName string
	UserAccessKeyId   string
	Status            string
}

type UpdateAccessKeyResponse struct {
	common.Response
}

func (client *ImsClient) UpdateAccessKey(args *UpdateAccessKeyRequest) (*UpdateAccessKeyResponse, error) {
	response := &UpdateAccessKeyResponse{}

	err := client.Invoke("UpdateAccessKey", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
