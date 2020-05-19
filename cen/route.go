package cen

import (
	"github.com/labstack/gommon/log"
	"github.com/denverdino/aliyungo/common"
)

type PublishRouteEntriesArgs struct {
	CenId  					string
	ChildInstanceId 		string
	ChildInstanceRegionId  	string
	ChildInstanceRouteTableId string
	ChildInstanceType         string
	DestinationCidrBlock      string
}

type DescribePublishedRouteEntriesArgs struct {
	common.Pagination
	CenId 			string
	ChildInstanceId string
	ChildInstanceRegionId 		string
	ChildInstanceType  			string
	ChildInstanceRouteTableId  	string
	DestinationCidrBlock 		string
}

type DescribePublishedRouteEntriesResponse struct {
	common.Response
	common.PaginationResult
	PublishedRouteEntries struct {
		PublishedRouteEntry []PublishedRouteEntry
	}
}

type ConflictStatus string
type NextHopType    string

const (
	ConflictStatusConflict   = ConflictStatus("conflict")
	ConflictStatusOverflow   = ConflictStatus("overflow")
	ConflictStatusProhibited = ConflictStatus("prohibited")
)

const (
	NextHopTypeInstance = NextHopType("Instance")
	NextHopTypeHaVip	= NextHopType("HaVip")
	NextHopTypeRouterInterface	= NextHopType("RouterInterface")
)

type PublishStatus string

const (
	PublishStatusPublished	= PublishStatus("Published")
	PublishStatusNotPublished = PublishStatus("NonPublished")
)

type RouteType string
const (
	RouteTypeSystem = RouteType("System")
	RouteTypeCustom = RouteType("Custom")
	RouteTypeBGP    = RouteType("BGP")
)

type PublishedRouteEntry struct {
	ChildInstanceRouteTableId  string
	Conflicts	struct{
		Conflict []Conflict
	}
	DestinationCidrBlock  string
	NextHopId		string

	NextHopType		string
	OperationalMode string
	PublishStatus   string
	RouteType		string
}

type Conflict struct {
	DestinationCidrBlock	string
	InstanceId		string
	InstanceType	string
	RegionId		string
	Status			string
}

// PublishRouteEntries publish route
//
// You can read doc at https://help.aliyun.com/document_detail/85470.html
func (client *Client) PublishRouteEntries(args *PublishRouteEntriesArgs) error {
	response := common.Response{}
	err := client.Invoke("PublishRouteEntries", &args, &response)
	if err != nil {
		log.Errorf("PublishRouteEntries: %s, %s",response.RequestId, err.Error())
	}
	return err
}


// DescribePublishedRouteEntries describe published route
//
// You can read doc at https://help.aliyun.com/document_detail/85472.html
func (client *Client) DescribePublishedRouteEntries(
	args *DescribePublishedRouteEntriesArgs,
) (response *DescribePublishedRouteEntriesResponse, err error) {

	response = &DescribePublishedRouteEntriesResponse{}

	err = client.Invoke("DescribePublishedRouteEntries", args, response)

	if err != nil {
		log.Errorf("DescribePublishedRouteEntries: %s, %s", response)
	}

	return response, err
}
