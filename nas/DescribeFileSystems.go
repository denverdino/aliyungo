package nas

import "github.com/denverdino/aliyungo/common"

type DescribeFileSystemsRequest struct {
	FileSystemId string
	PageSize     int
	PageNumber   int
	RegionId     common.Region
}

type DescribeFileSystemsResponse struct {
	common.Response
	FileSystems struct {
		FileSystem []FileSystem
	}
	TotalCount int
	PageSize   int
	PageNumber int
}

type FileSystem struct {
	StorageType  string
	ProtocolType string
	CreateTime   string
	Destription  string
	MountTargets struct {
		MountTarget []MountTarget
	}
	FileSystemId string
	RegionId     common.Region
	MeteredSize  int64
	Packages     struct {
		Package []Package
	}
}

type Package struct {
	PackageId string
}

func (client *Client) DescribeFileSystems(args *DescribeFileSystemsRequest) (resp DescribeFileSystemsResponse, err error) {
	response := DescribeFileSystemsResponse{}
	//args.Version = VERSION
	err = client.Invoke("DescribeFileSystems", args, &response)
	return response, err
}
