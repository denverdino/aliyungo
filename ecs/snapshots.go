package ecs

import (
	"github.com/denverdino/aliyungo/util"
)

type DescribeSnapshotsArgs struct {
	RegionId    Region
	InstanceId  string
	DiskId      string
	SnapshotIds []string //["s-xxxxxxxxx", "s-yyyyyyyyy", ..."s-zzzzzzzzz"]
	Pagination
}

type SnapshotType struct {
	SnapshotId     string
	SnapshotName   string
	Description    string
	Progress       string
	SourceDiskId   string
	SourceDiskSize string //GB Why it is string
	SourceDiskType string //enum for System | Data
	ProductCode    string
	CreationTime   util.ISO6801Time
}

type DescribeSnapshotsResponse struct {
	CommonResponse
	PaginationResult
	Snapshots struct {
		Snapshot []SnapshotType
	}
}

// DescribeSnapshots describe snapshots
func (client *Client) DescribeSnapshots(args *DescribeSnapshotsArgs) (snapshots []SnapshotType, pagination *PaginationResult, err error) {
	args.validate()
	response := DescribeSnapshotsResponse{}

	err = client.Invoke("DescribeSnapshots", args, &response)

	if err != nil {
		return nil, nil, err
	}
	return response.Snapshots.Snapshot, &response.PaginationResult, nil

}

type DeleteSnapshotArgs struct {
	SnapshotId string
}

type DeleteSnapshotResponse struct {
	CommonResponse
}

// DeleteSnapshot deletes snapshot
func (client *Client) DeleteSnapshot(snapshotId string) error {
	args := DeleteSnapshotArgs{SnapshotId: snapshotId}
	response := DeleteSnapshotResponse{}

	return client.Invoke("DeleteSnapshot", &args, &response)
}

type CreateSnapshotArgs struct {
	DiskId       string
	SnapshotName string
	Description  string
	ClientToken  string
}

type CreateSnapshotResponse struct {
	CommonResponse
	SnapshotId string
}

// CreateSnapshot creates a new snapshot
func (client *Client) CreateSnapshot(args *CreateSnapshotArgs) (snapshotId string, err error) {

	response := CreateSnapshotResponse{}

	err = client.Invoke("CreateSnapshot", args, &response)
	if err == nil {
		snapshotId = response.SnapshotId
	}
	return snapshotId, err
}

// Default timeout value for WaitForSnapShotReady method
const SnapshotDefaultTimeout = 120

// WaitForVSwitchAvailable waits for VSwitch to given status
func (client *Client) WaitForSnapShotReady(regionId Region, snapshotId string, strategy util.AttemptStrategy) (status interface{}, err error) {

	fn := func() (bool, interface{}, error) {

		args := DescribeSnapshotsArgs{
			RegionId:    regionId,
			SnapshotIds: []string{snapshotId},
		}

		snapshots, _, err := client.DescribeSnapshots(&args)
		if err != nil {
			return false, "N/A", err
		}
		if snapshots == nil || len(snapshots) == 0 {
			return false, "N/A", getECSErrorFromString("Not found")
		}
		if snapshots[0].Progress == "100%" {
			return true, snapshots[0].Progress, nil
		}

		return false, "N/A", nil
	}

	status, e1 := util.LoopCall(strategy, fn)

	return status, e1
}
