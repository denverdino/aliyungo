package ram

/*
	ringtail 2016/1/19
	All RAM apis provided
*/

type RamClientInterface interface {
	//ram user
	CreateUser(user UserRequest) (UserResponse, error)
	GetUser(username string) (UserResponse, error)
	UpdateUser(newUser UpdateUserRequest) (UserResponse, error)
	DeleteUser(username string) (RamCommonResponse, error)
	ListUser(listParams ListUserRequest) (ListUserResponse, error)

	//TODO login ram console
	CreateLoginProfile()
	GetLoginProfile()
	DeleteLoginProfile()
	UpdateLoginProfile()

	//ram ak
	CreateAccessKey(username string) (AccessKeyResponse, error)
	UpdateAccessKey(accessKeyRequest UpdateAccessKeyRequest) (RamCommonResponse, error)
	DeleteAccessKey(accessKeyRequest UpdateAccessKeyRequest) (RamCommonResponse, error)
	ListAccessKeys(username string) (AccessKeyListResponse, error)

	//TODO MFA
	CreateVirtualMFADevices()
	ListVirtualMFADevices()
	DeleteVirtualMFADevices()
	BindMFADevice()
	GetUserMFAInfo()

	//TODO group
	CreateGroup()
	GetGroup()
	UpdateGroup()
	ListGroup()
	DeleteGroup()
	AddUserToGroup()
	RemoveUserFromGroup()
	ListGroupsForUser()
	ListUsersForGroup()

	//TODO role
	CreateRole()
	GetRole()
	UpdateRole()
	ListRoles()
	DeleteRole()

	//policy

	CreatePolicy()
	GetPolicy()
	DeletePolicy()
	ListPolicies()
	CreatePolicyVersion()
	GetPolicyVersion()
	DeletePolicyVersion()
	ListPolicyVersions()
	SetDefaultPolicyVersion()
	AttachPolicyToUser()
	DetachPolicyFromUser()
	ListEnitiesForPolicy()
	ListPoliciesForUser()
	ListPoliciesForGroup()
	ListPoliciesForRole()

	//TODO security apis
	SetAccountAlias(accountAlias AccountAlias) (RamCommonResponse, error)
	GetAccountAlias() (AccountAliasResponse, error)
	ClearAccountAlias() (RamCommonResponse, error)
	SetPasswordPolicy(passwordPolicy PasswordPolicyRequest) (PasswordPolicyResponse, error)
	GetPasswordPolicy(accountAlias AccountAlias) (PasswordPolicyResponse, error)
}
