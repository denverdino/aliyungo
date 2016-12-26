package grandcanal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/denverdino/aliyungo/common"
)

type WorkflowDecider struct {
	Metadata *SequenceFlow
	Client   *Client
}

func (decider *WorkflowDecider) HandleEvent(response *PollForWorkflowTaskResponse) error {
	metadata := decider.Metadata
	client := decider.Client
	if metadata == nil || client == nil {
		return errors.New("metadata and grandcanal client can not be null when handle workflow event.")
	}

	fmt.Printf("TaskType = %s token: %s \n", response.TaskType, response.TaskToken )
	if strings.EqualFold(response.TaskType, "WORKFLOW_START") {
		return decider.handleWorkflowStartEvent(response)
	} else if strings.EqualFold(response.TaskType, "ACTIVITY_COMPLETE") {
		return decider.handleActivityCompleteEvent(response)
	} else if strings.EqualFold(response.TaskType, "WORKFLOW_RESUME") {
		return decider.handleWorkflowResumeEvent(response)
	}
	return nil
}

func (decider *WorkflowDecider) handleWorkflowStartEvent(response *PollForWorkflowTaskResponse) error {
	metadata := decider.Metadata
	client := decider.Client
	firstActivity := getFirstActivity(metadata)
	fmt.Printf("handle workflow start event, the first activity is %s\n", firstActivity.Name)

	taskToken := response.TaskToken
	pollTaskAttribute := response.TaskAttribute
	input := pollTaskAttribute.Input

	// respond workflow task, start an activity instance
	taskAttributes := RespondTaskAttributes{
		[]RespondTaskAttribute{
			{
				TaskType:            "ACTIVITY_START",
				ActivityTypeName:    firstActivity.Action,
				ActivityTypeVersion: firstActivity.Version,
				Name:                firstActivity.Name,
				Method:              "INVOKE",
				Input:               input,
				Timeout:             0,
				HeartbeatTimeout:    0,
				Delay:               0,
			},
		},
	}
	respondWorkflowTaskArgs := RespondWorkflowTaskArgs{
		TaskToken:      taskToken,
		Identity:       "127.0.0.1",
		TaskAttributes: taskAttributes,
	}
	_, err := client.RespondWorkflowTask(&respondWorkflowTaskArgs)
	if err != nil {
		fmt.Printf("Failed to respond workflow task %s: %v\n", taskToken, err)
		return err
	}
	return nil
}

func (decider *WorkflowDecider) queryActivityInstances(workflowId string) []ActivityInstance {
	totalPage := 1
	result := []ActivityInstance{}
	client := decider.Client
	for page := 1; page <= totalPage; page++ {
		queryActivityInstancesArgs := QueryActivityInstancesArgs{
			WorkflowId: workflowId,
			Pagination: common.Pagination{
				PageNumber: 1,
				PageSize:   100,
			},
		}
		response, err := client.QueryActivityInstances(&queryActivityInstancesArgs)
		if err != nil {
			fmt.Printf("Failed to query activity instance %s: %v", workflowId, err)
			return result
		}
		activities := response.ActivityInstances.ActivityInstance
		fmt.Printf("Activity number is %d\n", len(activities))
		result = append(result, activities...)
		totalPage = (response.PaginationResult.TotalCount / 100) + 1
		fmt.Printf("Total page is %d\n", totalPage)
	}
	return result
}

