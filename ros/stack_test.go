package ros

import (
	"testing"

	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	myTestTemplate = `
{
  "ROSTemplateFormatVersion": "2015-09-01",
  "Resources": {
    "string3": {
      "Type": "ALIYUN::RandomString",
      "DependsOn": [
        "string2"
      ],
      "Properties": {
        "sequence": "octdigits",
        "length": 100
      }
    },
    "string1": {
      "Test1": "Hello",
      "Type": "ALIYUN::RandomString"
    },
    "string2": {
      "Type": "ALIYUN::RandomString",
      "DependsOn": [
        "string1"
      ],
      "Properties": {
        "sequence": "octdigits",
        "length": 100
      }
    }
  }
}
	`
)

func TestClient_CreateStack(t *testing.T) {
	p := map[string]interface{}{
		"VpcId":              "vpc-uf63x7z49v37kcazrbmca",
		"VSwitchId":          "vsw-uf66bi1fvhf1kyv8v4kvl",
		"NatGateway":         false,
		"K8SExportAddress":   "139.224.136.43",
		"ImageId":            "centos_7",
		"MasterInstanceType": "ecs.n4.large",
		"WorkerInstanceType": "ecs.n4.large",
		"NumOfNodes":         2,
		"LoginPassword":      "Hello1234",
	}

	tmpl, err := loadK8s()
	if err != nil {
		t.Fatalf("Failed to load k8s file %++v", err)
	}

	//ps, _ := json.Marshal(p)
	args := &CreateStackRequest{
		Name:            fmt.Sprintf("my-k8s-test-stack-%d", time.Now().Unix()),
		Template:        tmpl,
		Parameters:      p,
		DisableRollback: false,
		TimeoutMins:     30,
	}

	response, err := debugClientForTestCase.CreateStack(TestRegionId, args)
	if err != nil {
		t.Fatalf("Failed to CreateStack %++v", err)
	} else {
		t.Logf("Success %++v", response)
	}
}

func TestClient_DeleteStack(t *testing.T) {
	stackName := os.Getenv("StackName")
	stackId := os.Getenv("StackId")

	response, err := debugClientForTestCase.DeleteStack(TestRegionId, stackId, stackName)
	if err != nil {
		t.Fatalf("Failed to DeleteStack %++v", err)
	} else {
		t.Logf("Success %++v", response)
	}
}

func TestClient_AbandonStack(t *testing.T) {
	stackName := os.Getenv("StackName")
	stackId := os.Getenv("StackId")

	response, err := debugClientForTestCase.AbandonStack(TestRegionId, stackId, stackName)
	if err != nil {
		t.Fatalf("Failed to AbandonStack %++v", err)
	} else {
		t.Logf("Success %++v", response)
	}
}

func TestClient_DescribeStacks(t *testing.T) {
	args := &DescribeStacksRequest{
		RegionId: TestRegionId,
	}

	stacks, err := debugClientForTestCase.DescribeStacks(args)
	if err != nil {
		t.Fatalf("Failed to DescribeStacks %++v", err)
	} else {
		t.Logf("Response is %++v", stacks)
	}
}

func TestClient_DescribeStack(t *testing.T) {
	stackName := os.Getenv("StackName")
	stackId := os.Getenv("StackId")

	response, err := debugClientForTestCase.DescribeStack(TestRegionId, stackId, stackName)
	if err != nil {
		t.Fatalf("Failed to DescribeStack %++v", err)
	} else {
		t.Logf("Success %++v", response)
	}
}

func TestClient_UpdateStack(t *testing.T) {
	stackName := os.Getenv("StackName")
	stackId := os.Getenv("StackId")

	p := map[string]interface{}{
		"VpcId":              "vpc-uf63x7z49v37kcazrbmca",
		"VSwitchId":          "vsw-uf66bi1fvhf1kyv8v4kvl",
		"NatGateway":         false,
		"K8SExportAddress":   "139.224.136.43",
		"ImageId":            "centos_7",
		"MasterInstanceType": "ecs.n4.large",
		"WorkerInstanceType": "ecs.n4.large",
		"NumOfNodes":         3,
		"LoginPassword":      "Hello1234",
	}

	tmpl, err := loadK8s()
	if err != nil {
		t.Fatalf("Failed to load k8s file %++v", err)
	}

	//ps, _ := json.Marshal(p)
	args := &UpdateStackRequest{
		Template:        tmpl,
		Parameters:      p,
		DisableRollback: false,
		TimeoutMins:     30,
	}

	response, err := debugClientForTestCase.UpdateStack(TestRegionId, stackId, stackName, args)
	if err != nil {
		t.Fatalf("Failed to UpdateStack %++v", err)
	} else {
		t.Logf("Success %++v", response)
	}
}

func loadK8s() (string, error) {
	b, err := ioutil.ReadFile("k8s.json")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
