package ecs

import (
	"testing"
	"time"

	"github.com/denverdino/aliyungo/common"
)

func TestVPCCreationAndDeletion(t *testing.T) {

	client := NewTestClient()

	instance, err := client.DescribeInstanceAttribute(TestInstanceId)
	if err != nil {
		t.Fatalf("Failed to describe instance %s: %v", TestInstanceId, err)
	}

	//client.SetDebug(true)

	regionId := instance.RegionId
	zoneId := instance.ZoneId

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

	_testECSSecurityGroupCreationAndDeletion(t, client, regionId, vpcId)

	//Test VSwitch
	vSwitchId, err := testCreateVSwitch(t, client, regionId, zoneId, vpcId, resp.VRouterId)
	if err != nil {
		t.Errorf("Failed to create VSwitch: %v", err)
	} else {
		if TestIAmRich {
			instanceId, sgId, err := testCreateInstanceVpc(t, client, regionId, vpcId, vSwitchId, instance.ImageId)

			if err == nil {
				testEipAddress(t, client, regionId, instanceId)

				//Test VRouter
				testVRouter(t, client, regionId, vpcId, resp.VRouterId, instanceId)

			}

			if instanceId != "" {
				err = client.StopInstance(instanceId, true)
				if err != nil {
					t.Errorf("Failed to stop instance %s: %v", instanceId, err)
				} else {
					err = client.WaitForInstance(instanceId, Stopped, 0)
					if err != nil {
						t.Errorf("Instance %s is failed to stop: %v", instanceId, err)
					}
					t.Logf("Instance %s is stopped successfully.", instanceId)
				}
				err = client.DeleteInstance(instanceId)

				if err != nil {
					t.Errorf("Failed to delete instance %s: %v", instanceId, err)
				} else {
					t.Logf("Instance %s is deleted successfully.", instanceId)
				}
			}
			if sgId != "" {
				//Wait the instance deleted completedly
				time.Sleep(10 * time.Second)
				err = client.DeleteSecurityGroup(regionId, sgId)
				if err != nil {
					t.Fatalf("Failed to delete security group %s: %v", sgId, err)
				}
				t.Logf("Security group %s is deleted successfully.", sgId)
			}
		}
	}

	if vSwitchId != "" {
		err = client.DeleteVSwitch(vSwitchId)
		if err != nil {
			t.Fatalf("Failed to delete VSwitch: %v", err)
		}
		t.Logf("VSwitch %s is deleted successfully.", vSwitchId)
	}

	time.Sleep(20 * time.Second)

	err = client.DeleteVpc(vpcId)
	if err != nil {
		t.Errorf("Failed to delete VPC: %v", err)
	}
	t.Logf("VPC %s is deleted successfully.", vpcId)

}

func testCreateInstanceVpc(t *testing.T, client *Client, regionId common.Region, vpcId string, vswitchId, imageId string) (instanceId string, sgId string, err error) {
	sgName := "test-security-group"
	args := CreateSecurityGroupArgs{
		RegionId:          regionId,
		VpcId:             vpcId,
		SecurityGroupName: sgName,
	}

	sgId, err = client.CreateSecurityGroup(&args)

	if err != nil {
		t.Errorf("Failed to create security group %s: %v", sgName, err)
		return "", "", err
	}

	createArgs := CreateInstanceArgs{
		RegionId:        regionId,
		ImageId:         imageId,
		InstanceType:    "ecs.t1.small",
		SecurityGroupId: sgId,
		VSwitchId:       vswitchId,
	}

	instanceId, err = client.CreateInstance(&createArgs)
	if err != nil {
		t.Errorf("Failed to create instance from Image %s: %v", imageId, err)
		return "", sgId, err
	}
	t.Logf("Instance %s is created successfully.", instanceId)
	instance, err := client.DescribeInstanceAttribute(instanceId)
	t.Logf("Instance: %++v  %v", instance, err)
	err = client.WaitForInstance(instanceId, Stopped, 0)

	err = client.StartInstance(instanceId)
	if err != nil {
		t.Errorf("Failed to start instance %s: %v", instanceId, err)
		return instanceId, sgId, err
	}
	err = client.WaitForInstance(instanceId, Running, 0)

	return instanceId, sgId, err
}