func (decider *WorkflowDecider) handleWorkflowResumeEvent(response *PollForWorkflowTaskResponse) error {
	metadata := decider.Metadata
	client := decider.Client
	taskToken := response.TaskToken
	workflowId := response.WorkflowId

	// query activity execution history
	activities := decider.queryActivityInstances(workflowId)

	// get the last activity from history
	lastActivity := activities[len(activities)-1]
	var input string
	var nextActivity *ActivityNode
	if strings.EqualFold(lastActivity.Status, "COMPLETED") {
		// if the status of the last activity is COMPLETED, then start next activity
		input = lastActivity.Output
		nextActivity = getNextActivity(metadata, lastActivity.ActivityName)
	} else {
		// otherwise, retry the last activity
		input = lastActivity.Input
		nextActivity = getCurrActivity(metadata, lastActivity.ActivityName)
	}

	// respond workflow task, start an activity instance
	taskAttributes := RespondTaskAttributes{
		[]RespondTaskAttribute{
			{
				TaskType:            "ACTIVITY_START",
				ActivityTypeName:    nextActivity.Action,
				ActivityTypeVersion: nextActivity.Version,
				Name:                nextActivity.Name,
				Method:              "INVOKE",
				Input:               input,
				Timeout:             0,
				HeartbeatTimeout:    0,
				Delay:               0,
			},
		},
	}
	respondWorkflowTaskArgs := RespondWorkflowTaskArgs{
		TaskToken:      taskToken,
		Identity:       "127.0.0.1",
		TaskAttributes: taskAttributes,
	}
	_, err2 := client.RespondWorkflowTask(&respondWorkflowTaskArgs)
	if err2 != nil {
		fmt.Printf("Failed to respond workflow task %s: %v\n", taskToken, err2)
		return err2
	}
	return nil
}

func (decider *WorkflowDecider) handleActivityCompleteEvent(response *PollForWorkflowTaskResponse) error {
	pollTaskAttribute := response.TaskAttribute
	activityState := pollTaskAttribute.ActivityState
	method := pollTaskAttribute.Method

	if strings.EqualFold(activityState, "COMPLETED") && strings.EqualFold(method, "INVOKE") {
		return decider.handleActivityInvokeCompleteEvent(response)
	} else if strings.EqualFold(activityState, "FAILED") && strings.EqualFold(method, "INVOKE") {
		return decider.handleActivityInvokeFailedEvent(response)
	} else if strings.EqualFold(method, "CANCEL") {
		return decider.handleActivityRollbackEvent(response)
	}

	return errors.New("Uncognized activity method: " + method)
}

func (decider *WorkflowDecider) handleActivityInvokeCompleteEvent(response *PollForWorkflowTaskResponse) error {
	metadata := decider.Metadata
	client := decider.Client
	pollTaskAttribute := response.TaskAttribute
	activityName := pollTaskAttribute.Name
	nextActivity := getNextActivity(metadata, activityName)
	taskToken := response.TaskToken

	if nextActivity != nil {
		// if there is a next activity, then start it
		// the output of the current activity will become the input of the next activity
		input := pollTaskAttribute.Result

		// respond workflow task, start an activity instance
		taskAttributes := RespondTaskAttributes{
			[]RespondTaskAttribute{
				{
					TaskType:            "ACTIVITY_START",
					ActivityTypeName:    nextActivity.Action,
					ActivityTypeVersion: nextActivity.Version,
					Name:                nextActivity.Name,
					Method:              "INVOKE",
					Input:               input,
					Timeout:             0,
					HeartbeatTimeout:    0,
					Delay:               0,
				},
			},
		}
		respondWorkflowTaskArgs := RespondWorkflowTaskArgs{
			TaskToken:      taskToken,
			Identity:       "127.0.0.1",
			TaskAttributes: taskAttributes,
		}
		_, err := client.RespondWorkflowTask(&respondWorkflowTaskArgs)
		if err != nil {
			fmt.Printf("Failed to respond workflow task %s: %v", taskToken, err)
			return err
		}
	} else {
		// if can not find next activity, then complete workflow
		workflowId := response.WorkflowId
		taskAttributes := RespondTaskAttributes{
			[]RespondTaskAttribute{
				{
					TaskType:      "WORKFLOW_COMPLETE",
					RunId:         workflowId,
					WorkflowState: "COMPLETED",
				},
			},
		}
		respondWorkflowTaskArgs := RespondWorkflowTaskArgs{
			TaskToken:      taskToken,
			Identity:       "127.0.0.1",
			TaskAttributes: taskAttributes,
		}
		_, err := client.RespondWorkflowTask(&respondWorkflowTaskArgs)
		if err != nil {
			fmt.Printf("Failed to response workflow task %s: %v", taskToken, err)
			return err
		}
	}
	return nil
}

