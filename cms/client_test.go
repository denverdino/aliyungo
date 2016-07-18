package cms

import (
	"github.com/denverdino/aliyungo/common"
	"testing"
)

const (
	AccessKeyId      = ""
	AccessKeySecret  = ""
	Region           = common.Hangzhou
)

func TestCresateAlert(t *testing.T)  {
	client := NewClient( AccessKeyId, AccessKeySecret )


	req := `
	{
    "actions":{
        "alertActions":[
        {
                "contactGroups":[],
                "httpNotifyParam":{
                    "type":"http",
                    "method":"GET",
                    "url":"https://cs.console.aliyun.com/hook/trigger?triggerUrl=YzcwMWYyZGE1MWVmZDRjNzg4MWRiMWEwZjIyMmZmZGMxfHdvcmRwcmVzcy1kZWZhdWx0fHNjYWxpbmd8MThucHU3dWo4Nml1Znx7InNlcnZpY2VfaWQiOiJ3b3JkcHJlc3MtZGVmYXVsdF93ZWIifQ==&secret=535651705256494f56325a386f5878466e218fdaa8a86e880b22f9003e801567&type=scale_out&step=1"
                },
                "level":4,

            }
        ],
        "effective":"* * 8-22 * * ?",
        "failure":{
            "contactGroups":["云账号报警联系人"],
            "id":"failActionID"
        },
        "ok":{
            "contactGroups":[]
        },
        "silence":"120"
    },
    "condition":{
        "metricName":"CpuUtilization",
        "project":"acs_containerservice",
        "sourceType":"METRIC",
        "dimensionKeys":["userId","clusterId","serviceId"]
    },
    "deepDives":[
        {
            "text":"您的站点信息如下："
        },
        {
            "condition":{
                "metricName":"CpuUtilization"
            }
        }
    ],
    "enable":true,
    "escalations":[
        {
            "expression":"$Average>0.7",
            "level":4,
            "times":1
        }
    ],
    "interval":120,
    "name":"test_alert2",
    "template":true
}
	`

	result, err := client.CreateAlert4Json("acs_custom_1047840662545713", req)
	if err != nil {
		t.Errorf("CreateAlert encounter error: %v \n", err)
	}
	t.Logf("CreateAlert result : %++v %v \n ", result, err)

	dimension := DimensionRequest{
		UserId      :"1047840662545713",
		AlertName   :"test_alert2",
		Dimensions  : "{\"userId\":\"1047840662545713\",\"clusterId\":\"cc2256150655c43d7bcd4054bfc256c45\",\"serviceId\":\"acsmonitoring_acs-monitoring-agent\"}",
	}
	result , err = client.CreateAlertDimension( "acs_custom_1047840662545713", dimension )
	if err != nil {
		t.Errorf("CreateAlertDimension encounter error: %v \n", err)
	}
	t.Logf("CreateAlertDimension result : %++v  \n ", result)

	result2, err2 := client.GetAlert("acs_custom_1047840662545713", "test_alert2")
	if err2 != nil {
		t.Errorf("GetAlertList encounter error: %v \n", err2)
	}
	t.Logf("GetAlert result : %++v %v \n ", result2, err2)
    
        
}

func TestDeleteAlert(t *testing.T) {
    client := NewClient(AccessKeyId, AccessKeySecret)
    
    result, err := client.GetDimensions( "acs_custom_1047840662545713", "test_alert2" )
    t.Logf("GetDimensionsRequest result : %++v %v \n ", result, err)
    //client.DeleteAlert( "acs_custom_1047840662545713",  "test_alert2")
}
