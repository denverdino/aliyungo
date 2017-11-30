package cms

import (
	"encoding/json"
	"testing"
)

//TODO 暂时不需要创建，参数转换为urlParams有些问题
//func TestCreateDimTable(t *testing.T) {
//	dimTable := DimTable{
//		Owner: "menglingwei@aliyun.com",
//	}

//	response, err := client.CreateDimTable(CreateDimTableRequest{
//		DimTableName: "my-dim-table",
//		DimTable:     dimTable,
//	})
//	if err != nil {
//		t.Errorf("Failed to invoke api, %v", err)
//	} else {
//		t.Logf("response = %v", response)
//	}
//}

func TestGetDimTable(t *testing.T) {
	client := NewTestClientForDebug()
	request := GetDimTableRequest{
		DimTableName: "zhishi_test",
	}
	response, err := client.GetDimTable(request)
	if err != nil {
		t.Errorf("Failed to invoke api, %v", err)
	} else {
		t.Logf("response = %v", response)
	}
}

func TestListDimTable(t *testing.T) {
	request := ListDimTableRequest{
		//DimTableName: "zhishi_test",
		Page:     1,
		PageSize: 10,
	}
	response, err := client.ListDimTable(request)
	if err != nil {
		t.Errorf("Failed to invoke api, %v", err)
	} else {
		t.Logf("response = %v", response)
	}
}

func TestPutDimTableData(t *testing.T) {
	dims := map[string]string{
		"dims": "cluster_id=c07c5c4d6cb6348a9ab2eda27ced6cb9a",
		"key":  "i-25ywxlcze",
	}

	body, _ := json.Marshal(dims)
	request := PutDimTableDataRequest{
		DimTableName: "zhishi_test",
		Body:         string(body),
	}
	response, err := client.PutDimTableData(request)
	if err != nil {
		t.Errorf("Failed to invoke api, %v", err)
	} else {
		t.Logf("response = %v", response)
	}
}

func TestBatchPutDimTableData(t *testing.T) {
	dims := make([]map[string]string, 0)
	dims = append(dims, map[string]string{
		"dims": "cluster_id=c07c5c4d6cb6348a9ab2eda27ced6cb9a",
		"key":  "i-25ywxlcze",
	})
	dims = append(dims, map[string]string{
		"dims": "cluster_id=c07c5c4d6cb6348a9ab2eda27ced6cb9a",
		"key":  "i-25ywxlcze-1",
	})
	dims = append(dims, map[string]string{
		"dims": "cluster_id=c07c5c4d6cb6348a9ab2eda27ced6cb9a",
		"key":  "i-25ywxlcze-2",
	})

	body, _ := json.Marshal(dims)
	request := BatchPutDimTableDataRequest{
		DimTableName: "zhishi_test",
		Body:         string(body),
	}
	response, err := client.BatchPutDimTableData(request)
	if err != nil {
		t.Errorf("Failed to invoke api, %v", err)
	} else {
		t.Logf("response = %v", response)
	}
}

func TestDeleteDimTableData(t *testing.T) {
	request := DeleteDimTableDataRequest{
		DimTableName: "zhishi_test",
		Key:          "i-25ywxlcze",
	}
	response, err := client.DeleteDimTableData(request)
	if err != nil {
		t.Errorf("Failed to invoke api, %v", err)
	} else {
		t.Logf("response = %v", response)
	}
}

func TestQueryDimTableData(t *testing.T) {
	request := QueryDimTableDataRequest{
		DimTableName: "zhishi_test",
		Key:          "i-25ywxlcze",
	}
	response, err := client.QueryDimTableData(request)
	if err != nil {
		t.Errorf("Failed to invoke api, %v", err)
	} else {
		t.Logf("response = %v", response)
	}
}
