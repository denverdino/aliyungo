package ims

import (
	"fmt"
	"testing"
	"time"
)

func TestImsClient_CreateUser(t *testing.T) {
	name := fmt.Sprintf("user_%d", time.Now().Unix())
	args := &CreateUserRequest{
		UserPrincipalName: fmt.Sprintf("%s@1312623533146984.onaliyun.com", name),
		DisplayName:       "acs_ram_account",
		Enabled:           true,
	}

	response, err := client.CreateUser(args)
	if err != nil {
		t.Fatalf("Failed to CreateUser %++v", err)
	} else {
		t.Logf("Response is %++v", response)
	}
}

func TestImsClient_ListUsers(t *testing.T) {
	args := &ListUsersRequest{}

	response, err := client.ListUsers(args)
	if err != nil {
		t.Fatalf("Failed to ListUsers %++v", err)
	} else {
		for index, user := range response.Users.User {
			t.Logf("users[%d] =  %++v", index, user)
		}
	}
}

func TestImsClient_GetUser(t *testing.T) {
	args := &GetUserRequest{
		UserId: "2860441599678535536",
	}

	response, err := client.GetUser(args)
	if err != nil {
		t.Fatalf("Failed to GetUser %++v", err)
	} else {
		t.Logf("Response is  %++v", response)
	}
}

func TestImsClient_DeleteUser(t *testing.T) {
	name := fmt.Sprintf("user-%d", time.Now().Unix())
	createArgs := &CreateUserRequest{
		UserPrincipalName: fmt.Sprintf("%s@1312623533146984.onaliyun.com", name),
		DisplayName:       name,
		Enabled:           true,
	}

	createResponse, err := client.CreateUser(createArgs)
	if err != nil {
		t.Fatalf("Failed to createUser %++v", err)
	} else {
		fmt.Printf("userId = %s", createResponse.User.UserId)
		args := &DeleteUserRequest{
			UserId: createResponse.User.UserId,
		}

		time.Sleep(5 * time.Second)
		fmt.Printf("args = %++v", args)
		_, err := client.DeleteUser(args)
		if err != nil {
			t.Fatalf("Failed to DeleteUser %++v", err)
		} else {
			t.Logf("DeleteUser successful")
		}
	}
}
