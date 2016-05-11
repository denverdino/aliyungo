package cms

type CommonMetricsRequest struct {
	ProjectName string
	IsPublic    int
}

type Metrics struct {
}

type CreateMetricsRequest struct {
	CommonMetricsRequest
	Metrics
}

type CreateMetricsResponse struct {
	Response
	Result int64
}

//创建Metrics
func (client *CmsClient) CreateMetrics(request CreateMetricsRequest) (CreateMetricsResponse, error) {
	var resp CreateMetricsResponse
	err := client.Invoke("CreateMetrics", request, &resp)
	if err != nil {
		return CreateMetricsResponse{}, err
	}
	return resp, nil
}

type CreateSQLMetricsRequest struct {
	CommonMetricsRequest
	Sql string
}

//创建SQLMetrics
func (client *CmsClient) CreateSQLMetrics(request CreateSQLMetricsRequest) (CreateMetricsResponse, error) {
	var resp CreateMetricsResponse
	err := client.Invoke("CreateSQLMetrics", request, &resp)
	if err != nil {
		return CreateMetricsResponse{}, err
	}
	return resp, nil
}

type BatchCreateSQLMetricsRequest struct {
	CommonMetricsRequest
	Sqls []string
}

type BatchCreateSQLMetricsResponse struct {
	Response
	Result []int64
}

//批量创建SQLMetrics
func (client *CmsClient) BatchCreateSqlMetrics(request BatchCreateSQLMetricsRequest) (BatchCreateSQLMetricsResponse, error) {
	var resp BatchCreateSQLMetricsResponse
	err := client.Invoke("BatchCreateSqlMetrics", request, &resp)
	if err != nil {
		return BatchCreateSQLMetricsResponse{}, err
	}
	return resp, nil
}

type UpdateSQLMetricsRequest struct {
	CommonMetricsRequest
	MetricName string
	Sql        string
}

//修改SQLMetrics
func (client *CmsClient) UpdateSQLMetrics(request UpdateSQLMetricsRequest) (Response, error) {
	var resp Response
	err := client.Invoke("UpdateSQLMetrics", request, &resp)
	if err != nil {
		return Response{}, err
	}
	return resp, nil
}

type GetSQLMetricsRequest struct {
	ProjectName string
	MetricName  string
}

type GetSQLMetricsResponse struct {
	Response
	Result string
}

//查询SQLMetrics
func (client *CmsClient) GetSQLMetrics(request GetSQLMetricsRequest) (GetSQLMetricsResponse, error) {
	var resp GetSQLMetricsResponse
	err := client.Invoke("GetSQLMetrics", request, &resp)
	if err != nil {
		return GetSQLMetricsResponse{}, err
	}
	return resp, nil
}

type ListSQLMetricsRequest struct {
	ProjectName string
	MetricName  string
	Offset      int
	Size        int
}

type ListSQLMetricsResponse struct {
	NextToken  int
	Datapoints []string
}

//查询SQLMetrics列表
func (client *CmsClient) ListSQLMetrics(request ListSQLMetricsRequest) (ListSQLMetricsResponse, error) {
	var resp ListSQLMetricsResponse
	err := client.Invoke("ListSQLMetrics", request, &resp)
	if err != nil {
		return ListSQLMetricsResponse{}, err
	}
	return resp, nil
}

type DeleteMetricsRequest struct {
	ProjectName string
	MetricName  string
}

//删除Metrics
func (client *CmsClient) DeleteMetrics(request DeleteMetricsRequest) (Response, error) {
	var resp Response
	err := client.Invoke("DeleteMetrics", request, &resp)
	if err != nil {
		return Response{}, err
	}
	return resp, nil
}

type GetMetricsMetaRequest struct {
	ProjectName string
	MetricName  string
}

type GetMetricsMetaResponse struct {
	Response
	Result []string
}

//查询Metric的Meta信息
func (client *CmsClient) GetMetricsMeta(request GetMetricsMetaRequest) (GetMetricsMetaResponse, error) {
	var resp GetMetricsMetaResponse
	err := client.Invoke("GetMetricsMeta", request, &resp)
	if err != nil {
		return GetMetricsMetaResponse{}, err
	}
	return resp, nil
}
