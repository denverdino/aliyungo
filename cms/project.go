package cms

import (
	"time"
)

type Project struct {
	ID              int64
	ProjectName     string
	ProjectDesc     string
	ProjectOwner    string
	GmtCreate       time.Time
	GmtModified     time.Time
	Creator         string
	Status          int
	PublicMetricNum int
	Run             RunStatus
	Msg             string
	GalaxyConf      GalaxyConf
	Galaxy          string
}

type GalaxyConf struct {
	InstanceNum int
	Coefficient float64
	WorkNum     int
	SourceNum   int
	HashNum     int
	GroupNum    int
	CrossNum    int
}

type CreateProjectResponse struct {
	Response
	Result int64
}

//创建project
func (client *CmsClient) CreateProject(project Project) (CreateProjectResponse, error) {
	var projectResp CreateProjectResponse
	err := client.Invoke("CreateProject", project, &projectResp)
	if err != nil {
		return CreateProjectResponse{}, err
	}

	return projectResp, nil
}

type CommonProjectRequest struct {
	ProjectName string
}

type GetProjectRequest struct {
	CommonProjectRequest
}

type GetProjectResponse struct {
	Response
	Result Project
}

//获取project详情
func (client *CmsClient) GetProject(request GetProjectRequest) (GetProjectResponse, error) {
	var projectResp GetProjectResponse
	err := client.Invoke("GetProject", request, &projectResp)
	if err != nil {
		return GetProjectResponse{}, err
	}
	return projectResp, nil
}

type UpdateProjectRequest struct {
	Project
	ProjectName string
}

//更新project
func (client *CmsClient) UpdateProject(request UpdateProjectRequest) (Response, error) {
	var resp Response
	err := client.Invoke("UpdateProject", request, &resp)
	if err != nil {
		return Response{}, err
	}
	return resp, nil
}

type DeleteProjectRequest struct {
	CommonProjectRequest
}

//删除project
func (client *CmsClient) DeleteProject(request DeleteProjectRequest) (Response, error) {
	var resp Response
	err := client.Invoke("DeleteProject", request, &resp)
	if err != nil {
		return Response{}, err
	}
	return resp, nil
}

type ListProjectsRequest struct {
	ProjectOwner string
	ProjectName  string
	Page         int
	PageSize     int
}

type ListProjectsResponse struct {
	Response
	NextToken  int
	Datapoints []Project
}

//获取project列表
func (client *CmsClient) ListProjects(request ListProjectsRequest) (ListProjectsResponse, error) {
	var resp ListProjectsResponse
	err := client.Invoke("ListProject", request, &resp)
	if err != nil {
		return ListProjectsResponse{}, err
	}
	return resp, nil
}

type StatusProjectRequest struct {
	CommonProjectRequest
}

type StatusProjectResponse struct {
	Response
	Result GalaxyStatus
}

type GalaxyStatus struct {
	Status RunStatus
	Msg    string
}

//查询project的运行状态
func (client *CmsClient) StatusProject(request StatusProjectRequest) (StatusProjectResponse, error) {
	var resp StatusProjectResponse
	err := client.Invoke("StatusProject", request, &resp)
	if err != nil {
		return StatusProjectResponse{}, err
	}
	return resp, nil
}

type StartProjectRequest struct {
	CommonProjectRequest
	Version string
	Type    string
}

//启动project
func (client *CmsClient) StartProject(request StartProjectRequest) (Response, error) {
	var resp Response
	err := client.Invoke("StartProject", request, &resp)
	if err != nil {
		return Response{}, err
	}
	return resp, nil
}

type StopProjectRequest struct {
	CommonProjectRequest
}

//停止project
func (client *CmsClient) StopProject(request StopProjectRequest) (Response, error) {
	var resp Response
	err := client.Invoke("StopProject", request, &resp)
	if err != nil {
		return Response{}, err
	}
	return resp, nil
}
