package ims

import "github.com/denverdino/aliyungo/common"

type User struct {
	UserId            string
	TenantId          string
	Type              string
	UserPrincipalName string
	DisplayName       string
	Enabled           bool

	Comments      string
	Source        int    //
	SourceText    string //
	CreateDate    string //
	UpdateDate    string //
	LastLoginDate string //

	MobilePhone  string
	Email        string
	WorkPhone    string
	HomeAddress  string
	WorkAddress  string
	EmployeeId   string
	EmployeeType string
	JobTitle     string
	ManagerId    string
	OrgUnitId    string
	DirectoryId  string
}

type CreateUserRequest struct {
	DirectoryId              string
	UserPrincipalName        string //true
	DisplayName              string //true
	Enabled                  bool
	MobilePhone              string
	Email                    string
	Comments                 string
	WorkPhone                string
	HomeAddress              string
	WorkAddress              string
	EmployeeId               string
	EmployeeType             string
	JobTitle                 string
	ManagerUserId            string
	ManagerUserPrincipalName string
	OrgUnitId                string
	OrgUnitPath              string
}

type CreateUserResponse struct {
	common.Response
	User User
}

func (client *ImsClient) CreateUser(args *CreateUserRequest) (*CreateUserResponse, error) {
	response := &CreateUserResponse{}
	err := client.Invoke("CreateUser", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type ListUsersRequest struct {
	DirectoryId string
	UserType    string
	ScopeType   string
	OrgUnit     string
	Marker      string
	MaxItems    int
}

type ListUsersResponse struct {
	common.Response
	Marker      string
	IsTruncated bool
	Users       struct {
		User []User
	}
}

func (client *ImsClient) ListUsers(args *ListUsersRequest) (*ListUsersResponse, error) {
	response := &ListUsersResponse{}
	err := client.Invoke("ListUsers", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type GetUserRequest struct {
	DirectoryId       string
	UserPrincipalName string
	UserId            string
}

type GetUserResponse struct {
	common.Response
	User User
}

func (client *ImsClient) GetUser(args *GetUserRequest) (*GetUserResponse, error) {
	response := &GetUserResponse{}
	err := client.Invoke("GetUser", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type DeleteUserRequest struct {
	DirectoryId       string
	UserPrincipalName string
	UserId            string
}

type DeleteUserResponse struct {
	common.Response
}

func (client *ImsClient) DeleteUser(args *DeleteUserRequest) (*DeleteUserResponse, error) {
	response := &DeleteUserResponse{}
	err := client.Invoke("DeleteUser", args, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
