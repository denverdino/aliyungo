package cms

import (
	"os"
	"testing"
)

var (
	accessKeyId       = os.Getenv("accessKeyId")
	accessKeySecret   = os.Getenv("accessKeySecret")
	defaultCmsVersion = "2017-03-01"
	defaultAlarmName  = "ACS_SYSTEM_LOAD"
)

func TestListAlarm(t *testing.T) {
	req := ListAlarmRequest{
		Namespace:  "acs_containerservice_dashboard",
		PageNumber: 1,
		PageSize:   10,
	}

	client := NewClientWithVersion(accessKeyId, accessKeySecret, defaultCmsVersion)
	resp, err := client.ListAlarm(req)
	if err != nil {
		t.Errorf("Failed to list alarm %v", err.Error())
		return
	}
	t.Logf("pass list alarm %v", resp)
}

func TestCreateAlarm(t *testing.T) {
	req := CreateAlarmRequest{
		Name:               defaultAlarmName,
		Namespace:          "acs_containerservice_dashboard",
		MetricName:         "GPUMemoryUsed",
		Dimensions:         "[{\"clusterId\":\"c6d1e7ccf550f420bb0926738d3411298\",\"instanceId\":\"i-uf6hx8v6lzadam4pfom3\"}]",
		Statistics:         "Average",
		ComparisonOperator: "<=",
		Threshold:          "30",
		ContactGroups:      "[\"云账号报警联系人\"]",
	}

	client := NewClientWithVersion(accessKeyId, accessKeySecret, defaultCmsVersion)
	resp, err := client.CreateAlarm(req)
	if err != nil {
		t.Error("Failed to create alaram:" + err.Error())
	}
	t.Logf("pass unit tests %v", resp)

}
