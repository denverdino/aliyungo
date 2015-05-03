package ecs

import (
	"github.com/denverdino/aliyungo/util"
)

// Describe Images
type DescribeImagesArgs struct {
	RegionId        string
	ImageId         string
	SnapshotId      string
	ImageName       string
	ImageOwnerAlias string //Enum system | self | others | marketplace
	Pagination
}

type DescribeImagesResponse struct {
	CommonResponse

	Pagination string
	PaginationResult
	Images struct {
		Image []ImageType
	}
}

type DiskDeviceMapping struct {
	SnapshotId string
	//Why Size Field is string-type.
	Size   string
	Device string
}

type ImageType struct {
	ImageId            string
	ImageVersion       string
	Architecture       string
	ImageName          string
	Description        string
	Size               int
	ImageOwnerAlias    string
	OSName             string
	DiskDeviceMappings struct {
		DiskDeviceMapping []DiskDeviceMapping
	}
	ProductCode  string
	IsSubscribed bool
	Progress     string
	Status       string
	CreationTime util.ISO6801Time
}

// Describe Images
func (client *Client) DescribeImages(args *DescribeImagesArgs) (images []ImageType, pagination *PaginationResult, err error) {

	args.validate()
	response := DescribeImagesResponse{}
	err = client.Invoke("DescribeImages", args, &response)
	if err != nil {
		return nil, nil, err
	}
	return response.Images.Image, &response.PaginationResult, nil
}

// Create Image
type CreateImageArgs struct {
	RegionId     string
	SnapshotId   string
	ImageName    string
	ImageVersion string
	Description  string
	ClientToken  string
}

type CreateImageResponse struct {
	CommonResponse

	ImageId string
}

func (client *Client) CreateImage(args *CreateImageArgs) (imageId string, err error) {
	response := &CreateImageResponse{}
	err = client.Invoke("CreateImage", args, &response)
	if err != nil {
		return "", err
	}
	return response.ImageId, nil
}

// Delete Image
type DeleteImageArgs struct {
	RegionId string
	ImageId  string
}

type DeleteImageResponse struct {
	CommonResponse
}

func (client *Client) DeleteImage(regionId string, imageId string) error {
	args := DeleteImageArgs{
		RegionId: regionId,
		ImageId:  imageId,
	}

	response := &DeleteImageResponse{}
	return client.Invoke("DeleteImage", &args, &response)
}
