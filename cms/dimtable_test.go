package cms

import (
	"encoding/json"
	"testing"
)

func TestCreateDimTable(t *testing.T) {
	dimTable := DimTable{
		DimTableName: "my-dim-talbe",
		Owner:        "menglingwei@aliyun.com",
	}

	response, err := client.CreateDimTable(CreateDimTableRequest{
		DimTable: dimTable,
	})
	if err != nil {
		t.Errorf("Failed to invoke api, %v", err)
	} else {
		t.Logf("response = %v", response)
	}
}

func TestGetDimTable(t *testing.T) {
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
