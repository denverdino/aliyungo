package grandcanal

import (
	"github.com/denverdino/aliyungo/common"
)

// CreateActivityTypeArgs represents arguments to create an activity type
type CreateActivityTypeArgs struct {
	Domain                  string
	ActivityTypeName        string
	ActivityTypeVersion     string
	Description             string
	DefaultTimeout          int64
	TimeoutPolicy           int32
	DefaultHeartbeatTimeout int64
	Priority                int32
}

type CreateActivityTypeResponse struct {
	common.Response
}

// CreateActivityType create a new activity type
func (client *Client) CreateActivityType(args *CreateActivityTypeArgs) (requestId string, err error) {
	response := &CreateActivityTypeResponse{}
	err = client.Invoke("CreateActivityType", args, &response)
	if err != nil {
		return "", err
	}
	return response.RequestId, nil
}

// DescribeActivityTypeArgs represents arguments to describe activity type
type DescribeActivityTypeArgs struct {
	Domain              string
	ActivityTypeName    string
	ActivityTypeVersion string
}

type DescribeActivityTypeResponse struct {
	common.Response
	ActivityType ActivityType
}

type ActivityType struct {
	ActivityTypeName        string
	ActivityTypeVersion     string
	Description             string
	Domain                  string
	CreateTime              string
	DisabledTime            string
	Status                  string
	DefaultTimeout          int64
	TimeoutPolicy           int32
	DefaultHeartbeatTimeout int64
	Priority                int32
}

// DescribeActivityType describe the detailed information of an activity type
func (client *Client) DescribeActivityType(args *DescribeActivityTypeArgs) (activityType *ActivityType, err error) {
	response := &DescribeActivityTypeResponse{}
	err = client.Invoke("DescribeActivityType", args, &response)
	if err != nil {
		return nil, err
	}
	return &response.ActivityType, nil
}

// DeleteActivityTypeArgs represents arguments to delete an activity type
type DeleteActivityTypeArgs struct {
	Domain              string
	ActivityTypeName    string
	ActivityTypeVersion string
}

type DeleteActivityTypeResponse struct {
	common.Response
}

// DeleteActivityType delete an activity type object
func (client *Client) DeleteActivityType(args *DeleteActivityTypeArgs) (requestId string, err error) {
	response := &DescribeActivityTypeResponse{}
	err = client.Invoke("DeleteActivityType", args, &response)
	if err != nil {
		return "", err
	}
	return response.RequestId, nil
}
