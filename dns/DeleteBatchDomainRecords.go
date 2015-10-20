package dns

import (
	"github.com/denverdino/aliyungo/common"
	"log"
)

type DeleteDomainRecordItem struct{
	Index string
	Domain string
	RecordId string	
}

type DeleteBatchDomainRecordsArgs struct {
	Records []DeleteDomainRecordItem
}

type DeleteBatchDomainRecordsResponse struct {
	common.Response
	InstanceId string
	TraceId   string
}

// DeleteBatchDomainRecords
//
// You can read doc at https://docs.aliyun.com/?spm=5176.100054.201.106.OeZ3dN#/pub/dns/api-reference/batch-related&DeleteBatchDomainRecords
func (client *Client) DeleteBatchDomainRecords(args *DeleteBatchDomainRecordsArgs) (response *DeleteBatchDomainRecordsResponse, err error) {
	action := "DeleteBatchDomainRecords"
	response = &DeleteBatchDomainRecordsResponse{}
	err = client.InvokePost(action, "Records",args, response)
	if err == nil {
		return response, nil
	} else {
		log.Fatalf("%s error, %v", action, err)
		return response, err
	}
}
