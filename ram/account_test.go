package ram

import (
	"strconv"
	"testing"
	"time"
)

var (
	offset   int64 = 100
	username       = strconv.FormatInt(time.Now().Unix(), 10)
	user           = UserRequest{
		User: User{
			UserName:    username,
			DisplayName: "unit_test_account",
			MobilePhone: "13500000000",
			Email:       "no-reply@alibaba-inc.com",
			Comments:    "no any comments",
		},
	}
	NewUserName = strconv.FormatInt((time.Now().Unix() + offset), 10)
	NewUser     = UpdateUserRequest{
		UserName:       username,
		NewUserName:    NewUserName,
		NewDisplayName: "unit_test_account_new",
	}
	listParams = ListUserRequest{}
)

func TestCreateUser(t *testing.T) {
	client := NewTestClient()
	resp, err := client.CreateUser(user)
	if err != nil {
		t.Errorf("Failed to CreateUser %v", err)
	}
	t.Logf("pass CreateUser %v", resp)
}

func TestGetUser(t *testing.T) {
	client := NewTestClient()
	resp, err := client.GetUser(username)
	if err != nil {
		t.Errorf("Failed to GetUser %v", err)
	}
	t.Logf("pass GetUser %v", resp)
}

func TestUpdateUsernewUser(t *testing.T) {
	client := NewTestClient()
	resp, err := client.UpdateUser(NewUser)
	if err != nil {
		t.Errorf("Failed to UpdateUser %v", err)
	}
	t.Logf("pass UpdateUser %v", resp)
}

func TestListUser(t *testing.T) {
	client := NewTestClient()
	resp, err := client.ListUser(listParams)
	if err != nil {
		t.Errorf("Failed to ListUser %v", err)
	}
	t.Logf("pass ListUser %v", resp)
}

func TestDeleteUser(t *testing.T) {
	client := NewTestClient()
	resp, err := client.DeleteUser(username)
	if err != nil {
		t.Errorf("Failed to DeletUser %v", err)
	}
	t.Logf("pass DeletUser %v", resp)
}
