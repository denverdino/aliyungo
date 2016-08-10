package ram

type RoleRequest struct {
	RoleName                 string
	AssumeRolePolicyDocument string
	Description              string
}

type Role struct {
	RoleId                   string
	RoleName                 string
	Arn                      string
	Description              string
	AssumeRolePolicyDocument AssumeRolePolicyDocument
	CreateDate               string
	UpdateDate               string
}

type RoleResponse struct {
	RamCommonResponse
	Role Role
}

type RoleQueryRequest struct {
	RoleName string
}

type UpdateRoleRequest struct {
	RoleName                    string
	NewAssumeRolePolicyDocument string
}

type ListRoleResponse struct {
	RamCommonResponse
	Roles struct {
		Role []Role
	}
}

func (client *RamClient) CreateRole(role RoleRequest) (RoleResponse, error) {
	var roleResponse RoleResponse
	err := client.Invoke("CreateRole", role, &roleResponse)
	if err != nil {
		return RoleResponse{}, err
	}
	return roleResponse, nil
}

func (client *RamClient) GetRole(roleQuery RoleQueryRequest) {
	var roleResponse RoleResponse
	err := client.Invoke("GetRole", userQuery, &roleResponse)
	if err != nil {
		return RoleResponse{}, nil
	}
	return roleResponse, nil
}

func (client *RamClient) UpdateRole(newRole UpdateRoleRequest) {
	var roleResponse RoleResponse
	err := client.Invoke("UpdateRole", newRole, &roleResponse)
	if err != nil {
		return RoleResponse{}, err
	}
	return roleResponse, nil
}

func (client *RamClient) ListRoles() {
	var roleList ListRoleResponse
	err := client.Invoke("ListRoles", nil, &roleList)
	if err != nil {
		return ListRoleResponse{}, err
	}
	return roleList, nil
}

func (client *RamClient) DeleteRole(roleQuery RoleQueryRequest) {
	var commonResp RamCommonResponse
	err := client.Invoke("DeleteRole", roleQuery, &commonResp)
	if err != nil {
		return RamCommonResponse{}, err
	}
	return commonResp, nil
}
