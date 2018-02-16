package pvtz

import (
	"github.com/denverdino/aliyungo/common"
	//"github.com/denverdino/aliyungo/util"
)

type DescribeZoneRecordsArgs struct {
	ZoneId       string
	Keyword      string
	Lang         string
	UserClientIp string
	common.Pagination
}

//
type ZoneRecordType struct {
	Rr       string
	Type     string
	Ttl      int
	Priority int
	Value    string
	Status   string
}

type DescribeZoneRecordsResponse struct {
	common.Response
	common.PaginationResult
	Records struct {
		Record []ZoneRecordType
	}
}

// DescribeZoneRecords describes zones
//
// You can read doc at https://help.aliyun.com/document_detail/66252.html
func (client *Client) DescribeZoneRecords(args *DescribeZoneRecordsArgs) (records []ZoneRecordType, err error) {

	result := []ZoneRecordType{}

	for {
		response := DescribeZoneRecordsResponse{}
		err = client.Invoke("DescribeZoneRecords", args, &response)

		if err != nil {
			return result, err
		}

		result = append(result, response.Records.Record...)

		nextPage := response.PaginationResult.NextPage()
		if nextPage == nil {
			break
		}
		args.Pagination = *nextPage
	}

	return result, nil
}

type AddZoneRecordArgs struct {
	ZoneName     string
	Rr           string
	Type         string
	Value        string
	ZoneId       string
	Lang         string
	Priority     int
	Ttl          int
	UserClientIp string
}

type AddZoneRecordResponse struct {
	common.Response
	Success  bool
	RecordId int64
}

// AddZoneRecord add zone record
//
// You can read doc at https://help.aliyun.com/document_detail/66248.html
func (client *Client) AddZoneRecord(args *AddZoneRecordArgs) (response *AddZoneRecordResponse, err error) {
	response = &AddZoneRecordResponse{}

	err = client.Invoke("AddZoneRecord", args, &response)

	return response, err
}


type DeleteZoneRecordArgs struct {
	RecordId     int64
	Lang         string
	UserClientIp string
}

type DeleteZoneRecordResponse struct {
	common.Response
	ZoneId   string
}

// DeleteZone delete zone
//
// You can read doc at https://help.aliyun.com/document_detail/66249.html
func (client *Client) DeleteZoneRecord(args *DeleteZoneRecordArgs) (err error) {
	response := &DeleteZoneRecordResponse{}
	err = client.Invoke("DeleteZoneRecord", args, &response)

	return err
}
