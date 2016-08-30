package vpc

import (
	"testing"

	"github.com/hdksky/aliyungo/common"
)

func TestVPCCreationAndDeletion(t *testing.T) {

	client := NewTestClient()

	regionId := common.Region("cn-beijing")
	args := CreateVpcArgs{
		RegionId:    regionId,
		VpcName:     "My_AliyunGO_test_VPC",
		Description: "My AliyunGO test VPC",
		CidrBlock:   "172.16.0.0/16",
		ClientToken: client.GenerateClientToken(),
	}

	resp, err := client.CreateVpc(&args)
	if err != nil {
		t.Fatalf("Failed to create VPC: %v", err)
	}
	t.Logf("VPC is created successfully: %++v", resp)

	vpcId := resp.VpcId
	newName := args.VpcName + "_update"
	newDesc := args.Description + "_update"
	modifyArgs := ModifyVpcAttributeArgs{
		VpcId:       vpcId,
		VpcName:     newName,
		Description: newDesc,
	}
	err = client.ModifyVpcAttribute(&modifyArgs)
	if err != nil {
		t.Errorf("Failed to modify VPC: %v", err)
	}

	describeArgs := DescribeVpcsArgs{
		VpcId:    vpcId,
		RegionId: regionId,
	}
	vpcs, _, err := client.DescribeVpcs(&describeArgs)
	if err != nil {
		t.Errorf("Failed to describe VPCs: %v", err)
	}
	t.Logf("VPCs: %++v", vpcs)
	if vpcs[0].VpcName != newName {
		t.Errorf("Failed to modify VPC with new name: %s", newName)
	}

	err = client.WaitForVpcAvailable(regionId, vpcId, 60)
	if err != nil {
		t.Errorf("Failed to wait VPC to available: %v", err)
	}

	err = client.DeleteVpc(vpcId)
	if err != nil {
		t.Errorf("Failed to delete VPC: %v", err)
	}
	t.Logf("VPC %s is deleted successfully.", vpcId)

}
