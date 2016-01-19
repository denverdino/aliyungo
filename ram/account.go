package ram

type UserRequest struct {
	User
}

type UserResponse struct {
	RamCommonResponse
	User User
}

type UpdateUserRequest struct {
	UserName       string
	NewUserName    string
	NewDisplayName string
	NewMobilePhone string
	NewEmail       string
	NewComments    string
}

type ListUserRequest struct {
	Marker   string
	MaxItems int8
}

type ListUserResponse struct {
	RamCommonResponse
	IsTruncated bool
	Marker      string
	Users       []User
}

func (client *RamClient) CreateUser(user UserRequest) (UserResponse, error) {
	return UserResponse{}, nil
}

func (client *RamClient) GetUser(username string) (UserResponse, error) {
	return UserResponse{}, nil
}

func (client *RamClient) UpdateUser(newUser UpdateUserRequest) (UserResponse, error) {
	return UserResponse{}, nil
}

func (client *RamClient) DeleteUser(username string) (RamCommonResponse, error) {
	return RamCommonResponse{}, nil
}

func (client *RamClient) ListUser(listParams ListUserRequest) (ListUserResponse, error) {
	return ListUserResponse{}, nil
}