func TestClient_DescribeVpcs(t *testing.T) {
	client := NewTestClientForDebug()
	client.SetSecurityToken(TestSecurityToken)

	args := &DescribeVpcsArgs{
		RegionId: TestRegionID,
		Pagination: common.Pagination{
			PageNumber: 1,
			PageSize:   100,
		},
	}

	vpcs, _, err := client.DescribeVpcs(args)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Result = %++v", vpcs)
	}
}

func TestClient_CreateVpc(t *testing.T) {
	client := NewVpcTestClientForDebug()

	for i := 0; i < 1; i++ {
		args := &CreateVpcArgs{
			RegionId:    common.Qingdao,
			CidrBlock:   "172.16.0.0/16",
			VpcName:     "vpc-quota-test",
			Description: "vpc-quota-test",
			ClientToken: client.GenerateClientToken(),
		}

		response, err := client.CreateVpc(args)
		if err != nil {
			t.Fatalf("Error %++v", err)
		} else {
			t.Logf("Result %++v", response)
		}
		time.Sleep(1 * time.Second)
	}
}

func TestClient_AllocateEipAddress2(t *testing.T) {
	client := NewVpcTestClientForDebug()
	for i := 0; i < 200; i++ {
		args := &AllocateEipAddressArgs{
			RegionId:           common.Beijing,
			Bandwidth:          100,
			InternetChargeType: common.PayByTraffic,
			ClientToken:        client.GenerateClientToken(),
		}

		ip, id, err := client.AllocateEipAddress(args)
		if err != nil {
			t.Fatalf("Error %++v", err)
		} else {
			t.Logf("Eip = %s, Id=%s", ip, id)
		}
	}
}

func TestClient_DeleteVpc(t *testing.T) {
	client := NewVpcTestClientForDebug()

	args := &DescribeVpcsArgs{
		RegionId: common.Beijing,
		Pagination: common.Pagination{
			PageNumber: 1,
			PageSize:   100,
		},
	}

	vpcs, _, err := client.DescribeVpcs(args)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		for _, vpc := range vpcs {
			if vpc.VpcName == "vpc-quota-test" {
				client.DeleteVpc(vpc.VpcId)
			}
		}
	}
}

func Test_DescribeNatGateways(t *testing.T) {
	client := NewVpcTestClientForDebug()
	args := &DescribeNatGatewaysArgs{
		RegionId: common.Shanghai,
		VpcId:    TestVpcId,
		Pagination: common.Pagination{
			PageNumber: 1,
			PageSize:   50,
		},
	}

	ngws, _, err := client.DescribeNatGateways(args)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("NGW = %++v", ngws)
	}
}

func TestClient_DescribeNetworkQuotas(t *testing.T) {
	client := NewVpcTestClientForDebug()
	args := &DescribeNetworkQuotasArgs{
		Product:  "vpc",
		RegionId: TestRegionID,
	}

	response, err := client.DescribeNetworkQuotas(args)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("quota = %++v", response)
	}
}

func TestClient_DescribeEipAddressesWithRaw(t *testing.T) {
	client := NewVpcTestClientForDebug()
	regions, err := client.DescribeRegions()
	if err == nil {
		total := 0
		for _, region := range regions {
			args := &DescribeEipAddressesArgs{
				RegionId: region.RegionId,
				Pagination: common.Pagination{
					PageNumber: 1,
					PageSize:   50,
				},
			}

			eips, page, err := client.DescribeEipAddresses(args)
			if err != nil {
				t.Fatalf("Error %++v", err)
			} else {
				t.Logf("eips = %++v", eips)
				t.Logf("page = %++v", page)
				total += page.TotalCount
			}
		}
		t.Logf("total Eips = %d", total)
	}
}

func TestClient_DescribeVpcsWithRaw(t *testing.T) {
	client := NewVpcTestClientForDebug()
	args := &DescribeVpcsArgs{
		RegionId: common.Beijing,
		Pagination: common.Pagination{
			PageNumber: 1,
			PageSize:   100,
		},
	}

	vpcs, page, err := client.DescribeVpcs(args)
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("vpcs = %++v", vpcs)
		t.Logf("page = %++v", page)
	}
}

func TestClient_DescribeRegions(t *testing.T) {
	client := NewVpcTestClientForDebug()
	regions, err := client.DescribeRegions()
	if err != nil {
		t.Fatalf("Error %++v", err)
	} else {
		t.Logf("Regions = %++v", regions)
	}
}
