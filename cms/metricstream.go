package cms

type CommonMetricStreamRequest struct {
	ProjectName      string
	MetricStreamName string
}

type CreateMetricStreamRequest struct {
	ProjectName  string
	MetricStream MetricStream
}
type MetricStream struct {
}

type CreateMetricStreamResponse struct {
	Response
	Result int64
}

//创建MetricStream
func (client *CmsClient) CreateMetricStream(request CreateMetricStreamRequest) (CreateMetricStreamResponse, error) {
	var resp CreateMetricStreamResponse
	err := client.Invoke("CreateMetricStream", request, &resp)
	if err != nil {
		return CreateMetricStreamResponse{}, err
	}
	return resp, nil
}

type UpdateMetricStreamRequest struct {
	CommonMetricStreamRequest
	MetricStream MetricStream
}

type UpdateMetricStreamResponse struct {
	Response
	MetricStream MetricStream
}

//更新MetricStream
func (client *CmsClient) UpdateMetricStream(request UpdateMetricStreamRequest) (UpdateMetricStreamResponse, error) {
	var resp UpdateMetricStreamResponse
	err := client.Invoke("UpdateMetricStream", request, &resp)
	if err != nil {
		return UpdateMetricStreamResponse{}, err
	}
	return resp, nil
}

type DeleteMetricStreamRequest struct {
	CommonMetricStreamRequest
}

//删除MetricStream
func (client *CmsClient) DeleteMetricStream(request DeleteMetricStreamRequest) (Response, error) {
	var resp Response
	err := client.Invoke("DeleteMetricStream", request, &resp)
	if err != nil {
		return Response{}, err
	}
	return resp, nil
}

type GetMetricStreamRequest struct {
	CommonMetricStreamRequest
}

type GetMetricStreamResponse struct {
	Response
	MetricStream
}

//获取MetricStream
func (client *CmsClient) GetMetricStream(request GetMetricStreamRequest) (GetMetricStreamResponse, error) {
	var resp GetMetricStreamResponse
	err := client.Invoke("GetMetricStream", request, &resp)
	if err != nil {
		return GetMetricStreamResponse{}, err
	}
	return resp, nil
}

type ListMetricStreamsRequest struct {
	CommonMetricStreamRequest
	Page     int
	PageSize int
}

type ListMetricStreamsResponse struct {
	Response
	NextToken  int
	Datapoints []MetricStream
}

//查询MetricStream列表
func (client *CmsClient) ListMetricStreams(request ListMetricStreamsRequest) (ListMetricStreamsResponse, error) {
	var resp ListMetricStreamsResponse
	err := client.Invoke("ListMetricStream", request, &resp)
	if err != nil {
		return ListMetricStreamsResponse{}, err
	}
	return resp, nil
}
