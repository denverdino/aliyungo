package ecs

import (
	"github.com/denverdino/aliyun-go/util"
)

// Describe `DescribeSnapshots`
type DescribeSnapshotsArgs struct {
	RegionId    string
	InstanceId  string
	DiskId      string
	SnapshotIds []string //Json Array：["s-xxxxxxxxx", "s-yyyyyyyyy", ..."s-zzzzzzzzz"]
	Pagination
}

type SnapshotType struct {
	SnapshotId     string
	SnapshotName   string
	Description    string
	Progress       string
	SourceDiskId   string
	SourceDiskSize string //GB Why it is string
	SourceDiskType string //System | Data
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

func (client *Client) DescribeSnapshots(args *DescribeSnapshotsArgs) (snapshots []SnapshotType, pagination *PaginationResult, err *ECSError) {
	args.validate()
	response := DescribeSnapshotsResponse{}

	err = client.Invoke("DescribeSnapshots", args, &response)

	if err != nil {
		return nil, nil, err
	}
	return response.Snapshots.Snapshot, &response.PaginationResult, nil

}

// Describe `DescribeSnapshots`
type DeleteSnapshotArgs struct {
	SnapshotId string
}

type DeleteSnapshotResponse struct {
	CommonResponse
}

func (client *Client) DeleteSnapshot(snapshotId string) *ECSError {
	args := DeleteSnapshotArgs{SnapshotId: snapshotId}
	response := DeleteSnapshotResponse{}

	return client.Invoke("DeleteSnapshot", &args, &response)
}

// Describe `DescribeSnapshots`
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

func (client *Client) CreateSnapshot(args *CreateSnapshotArgs) (snapshotId string, err *ECSError) {

	response := CreateSnapshotResponse{}

	err = client.Invoke("CreateSnapshot", args, &response)
	if err == nil {
		snapshotId = response.SnapshotId
	}
	return snapshotId, err
}
