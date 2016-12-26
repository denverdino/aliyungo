package grandcanal

import (
	"encoding/xml"
	"strings"
)

type Workflows struct {
	XMLName xml.Name       `xml:"workflows"`
	Flows   []SequenceFlow `xml:"sequence-flow"`
}

type SequenceFlow struct {
	XMLName         xml.Name       `xml:"sequence-flow"`
	Name            string         `xml:"name,attr"`
	Version         string         `xml:"version,attr"`
	EnableHumanTask bool           `xml:"enableHumanTask,attr"`
	Activities      []ActivityNode `xml:"activity"`
}

type ActivityNode struct {
	XMLName       xml.Name `xml:"activity"`
	Name          string   `xml:"name,attr"`
	Action        string   `xml:"action,attr"`
	Version       string   `xml:"version,attr"`
	MaxRetry      int32    `xml:"maxRetry,attr"`
	RetryInterval int32    `xml:"retryInterval,attr"`
}

func getFirstActivity(flow *SequenceFlow) (activity *ActivityNode) {
	if len(flow.Activities) > 0 {
		return &flow.Activities[0]
	}
	return nil
}

func getCurrActivity(flow *SequenceFlow, activityName string) (activity *ActivityNode) {
	if len(flow.Activities) == 0 {
		return nil
	}
	for _, activity := range flow.Activities {
		if strings.EqualFold(activity.Name, activityName) {
			return &activity
		}
	}
	return nil
}

func getNextActivity(flow *SequenceFlow, activityName string) (activity *ActivityNode) {
	if len(flow.Activities) == 0 {
		return nil
	}
	pos := 0
	for i, activity := range flow.Activities {
		if strings.EqualFold(activity.Name, activityName) {
			pos = i
			break
		}
	}
	if pos < (len(flow.Activities) - 1) {
		return &flow.Activities[pos+1]
	}
	return nil
}

func getPrevActivity(flow *SequenceFlow, activityName string) (activity *ActivityNode) {
	if len(flow.Activities) == 0 {
		return nil
	}
	pos := 0
	for i, activity := range flow.Activities {
		if strings.EqualFold(activity.Name, activityName) {
			pos = i
			break
		}
	}
	if pos > 0 {
		return &flow.Activities[pos-1]
	}
	return nil
}
