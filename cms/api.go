package cms

type CmsClientInterface interface {
	//project
	CreateProject(project Project) (CreateProjectResponse, error)
	UpdateProject(request UpdateProjectRequest) (Response, error)
	DeleteProject(request DeleteProjectRequest) (Response, error)
	GetProject(request GetProjectRequest) (GetProjectResponse, error)
	ListProjects(request ListProjectsRequest) (ListProjectsResponse, error)
	StatusProject(request StatusProjectRequest) (StatusProjectResponse, error)
	StartProject(request StartProjectRequest) (Response, error)
	StopProject(request StopProjectRequest) (Response, error)

	//MetricStream
	CreateMetricStream(request CreateMetricStreamRequest) (CreateMetricStreamResponse, error)
	UpdateMetricStream(request UpdateMetricStreamRequest) (UpdateMetricStreamResponse, error)
	DeleteMetricStream(request DeleteMetricStreamRequest) (Response, error)
	GetMetricStream(request GetMetricStreamRequest) (GetMetricStreamResponse, error)
	ListMetricStreams(request ListMetricStreamsRequest) (ListMetricStreamsResponse, error)

	//Metrics
	CreateMetrics(request CreateMetricsRequest) (CreateMetricsResponse, error)
	CreateSQLMetrics(request CreateSQLMetricsRequest) (CreateMetricsResponse, error)
	BatchCreateSqlMetrics(request BatchCreateSQLMetricsRequest) (BatchCreateSQLMetricsResponse, error)
	UpdateSQLMetrics(request UpdateSQLMetricsRequest) (Response, error)
	GetSQLMetrics(request GetSQLMetricsRequest) (GetSQLMetricsResponse, error)
	ListSQLMetrics(request ListSQLMetricsRequest) (ListSQLMetricsResponse, error)
	DeleteMetrics(request DeleteMetricsRequest) (Response, error)
	GetMetricsMeta(request GetMetricsMetaRequest) (GetMetricsMetaResponse, error)

	//DimTable
	CreateDimTable(request CreateDimTableRequest) (CreateDimTableResponse, error)
	UpdateDimTable(request UpdateDimTableRequest) (Response, error)
	DeleteDimTable(request DeleteDimTableRequest) (Response, error)
	GetDimTable(request GetDimTableRequest) (GetDimTableResponse, error)
	ListDimTable(request ListDimTableRequest) (ListDimTableResponse, error)
	PutDimTableData(request PutDimTableDataRequest) (Response, error)
	BatchPutDimTableData(request BatchPutDimTableDataRequest) (Response, error)
	QueryDimTableData(request QueryDimTableDataRequest) (QueryDimTableDataResponse, error)
	DeleteDimTableData(request DeleteDimTableDataRequest) (Response, error)
	QueryDimTableDataByKey(request QueryDimTableDataByKeyRequest) (QueryDimTableDataByKeyResponse, error)
	DeleteDimTableDataByKey(request DeleteDimTableDataByKeyRequest) (Response, error)

	//MetricQuery
	QueryMetricData(request QueryMetricDataRequest) (QueryMetricDataResponse, error)
	QueryMetricLast(request QueryMetricLastRequest) (QueryMetricLastResponse, error)
	QueryMetricTop(request QueryMetricTopRequest) (QueryMetricTopResponse, error)
	DescribeMetric(request DescribeMetricRequest) (DescribeMetricResponse, error)
	QueryMetric(request QueryMetricRequest) (QueryMetricResponse, error)
	BatchQueryMetric(request BatchQueryMetricRequest) (QueryMetricResponse, error)
	QueryListMetric(request QueryListMetricRequest) (QueryMetricResponse, error)
}
