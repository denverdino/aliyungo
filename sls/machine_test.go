package sls

import (
	"fmt"
	"testing"
)

func TestListMachineGroups(t *testing.T) {
	p := DefaultProject()
	groups, err := p.MachineGroups(0, 100)
	if err != nil {
		fmt.Printf("Error in list groups %s \n", err)
		t.FailNow()
	}

	fmt.Println(groups)
}

func TestMachineGroup(t *testing.T) {
	p := DefaultProject()
	group, err := p.MachineGroup("qunqi")
	if err != nil {
		fmt.Printf("Error in TestMachineGroup %s \n", err)
		t.FailNow()
	}

	fmt.Println(group)
}

func TestCreateMachineGroup(t *testing.T) {
	p := DefaultProject()
	groupName := "testGroup"
	group := &MachineGroup{
		Name:                groupName,
		MachineIdentifyType: "ip",
		Attribute:           &GroupAttribute{},
		MachineList: []string{
			"127.0.0.1",
			"127.0.0.2",
		},
	}
	err := p.CreateMachineGroup(group)
	if err != nil {
		t.Fatalf("Create machine error: %v", err)
	}

	mg, err := p.MachineGroup(groupName)
	if err != nil {
		t.Fatalf("Find machine error: %v", err)
	}

	if err = mg.Delete(); err != nil {
		t.Fatalf("Delete machine error: %v", err)
	}
}

func TestListMachines(t *testing.T) {
	p := DefaultProject()
	mg, err := p.MachineGroup("qunqi")
	if err != nil {
		t.Fatalf("Find machineGroup error %v\n", err)
	}

	_, err = mg.ListMachines(0, 100)
	if err != nil {
		t.Fatalf("List machine error %v\n", err)
	}
}
