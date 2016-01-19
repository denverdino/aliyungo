package ram

/*
	CreateAccessKey()
	UpdateAccessKey()
	DeleteAccessKey()
	ListAccessKeys()
*/
type State string

type AccessKeyResponse struct {
	RamCommonResponse
	AccessKey AccessKey
}

type UpdateAccessKeyRequest struct {
	AccessKeyId string
	Status      State
	username    string
}

type AccessKeyListResponse struct {
	RamCommonResponse
	AccessKeys []AccessKey
}

func (client *RamClient) CreateAccessKey(username string) (AccessKeyResponse, error) {
	return AccessKeyResponse{}, nil
}

func (client *RamClient) UpdateAccessKey(accessKeyRequest UpdateAccessKeyRequest) (RamCommonResponse, error) {
	return RamCommonResponse{}, nil
}

func (client *RamClient) DeleteAccessKey(accessKeyRequest UpdateAccessKeyRequest) (RamCommonResponse, error) {
	return RamCommonResponse{}, nil
}

func (client *RamClient) ListAccessKeys(username string) (AccessKeyListResponse, error) {
	return AccessKeyListResponse{}, nil
}