func (decider *WorkflowDecider) listActivityHistory(workflowId string) []ActivityHistory {
	totalPage := 1
	result := []ActivityHistory{}
	client := decider.Client
	for page := 1; page <= totalPage; page++ {
		listAvailableActivityHistoryArgs := ListAvailableActivityHistoryArgs{
			WorkflowId: workflowId,
			Pagination: common.Pagination{
				PageNumber: page,
				PageSize:   100,
			},
		}
		response, err := client.ListAvailableActivityHistory(&listAvailableActivityHistoryArgs)
		if err != nil {
			fmt.Printf("Failed to list available activity history %s: %v", workflowId, err)
			return result
		}
		activities := response.Activities.Activity
		fmt.Printf("Activity number is %d\n", len(activities))
		result = append(result, activities...)
		totalPage = (response.PaginationResult.TotalCount / 100) + 1
		fmt.Printf("Total page is %d\n", totalPage)
	}
	return result
}

func (decider *WorkflowDecider) handleActivityInvokeFailedEvent(response *PollForWorkflowTaskResponse) error {
	metadata := decider.Metadata
	client := decider.Client
	pollTaskAttribute := response.TaskAttribute
	activityName := pollTaskAttribute.Name
	taskToken := response.TaskToken
	workflowId := response.WorkflowId

	currentActivity := getCurrActivity(metadata, activityName)
	if currentActivity == nil {
		return errors.New("can not find activity " + activityName + " in workflow instance " + workflowId)
	}
	// get max retry number
	maxRetry := currentActivity.MaxRetry
	// count execution number of the activity
	countActivityInstanceArgs := CountActivityInstanceArgs{
		WorkflowId: workflowId,
		Name:       activityName,
	}
	count, err := client.CountActivityInstance(&countActivityInstanceArgs)
	if err != nil {
		fmt.Printf("Failed to count activity instance number %s: %v", activityName, err)
		return err
	}
	if (maxRetry == 0) || (count >= maxRetry+1) {
		// maxRetry equals 0 or retry number larger than maxRetry means rollback
		enableHumanTask := metadata.EnableHumanTask
		if enableHumanTask {
			// if enableHumanTask set to true, then suspend workflow instance
			suspendWorkflowArgs := SuspendWorkflowArgs{
				WorkflowId: workflowId,
			}
			_, err2 := client.SuspendWorkflow(&suspendWorkflowArgs)
			if err2 != nil {
				fmt.Printf("Failed to suspend workflow instance %s: %v\n", workflowId, err2)
				return err2
			}
		} else {
			// otherwise, find next rollback activity from history
			activities := decider.listActivityHistory(workflowId)
			cancelled := make(map[string]bool)
			var prevActivity *ActivityHistory
			for _, activity := range activities {
				activityName := activity.Name
				method := activity.Method
				if strings.EqualFold(method, "CANCEL") {
					cancelled[activityName] = true
				} else if _, present := cancelled[activityName]; !present {
					prevActivity = &activity
					break
				}
			}
			if prevActivity == nil {
				fmt.Printf("can not find COMPLETED activity in history, so complete workflow directly.\n")
				// can not find previous activity, complete workflow as FAILED
				workflowId := response.WorkflowId
				taskAttributes := RespondTaskAttributes{
					[]RespondTaskAttribute{
						{
							TaskType:      "WORKFLOW_COMPLETE",
							RunId:         workflowId,
							WorkflowState: "FAILED",
						},
					},
				}
				respondWorkflowTaskArgs := RespondWorkflowTaskArgs{
					TaskToken:      taskToken,
					Identity:       "127.0.0.1",
					TaskAttributes: taskAttributes,
				}
				_, err := client.RespondWorkflowTask(&respondWorkflowTaskArgs)
				if err != nil {
					fmt.Printf("Failed to response workflow task %s: %v", taskToken, err)
					return err
				}
			} else {
				// rollback the prev activity
				// the input of the current activity will become the input of the previous rollback activity
				fmt.Printf("find previous activity %s.\n", prevActivity.Name)
				input := prevActivity.Input

				// respond workflow task, rollback the previous activity instance
				taskAttributes := RespondTaskAttributes{
					[]RespondTaskAttribute{
						{
							TaskType:         "ACTIVITY_CANCEL",
							RunId:            prevActivity.RunId,
							Method:           "CANCEL",
							Input:            input,
							Timeout:          0,
							HeartbeatTimeout: 0,
							Delay:            0,
						},
					},
				}
				respondWorkflowTaskArgs := RespondWorkflowTaskArgs{
					TaskToken:      taskToken,
					Identity:       "127.0.0.1",
					TaskAttributes: taskAttributes,
				}
				_, err := client.RespondWorkflowTask(&respondWorkflowTaskArgs)
				if err != nil {
					fmt.Printf("Failed to respond workflow task %s: %v", taskToken, err)
					return err
				}
			}
		}
	} else {
		// otherwise reinitiate a retry request
		// the input of the current activity will become the input of the next retry activity
		input := pollTaskAttribute.Input

		// respond workflow task, rollback the previous activity instance
		taskAttributes := RespondTaskAttributes{
			[]RespondTaskAttribute{
				{
					TaskType:            "ACTIVITY_START",
					ActivityTypeName:    currentActivity.Action,
					ActivityTypeVersion: currentActivity.Version,
					Name:                currentActivity.Name,
					Method:              "INVOKE",
					Input:               input,
					Timeout:             0,
					HeartbeatTimeout:    0,
					Delay:               0,
				},
			},
		}
		respondWorkflowTaskArgs := RespondWorkflowTaskArgs{
			TaskToken:      taskToken,
			Identity:       "127.0.0.1",
			TaskAttributes: taskAttributes,
		}
		_, err := client.RespondWorkflowTask(&respondWorkflowTaskArgs)
		if err != nil {
			fmt.Printf("Failed to respond workflow task %s: %v", taskToken, err)
			return err
		}
	}

	return nil
}

