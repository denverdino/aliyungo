package cms

type DimTable struct {
	Id           int64
	DimTableName string
	Owner        string
	ExtendInfo   map[string]interface{} //JSONOBJECT
	GmtCreate    int64
	GmtModified  int64
}

type CreateDimTableRequest struct {
	DimTableName string
	DimTable     DimTable
}

type CreateDimTableResponse struct {
	Response
	Result int64
}

//创建DimTable
func (client *CmsClient) CreateDimTable(request CreateDimTableRequest) (CreateDimTableResponse, error) {
	var resp CreateDimTableResponse
	err := client.Invoke("CreateDimTable", request, &resp)
	if err != nil {
		return CreateDimTableResponse{}, err
	}
	return resp, nil
}

type UpdateDimTableRequest struct {
	DimTableName string
	DimTable     DimTable
}

//修改DimTable
func (client *CmsClient) UpdateDimTable(request UpdateDimTableRequest) (Response, error) {
	var resp Response
	err := client.Invoke("UpdateDimTable", request, &resp)
	if err != nil {
		return Response{}, err
	}
	return resp, nil
}

type DeleteDimTableRequest struct {
	DimTableName string
}

//删除DimTable
func (client *CmsClient) DeleteDimTable(request DeleteDimTableRequest) (Response, error) {
	var resp Response
	err := client.Invoke("DeleteDimTable", request, &resp)
	if err != nil {
		return Response{}, err
	}
	return resp, nil
}

type GetDimTableRequest struct {
	DimTableName string
}

type GetDimTableResponse struct {
	Response
	Result DimTable
}

//查询DimTable
func (client *CmsClient) GetDimTable(request GetDimTableRequest) (GetDimTableResponse, error) {
	var resp GetDimTableResponse
	err := client.Invoke("GetDimTable", request, &resp)
	if err != nil {
		return GetDimTableResponse{}, err
	}
	return resp, nil
}

type ListDimTableRequest struct {
	DimTableName string
	Page         int
	PageSize     int
}

type ListDimTableResponse struct {
	Response
	NextToken  int
	Datapoints Datapoints
}

type Datapoints struct {
	Datapoint []DimTable
}

//查询DimTable列表
func (client *CmsClient) ListDimTable(request ListDimTableRequest) (ListDimTableResponse, error) {
	var resp ListDimTableResponse
	err := client.Invoke("ListDimTable", request, &resp)
	if err != nil {
		return ListDimTableResponse{}, err
	}
	return resp, nil
}

type PutDimTableDataRequest struct {
	DimTableName string
	Body         string //JSONString
}

//上传DimTable数据
func (client *CmsClient) PutDimTableData(request PutDimTableDataRequest) (Response, error) {
	var resp Response
	err := client.Invoke("PutDimTableData", request, &resp)
	if err != nil {
		return Response{}, err
	}
	return resp, nil
}

type BatchPutDimTableDataRequest struct {
	DimTableName string
	Body         string //JSONArray
}

//批量上传DimTable数据
func (client *CmsClient) BatchPutDimTableData(request BatchPutDimTableDataRequest) (Response, error) {
	var resp Response
	err := client.Invoke("BatchPutDimTableData", request, &resp)
	if err != nil {
		return Response{}, err
	}
	return resp, nil
}

type QueryDimTableDataRequest struct {
	DimTableName string
	Key          string
}

type QueryDimTableDataResponse struct {
	Response
	Result map[string]interface{} //JSONObject
}

//查询DimTable数据
func (client *CmsClient) QueryDimTableData(request QueryDimTableDataRequest) (QueryDimTableDataResponse, error) {
	var resp QueryDimTableDataResponse
	err := client.Invoke("QueryDimTableData", request, &resp)
	if err != nil {
		return QueryDimTableDataResponse{}, err
	}
	return resp, nil
}

type DeleteDimTableDataRequest struct {
	DimTableName string
	Key          string
}

//删除DimTable的data
func (client *CmsClient) DeleteDimTableData(request DeleteDimTableDataRequest) (Response, error) {
	var resp Response
	err := client.Invoke("DeleteDimTableData", request, &resp)
	if err != nil {
		return Response{}, err
	}
	return resp, nil
}

type QueryDimTableDataByKeyRequest struct {
	DimTableName string
	Key          string
	value        string
	Page         int
	PageSize     int
}

type QueryDimTableDataByKeyResponse struct {
	Response
	NextToken  int
	Datapoints []string
}

//根据dims(key:value)查询DimTable key列表
func (client *CmsClient) QueryDimTableDataByKey(request QueryDimTableDataByKeyRequest) (QueryDimTableDataByKeyResponse, error) {
	var resp QueryDimTableDataByKeyResponse
	err := client.Invoke("QueryDimTableDataByKey", request, &resp)
	if err != nil {
		return QueryDimTableDataByKeyResponse{}, err
	}
	return resp, nil
}

type DeleteDimTableDataByKeyRequest struct {
	DimTableName string
	Key          string
	value        string
}

//删除Dims
func (client *CmsClient) DeleteDimTableDataByKey(request DeleteDimTableDataByKeyRequest) (Response, error) {
	var resp Response
	err := client.Invoke("DeleteDimTableDataByKey", request, &resp)
	if err != nil {
		return Response{}, err
	}
	return resp, nil
}
