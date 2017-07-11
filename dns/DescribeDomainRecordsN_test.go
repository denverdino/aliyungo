package dns

import (
	"testing"
)

func TestDescribeDomainRecordsN(t *testing.T) {
	//prepare
	client := NewTestClientNew()
	describeArgs := DescribeDomainRecordsNArgs{
		DomainName: TestDomainName,
	}
	describeArgs.PageSize = 100

	describeResponse, err := client.DescribeDomainRecordsN(&describeArgs)
	if err == nil {
		t.Logf("DescribeDomainRecords success: TotalCount:%d ", describeResponse.TotalCount)
	} else {
		t.Errorf("Failed to DescribeDomainRecords: %s", describeArgs.DomainName)
	}
}
