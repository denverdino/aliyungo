package dns

import (
	"testing"
)

func TestDescribeDomainRecordInfoN(t *testing.T) {
	//prepare
	client := NewTestClientNew()
	describeArgs := DescribeDomainRecordsNArgs{
		DomainName: TestDomainName,
	}
	describeArgs.PageSize = 100

	describeResponse, err := client.DescribeDomainRecordsN(&describeArgs)
	if err == nil {
		record := describeResponse.DomainRecords.Record[0]
		arg := DescribeDomainRecordInfoNArgs{
			RecordId: record.RecordId,
		}
		response, err := client.DescribeDomainRecordInfoN(&arg)
		if err == nil {
			t.Logf("DescribeDomainRecordInfo success: %v", response)
		} else {
			t.Errorf("Failed to DescribeDomainRecordInfo: %s", describeArgs.DomainName)
		}
	} else {
		t.Errorf("Failed to DescribeDomainRecords: %s", describeArgs.DomainName)
	}
}
