package cms

var client = NewTestClient()

//TODO project暂时不需要创建
//func TestCreateProject(t *testing.T) {
//	galaxyConf := GalaxyConf{
//		InstanceNum: 100,
//		Coefficient: 1000,
//	}
//	project := Project{
//		ProjectName:  "mytest-project-001",
//		ProjectDesc:  "mytestproject",
//		ProjectOwner: "menglingwei@aliyun.com",
//		GalaxyConf:   galaxyConf,
//	}

//	response, err := client.CreateProject(project)
//	if err != nil {
//		t.Errorf("Failed to create project, %v", err)
//	} else {
//		t.Logf("response = %v", response)
//	}
//}
