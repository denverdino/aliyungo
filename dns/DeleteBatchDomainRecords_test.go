package dns

import (
	"testing"
)

func TestDeleteBatchDomainRecords(t *testing.T) {
/*
	//prepare
	recordNum := 0
	RRPattern := "testdeletebatchdomainrecords"
	client := NewTestClient()
	
	for i:=0;i<recordNum;i++ {
		addDomainRecordArgs := AddDomainRecordArgs{
			DomainName: TestDomainName,
			RR:         RRPattern + strconv.Itoa(i),
			Type:       ARecord,
			Value:      "8.8.8.8",
		}
		client.AddDomainRecord(&addDomainRecordArgs)
	}
	
	//prepare
	describeArgs := DescribeDomainRecordsArgs{
		DomainName: TestDomainName,
		RRKeyWord: RRPattern,
		PageSize:   500,
	}
	describeResponse, err := client.DescribeDomainRecords(&describeArgs)
	if (err != nil) || ( describeResponse.TotalCount <= 0 ) {
		t.Fatalf("Failed to DescribeDomainRecords: %v", describeResponse)
	}
	
	deleteArgs := DeleteBatchDomainRecordsArgs{}
	deleteArgs.Records = make([]DeleteDomainRecordItem, describeResponse.TotalCount)
	for i:=int32(0);i<describeResponse.TotalCount;i++{
		deleteArgs.Records[i].Domain = TestDomainName;
		deleteArgs.Records[i].RecordId = describeResponse.DomainRecords.Record[i].RecordId
	}
	response, err := client.DeleteBatchDomainRecords(&deleteArgs)
	if err == nil {
		t.Logf("DeleteBatchDomainRecords: %v", response)
	} else {
		t.Fatalf("Failed to DeleteBatchDomainRecords: %v", deleteArgs)
	}
*/
}