func (decider *WorkflowDecider) handleActivityRollbackEvent(response *PollForWorkflowTaskResponse) error {
	client := decider.Client
	pollTaskAttribute := response.TaskAttribute
	taskToken := response.TaskToken
	workflowId := response.WorkflowId

	// find next rollback activity
	listAvailableActivityHistoryArgs := ListAvailableActivityHistoryArgs{
		WorkflowId: workflowId,
		Pagination: common.Pagination{
			PageNumber: 1,
			PageSize:   100,
		},
	}
	listResponse, err2 := client.ListAvailableActivityHistory(&listAvailableActivityHistoryArgs)
	if err2 != nil {
		fmt.Printf("Failed to list available activity history %s: %v", workflowId, err2)
		return err2
	}
	activities := listResponse.Activities.Activity
	cancelled := make(map[string]bool)
	var prevActivity *ActivityHistory
	for _, activity := range activities {
		activityName := activity.Name
		method := activity.Method
		if strings.EqualFold(method, "CANCEL") {
			cancelled[activityName] = true
		} else if _, present := cancelled[activityName]; !present {
			prevActivity = &activity
			break
		}
	}

	if prevActivity == nil {
		// If can not find previous activity, complete workflow as FAILED
		taskAttributes := RespondTaskAttributes{
			[]RespondTaskAttribute{
				{
					TaskType:      "WORKFLOW_COMPLETE",
					RunId:         workflowId,
					WorkflowState: "FAILED",
				},
			},
		}
		respondWorkflowTaskArgs := RespondWorkflowTaskArgs{
			TaskToken:      taskToken,
			Identity:       "127.0.0.1",
			TaskAttributes: taskAttributes,
		}
		_, err := client.RespondWorkflowTask(&respondWorkflowTaskArgs)
		if err != nil {
			fmt.Printf("Failed to response workflow task %s: %v", taskToken, err)
			return err
		}
	} else {
		// otherwise, rollback previous activity
		// the input of the current activity will become the input of the previous rollback activity
		input := pollTaskAttribute.Input

		// respond workflow task, rollback the previous activity instance
		taskAttributes := RespondTaskAttributes{
			[]RespondTaskAttribute{
				{
					TaskType:         "ACTIVITY_CANCEL",
					RunId:            prevActivity.RunId,
					Method:           "CANCEL",
					Input:            input,
					Timeout:          0,
					HeartbeatTimeout: 0,
					Delay:            0,
				},
			},
		}
		respondWorkflowTaskArgs := RespondWorkflowTaskArgs{
			TaskToken:      taskToken,
			Identity:       "127.0.0.1",
			TaskAttributes: taskAttributes,
		}
		_, err := client.RespondWorkflowTask(&respondWorkflowTaskArgs)
		if err != nil {
			fmt.Printf("Failed to respond workflow task %s: %v", taskToken, err)
			return err
		}
	}
	return nil
}
