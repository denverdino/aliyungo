package ram_test

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/denverdino/aliyungo/ram"
)

/*
  Please also set account id in env so that roles could be created test
	 AccessKeyId=YourAccessKeyId AccessKeySecret=YourAccessKeySecret AccountId=111111111 go test -v -run=Role
*/
var (
	accountId = os.Getenv("AccountId")
	roleName  = strconv.FormatInt(time.Now().Unix(), 10)

	princpal = ram.AssumeRolePolicyPrincpal{RAM: []string{"acs:ram::" + accountId + ":root"}}

	policyDocument = ram.AssumeRolePolicyDocument{
		Statement: []ram.AssumeRolePolicyItem{
			ram.AssumeRolePolicyItem{Action: "sts:AssumeRole", Effect: "Allow", Principal: princpal},
		},
		Version: "1"}

	newPolicyDocument = ram.AssumeRolePolicyDocument{
		Statement: []ram.AssumeRolePolicyItem{
			ram.AssumeRolePolicyItem{Action: "sts:AssumeRole", Effect: "Deny", Principal: princpal},
		},
		Version: "1"}

	RoleReq = ram.RoleRequest{
		RoleName:                 roleName,
		AssumeRolePolicyDocument: getAssumeRolePolicyDocumentStr(),
		Description:              "this is a role for unit test purpose",
	}

	updateRoleRequest = ram.UpdateRoleRequest{
		RoleName:                    roleName,
		NewAssumeRolePolicyDocument: getNewAssumeRolePolicyDocumentStr(),
	}

	roleQuery = ram.RoleQueryRequest{RoleName: roleName}
)

func getAssumeRolePolicyDocumentStr() string {
	b, _ := json.Marshal(policyDocument)
	return string(b)
}

func getNewAssumeRolePolicyDocumentStr() string {
	b, _ := json.Marshal(newPolicyDocument)
	return string(b)
}

func TestCreateRole(t *testing.T) {
	client := ram.NewTestClient()
	resp, err := client.CreateRole(RoleReq)
	if err != nil {
		t.Errorf("Failed to CreateRole %v", err)
	}
	t.Logf("pass CreateRole %v", resp)
}

func TestGetRole(t *testing.T) {
	client := ram.NewTestClient()
	resp, err := client.GetRole(roleQuery)
	if err != nil {
		t.Errorf("Failed to GetRole %v", err)
	}
	t.Logf("pass GetRole %v", resp)
}

func TestUpdateRole(t *testing.T) {
	client := ram.NewTestClient()
	resp, err := client.UpdateRole(updateRoleRequest)
	if err != nil {
		t.Errorf("Failed to UpdateRole %v", err)
	}
	t.Logf("pass UpdateRole %v", resp)
}

func TestListRoles(t *testing.T) {
	client := ram.NewTestClient()
	resp, err := client.ListRoles()
	if err != nil {
		t.Errorf("Failed to ListRoles %v", err)
	}
	t.Logf("pass ListRoles %v", resp)
}

func TestDeleteRole(t *testing.T) {
	client := ram.NewTestClient()
	resp, err := client.DeleteRole(ram.RoleQueryRequest{RoleName: roleName})
	if err != nil {
		t.Errorf("Failed to DeleteRole %v", err)
	}
	t.Logf("pass DeleteRole %v", resp)
}
