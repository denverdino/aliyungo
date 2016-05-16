package cms

type MetricQuery struct {
}

type QueryMetricDataRequest struct {
	Project    string
	Metric     string
	Period     string
	StartTime  string
	EndTime    string
	Dimensions string
	Express    string
	Length     string
}

type QueryMetricDataResponse struct {
	Response
	Period     string
	Datapoints []map[string]interface{} //List<JSONObject>
}

//查询MetricData数据
func (client *CmsClient) QueryMetricData(request QueryMetricDataRequest) (QueryMetricDataResponse, error) {
	var resp QueryMetricDataResponse
	err := client.Invoke("QueryMetricData", request, &resp)
	if err != nil {
		return QueryMetricDataResponse{}, err
	}
	return resp, nil
}

type QueryMetricLastRequest struct {
	Project    string
	Metric     string
	Period     string
	StartTime  string
	EndTime    string
	Dimensions string
	Express    string
	Page       string
	Length     string
	Cursor     string
}

type QueryMetricLastResponse struct {
	Response
	Period     string
	Cursor     string
	Size       int
	NextToken  int
	Datapoints []map[string]interface{} //List<JSONObject>
}

//查询Metric下各Dimension的最新一条数据
func (client *CmsClient) QueryMetricLast(request QueryMetricLastRequest) (QueryMetricLastResponse, error) {
	var resp QueryMetricLastResponse
	err := client.Invoke("QueryMetricLast", request, &resp)
	if err != nil {
		return QueryMetricLastResponse{}, err
	}
	return resp, nil
}

type QueryMetricTopRequest struct {
	Project    string
	Metric     string
	Period     string
	StartTime  string
	EndTime    string
	Dimensions string
	Express    string
	OrderBy    string
}

type QueryMetricTopResponse struct {
	Period     string
	Datapoints []map[string]interface{} //List<JSONObject>
}

//查询Metric下各Dimension按指定字段排序后顶部的数据
func (client *CmsClient) QueryMetricTop(request QueryMetricTopRequest) (QueryMetricTopResponse, error) {
	var resp QueryMetricTopResponse
	err := client.Invoke("QueryMetricTop", request, &resp)
	if err != nil {
		return QueryMetricTopResponse{}, err
	}
	return resp, nil
}

type DescribeMetricRequest struct {
	Project    string
	Metric     string
	Period     string
	StartTime  string
	EndTime    string
	Dimensions string
}

type DescribeMetricResponse struct {
	Response
	Datapoints []map[string]interface{} //List<JSONObject>
}

//查询Metric下各Dimension的描述信息
func (client *CmsClient) DescribeMetric(request DescribeMetricRequest) (DescribeMetricResponse, error) {
	var resp DescribeMetricResponse
	err := client.Invoke("DescribeMetric", request, &resp)
	if err != nil {
		return DescribeMetricResponse{}, err
	}
	return resp, nil
}

type QueryMetricRequest struct {
	Project    string
	Metric     string
	Period     string
	StartTime  string
	EndTime    string
	Dimensions string
	Page       string
	Length     string
	Extend     string
}

type QueryMetricResponse struct {
	Response
	NextToken  string
	Datapoints []map[string]interface{} //List<JSONObject>
	Datapoint  string
}

//QueryMetric-查询Metric数据
func (client *CmsClient) QueryMetric(request QueryMetricRequest) (QueryMetricResponse, error) {
	var resp QueryMetricResponse
	err := client.Invoke("QueryMetric", request, &resp)
	if err != nil {
		return QueryMetricResponse{}, err
	}
	return resp, nil
}

type BatchQueryMetricRequest struct {
	Project    string
	Metric     string
	Period     string
	StartTime  string
	EndTime    string
	Dimensions string
	Page       string
	Length     string
}

//BatchQueryMetric-批量查询Metric数据
func (client *CmsClient) BatchQueryMetric(request BatchQueryMetricRequest) (QueryMetricResponse, error) {
	var resp QueryMetricResponse
	err := client.Invoke("BatchQueryMetric", request, &resp)
	if err != nil {
		return QueryMetricResponse{}, err
	}
	return resp, nil
}

type QueryListMetricRequest struct {
	Project    string
	Metric     string
	Period     string
	StartTime  string
	EndTime    string
	Dimensions string
	Page       string
	Length     string
	Extend     string
	Filter     string
}

//ListMetric-查询Metric下各Dimension的最新数据一条
func (client *CmsClient) QueryListMetric(request QueryListMetricRequest) (QueryMetricResponse, error) {
	var resp QueryMetricResponse
	err := client.Invoke("QueryListMetric", request, &resp)
	if err != nil {
		return QueryMetricResponse{}, err
	}
	return resp, nil
}
